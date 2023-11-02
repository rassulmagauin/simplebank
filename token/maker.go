package token

import "time"

// Maker is an Interface for managing tokens
type Maker interface {
	//CreateToken creates token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//VerifyToken checks token fo validity
	VerifyToken(token string) (*Payload, error)
}
