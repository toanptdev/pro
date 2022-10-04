package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"rest-api/component/tokenprovider"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type MyClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
	})

	myToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:     myToken,
		CreatedAt: time.Now(),
		Expiry:    expiry,
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*MyClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return &claims.Payload, nil
}
