package repositories

import (
	"log"
	"ocg-be/database"
	"ocg-be/models"
	dto "ocg-be/models/displayOder"
)

func GetAllOrder() []dto.DetailUser {

	qtext := "SELECT od.id,od.create_at,od.address, od.phone_number, od.order_notes, od.status,od.total_price,"+ 
	"u.first_name, u.last_name, u.email FROM orders od JOIN users u ON u.id = od.user_id WHERE u.role_id = 2 LIMIT 25 OFFSET 1;"

	rows, err := database.DB.Query(qtext)
	
	if err != nil {
		log.Println(err)
		return nil
	}
	var users []dto.DetailUser
	for rows.Next() {
		var user dto.DetailUser
		rows.Scan(&user.Id,&user.CreatedAt, &user.Address, &user.PhoneNumber, &user.OrderNote, &user.Status,&user.TotalPrice, &user.First_name, &user.Last_name,&user.Email)
		users = append(users, user)
	}
	return users

}

func GetOrderByID(id int64) []dto.DetailOrder {

	rows, _ := database.DB.Query("SELECT p.name, p.price, oi.quantity, o.status FROM orders o JOIN order_items oi ON o.id = oi.order_id JOIN products p ON p.id = oi.product_id WHERE o.id = ?;", id)
	var orderItems []dto.DetailOrder
	for rows.Next() {
		var order_item dto.DetailOrder
		rows.Scan(&order_item.Name, &order_item.Price, &order_item.Quantity, &order_item.Status)
		
		orderItems = append(orderItems, order_item)
	}
	return orderItems
}

func UpdateStatusOrder(order models.Order) error {

	_, err := database.DB.Exec("UPDATE orders "+"SET status = ? "+
		"WHERE id = ?", order.Status, order.Id)

	if err != nil {
		return err
	}

	return nil
}

func CreateOrder(order models.Order) error {

	_, err := database.DB.Exec("INSERT INTO orders (create_at,user_id,address,phone_number,total_price,order_notes,status)"+
		"VALUES (NOW(),?,?,?,?,?,0)",order.UserId, order.Address, order.PhoneNumber, order.TotalPrice,order.OrderNote)
	if err != nil {
		return err
	}

	var id int
	row, err := database.DB.Query("SELECT MAX(id) FROM orders")
	if err != nil {
		return err
	}
	if row.Next() {
		row.Scan(&id)
	}

	for _, product := range order.OrderItems {
		sqlStr, _ := database.DB.Prepare("INSERT INTO order_items (order_id,product_id,quantity) VALUES (?, ? ,?)")
		_,err := sqlStr.Exec(id,product.ProductId, product.Quantity)

		if err != nil {
			return err
		}
	}
	return nil
}

func GetRevenueByDay(offset int) (interface{}, error) {

	rows, err := database.DB.Query("SELECT o.create_at, IFNULL(SUM(o.total_price),0) FROM orders o JOIN order_items oi ON oi.order_id = o.id GROUP BY create_at;")

	if err != nil {
		log.Println(err)
		return 0, err
	}

	type Rev struct {
		Date          string `json:"date"`
		Revenue       int64  `json:"revenue"`
		// TotalProducts int64  `json:"total_products"`
	}
	var revs []Rev

	for rows.Next() {
		var rev Rev
		rows.Scan(&rev.Date, &rev.Revenue)
		revs = append(revs, rev)
	}
	return revs, nil
}
