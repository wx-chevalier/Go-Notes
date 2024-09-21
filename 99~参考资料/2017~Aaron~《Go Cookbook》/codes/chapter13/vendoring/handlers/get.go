package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetHandler returns the current score in a resp object
func (c *Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	payload := resp{Status: "error"}
	oldScore, err := c.db.GetScore()
	if err != nil {
		logrus.WithField("error", err).Error("failed to get the score")
		w.WriteHeader(http.StatusInternalServerError)
		enc.Encode(&payload)
		return
	}
	w.WriteHeader(http.StatusOK)
	payload.Value = oldScore
	payload.Status = "ok"
	enc.Encode(&payload)
}
