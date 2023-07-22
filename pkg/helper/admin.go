package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
)

type authCustomClaimsAdmin struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {

	claims := &authCustomClaimsAdmin{
		Name:  admin.Name,
		Email: admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("132457689"))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
