package repositories

import (
	"ocg-be/database"
	"ocg-be/models"
)

func GetAllOrders() []models.Order {

	rows, _ := database.DB.Query("SELECT id,create_at,user_id,address,phone_number,total_price,order_notes " +
		"FROM orders")
	var orders []models.Order
	for rows.Next() {
		var order models.Order
		rows.Scan(&order.Id, &order.CreatedAt, &order.UserId, &order.Address, &order.PhoneNumber, &order.TotalPrice, &order.OrderNote)
		order.OrderItems = getOrderItem(order.Id)
		orders = append(orders, order)
	}
	return orders
}

func getOrderItem(id int64) []models.OderItem {
	rows, _ := database.DB.Query("SELECT product_id,quantity FROM order_items "+
		"WHERE order_id = ?", id)
	var orderItems []models.OderItem
	for rows.Next() {
		var orderItem models.OderItem
		rows.Scan(&orderItem.ProductId, &orderItem.Quantity)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems
}

func GetOrderByID(id int64) models.Order {

	rows, _ := database.DB.Query("SELECT id,create_at,address,phone_number,total_price,order_notes,status FROM orders "+
		"WHERE id = ?", id)
	var order models.Order
	if rows.Next() {
		rows.Scan(&order.Id, &order.CreatedAt, &order.Address, &order.PhoneNumber, &order.TotalPrice, &order.OrderNote, &order.Status)
	}
	order.OrderItems = getOrderItem(order.Id)
	return order
}

func UpdateStatusOrder(order models.Order) error {

	_, err := database.DB.Exec("UPDATE orders "+
		"SET status = ? "+
		"WHERE id = ?", order.Status, order.Id)

	if err != nil {
		panic(err.Error())
		// return err
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
