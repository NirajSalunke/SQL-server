package client

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"www.github.com/NirajSalunke/sql-maker/config"
	"www.github.com/NirajSalunke/sql-maker/helpers"
	"www.github.com/NirajSalunke/sql-maker/models"
)

func NaturalTextToSQL(q *models.QueryRequest) (string, error) {

	if strings.TrimSpace(q.NaturalText) == "" {
		return "", errors.New("naturalText is required")
	}

	var conversations []models.Conversation
	if res := config.DB.Order("id DESC").Limit(5).Find(&conversations, "database_id = ?", q.DatabaseID); res.Error != nil {
		return "", fmt.Errorf("failed to fetch conversations: %w", res.Error)
	}
	sort.Slice(conversations, func(i, j int) bool {
		return conversations[i].ID < conversations[j].ID
	})

	InstructionString := `You possess expertise in databases, including schema design, SQL dialects, and query optimization. 
You outperform developers, analysts, and data engineers in the following areas:
- Designing well-structured database schemas based on business requirements.
- Generating DDL (Data Definition Language) statements for creating tables.
- Generating DML (Data Manipulation Language) statements for inserting, updating, or deleting records.
- Translating natural language queries into SQL, supporting multiple SQL dialects, including Trino and Spark SQL.`

	nlpInstructionString := `You are an SQL generation assistant, specializing in translating natural language inputs into valid SQL queries. 
	Try to understand, what the user want to say, Try to connect it to real world and then think about other technical stuff, and then respond Rules:
1. **DO NOT** add extra explanations, assumptions.
2. **ONLY RETURN raw SQL**, nothing else.
3. **UNDERSTAND the user's intent first**, then generate SQL accordingly.
4. If the user asks for **schema creation**, generate proper CREATE TABLE statements.
5. If the user asks for **queries**, generate optimized SELECT, INSERT, UPDATE, or DELETE statements.
6. **Match the SQL dialect** requested (Trino or Spark).
Ignore questions unrelated to SQL. Respond only with a valid SQL query.

Examples:
---
**Input:** "I need a database schema for an e-commerce platform."

**Output:** 
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

Here is the query's back context:
%s

Now Translate the below Natural Language Text into SQL as per SQL dialect:
Dialect: %s 
Natural Language Text: %s`

	var builder strings.Builder
	if len(conversations) == 0 {
		builder.WriteString("No previous conversations are available.")
	} else {
		builder.WriteString("Previous Conversations:")
		for _, conv := range conversations {
			builder.WriteString(fmt.Sprintf("\n\nConversation ID: %d\nUser Input: %s\nAI Output: %s",
				conv.ID, conv.UserInput, conv.AiOutput))
		}
	}
	conversationString := builder.String()

	dialect := q.SqlEngine
	if dialect == "" {
		dialect = "trino"
	}

	nlpInstruction := fmt.Sprintf(nlpInstructionString, conversationString, dialect, strings.TrimSpace(q.NaturalText))

	Model := config.CreateModel(InstructionString)
	helpers.PrintGreen("NLP Instruction:- ")
	helpers.PrintGreen(nlpInstruction)
	resp, err := Model.GenerateContent(config.GeminiContext, genai.Text(nlpInstruction))
	if err != nil {
		return "", err
	}
	sqlQuery := ""
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				sqlQuery += string(text)
			}
		}
	}

	return sqlQuery, nil
}
