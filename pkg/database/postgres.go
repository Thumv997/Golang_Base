package database

import (
	"fmt"
	config "lore_project/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB represents the database connection
type DB struct {
	conn *gorm.DB
}

// NewDB creates a new database connection
func NewDB(cfg config.Config) (*DB, error) {
	// Tạo chuỗi kết nối PostgreSQL
	pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName)
	db, err := gorm.Open(postgres.Open(pgConnStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

func (db *DB) AutoMigrate(models ...interface{}) error {
	err := db.conn.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	dbClose, err := db.conn.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying DB: %w", err)
	}
	if err := dbClose.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}

// Query executes a raw SQL query and returns the result
func (db *DB) Query(query string, args ...interface{}) (*gorm.DB, error) {
	result := db.conn.Raw(query, args...)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to execute query: %w", result.Error)
	}
	return result, nil
}

// Create creates a new record in the database
func (db *DB) Create(data interface{}) error {
	result := db.conn.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates an existing record in the database
func (db *DB) Update(data interface{}) error {
	result := db.conn.Save(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete deletes a record from the database
func (db *DB) Delete(data interface{}) error {
	result := db.conn.Delete(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
