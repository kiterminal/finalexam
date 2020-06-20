package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/kiterminal/finalexam/database"
	"net/http"
)

func createCustomer(c *gin.Context) {
	var reqBody Customer
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connector := database.Connect()
	var customer Customer
	if err := createCustomerService(connector, reqBody, &customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func createCustomerService(connector database.Connector, reqBody Customer, customer *Customer) error {
	var err error
	if customer.ID, customer.Name, customer.Email, customer.Status, err = connector.CreateCustomer(reqBody.Name, reqBody.Email, reqBody.Status); err != nil {
		return err
	}

	return nil
}
