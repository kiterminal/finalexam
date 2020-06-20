package customer

import (
	"errors"
	"github.com/kiterminal/finalexam/database"
	"testing"
)

type mockConnector struct {
	database.Connector
}

func (m mockConnector) DeleteCustomerById(id string) error {
	return errors.New("can't prepare delete statement")
}

func TestCreateCustomerService(t *testing.T) {
	m := mockConnector{}
	err := deleteCustomerService(m, "1")

	if err == nil {
		t.Error("expect should return an error")
	}
}
