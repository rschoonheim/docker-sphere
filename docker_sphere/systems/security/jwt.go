package security

import "github.com/golang-jwt/jwt/v5"

// JwtCreate - creates a JWT token for a user
func JwtCreate(username string) string {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("example")
	t = jwt.New(jwt.SigningMethodHS256)
	s, _ = t.SignedString(key)

	return s
}
