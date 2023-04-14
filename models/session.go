package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/sxc/oishifood/rand"
)

const (
	// The minimum number of bytes a token should be used.
	MinBytesPerToken = 32
)

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
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// 1. Create a new session token
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
		// TODO: Hash the token
	}

	// Store the session in the database
	// 1. Query for a user's session
	// 2. If the user has a session, update it
	// 3. If the user doesn't have a session, create it
	// 4. Return the session

	// 1. try to update the user's session
	// 2. if err, create a new session
	// 3. return the session

	row := ss.DB.QueryRow(`
	update sessions set token_hash = $2
	where user_id = $1
	returning id
	;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)

	if err == sql.ErrNoRows {
		row := ss.DB.QueryRow(`
		insert into sessions (user_id, token_hash) 
		values ($1, $2)
		returning id;
		`, session.UserID, session.TokenHash)
		err = row.Scan(&session.ID)
	}

	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	// Implement this function
	// return &session, nil
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement this function
	// 1. hash the session tokens
	tokenHash := ss.hash(token)
	var userID int
	row := ss.DB.QueryRow(`
		select user_id 
		from sessions
		where token_hash = $1;`, tokenHash)

	err := row.Scan(&userID)
	if err == nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	var user User
	row = ss.DB.QueryRow(`
	select email, password_hash
	from users
	where id = $1;`, userID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
	delete from sessions
	where token_hash = $1;`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

// type TokenManager struct { }

// func (tm TokenManager)
