package helpers

import "net/http"

func HttpError(w http.ResponseWriter, error string, code int) {

	w.Header().Set("IsError", "true")
	http.Error(w, error, code)
}
