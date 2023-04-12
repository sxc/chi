package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session. When look up a session
	// this will be empty.
	//can not reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// 1. Create a new session token
	// Implement this function
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
