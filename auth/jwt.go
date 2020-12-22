package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/pkg/errors"
)

type JwtHandler struct {
	TokenKeyURI string
}

func (h *JwtHandler) ParseAuthToken(tokenString string) (map[string]interface{}, error) {
	keySet, err := h.getTokenKey()
	if err != nil {
		return nil, errors.Wrap(err, "error getting key set")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}
		keys := keySet.LookupKeyID(kid)
		if len(keys) == 0 {
			return nil, errors.New("key %v not found")
		}
		var raw interface{}
		return raw, keys[0].Raw(&raw)
	})
	if err != nil {
		return nil, errors.Wrap(err, "token error")
	}
	if token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	} else {
		return nil, errors.New("error decoding claims")
	}
}

// getTokenKey returns the keyset of JWT token from AWS
func (h *JwtHandler) getTokenKey() (*jwk.Set, error) {
	keySet, err := jwk.Fetch(h.TokenKeyURI)
	if err != nil {
		return nil, errors.Wrap(err, "error loading jwt key")
	}
	return keySet, nil
}
