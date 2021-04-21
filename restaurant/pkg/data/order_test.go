package data

import (
	"database/sql"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func NewGorm(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	})
	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gormDb
}

var testFullOrder = &FullOrder{
	Id:           6,
	Date:         time.Now(),
	Table_number: 1,
	Waiter:       "Mark",
	Price:        124,
	Payment:      true,
}

var testOrder = Order{
	Id:           1,
	Date:         time.Now(),
	Table_number: 1,
	WaiterId:     1,
	Price:        124,
	Payment:      true,
}

func TestReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	rows := sqlmock.NewRows([]string{"id", "Date", "Table_number", "Waiter", "Price", "Payment"}).
		AddRow(testFullOrder.Id, testFullOrder.Date, testFullOrder.Table_number, testFullOrder.Waiter, testFullOrder.Price, testFullOrder.Payment)
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
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectQuery(readAllOrdersQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(users)
}

func TestRead(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	rows := sqlmock.NewRows([]string{"id", "Date", "Table_number", "Waiter", "Price", "Payment"}).
		AddRow(testFullOrder.Id, testFullOrder.Date, testFullOrder.Table_number, testFullOrder.Waiter, testFullOrder.Price, testFullOrder.Payment)
	mock.ExpectQuery(readOrdersQuery).WithArgs(testFullOrder.Id).WillReturnRows(rows)
	orders, err := data.Read(testFullOrder.Id)
	assert.NoError(err)
	assert.Equal(orders, *testFullOrder)
}

func TestReadErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectQuery(readOrdersQuery).WithArgs(testFullOrder.Id).
		WillReturnError(errors.New("something went wrong..."))
	users, err := data.Read(testFullOrder.Id)
	assert.Error(err)
	assert.Empty(users)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(deleteOrderQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.Delete(1)
	assert.NoError(err)
}

func TestDeleteErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(deleteOrderQuery).
		WithArgs(1).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	err := data.Delete(1)
	assert.Error(err)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(updateOrderQuery).
		WithArgs(testOrder.Payment, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.Update(1, testOrder.Payment)
	assert.NoError(err)
}

func TestUpdateErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewOrderData(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(updateOrderQuery).
		WithArgs(testOrder.Payment, 1).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	err := data.Update(1, testOrder.Payment)
	assert.Error(err)
}
