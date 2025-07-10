package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"api-go/internal/domain"
	"api-go/internal/usecase"
)

// MockJWTProvider es una implementación mock de usecase.JWTProvider para pruebas
type MockJWTProvider struct {
	mock.Mock
}

func (m *MockJWTProvider) GenerateToken(username, secret string) (string, error) {
	args := m.Called(username, secret)
	return args.String(0), args.Error(1)
}

func (m *MockJWTProvider) ValidateToken(tokenString, secret string) error {
	args := m.Called(tokenString, secret)
	return args.Error(0)
}

func TestAuthUsecaseAuthenticate(t *testing.T) {
	mockSecret := "testsecret"

	tests := []struct {
		name              string
		username          string
		password          string
		mockGenerateToken func(*MockJWTProvider)
		expectedToken     string
		expectedError     error
	}{
		{
			name:     "Successful authentication",
			username: "admin",
			password: "password",
			mockGenerateToken: func(m *MockJWTProvider) {
				m.On("GenerateToken", "admin", mockSecret).Return("mock_jwt_token", nil).Once()
			},
			expectedToken: "mock_jwt_token",
			expectedError: nil,
		},
		{
			name:     "Invalid credentials",
			username: "wronguser",
			password: "wrongpassword",
			mockGenerateToken: func(m *MockJWTProvider) {
				// No se llama a GenerateToken si las credenciales son inválidas
			},
			expectedToken: "",
			expectedError: domain.ErrInvalidCredentials,
		},
		{
			name:     "Failed to generate token",
			username: "admin",
			password: "password",
			mockGenerateToken: func(m *MockJWTProvider) {
				m.On("GenerateToken", "admin", mockSecret).Return("", errors.New("jwt error")).Once()
			},
			expectedToken: "",
			expectedError: domain.ErrFailedToGenerateToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockJWTProvider := new(MockJWTProvider)
			if tt.mockGenerateToken != nil {
				tt.mockGenerateToken(mockJWTProvider)
			}

			uc := usecase.NewAuthUsecase(mockJWTProvider, mockSecret)
			token, err := uc.Authenticate(tt.username, tt.password)

			assert.Equal(t, tt.expectedToken, token)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tt.expectedError), "Expected error %v, got %v", tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}

			mockJWTProvider.AssertExpectations(t)
		})
	}
}