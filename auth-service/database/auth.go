package database

import (
	"gorm.io/gorm"
	"github.com/pkg/errors"
)

type Auth struct {
	ID       	int    `json:"-" gorm:"primary_key"`
	Username 	string `json:"username"`
	Password    string    `json:"password"`
	Token 		string `json:"token"`
}

//func  ValidateAuth(token string,db *gorm.DB) (*Auth,error)  {
//	var auth Auth
//
//	if err := db.Where(&Auth{Token: token}).First(&auth).Error; err != nil {
//		return nil, errors.Errorf("invalid Token")
//	}
//
//}

func (auth *Auth) SignUp (db *gorm.DB) error {
	// SELECT * FROM AUTH WHERE username = "fadhlan@gmail.com"
	if err:= db.Where(&Auth{Username: auth.Username}).First(auth).Error;err != nil {
		if err == gorm.ErrRecordNotFound{
			if err = db.Create(auth).Error; err != nil {
				return err
			}
		}
	}else{
		return errors.Errorf("Duplicate Email")
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
