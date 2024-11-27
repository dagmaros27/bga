package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type PasswordTokenRepository struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

// CreatePasswordResetToken implements interfaces.PasswordTokenRepository.
func (p PasswordTokenRepository) CreatePasswordResetToken(token *entities.PasswordResetToken) (*entities.PasswordResetToken, error) {
	_, err := p.Collection.InsertOne(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// DeletePasswordResetTokenByUserId implements interfaces.PasswordTokenRepository.
func (p PasswordTokenRepository) DeletePasswordResetTokenByUserId(userId string) error {
	_, err := p.Collection.DeleteOne(context.Background(), map[string]string{"userId": userId})
	if err != nil {
		return err
	}

	return nil
}

// FindPasswordResetTokenByUserId implements interfaces.PasswordTokenRepository.
func (p PasswordTokenRepository) FindPasswordReset(tok string) (*entities.PasswordResetToken, error) {
	var token entities.PasswordResetToken
	err := p.Collection.FindOne(context.Background(), map[string]string{"token": tok}).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func NewPasswordTokenRepository(db *mongo.Database) interfaces.PasswordTokenRepository {
	return PasswordTokenRepository{Database: db, Collection: db.Collection("password_reset_tokens")}
}
