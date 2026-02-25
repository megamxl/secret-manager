package authentication

import (
	"net"
	"net/http"
)

func NoAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// RemoteAddr is "ip:port", strip the port
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr // fallback if parsing fails somehow
		}

		id := Identity{
			Method:  AuthMethodNone,
			Subject: ip,
		}

		next.ServeHTTP(w, WithIdentity(r, id))
	})
}
