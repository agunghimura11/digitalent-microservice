package database

import (
	"gorm.io/gorm"
	"github.com/pkg/errors"
)

type Auth struct {
	ID       	int    `json:"-" gorm:"primary_key"`
	Username 	string `json:"username,omitempty"`
	Password    string    `json:"password,omitempty"`
	Token 		string `json:"token,omitempty"`
}
// validasi token
func  ValidateAuth(token string,db *gorm.DB) (*Auth,error)  {
	var auth Auth

	// cek jika token match
	if err := db.Where(&Auth{Token: token}).First(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, errors.Errorf("invalid Token")
		}
	}

	return &auth, nil

}

func (auth *Auth) SignUp (db *gorm.DB) error {

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
