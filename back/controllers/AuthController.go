package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"back/database"
	"back/models"
	"back/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if data["password"] != data["password_confirm"] {
		http.Error(w, "password match incorrect", http.StatusBadRequest)
		return
	}

	user := models.User{
		Last_name:  data["last_name"],
		First_name: data["first_name"],
		Email:      data["email"],
		Role:       1,
	}

	user.SetPassword(data["password"])

	strQuery, err := database.DB.Prepare("insert into users (last_name, first_name, email, password,role_id)" +
		"values (?, ?, ?, ?,?)")

	if err != nil {
		panic("Could not create new user")
	}

	strQuery.Exec(user.Last_name, user.First_name, user.Email, user.Password, user.Role)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Register successfully",
		"user" : user,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User

	row, _ := database.DB.Query("SELECT * FROM users "+
		"WHERE email = ?", data["email"])

	if row.Next() {
		row.Scan(&user.Id, &user.Last_name, &user.First_name, &user.Email, &user.Password, &user.Role)
	}
	if user.Id == 0 {
		http.Error(w, "email incorrect", http.StatusBadRequest)
		return
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		http.Error(w, "password incorrect", http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "login successfully",
	})

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "logout successfully",
	})
}
