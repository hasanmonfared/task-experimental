package mysqldelayreport

import "gameapp/adapter/mysql"

type DB struct {
	adapter mysql.Adapter
}

func New(adapter mysql.Adapter) DB {
	return DB{adapter: adapter}
}
