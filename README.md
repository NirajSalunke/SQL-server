# SQL-Server

A Go-based HTTP API that converts natural-language questions into SQL queries using Google’s Gemini model.

---

## ✨ What It Does

### 🔄 Natural-Language → SQL

* Transforms plain-English questions into syntactically valid SQL queries.

### 📁 RESTful Endpoint

* **Endpoint:** `/api/v1/query`
* **Request Example:**

  ```json
  {
    "question": "Find all orders over $100 placed in the last week"
  }
  ```
* **Response Example:**

  ```json
  {
    "sql": "SELECT * FROM orders WHERE amount > 100 AND order_date >= NOW() - INTERVAL '7 days';"
  }
  ```

### ⚖️ Dialect Flexibility

* Generates ANSI-compliant SQL that can be adapted for MySQL, PostgreSQL, SQL Server, and other database systems.

---

## 🔧 What It Uses

### 💀 Go (1.19+)

* Fast, statically compiled backend server.

### 🌐 Gin Web Framework

* Lightweight and efficient routing with middleware support.

### 🌀 Google Gemini API

* Powered by a cutting-edge large-language model for natural-language understanding and SQL generation.

### 🔠 YAML Configuration

* Flexible configuration for server settings, Gemini model parameters, timeouts, and logging.

### 🔄 Modular Code Structure

* **`controllers/`**: Handles HTTP requests.
* **`helpers/`**: Utilities for prompt building and Gemini client interaction.
* **`models/`**: Data models for requests and responses.
* **`routes/`**: API endpoint wiring.

---

## 🚀 Use Cases

### 🔄 BI Dashboards

* Enable users to query datasets in plain English, generating SQL for direct use in visualization tools.

### 🔐 Data Exploration Tools

* Speed up query prototyping for analysts without requiring SQL knowledge.

### 🖬 Chatbot Integrations

* Enhance conversational UIs with natural-language-to-SQL translation capabilities.

### 🔨 Learning & Prototyping

* Demonstrate the relationship between natural-language inputs and SQL patterns, aiding education and experimentation.

---

## 🔖 License

Distributed under the **MIT License**.
