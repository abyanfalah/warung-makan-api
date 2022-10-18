package utils

const (
	MENU_GET_ALL           = "SELECT id, name, price, stock, image FROM menu"
	MENU_GET_ALL_PAGINATED = MENU_GET_ALL + " limit $1 offset $2"
	MENU_GET_BY_ID         = MENU_GET_ALL + " WHERE id = $1"
	MENU_GET_BY_NAME       = MENU_GET_ALL + " WHERE name like $1"

	MENU_INSERT       = "INSERT INTO menu(id, name, price, stock) VALUES (:id, :name, :price, :stock)"
	MENU_UPDATE       = "UPDATE menu SET name=:name, price=:price, stock=:stock where id=:id"
	MENU_UPDATE_STOCK = "UPDATE menu SET stock=stock-:qty where id=:menu_id"
	MENU_DELETE       = "DELETE from menu WHERE id=$1"
	// ===========================================================

	USER_GET_ALL            = "SELECT id, name, username, image  FROM users"
	USER_GET_ALL_PAGINATED  = USER_GET_ALL + " limit $1 offset $2"
	USER_GET_BY_ID          = USER_GET_ALL + " WHERE id = $1"
	USER_GET_BY_NAME        = USER_GET_ALL + " WHERE name like $1"
	USER_GET_BY_CREDENTIALS = USER_GET_ALL + " WHERE username=$1 AND password=$2"

	USER_INSERT = "INSERT INTO users(id, name, username, password) VALUES (:id, :name, :username, :password)"
	USER_UPDATE = "UPDATE users SET name=:name, username=:username, password=:password where id=:id"
	USER_DELETE = "DELETE from users WHERE id=$1"
	// ===========================================================

	TRANSACTION_GET_ALL           = "SELECT id, total_price, created_at, updated_at FROM transaction "
	TRANSACTION_GET_ALL_PAGINATED = TRANSACTION_GET_ALL + " limit $1 offset $2"
	TRANSACTION_GET_BY_ID         = TRANSACTION_GET_ALL + " WHERE id = $1"
	TRANSACTION_GET_LAST_ID       = "SELECT id from transaction order by id desc limit 1"

	TRANSACTION_INSERT = "INSERT INTO transaction(id, total_price) VALUES (:id, :total_price)"
	TRANSACTION_UPDATE = "UPDATE transaction set total_price=:total_price where id=:id"
	TRANSACTION_DELETE = "DELETE from transaction WHERE id=$1"
	// ==============================================================

	TRANSACTION_DETAIL_INSERT                = "INSERT INTO transaction_detail(transaction_id, menu_id, qty, subtotal) VALUES ( :transaction_id, :menu_id, :qty, :subtotal)"
	TRANSACTION_DETAIL_GET_ALL               = "SELECT transaction_id, menu_id, qty, subtotal from transaction_detail "
	TRANSACTION_DETAIL_GET_BY_ID_TRANSACTION = TRANSACTION_DETAIL_GET_ALL + " where transaction_id = $1"
	// ==============================================================

	GET_DAILY_REPORT   = "SELECT date, COUNT(id) as transaction, SUM(total_price) as income from transaction group by date"
	GET_MONTHLY_REPORT = "SELECT date, COUNT(id) as transaction, SUM(total_price) as income from transaction where date between $1 and $2"

	GET_OVERALL_REPORT = "SELECT date, COUNT(id) as transaction, SUM(total_price) as income from transaction"
)
