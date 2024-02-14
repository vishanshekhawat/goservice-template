package repo

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"

	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

type MySQLDB struct {
	*sql.DB
}

// Connect initializes a connection to the MySQL database.
func (m *MySQLDB) Connect(cfg models.Config) error {
	sslMode := "require"
	// if cfg.DisableTLS {
	// 	sslMode = "disable"
	// }

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "mysql",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     "127:0:0:1:3306",
		Path:     "sale_api",
		RawQuery: q.Encode(),
	}
	fmt.Println(u.String())
	db, err := sql.Open("mysql", "user:"+cfg.Password+"@tcp("+cfg.HostPort+":3306)/users_db")

	// db, err := sql.Open("mysql", u.String())
	if err != nil {
		//panic(err.Error())
		return err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	m.DB = db
	return nil

}

// Close closes the connection to the MySQL database.
func (m *MySQLDB) Close() error {

	fmt.Println("Closed MySQL database connection")
	return m.DB.Close()
}

func (m *MySQLDB) GetConn() *sql.DB {
	return m.DB
}
