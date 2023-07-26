package helper

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"html"
	"os"
	"strings"
	"user-service/model"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func Encode(input string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

func VerifyPassword(hashPassword string, pass string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(pass))
	if err != nil {
		return "PASSWORD NOT MATCHED!", err
	}
	return "VERIFIED", nil
}

func TrimString(str string) string {
	return html.EscapeString(strings.TrimSpace(str))
}

//func GenerateJWT(user service.User) (string, error) {
//	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"id":  user.Id,
//		"iat": time.Now().Unix(),
//		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
//	})
//	return token.SignedString(privateKey)
//}

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func CurrentUser(context *gin.Context) (model.UserEntity, error) {
	//err := ValidateJWT(context)
	//if err != nil {
	//	return service.User{}, err
	//}
	//
	//token, _ := getToken(context)
	//claims, _ := token.Claims.(jwt.MapClaims)
	//userId := int(claims["id"].(int))

	//user := service.User{Id: userId}
	//u, err := user.GetUserById()

	//if err != nil {
	//	return service.User{}, err
	//}

	return model.UserEntity{}, nil
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
