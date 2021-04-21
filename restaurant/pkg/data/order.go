package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id           int
	Date         time.Time
	Table_number int
	WaiterId     int
	Price        int
	Payment      bool
}

type FullOrder struct {
	Id           int
	Date         time.Time
	Table_number int
	Waiter       string
	Price        int
	Payment      bool
}

func (o FullOrder) String() string {
	return fmt.Sprintf("(%d %s %d %s %d %t)", o.Id, o.Date, o.Table_number, o.Waiter, o.Price, o.Payment)
}

type OrderData struct {
	db *gorm.DB
}

func NewOrderData(db *gorm.DB) *OrderData {
	return &OrderData{db: db}
}

func (o OrderData) ReadAll() ([]FullOrder, error) {
	var orders []FullOrder
	err := o.db.Table("orders").
		Select("orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment").
		Joins("RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id").
		Find(&orders)
	if err.Error != nil {
		return nil, err.Error
	}
	return orders, nil
}

func (o OrderData) Read(id int) (FullOrder, error) {
	var order FullOrder
	err := o.db.Table("orders").
		Where("orders.id = ?", id).
		Select("orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment").
		Joins("RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id").
		Find(&order)

	if err.Error != nil {
		return FullOrder{}, err.Error
	}
	return order, nil
}

func (o OrderData) Create(order Order) error {
	err := o.db.Create(&order)
	if err.Error != nil {
		return fmt.Errorf("error: %s", err.Error)
	}
	return nil
}

func (o OrderData) Update(id int, payment bool) error {
	err := o.db.Table("orders").Where("orders.id = ?", id).Update("payment", payment)
	if err.Error != nil {
		return fmt.Errorf("error: %s", err.Error)
	}
	return nil
}

func (o OrderData) Delete(id int) error {
	err := o.db.Where("orders.id = ?", id).Delete(&Order{})
	if err.Error != nil {
		return fmt.Errorf("error: %s", err.Error)
	}
	return nil
}
