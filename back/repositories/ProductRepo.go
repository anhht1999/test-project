package repositories

import (
	"back/database"
	"back/models"
)

func GetAllProducts() []models.Product{
	rows, _ := database.DB.Query("SELECT * " +
		"FROM products")
	var products []models.Product
	for rows.Next(){
		var product models.Product
		rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Trademark)
		product.Images = getImages(product.Id)
		products = append(products, product)
	}
	return products
}

func getImages(id int64) []models.Image {
	rows, _ := database.DB.Query("SELECT url FROM images " +
		"WHERE product_id = ?", id)
	var images []models.Image
	for rows.Next(){
		var image models.Image
		rows.Scan(&image.Url)
		images = append(images, image)
	}
	return images
}

func GetProductByID(id int64) models.Product{
	
	rows, _ := database.DB.Query("SELECT * FROM products " +
		"WHERE id = ?", id)
	var product models.Product
	if rows.Next(){
		rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Trademark)
	}
	product.Images = getImages(product.Id)
	return product
}

func UpdateProductByID( product models.Product) error {
	
	strQuery, err := database.DB.Prepare("UPDATE products " +
		"SET name = ?, description = ?, price = ?, category_id = ?, trade_mark = ? " +
		"WHERE id = ?")

	if err != nil{
		panic(err.Error())
	}

	strQuery.Exec(product.Name, product.Description, product.Price, product.CategoryId, product.Trademark, product.Id)
	return nil
}

func CreateProduct(product models.Product) error{

	row, err := database.DB.Query("SELECT MAX(id) FROM products")

	if err != nil{
		panic(err.Error())
	}
	if row.Next(){
		row.Scan(&product.Id)
	}else {
		product.Id = 0
	}
	product.Id += 1
	strQuery, err := database.DB.Prepare("INSERT INTO products " +
		"(name, description, price, category_id, trade_mark) " +
		"VALUES (?, ?, ?, ?, ?)")

	for _, image := range product.Images{
		//fmt.Println(image.ImageUrl)
		_, err = database.DB.Query("INSERT INTO images VALUES (?, ?)", product.Id, image.Url)
	}
	if err != nil{
		panic(err.Error())
	}

	strQuery.Exec(product.Name, product.Description, product.Price, product.CategoryId, product.Trademark)
	return nil
}

func DeleteProduct(id int64) error {
	
	strQuery, err := database.DB.Prepare("DELETE FROM products " +
		"where id = ?")
	if err != nil{
		panic(err.Error())
	}
	strQuery.Exec(id)
	return nil
}