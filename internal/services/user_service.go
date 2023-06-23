package services

import (
	"fmt"
	"lore_project/internal/models"
	"lore_project/pkg/database"
)

// UserService represents the service for user-related operations.
type UserService struct {
	db *database.DB
}

// NewUserService creates a new instance of UserService.
func NewUserService(db *database.DB) *UserService {

	return &UserService{
		db: db,
	}
}

// CreateUser creates a new user in the database.
func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user)
}

// GetUserByID retrieves a user from the database based on the ID.
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user := new(models.User)
	query := "SELECT * FROM users WHERE id = ?"
	result, _ := s.db.Query(query, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to execute query: %w", result.Error)
	}
	defer s.db.Close()

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	err := result.Scan(user)
	if err != nil {
		return nil, fmt.Errorf("failed to scan user: %v", err)
	}

	return user, nil
}

// GetUserByEmail retrieves a user from the database based on the email.
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	query := "SELECT * FROM users WHERE email = ?"
	result, err := s.db.Query(query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	err = result.Scan(user).Error
	if err != nil {

		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// GetUserEmailNotExist retrieves users from the database whose email does not exist.
func (s *UserService) CheckEmailNotExist(email string) (uint64, error) {
	var count uint64
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	result, err := s.db.Query(query, email)
	if err != nil {
		return 0, fmt.Errorf("Failed to execute query: %w", err)
	}

	err = result.Scan(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to scan user: %w", err)

	}
	return count, nil
}

// UpdateUser updates an existing user in the database.
func (s *UserService) UpdateUser(user *models.User) error {
	return s.db.Update(user)
}

// DeleteUser deletes a user from the database.
func (s *UserService) DeleteUser(user *models.User) error {
	return s.db.Delete(user)
}
