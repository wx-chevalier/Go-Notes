package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

// SetHandler Sets the value, and returns it in a resp
func (c *Controller) SetHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	payload := resp{Status: "error"}
	r.ParseForm()
	val := r.FormValue("score")
	score, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		logrus.WithField("error", err).Error("failed to parse input")
		w.WriteHeader(http.StatusBadRequest)
		enc.Encode(&payload)
		return
	}

	if err := c.db.SetScore(score); err != nil {
		logrus.WithField("error", err).Error("failed to set the score")
		w.WriteHeader(http.StatusInternalServerError)
		enc.Encode(&payload)
		return
	}
	w.WriteHeader(http.StatusOK)
	payload.Value = score
	payload.Status = "ok"
	enc.Encode(&payload)
}
