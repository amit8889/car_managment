package main

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strings"

	"github.com/amit8889/car_managemnt/driver"
	"github.com/amit8889/car_managemnt/router"
	"github.com/amit8889/car_managemnt/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// start := time.Now()
	// numUUIDs := 10000000 // Number of UUIDs to generate

	// for i := 0; i < numUUIDs; i++ {
	// 	// Generate a random UUID
	// 	uuid.New()
	// }

	// duration := time.Since(start)
	// fmt.Printf("Generated %d UUIDs in %v, which is %f UUIDs per second\n", numUUIDs, duration, float64(numUUIDs)/duration.Seconds())
	//load env
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", "")

	}
	//init database
	driver.InitDB()
	db := driver.GetDB()

	//make table
	schemaFile := "./store/schema.sql"
	if executeSchemaFile(db, schemaFile) != nil {
		logger.Error("Error executing schema file", "")
	}
	r := gin.Default()
	router.InitRouter(db, r)
	//get Port
	port := os.Getenv("PORT")
	if port == "" {
		logger.Error("$PORT must be set", "")
	}
	logger.Success("Server is running on port", port)
	r.Run(":" + port)
}

func executeSchemaFile(db *sql.DB, sqlFile string) error {
	// Read the SQL file content
	sqlContent, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		logger.Error("Failed to read the SQL file: %v", err)
	}
	sqlStatements := strings.Split(string(sqlContent), ";")
	for _, stmt := range sqlStatements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		// Execute each statement
		_, err := db.Exec(stmt)
		if err != nil {
			// Log the error and return
			logger.Warn("Error executing SQL statement: %v", err)
		}
	}
	// Success log
	logger.Info("Tables created successfully.", "")
	return nil
}
