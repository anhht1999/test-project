package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"ocg-be/database"
	"ocg-be/models"
	// "ocg-be/util"
	// "reflect"
)

type ProductStorage struct {
}

func (p *ProductStorage) Count(categoryId int, search string) int64 {
	var rows *sql.Rows
	var err error

	if categoryId == 0 {
		rows, err = database.DB.Query("SELECT COUNT(*) FROM products WHERE name LIKE ?", search)
		if err != nil {
			log.Println(err)
		}
	} else {
		rows, err = database.DB.Query("SELECT COUNT(*) FROM products WHERE category_id=? AND name LIKE ?", categoryId, search)
		if err != nil {
			log.Println(err)
		}
	}
	var total int64
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}
func (p *ProductStorage) TakeByParams(limit int, offset int, categoryId int, orderBy string, sort string, search string) interface{} {
	var products []models.Product
	var rows *sql.Rows
	var err error

	qtext := fmt.Sprintf("SELECT p.id,p.name,p.description,p.price,p.trade_mark,p.category_id,cate.name_cate FROM products p "+ 
	"INNER JOIN categories cate ON cate.id = p.category_id WHERE name LIKE ? ORDER BY %s %s LIMIT ? OFFSET ? ", orderBy, sort)

	qtextCate := fmt.Sprintf("SELECT p.id,p.name,p.description,p.price,p.trade_mark,,p.category_id,c.name_cate FROM products p JOIN categories c ON c.id = p.category_id WHERE name LIKE ? AND category_id = ? ORDER BY %s %s LIMIT ? OFFSET ? ", orderBy, sort)

	if categoryId == 0 {
		rows, err = database.DB.Query(qtext, search, limit, offset)
	} else {
		rows, err = database.DB.Query(qtextCate, search, categoryId, limit, offset)
	}
	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Trademark, &product.CategoryId, &product.Category)

		if err != nil {
			log.Println(err)
			return nil
		}

		var image ImageStorage
		product.Image = image.Take(int64(product.Id))
		products = append(products, product)
	}
	return products

}

func (p *ProductStorage) Add(product *models.Product) int {

	_, err := database.DB.Query("INSERT INTO products VALUES (?,?, ?, ?, ?, ?)", product.Id, product.Name, product.Description, product.Price, product.CategoryId, product.Trademark)

	if err != nil {
		product = nil
		fmt.Println(err)
		return 0
	}
	rows, _ := database.DB.Query("SELECT id FROM products ORDER BY id DESC LIMIT 1")

	if rows.Next() {
		err = rows.Scan(&product.Id)
		if err != nil {
			product = nil
			fmt.Println(err)
			return 0
		}
	}
	for _, image := range product.Image {
		_, err = database.DB.Exec("INSERT INTO images VALUES (?, ?)", product.Id, image.Url)
	}
	if err != nil {
		panic(err)
	}
	return product.Id
}

func (p *ProductStorage) GetById(id int) models.Product {
	var product models.Product
	rows, err := database.DB.Query("SELECT * from products WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)

		return product
	}
	if rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Trademark)
		if err != nil {
			fmt.Println(err)
			return product
		}
	}
	var image ImageStorage

	product.Image = image.Take(int64(product.Id))

	return product
}

func (p *ProductStorage) DeleteById(id int) {
	_, err := database.DB.Query("DELETE FROM products WHERE id=?", id)
	if err != nil {
		log.Println(err)
	}
}

func (p *ProductStorage) UpdateProduct(product models.Product) (err error) {

	strQuery, err := database.DB.Prepare("UPDATE products SET name= ?, description = ?,price= ?,category_id= ? ,trade_mark= ? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	strQuery.Exec(product.Name, product.Description, product.Price, product.CategoryId, product.Trademark, product.Id)
	return err
}