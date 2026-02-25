package authentication

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// AuthMethod tells downstream handlers how the caller was authenticated.
type AuthMethod string

const (
	AuthMethodNone  AuthMethod = "none"
	AuthMethodMTLS  AuthMethod = "mtls"
	AuthMethodJWT   AuthMethod = "jwt"
	AuthMethodSPIRE AuthMethod = "spire"
)

// Identity is the normalised caller identity regardless of auth method.
type Identity struct {
	Method AuthMethod
	// mTLS / SPIRE: DN of the leaf cert  e.g. "CN=test-client,O=Test"
	// JWT:          "sub" claim
	Subject string
	// Extra method-specific fields
	SPIFFEID string        // SPIRE only
	Claims   jwt.MapClaims // JWT only — raw claims for handlers that need them
}

type contextKey struct{}

// WithIdentity stores the identity in the request context.
func WithIdentity(r *http.Request, id Identity) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), contextKey{}, id))
}

// IdentityFrom retrieves the identity from the context.
func IdentityFrom(ctx context.Context) (Identity, error) {
	id, ok := ctx.Value(contextKey{}).(Identity)
	if !ok {
		return Identity{}, errors.New("no identity in context this is required for each action")
	}

	return id, nil
}
