package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/kiterminal/finalexam/database"
	"github.com/kiterminal/finalexam/middleware"
	"net/http"
)

func getCustomers(c *gin.Context) {
	stmt, err := database.Conn().Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	rows, _ := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
}

func getCustomer(c *gin.Context) {
	id := c.Param("id")

	stmt, err := database.Conn().Prepare("SELECT id, name, email, status FROM customers WHERE id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	row := stmt.QueryRow(id)
	var customer Customer
	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find customer " + id})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func createCustomer(c *gin.Context) {
	var reqBody Customer
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer Customer
	err := database.Conn().QueryRow("INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id, name, email, status", reqBody.Name, reqBody.Email, reqBody.Status).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func updateCustomer(c *gin.Context) {
	id := c.Param("id")

	var reqBody Customer
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer Customer
	err := database.Conn().QueryRow("UPDATE customers SET name=$2,email=$3,status=$4 WHERE id=$1 RETURNING id, name, email, status;", id, reqBody.Name, reqBody.Email, reqBody.Status).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot find customer " + id})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")

	stmt, err := database.Conn().Prepare("DELETE FROM customers WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err = stmt.Exec(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer has been deleted"})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Authorization)

	r.GET("/customers", getCustomers)
	r.POST("/customers", createCustomer)
	r.GET("/customers/:id", getCustomer)
	r.PUT("/customers/:id", updateCustomer)
	r.DELETE("/customers/:id", deleteCustomer)

	return r
}
