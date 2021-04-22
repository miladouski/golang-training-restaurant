package data

const readAllOrdersQuery = `SELECT orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment FROM "orders" RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id`
const readOrdersQuery = `SELECT orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment FROM "orders" RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id WHERE orders.id = $1`
const createOrderQuery = `INSERT INTO "orders" VALUES ($1, $2, $3, $4, $5, $6)`
const updateOrderQuery = `UPDATE "orders" SET "payment"=$1 WHERE orders.id = $2`
const deleteOrderQuery = `DELETE FROM "orders" WHERE orders.id = $1`
const allOrders = "orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment"
const allOrdersJoin = "RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id"
const ordersTable = "orders"
const readWhere = "orders.id = ?"
const readOrder = "orders.id, orders.date, orders.table_number, waiters.full_name, orders.price, orders.payment"
const readOrderJoin = "RIGHT JOIN waiters on waiters.waiter_id = orders.waiter_id"
const updateColumn = "payment"
