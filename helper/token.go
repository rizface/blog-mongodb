package helper

import (
	"blog-mongo/app/config"
	"blog-mongo/app/exception"
	"errors"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"time"
)

var jwtKey string

type Claims struct {
	Id       interface{} `json:"id"`
	Username interface{} `json:"username"`
	Email    interface{} `json:"email"`
	jwt.StandardClaims
}

func init() {
	err := loadConfig()
	if err != nil {
		jwtKey = config.DefaultConfig["jwtSecret"]
	} else {
		jwtKey = os.Getenv("JWT_SECRET")
	}
}

func GenerateToken(user bson.M) (string,error) {
	claims := Claims{
		Id:       user["_id"].(primitive.ObjectID).Hex(),
		Username: user["username"],
		Email:    user["email"],
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString([]byte(jwtKey))
	return tokenString,err
}

func TokenValidation(tokenString string) interface{}{
	claims := &Claims{}
	tkn,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey),nil
	})
	PanicCustomException(exception.Unauthorized{Err:errors.New("token kamu tidak valid")},err != nil && errors.Is(err,jwt.ErrSignatureInvalid))
	PanicCustomException(exception.Unauthorized{Err:errors.New("token kamu tidak valid")},tkn.Valid == false)
	return claims
}

func GetClaimsValue(r *http.Request,ctxKey string,dataKey string) interface{} {
	user := r.Context().Value(ctxKey)
	userData,ok := user.(*Claims)
	if ok {
		switch dataKey {
		case "username" :
			return userData.Username
			break
		case "email" :
			return userData.Email
			break
		default :
			return "data tidak ditemukan"
		}
	}
	return nil
}
