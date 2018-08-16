package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/convee/goboot/conf"
	_ "github.com/go-sql-driver/mysql"
)

func New(name string) *sql.DB {
	configs := conf.Get().Mysql
	return newDb(configs[name])
}

func newDb(config conf.MysqlConfig) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.Username,
		config.Password,
		config.Ip,
		config.Port,
		config.Database,
		config.Charset,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(dsn)
	// defer db.Close()
	db.SetMaxIdleConns(config.MaxIdle)                                     //设置闲置的连接数
	db.SetMaxOpenConns(config.MaxOpen)                                     //设置最大打开的连接数，默认值0表示不限制
	db.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second) //设置长连接的最长使用时间（从创建时开始计算），超过该时间go会自动关闭该连接
	return db
}
