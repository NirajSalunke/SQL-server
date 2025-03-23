package config

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"www.github.com/NirajSalunke/sql-maker/helpers"
)

var Client *genai.Client
var GeminiContext context.Context
var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		helpers.PrintRed("Error loading .env file")
	}
	helpers.PrintGreen("Env Variables loaded!")
}

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		helpers.PrintRed("Error:- Failed to connect to database:- ")
		log.Fatal(err.Error())
	}
	helpers.PrintGreen("Connected to Database Successfully.")
}

func SetupGeminiClient() {

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		helpers.PrintRed("GEMINI_API_KEY environment variable is not set or empty")
	}
	ctx := context.Background()
	var err error
	Client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		helpers.PrintRed("Failed to create Gemini client: ")
		log.Fatal(err)
	}
	GeminiContext = ctx
	helpers.PrintGreen("Created Client Successfully")
}

func CreateModel(instructionString string) *genai.GenerativeModel {
	modelName := "gemini-1.5-flash-001"

	if instructionString == "" {
		panic("Instruction string cannot be empty")
	}

	ModelAI := Client.GenerativeModel(modelName)
	ModelAI.SystemInstruction = genai.NewUserContent(genai.Text(instructionString))

	return ModelAI
}
