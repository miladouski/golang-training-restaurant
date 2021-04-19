package data

const readAllOrdersQuery = "SELECT id, date, table_number, waiters.full_name, price, payment from orders join waiters on waiters.waiter_id = orders.waiter_id"
const readOrdersQuery = "SELECT id, date, table_number, waiters.full_name, price, payment from orders join waiters on waiters.waiter_id = orders.waiter_id WHERE orders.id = $1"
const createOrderQuery = "INSERT INTO orders VALUES ($1, $2, $3, $4, $5, $6)"
const updateOrderQuery = "Update orders SET price = $1, payment = $2 WHERE id = $3"
const deleteOrderQuery = "DELETE from orders where id = $1"
