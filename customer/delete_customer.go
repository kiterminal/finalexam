package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/kiterminal/finalexam/database"
	"net/http"
)

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")

	if err := deleteCustomerService(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}

func deleteCustomerService(id string) error {
	if err := database.DeleteById("customers", id); err != nil {
		return err
	}

	return nil
}
