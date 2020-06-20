package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/kiterminal/finalexam/database"
	"net/http"
)

func deleteCustomer(c *gin.Context) {
	connector := database.Connect()
	id := c.Param("id")

	if err := deleteCustomerService(connector, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}

func deleteCustomerService(connector database.Connector, id string) error {
	if err := connector.DeleteCustomerById(id); err != nil {
		return err
	}

	return nil
}
