package client

import (
	"fmt"

	"entgo.io/ent/dialect"
)

func NewMemorySQLiteDialect() (driver, dataSource string) {
	return dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1"
}

type MysqlDialectOpts struct {
	User string
	Pass string
	Host string
	Port int
	DB   string
}

func NewMySQLDialect(opts MysqlDialectOpts) (driver, dataSource string) {
	return dialect.MySQL, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=True",
		opts.User, opts.Pass, opts.Host, opts.Port, opts.DB,
	)
}
