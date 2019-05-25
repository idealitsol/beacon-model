package iam

import (
	"fmt"
	"time"

	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// Model Constants
const (
	UniqueConstraintUsername = "auth_user_username_key"
	UniqueConstraintEmail    = "uix_auth_user_email"
)

// AdminUser model
type AdminUser struct {
	ID             string     `json:"id,omitempty" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Username       string     `json:"username,omitempty" gorm:"unique;not null"`
	Password       string     `json:"-"`
	Fullname       string     `json:"fullname,omitempty" gorm:"not null"`
	Email          string     `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	AccountAccess  bool       `json:"accountAccess,omitempty" gorm:"default:true"`
	LoginCounter   int        `json:"loginCounter,omitempty" gorm:"default:0"`
	LastLogin      *time.Time `json:"lastLogin,omitempty" gorm:"default:null"`
	AccountExpiry  *time.Time `json:"accountExpiry,omitempty"`
	Photo          *string    `json:"photo,omitempty"`
	PwdExpiry      bool       `json:"pwdExpiry,omitempty" gorm:"default:false"`
	PwdExpiryTime  *time.Time `json:"pwdExpiryTime,omitempty"`
	PwdLifeInDays  int        `json:"pwdLifeInDays,omitempty" gorm:"default:0"`
	ForcePWDChange bool       `json:"forcePwdChange,omitempty" gorm:"default:false"`
	Institution    string     `json:"-" gorm:"type:UUID"`
}

// AdminUsers is an array of AdminUser objects
type AdminUsers []AdminUser

// BeforeCreate hook   http://gorm.io/docs/hooks.html
func (o *AdminUser) BeforeCreate(scope *gorm.Scope) error {
	// if u.IsValid() {
	// 	err = errors.New("can't save invalid data")
	// }
	scope.SetColumn("ID", util.TimeUUID().String())
	scope.SetColumn("Password", util.HashAndSalt([]byte(o.Password)))
	// o.CreatedAt = time.Now()
	// o.UpdatedAt = nil
	return nil
}

// AfterCreate hook
func (o *AdminUser) AfterCreate(scope *gorm.Scope) error {

	return nil
}

// BeforeUpdate hook
func (o *AdminUser) BeforeUpdate() (err error) {
	return
}

// ConstraintError handles all the database constrainst defined in a model
func (o *AdminUser) ConstraintError(err error) error {
	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Username '%s' already exists", o.Username), UniqueConstraintUsername); ok {
		return err
	}

	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Email '%s' already exists", o.Email), UniqueConstraintEmail); ok {
		return err
	}

	return nil
}

// AllowLogin Checks whether user should be allowed to login
func (o *AdminUser) AllowLogin() error {
	if o.AccountAccess == false {
		// return util.util.ErrLoginAccessDenied
		// return utils.NewError("Your account does not have access to this system", "LOGIN_ERROR")
		return fmt.Errorf("Sorry! Account does not have access to this system")
	}

	return nil
}

// ToString is a somewhat generic ToString method.
func (o *AdminUser) ToString() string {
	return o.ID + " " + o.Fullname
}
