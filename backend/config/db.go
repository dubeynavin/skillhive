package config

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() *mongo.Database {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get URI from .env
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env")
	}

	// Mask password for logging (do not print credentials)
	masked := maskMongoURI(mongoURI)
	fmt.Println("Using MONGO_URI:", masked)

	// Set client options with a server selection timeout to get faster, clearer errors
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(30 * time.Second)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Cannot connect to MongoDB:")
		log.Println(err)
		// Give a hint for common Atlas issues
		log.Println("Hints: ensure your Atlas IP access list allows this machine's IP, the user/password are correct (URL-encode special characters), and the connection string uses the correct SRV or direct format.")
		log.Fatal(err)
	}

	fmt.Println("✅ Connected to MongoDB successfully!")

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "skillhiveai"
	}
	DB = client.Database(dbName)
	return DB
}

// maskMongoURI returns a masked version of the Mongo URI for safe logging.
func maskMongoURI(raw string) string {
	// Try to parse as URL and mask password
	if strings.HasPrefix(raw, "mongodb+srv://") || strings.HasPrefix(raw, "mongodb://") {
		// replace mongodb+srv:// with http scheme for parsing convenience
		u := raw
		// ensure proper scheme for url.Parse
		parsed, err := url.Parse(u)
		if err == nil && parsed.User != nil {
			username := parsed.User.Username()
			if pass, ok := parsed.User.Password(); ok && pass != "" {
				parsed.User = url.UserPassword(username, "*****")
			} else {
				parsed.User = url.User(username)
			}
			return parsed.String()
		}
	}
	// fallback: redact between @ and //@ or after // up to next /
	if at := strings.Index(raw, "@"); at != -1 {
		// find start of host portion
		start := strings.Index(raw, "//")
		if start != -1 && start < at {
			return raw[:start+2] + "*****@" + raw[at+1:]
		}
	}
	return raw
}
