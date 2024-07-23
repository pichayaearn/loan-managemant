package auth

import (
	"context"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pichayaearn/loan-management/pkg/admin/model"
	adminModel "github.com/pichayaearn/loan-management/pkg/admin/model"
	authModel "github.com/pichayaearn/loan-management/pkg/auth/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService adminModel.UserService
	secretKey   string
}

type NewAuthServiceCfgs struct {
	UserService adminModel.UserService
	SecretKey   string
}

type MyCustomClaims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func NewAuthService(cfg NewAuthServiceCfgs) authModel.AuthService {
	return &AuthService{
		userService: cfg.UserService,
		secretKey:   cfg.SecretKey,
	}
}

func (aSvc AuthService) Login(email, password string) (string, error) {
	ctx := context.Background()
	userExist, err := aSvc.userService.GetUser(model.GetUserOpts{
		Email:  email,
		Status: model.UserStatusActived,
	}, ctx)
	if err != nil {
		return "", err
	}

	if userExist == nil {
		return "", fmt.Errorf("email %s not found", email)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userExist.Password()), []byte(password)); err != nil {
		return "", err
	}

	return aSvc.createToken(userExist.UserID())
}

func (aSvc AuthService) createToken(userID uuid.UUID) (string, error) {
	secretKey := []byte(aSvc.secretKey)

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		userID.String(),
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
