package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var JwtSecret = []byte(viper.GetString("jwt.secret"))

const (
	UUIDKey      = "uuid"
	SaasIDKey    = "SaasId"
	PhoneKey     = "phone"
	NameKey      = "name"
	CompanyIDKey = "CompanyIDKey"
)

type Claims struct {
	UUID      string `json:"uid"`
	SaasId    int    `json:"saas_id"`
	CompanyId int    `json:"company_id"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(uuid, phone, name string, SaasId, companyId int) (string, error) {
	expireTime := time.Now().Add(viper.GetDuration("jwt.expire") * time.Second)

	claims := Claims{
		uuid,
		SaasId,
		companyId,
		phone,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jc-legal",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

//解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
