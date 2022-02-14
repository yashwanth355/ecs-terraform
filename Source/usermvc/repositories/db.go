package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
	Config "usermvc/config"
)

func NewDb() (*gorm.DB, error) {
	conf := Config.LoadConfig()

	var (
		host     = conf.DbConfig.Host
		port     = conf.DbConfig.Port
		user     = conf.DbConfig.User
		password = conf.DbConfig.Password
		dbname   = conf.DbConfig.Name
	)

	fmt.Println("prting the user and password", conf.DbConfig)
	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     dbname,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	fmt.Println(password)
	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		return nil, err
	}
	return db, nil
}
