package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golangpkg/pkg/util"
	"time"
)

var (
	jwtSecret string
)

func InitJwt(secretKey string) {
	jwtSecret = secretKey
}

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		util.EncodeMD5(username),
		util.EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseMapToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 创建token
func CreateMapToken(data map[string]string) string {
	if len(jwtSecret) == 0 {
		panic("jwtSecret Need InitJwt")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for index, val := range data {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(jwtSecret))
	return tokenString
}

// 解析token
func ParseMapToken(tokenString string) (map[string]string, bool) {
	if len(jwtSecret) == 0 {
		panic("jwtSecret Need InitJwt")
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapData := make(map[string]string)
		for index, val := range claims {
			mapData[index] = fmt.Sprintf("%v", val)
		}
		return mapData, true
	} else {
		return nil, false
	}
}
