package data

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

var testFullOrder = &FullOrder{
	Id:      6,
	Date:    time.Now(),
	Table:   1,
	Waiter:  "Mark",
	Price:   124,
	Payment: true,
}

var testOrder = Order{
	Id:       1,
	Date:     time.Now(),
	Table:    1,
	WaiterId: 1,
	Price:    124,
	Payment:  true,
}

func TestReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderDate(db)
	rows := sqlmock.NewRows([]string{"id", "Date", "table_number", "waiters.full_name", "Price", "Payment"}).AddRow(testFullOrder.Id, testFullOrder.Date, testFullOrder.Table, testFullOrder.Waiter, testFullOrder.Price, testFullOrder.Payment)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnRows(rows)
	orders, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(orders)
	assert.Equal(orders[0], *testFullOrder)
	assert.Len(orders, 1)
}

func TestReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderDate(db)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(users)
}

func TestRead(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderDate(db)
	rows := sqlmock.NewRows([]string{"id", "Date", "table_number", "waiters.full_name", "Price", "Payment"}).AddRow(testFullOrder.Id, testFullOrder.Date, testFullOrder.Table, testFullOrder.Waiter, testFullOrder.Price, testFullOrder.Payment)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnRows(rows)
	orders, err := data.Read(1)
	assert.NoError(err)
	assert.Equal(orders, *testFullOrder)
}

func TestReadErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderDate(db)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.Read(1)
	assert.Error(err)
	assert.Empty(users)
}
