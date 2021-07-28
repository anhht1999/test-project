package middlewares

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ocg-be/database"
	"ocg-be/models"
	"ocg-be/util"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("userId")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = util.ParseJwt(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}


func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("userId")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, _ := util.ParseJwt(cookie.Value)
		userId, _ := strconv.Atoi(id)
		user := models.User{
			Id: int64(userId),
		}

		row, _ := database.DB.Query("select u.role_id from users u "+
			"join roles r on r.id = u.role_id where u.id = ?", user.Id)
		if row.Next() {
			row.Scan(&user.Role)
		}
		if user.Role != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "access denied",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CommonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, "+
			"Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, "+
			"Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, "+
			"Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}