package authentication

import (
	"fmt"
	"net/http"
	"time"

	"secret-manager/internal/config"
	"secret-manager/internal/logging"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type jwtMiddleware struct {
	cfg    config.JWTConfig
	jwks   *keyfunc.JWKS
	secret []byte
}

func NewJWTMiddleware(cfg config.JWTConfig) (func(http.Handler) http.Handler, error) {
	m := &jwtMiddleware{cfg: cfg}

	switch {
	case cfg.JWKSURL != "":
		jwks, err := keyfunc.Get(cfg.JWKSURL, keyfunc.Options{
			RefreshInterval: 1 * time.Hour,
		})
		if err != nil {
			return nil, fmt.Errorf("jwks: %w", err)
		}
		m.jwks = jwks
	case cfg.Secret != "":
		m.secret = []byte(cfg.Secret)
	default:
		return nil, fmt.Errorf("jwt auth requires either secret or jwksUrl")
	}

	return m.middleware, nil
}

func (m *jwtMiddleware) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := bearerToken(r)
		if raw == "" {
			http.Error(w, "jwt: missing bearer token", http.StatusUnauthorized)
			return
		}

		var keyFunc jwt.Keyfunc
		if m.jwks != nil {
			keyFunc = m.jwks.Keyfunc
		} else {
			keyFunc = func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return m.secret, nil
			}
		}

		opts := []jwt.ParserOption{jwt.WithExpirationRequired()}
		if m.cfg.Issuer != "" {
			opts = append(opts, jwt.WithIssuer(m.cfg.Issuer))
		}
		if m.cfg.Audience != "" {
			opts = append(opts, jwt.WithAudience(m.cfg.Audience))
		}

		token, err := jwt.Parse(raw, keyFunc, opts...)
		if err != nil || !token.Valid {
			logging.L.App.Warn("jwt: invalid token", zap.Error(err))
			http.Error(w, "jwt: unauthorized", http.StatusUnauthorized)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		sub, _ := claims["sub"].(string)

		id := Identity{
			Method:  AuthMethodJWT,
			Subject: sub,
			Claims:  claims,
		}

		next.ServeHTTP(w, WithIdentity(r, id))
	})
}
