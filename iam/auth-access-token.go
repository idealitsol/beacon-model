package iam

import (
	"time"

	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// AuthAccessToken model
type AuthAccessToken struct {
	ID   string    `json:"id" gorm:"primary_key"`
	UserID   string    `json:"uuid" gorm:"column:userid"`
	TTL        int32    `json:"ttl"`
	IP       string      `json:"ip" gorm:"varchar(30)"`
	PrincipalType string    `json:"principalType" gorm:"column:principaltype;type:varchar(30)"`
	Platform string    `json:"platform" gorm:"varchar(10)"`
	Created       time.Time  `json:"created"`
}

// AuthAccessTokens is an array of AuthRoleMapping
type AuthAccessTokens []AuthAccessToken

// BeforeCreate hook
func (o *AuthAccessToken) BeforeCreate(scope *gorm.Scope) error {
	// if len(strconv.Itoa(o.ID)) < 10 {
	// 	return fmt.Errorf("Invalid Id length: Id must be a 10 digit number")
	// }

	return nil
}

// ConstraintError handles all the database constrainst defined in a model
func (o *AuthAccessToken) ConstraintError(err error) error {
	if ok, err := util.IsConstraintError(err, "", ""); ok {
		return err
	}

	return nil
}
