package data

const (
	readAllOrdersQuery = "SELECT id, date, table_number, waiters.full_name, price, payment FROM orders JOIN waiters ON waiters.waiter_id = orders.waiter_id"
	readOrdersQuery    = "SELECT id, date, table_number, waiters.full_name, price, payment FROM orders JOIN waiters ON waiters.waiter_id = orders.waiter_id WHERE orders.id = $1"
	createOrderQuery   = "INSERT INTO orders VALUES ($1, $2, $3, $4, $5, $6)"
	updateOrderQuery   = "UPDATE orders SET price = $1, payment = $2 WHERE id = $3"
	deleteOrderQuery   = "DELETE FROM orders WHERE id = $1"
)
