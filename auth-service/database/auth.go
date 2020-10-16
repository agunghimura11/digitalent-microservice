package database

import (
	"gorm.io/gorm"
	"github.com/pkg/errors"
)

type Auth struct {
	ID       	int    `json:"-" gorm:"primary_key"`
	Username 	string `json:"username"`
	Password    int    `json:"password"`
	Token 		string `json:"token"`
}

func (auth *Auth) SignUp (db *gorm.DB) error {
	// SELECT * FROM AUTH WHERE username = "fadhlan@gmail.com"
	if err:= db.Where(&Auth{Username: auth.Username}).First(auth).Error;err != nil {
		if err == gorm.ErrRecordNotFound{
			if err = db.Create(auth).Error; err != nil {
				return err
			}
		}

		return err
	}

	return nil
}

func (auth *Auth) Login (db *gorm.DB) (*Auth, error)  {
	if err := db.Where(&Auth{Username: auth.Username, Password: auth.Password}).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Errorf("incorect email")
		}
	}

	return auth, nil
}
