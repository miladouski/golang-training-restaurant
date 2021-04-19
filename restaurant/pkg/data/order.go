package data

import (
	"fmt"
	"time"

	"database/sql"
)

type Order struct {
	Id       int
	Date     time.Time
	Table    int
	WaiterId int
	Price    int
	Payment  bool
}

type FullOrder struct {
	Id      int
	Date    time.Time
	Table   int
	Waiter  string
	Price   int
	Payment bool
}

func (o FullOrder) String() string {
	return fmt.Sprintf("(%d %s %d %s %d %t)", o.Id, o.Date, o.Table, o.Waiter, o.Price, o.Payment)
}

type OrderData struct {
	db *sql.DB
}

func NewOrderDate(db *sql.DB) *OrderData {
	return &OrderData{db: db}
}

func (o OrderData) ReadAll() ([]FullOrder, error) {
	var orders []FullOrder
	rows, err := o.db.Query(readAllOrdersQuery)
	if err != nil {
		return nil, fmt.Errorf("can't get orders from database, error:%w", err)
	}
	for rows.Next() {
		var tempUser FullOrder
		err = rows.Scan(&tempUser.Id, &tempUser.Date, &tempUser.Table, &tempUser.Waiter, &tempUser.Price, &tempUser.Payment)
		if err != nil {
			return nil, fmt.Errorf("can't scan orders from database, error:%w", err)
		}
		orders = append(orders, tempUser)
	}
	return orders, nil
}

func (o OrderData) Read(id int) (FullOrder, error) {
	var order FullOrder
	rows, err := o.db.Query(readOrdersQuery, id)
	if err != nil {
		return FullOrder{}, fmt.Errorf("can't get order from database, error:%w", err)
	}
	rows.Next()
	err = rows.Scan(&order.Id, &order.Date, &order.Table, &order.Waiter, &order.Price, &order.Payment)
	if err != nil {
		return FullOrder{}, fmt.Errorf("can't scan order from database, error:%w", err)
	}
	return order, nil
}

func (o OrderData) Create(order Order) error {
	_, err := o.db.Exec(createOrderQuery, order.Id, order.Date, order.Table, order.WaiterId, order.Price, order.Payment)
	if err != nil {
		return fmt.Errorf("can't inser order to database, error: %w", err)
	}
	return nil
}

func (o OrderData) Update(id, price int, payment bool) error {
	_, err := o.db.Exec(updateOrderQuery, price, payment, id)
	if err != nil {
		return fmt.Errorf("can't update order, error: %w", err)
	}
	return nil
}

func (o OrderData) Delete(id int) error {
	_, err := o.db.Exec(deleteOrderQuery, id)
	if err != nil {
		return fmt.Errorf("can't delete order, error: %w", err)
	}
	return nil
}
