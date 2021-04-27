package data

import (
	"database/sql"
	"errors"
	"log"
	"regexp"
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
	data := NewOrderData(db)
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
	data := NewOrderData(db)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(users)
}

func TestRead(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderData(db)
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
	data := NewOrderData(db)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.Read(1)
	assert.Error(err)
	assert.Empty(users)
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderData(db)
	mock.ExpectExec(regexp.QuoteMeta(createOrderQuery)).
		WithArgs(testOrder.Id, testOrder.Date, testOrder.Table, testOrder.WaiterId, testOrder.Price, testOrder.Payment).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.Create(testOrder)
	assert.NoError(err)
}

func TestCreateErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderData(db)
	mock.ExpectExec(regexp.QuoteMeta(createOrderQuery)).
		WithArgs(testOrder.Id, testOrder.Date, testOrder.Table, testOrder.WaiterId, testOrder.Price, testOrder.Payment).
		WillReturnError(errors.New("something went wrong..."))
	err := data.Create(testOrder)
	assert.Error(err)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateOrderQuery)).
		WithArgs(testOrder.Price, testOrder.Payment, testOrder.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.Update(testOrder.Id, testOrder.Price, testOrder.Payment)
	assert.NoError(err)
}

func TestUpdateErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderData(db)
	mock.ExpectExec(regexp.QuoteMeta(updateOrderQuery)).
		WithArgs(testOrder.Price, testOrder.Payment, testOrder.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.Update(testOrder.Id, testOrder.Price, testOrder.Payment)
	assert.Error(err)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderData(db)
	mock.ExpectExec(regexp.QuoteMeta(deleteOrderQuery)).
		WithArgs(testOrder.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := data.Delete(testOrder.Id)
	assert.NoError(err)
}

func TestDeleteErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewOrderData(db)
	mock.ExpectExec(regexp.QuoteMeta(deleteOrderQuery)).
		WithArgs(testOrder.Id).
		WillReturnError(errors.New("something went wrong..."))
	err := data.Delete(testOrder.Id)
	assert.Error(err)
}
