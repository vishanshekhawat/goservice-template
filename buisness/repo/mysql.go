package repo

import (
	"database/sql"
	"fmt"
	"net/url"
)

type MySQLDB struct {
	*sql.DB
}

// Connect initializes a connection to the MySQL database.
func (m *MySQLDB) Connect() error {
	sslMode := "require"
	// if cfg.DisableTLS {
	// 	sslMode = "disable"
	// }

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("cfg.User", "cfg.Password"),
		Host:     "localhost",
		Path:     "user_service_db_dev",
		RawQuery: q.Encode(),
	}

	db, err := sql.Open("pgx", u.String())
	if err != nil {
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
