//go:build unit
// +build unit

package repository

import (
	"context"
	"testing"
	"time"

	"example.com/internal-service/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockCollection é um mock da collection do MongoDB
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	if result := args.Get(0); result != nil {
		return result.(*mongo.SingleResult)
	}
	return nil
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func (m *MockCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(int64), args.Error(1)
}

// MockSingleResult é um mock do SingleResult do MongoDB
type MockSingleResult struct {
	mock.Mock
}

func (m *MockSingleResult) Decode(v interface{}) error {
	args := m.Called(v)
	return args.Error(0)
}

func TestUserMongoRepository_Create(t *testing.T) {
	tests := []struct {
		name          string
		user          *user.User
		setupMock     func(*MockCollection)
		expectedError bool
	}{
		{
			name: "successful creation",
			user: &user.User{
				ID:        "test-id",
				Name:      "Test User",
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			setupMock: func(mockCol *MockCollection) {
				mockCol.On("InsertOne", mock.Anything, mock.AnythingOfType("userDocument"), mock.Anything).Return(&mongo.InsertOneResult{}, nil)
			},
			expectedError: false,
		},
		{
			name: "database error",
			user: &user.User{
				ID:        "test-id",
				Name:      "Test User",
				Email:     "test@example.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			setupMock: func(mockCol *MockCollection) {
				mockCol.On("InsertOne", mock.Anything, mock.AnythingOfType("userDocument"), mock.Anything).Return(&mongo.InsertOneResult{}, assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCollection := &MockCollection{}
			tt.setupMock(mockCollection)

			repo := NewUserMongoRepositoryWithCollection(mockCollection)

			err := repo.Create(context.Background(), tt.user)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockCollection.AssertExpectations(t)
		})
	}
}
