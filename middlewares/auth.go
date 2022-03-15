package middlewares

import (
	"context"
	"errors"
	"github.com/dominhchi/gglcoremicro/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func VerifyToken(token string) (*User, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(config.Cfg.SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &User{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*User)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

func AuthMiddleware(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		auth := ctx.Request().Header.Get("Authorization")
		if auth == "" {
			return fn(ctx)
		}
		bearer := "Bearer "
		token := auth[len(bearer):]

		user, err := VerifyToken(token)
		if err != nil || user.ID == "" {
			return ctx.JSON(http.StatusUnauthorized, "You are not authorized!")
		}
		ct := context.WithValue(ctx.Request().Context(), "auth", user)
		ctx.SetRequest(ctx.Request().WithContext(ct))
		return fn(ctx)
	}
}

func CtxUser(ctx context.Context) *User {
	raw, _ := ctx.Value("auth").(*User)
	return raw
}

//func AuthMiddleware1(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		auth := r.Header.Get("Authorization")
//
//		if auth == "" {
//			next.ServeHTTP(w, r)
//			return
//		}
//		auth_byte :=[]byte(auth)
//
//		var user User
//		err := json.Unmarshal(auth_byte, &user)
//		if err != nil || user.ID == "" {
//			http.Error(w, "Invalid token", http.StatusForbidden)
//			return
//		}
//
//		ctx := context.WithValue(r.Context(), "auth", user)
//
//		r = r.WithContext(ctx)
//		next.ServeHTTP(w, r)
//	})
//}
