package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitiateDB sets up the database connection and runs migrations.
func InitiateDB() (*gorm.DB, error) {
	// Get the database configuration details
	dbConfig := newDbConfig()

	// Run database migrations (make sure your migrations are defined properly)
	err := dbMigrationUP(dbConfig)
	if err != nil {
		// If migrations fail, return the error
		return nil, fmt.Errorf("error running migrations: %v", err)
	}

	// Open the database connection with GORM
	db, err := gorm.Open(postgres.Open(dbConfig.connString), &gorm.Config{
		// Set logging level (silent for production, info for debugging)
		Logger: logger.Default.LogMode(logger.Silent), // Change to logger.Info for debugging

		// Set a custom naming strategy (for schema and table names)
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "student_schema.", // Prefix for all table names (schema name)
			SingularTable: false,            // Pluralize table names (students)
		},
	})

	// If the database connection fails, return the error
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Log success and return the DB connection
	log.Println("Connected to the database successfully!")
	return db, nil
}
