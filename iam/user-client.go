package iam

import (
	"fmt"
	"time"

	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// ClientUser model
type ClientUser struct {
	Id             string            `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username       string            `json:"username" gorm:"unique;not null"`
	Password       string            `json:"-"`
	PIN            *int              `json:"-" gorm:"type:int(6)"`
	Fullname       string            `json:"fullname" gorm:"not null"`
	Email          string            `json:"email" gorm:"type:varchar(100);unique_index"`
	AccountAccess  bool              `json:"accountAccess" gorm:"default:true"`
	LoginCounter   int               `json:"loginCounter" gorm:"default:0"`
	LastLogin      *time.Time        `json:"lastLogin" gorm:"default:null"`
	AccountExpiry  *time.Time        `json:"accountExpiry"`
	Photo          *string           `json:"photo"`
	PwdExpiry      bool              `json:"pwdExpiry" gorm:"default:false"`
	PwdExpiryTime  *time.Time        `json:"pwdExpiryTime"`
	PwdLifeInDays  int               `json:"pwdLifeInDays" gorm:"default:0"`
	ForcePwdChange bool              `json:"forcePwdChange" gorm:"default:false"`
	Institution    string            `json:"-" gorm:"type:UUID"`
	Roles          []AuthRoleMapping `json:"-" gorm:"foreignkey:User"`

	util.ModelCUD
}

// ClientUsers is an array of AuthUsers objects
type ClientUsers []ClientUser

// BeforeCreate hook   http://gorm.io/docs/hooks.html
func (o *ClientUser) BeforeCreate(scope *gorm.Scope) error {
	// salt = 2FFDBBD2051702898CC1150C66FD41F649ACF020F81E64AAEBD7B (default)
	// var salt = viper.GetString("app.auth.client.add.password.salt")
	// scope.SetColumn("Email", o.Username+"@"+viper.GetString("app.auth.client.add.email.domain"))
	scope.SetColumn("AccountAccess", true)
	scope.SetColumn("ForcePWDChange", true)
	// scope.SetColumn("Password", util.HashAndSalt([]byte(o.Password+salt)))

	return nil
}

// AfterCreate hook
func (o *ClientUser) AfterCreate(scope *gorm.Scope) error {

	return nil
}

// BeforeUpdate hook
func (o *ClientUser) BeforeUpdate() (err error) {
	return
}

// ConstraintError handles all the database constrainst defined in a model
func (o *ClientUser) ConstraintError(err error) error {
	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Username '%s' already exists", o.Username), UniqueConstraintUsername); ok {
		return err
	}

	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Email '%s' already exists", o.Email), UniqueConstraintEmail); ok {
		return err
	}

	return nil
}

// AllowLogin Checks whether user should be allowed to login
func (o *ClientUser) AllowLogin() error {
	if o.AccountAccess == false {
		// return util.ErrLoginAccessDenied
		// return utils.NewError("Your account does not have access to this system", "LOGIN_ERROR")
		return fmt.Errorf("Sorry! Account does not have access to this system")
	}

	return nil
}

// ToString is a somewhat generic ToString method.
func (o *ClientUser) ToString() string {
	return o.Id + " " + o.Fullname
}
