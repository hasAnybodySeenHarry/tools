package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	api "harryd.com/tools/app/models/api"
)

var jwtKey []byte
var jwtKeyInitialized int32

func initialize() {
	if atomic.CompareAndSwapInt32(&jwtKeyInitialized, 0, 1) {
		var err error
		jwtKey, err = loadKey("app/secret/jwtSecret.json")
		if err != nil {
			log.Fatal("Failed to load the key:", err)
		}
	}
}

func loadKey(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var keyData struct {
		Key string `json:"key"`
	}

	err = json.Unmarshal(content, &keyData)
	if err != nil {
		return nil, err
	}

	return []byte(keyData.Key), nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		initialize()
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			response := api.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !token.Valid {
			response := api.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
