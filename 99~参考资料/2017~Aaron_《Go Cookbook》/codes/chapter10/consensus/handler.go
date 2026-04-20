package consensus

import (
	"net/http"
	"time"
)

// Handler grabs the get param ?next= and tries
// to transition to the state contained there
func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, rf := range rafts {
		if k == rf.Leader() {
			state := r.FormValue("next")
			result := rf.Apply([]byte(state), 1*time.Second)
			if result.Error() != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			newState, ok := result.Response().(string)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if newState != state {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("invalid transition"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result.Response().(string)))
			return
		}
	}
}
