//go:build unit
// +build unit

package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		userEmail   string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "should create user successfully",
			userName:    "John Doe",
			userEmail:   "john.doe@example.com",
			expectError: false,
		},
		{
			name:        "should fail with empty name",
			userName:    "",
			userEmail:   "john.doe@example.com",
			expectError: true,
			errorMsg:    "name cannot be empty",
		},
		{
			name:        "should fail with empty email",
			userName:    "John Doe",
			userEmail:   "",
			expectError: true,
			errorMsg:    "email cannot be empty",
		},
		{
			name:        "should fail with name too short",
			userName:    "A",
			userEmail:   "john.doe@example.com",
			expectError: true,
			errorMsg:    "name must be at least 2 characters long",
		},
		{
			name:        "should fail with name too long",
			userName:    "This is a very long name that exceeds the maximum allowed length of one hundred characters and should cause an error",
			userEmail:   "john.doe@example.com",
			expectError: true,
			errorMsg:    "name cannot exceed 100 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.userName, tt.userEmail)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.NotEmpty(t, user.ID)
				assert.Equal(t, tt.userName, user.Name)
				assert.Equal(t, tt.userEmail, user.Email)
				assert.False(t, user.CreatedAt.IsZero())
				assert.False(t, user.UpdatedAt.IsZero())
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	// Create a valid user first
	user, err := NewUser("John Doe", "john.doe@example.com")
	require.NoError(t, err)
	require.NotNil(t, user)

	originalUpdatedAt := user.UpdatedAt

	// Wait a bit to ensure time difference
	time.Sleep(1 * time.Millisecond)

	tests := []struct {
		name        string
		newName     string
		newEmail    string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "should update user successfully",
			newName:     "Jane Doe",
			newEmail:    "jane.doe@example.com",
			expectError: false,
		},
		{
			name:        "should fail with empty name",
			newName:     "",
			newEmail:    "jane.doe@example.com",
			expectError: true,
			errorMsg:    "name cannot be empty",
		},
		{
			name:        "should fail with empty email",
			newName:     "Jane Doe",
			newEmail:    "",
			expectError: true,
			errorMsg:    "email cannot be empty",
		},
		{
			name:        "should fail with name too short",
			newName:     "A",
			newEmail:    "jane.doe@example.com",
			expectError: true,
			errorMsg:    "name must be at least 2 characters long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh user for each test
			user, err := NewUser("John Doe", "john.doe@example.com")
			require.NoError(t, err)

			err = user.Update(tt.newName, tt.newEmail)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				// User should remain unchanged
				assert.Equal(t, "John Doe", user.Name)
				assert.Equal(t, "john.doe@example.com", user.Email)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newName, user.Name)
				assert.Equal(t, tt.newEmail, user.Email)
				assert.True(t, user.UpdatedAt.After(originalUpdatedAt))
			}
		})
	}
}

func TestUser_ValidateUserData(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		userEmail   string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "should validate valid data",
			userName:    "John Doe",
			userEmail:   "john.doe@example.com",
			expectError: false,
		},
		{
			name:        "should fail with empty name",
			userName:    "",
			userEmail:   "john.doe@example.com",
			expectError: true,
			errorMsg:    "name cannot be empty",
		},
		{
			name:        "should fail with empty email",
			userName:    "John Doe",
			userEmail:   "",
			expectError: true,
			errorMsg:    "email cannot be empty",
		},
		{
			name:        "should fail with name too short",
			userName:    "A",
			userEmail:   "john.doe@example.com",
			expectError: true,
			errorMsg:    "name must be at least 2 characters long",
		},
		{
			name:        "should fail with name too long",
			userName:    "This is a very long name that exceeds the maximum allowed length of one hundred characters and should cause an error",
			userEmail:   "john.doe@example.com",
			expectError: true,
			errorMsg:    "name cannot exceed 100 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateUserData(tt.userName, tt.userEmail)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
