package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo/v4"
)

type Product struct {
	ID string
	Name string
	Price float64
}

func main() {
		fmt.Println("Server run on port 3000")

		product := Product{
			ID: "1",
			Name: "Product 1",
			Price: 100,
		}
		err := SaveProduct(product)
		if err != nil {
			panic(err)
		}

		// http.HandleFunc("/", homeHandler)
		// http.ListenAndServe(":3000", nil) // GO ROUTINE
		e := echo.New()
		e.POST("/products", createProduct)
		e.Logger.Fatal(e.Start(":3000"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
}

func createProduct(c echo.Context) error {
		product := Product{}
		if err := c.Bind(&product); err != nil { // Bind JSON to struct
			return err
		}
		err := SaveProduct(product)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, product)
}

func SaveProduct(product Product) error{
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}