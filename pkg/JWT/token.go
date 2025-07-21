package jwt

import (
	"exercise/global"
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwt"
)

func GenerateToken(username string) (string, error) {
	now := time.Now()

	claims:= map[string]interface{}{
		"issuer": global.JWTSetting.Issuer,
		"issued_at": now,
		"expires_at": now.Add(global.JWTSetting.Expire),
		"username": username,
	}

	token,err := jwt.NewBuilder().
	Issuer(global.JWTSetting.Issuer).
	IssuedAt(now).
	Expiration(now.Add(global.JWTSetting.Expire)).
	Build()

	if err != nil {
		return "",fmt.Errorf("package jwt Generate error: %v",err)
	}
	for k , v := range claims {
		err := token.Set(k,v)
		if err != nil {
			return "",fmt.Errorf("package jwt Set(k,v) error:%v",err)
		}
	}
	signedToken,err := jwt.Sign(token,
	jwt.WithKey(global.JWTSetting.Algorithm,[]byte(global.JWTSetting.Secret)))
	if err != nil {
		return "",fmt.Errorf("package jwt Sign error:%v",err)
	}
	return string(signedToken),nil
}

func ValidateToken(tokenStr string) (jwt.Token,error){
	if len(tokenStr) == 0{
		return nil,fmt.Errorf("package jwt input tokenStr is empty")
	}
	token,err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKey(global.JWTSetting.Algorithm,[]byte(global.JWTSetting.Secret)),
		jwt.WithValidate(true),
	)
	if err != nil {
		return nil,fmt.Errorf("package jwt Parse error:%v",err)
	}
	return token,nil
}