package login

import (
	"github.com/aiyaruch1320/go-login-jwt/user"

	"github.com/golang-jwt/jwt"
)

// JwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	UserID   string    `json:"id"`
	Username string    `json:"username"`
	Role     user.Role `json:"role"`
	ClientId string    `json:"client_id"`
	Sub      string    `json:"sub"`
	AuthTime int64     `json:"auth_time"`
	Idp      string    `json:"ldp"`
	Iat      int64     `json:"iat"`
	Scope    []string  `json:"scope"`
	Arm      []string  `json:"amr"`
	jwt.StandardClaims
}
