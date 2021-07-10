package repositories

import (
	"back/database"
	"back/models"
)

func CreateNewUser(user models.User) models.User {
	
	strQuery, err := database.DB.Prepare("insert into users (last_name, first_name, email, password)"+
		"values (?, ?, ?, ?)")

	if err != nil {
		panic("Could not create new user")
	}

	strQuery.Exec(user.Last_name, user.First_name, user.Email, user.Password)

	return user
}

func GetAllUsers() []models.User{
	rows, _ := database.DB.Query("SELECT * " + "FROM users")
	var users []models.User
	for rows.Next(){
		var user models.User
		rows.Scan(&user.Id, &user.Last_name, &user.First_name, &user.Email, &user.Password)
		users = append(users, user)
	}
	return users
}

func GetUserById(id int64) models.User {
	var user models.User
	row, _ := database.DB.Query("select * " + "from users where id = ? ", id)
	if row.Next() {
		row.Scan(&user.Id, &user.Last_name, &user.First_name, &user.Email, &user.Password)
	}
	return user
}

func UpdateUserById(user models.User) error {
	// var roleId int64
	// if user.Role == "ADMIN" {
	// 	roleId = 1
	// } else {
	// 	roleId = 2
	// }
	strQuery, err := database.DB.Prepare("UPDATE users " +
		"SET last_name = ?, first_name = ?, email = ?, password = ?" + "WHERE id = ?")

	if err != nil{
		panic(err.Error())
	}

	strQuery.Exec(user.Last_name, user.First_name, user.Email, user.Password, user.Id)
	return nil
}

func DeleteUserById(id int64) error {

	strQuery, err := database.DB.Prepare("DELETE FROM users " + "where id = ?")
	if err != nil{
		panic(err.Error())
	}
	strQuery.Exec(id)
	return nil

}