package data

const (
	readAllOrdersQuery = `SELECT orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment FROM "orders" RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id`
	readOrdersQuery    = `SELECT orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment FROM "orders" RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id WHERE orders.id = $1`
	updateOrderQuery   = `UPDATE "orders" SET "payment"=$1 WHERE orders.id = $2`
	deleteOrderQuery   = `DELETE FROM "orders" WHERE orders.id = $1`
	allOrders          = "orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment"
	allOrdersJoin      = "RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id"
	ordersTable        = "orders"
	readWhere          = "orders.id = ?"
	readOrder          = "orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment"
	readOrderJoin      = "RIGHT JOIN waiters ON waiters.waiter_id = orders.waiter_id"
	updateColumn       = "payment"
)
