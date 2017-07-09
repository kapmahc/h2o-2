package nut

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

// Home home url
func Home() string {
	return beego.AppConfig.String("homeurl")
}

var (
	jwtMethod = crypto.SigningMethodHS512
	jwtKey    = []byte(beego.AppConfig.String("jwtkey"))
)

//JwtValidate validate jwt token
func JwtValidate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(jwtKey, jwtMethod); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

// JwtParse parse and validate token from http request
func JwtParse(r *http.Request) (jwt.Claims, error) {
	tk, err := jws.ParseJWTFromRequest(r)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(jwtKey, jwtMethod); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

//JwtSum create jwt token
func JwtSum(cm jws.Claims, exp time.Duration) ([]byte, error) {
	kid := uuid.New().String()
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.Add(exp))
	cm.Set("kid", kid)
	//TODO using kid

	jt := jws.NewJWT(cm, jwtMethod)
	return jt.Serialize(jwtKey)
}
