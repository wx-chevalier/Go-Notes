package views

import (
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	pkg "github.com/L04DB4L4NC3R/jobs-mhrd/pkg"
)

type ErrView struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var (
	ErrMethodNotAllowed = errors.New("Error: Method is not allowed")
	ErrInvalidToken     = errors.New("Error: Invalid Authorization token")
	ErrUserExists       = errors.New("User already exists")
)

var ErrHTTPStatusMap = map[string]int{
	pkg.ErrNotFound.Error():     http.StatusNotFound,
	pkg.ErrInvalidSlug.Error():  http.StatusBadRequest,
	pkg.ErrExists.Error():       http.StatusConflict,
	pkg.ErrNoContent.Error():    http.StatusNotFound,
	pkg.ErrDatabase.Error():     http.StatusInternalServerError,
	pkg.ErrUnauthorized.Error(): http.StatusUnauthorized,
	pkg.ErrForbidden.Error():    http.StatusForbidden,
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	ErrInvalidToken.Error():     http.StatusBadRequest,
	ErrUserExists.Error():       http.StatusConflict,
}

func Wrap(err error, w http.ResponseWriter) {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	// If error code is not found
	// like a default case
	if code == 0 {
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)

	errView := ErrView{
		Message: msg,
		Status:  code,
	}
	log.WithFields(log.Fields{
		"message": msg,
		"code":    code,
	}).Error("Error occurred")

	json.NewEncoder(w).Encode(errView)
}
