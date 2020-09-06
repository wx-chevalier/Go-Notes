package handler

import (
	"encoding/json"
	"net/http"

	"github.com/L04DB4L4NC3R/jobs-mhrd/api/middleware"
	"github.com/L04DB4L4NC3R/jobs-mhrd/api/views"
	"github.com/L04DB4L4NC3R/jobs-mhrd/pkg/admin"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func registerAdmin(svc admin.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}

		var admin admin.Admin
		if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
			views.Wrap(err, w)
			return
		}

		u, err := svc.Register(r.Context(), admin.Email, admin.Password, admin.PhoneNumber)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		w.WriteHeader(http.StatusCreated)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": u.Email,
			"id":    u.ID,
			"role":  "admin",
		})
		tokenString, err := token.SignedString([]byte(viper.GetString("jwt_secret")))
		if err != nil {
			views.Wrap(err, w)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
			"admin": u,
		})
		return
	})
}

func loginAdmin(svc admin.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
		var admin admin.Admin
		if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
			views.Wrap(err, w)
			return
		}

		u, err := svc.Login(r.Context(), admin.Email, admin.Password)
		if err != nil {
			views.Wrap(err, w)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": u.Email,
			"id":    u.ID,
			"role":  "admin",
		})
		tokenString, err := token.SignedString([]byte(viper.GetString("jwt_secret")))
		if err != nil {
			views.Wrap(err, w)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
			"admin": u,
		})
		return
	})
}

func profileAdmin(svc admin.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// @protected
		// @description build profile
		if r.Method == http.MethodPost {
			var admin admin.Admin
			if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
				views.Wrap(err, w)
				return
			}

			claims, err := middleware.ValidateAndGetClaims(r.Context(), "admin")
			if err != nil {
				views.Wrap(err, w)
				return
			}
			admin.Email = claims["email"].(string)
			u, err := svc.BuildProfile(r.Context(), &admin)
			if err != nil {
				views.Wrap(err, w)
				return
			}

			json.NewEncoder(w).Encode(u)
			return
		} else if r.Method == http.MethodGet {

			// @description view profile
			claims, err := middleware.ValidateAndGetClaims(r.Context(), "admin")
			if err != nil {
				views.Wrap(err, w)
				return
			}
			u, err := svc.GetAdminProfile(r.Context(), claims["email"].(string))
			if err != nil {
				views.Wrap(err, w)
				return
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Admin profile",
				"data":    u,
			})
			return
		} else {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
	})
}

func changePasswordAdmin(svc admin.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var u admin.Admin
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				views.Wrap(err, w)
				return
			}

			claims, err := middleware.ValidateAndGetClaims(r.Context(), "admin")
			if err != nil {
				views.Wrap(err, w)
				return
			}
			if err := svc.ChangePassword(r.Context(), claims["email"].(string), u.Password); err != nil {
				views.Wrap(err, w)
				return
			}
			return
		} else {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
	})
}

// expose handlers
func MakeAdminHandler(r *http.ServeMux, svc admin.Service) {
	r.Handle("/api/v1/admin/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	}))

	r.Handle("/api/v1/admin/register", registerAdmin(svc))
	r.Handle("/api/v1/admin/login", loginAdmin(svc))
	r.Handle("/api/v1/admin/profile", middleware.Validate(profileAdmin(svc)))
	r.Handle("/api/v1/admin/pwd", middleware.Validate(changePasswordAdmin(svc)))
}
