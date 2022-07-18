package jwt

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	j "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GenerateToken(user_id int64, isRefresh bool) (string, error) {
	signingKey := []byte(viper.GetString("SECRET"))
	var err error
	var token_lifespan int
	if isRefresh {
		token_lifespan, err = strconv.Atoi(viper.GetString("REFRESH_TOKEN_HOUR_LIFESPAN"))
		if err != nil {
			return "", err
		}
	} else {
		token_lifespan, err = strconv.Atoi(viper.GetString("TOKEN_HOUR_LIFESPAN"))
		if err != nil {
			return "", err
		}
	}

	claims := j.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()

	token := j.NewWithClaims(j.SigningMethodHS256, claims)

	return token.SignedString(signingKey)
}

func ValidateToken(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := j.Parse(tokenString, func(token *j.Token) (interface{}, error) {
		if _, ok := token.Method.(*j.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenID(c *gin.Context) (int64, error) {
	tokenString := ExtractToken(c)
	token, err := j.Parse(tokenString, func(token *j.Token) (interface{}, error) {
		if _, ok := token.Method.(*j.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("SECRET")), nil
	})
	log.Println(err)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(j.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return int64(uid), nil
	}
	return 0, nil
}

func GetUserFromToken(tokenString string) (int64, error) {
	token, err := j.Parse(tokenString, func(token *j.Token) (interface{}, error) {
		if _, ok := token.Method.(*j.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("SECRET")), nil
	})
	log.Println(err)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(j.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return int64(uid), nil
	}
	return 0, nil
}
