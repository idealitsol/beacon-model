package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// AuthRole model

// Model Constants
const (
// AuthRoleUniqueConstraintName = "auth_role_name_key"
)

// AuthAccessToken model
type AuthAccessToken struct {
	ID            string    `gorm:"primary_key"`
	ttl           int32     `gorm:"varchar(30);not null"`
	UserID        string    `gorm:"column:userid;type:uuid"`
	IP            string    `gorm:"column:ip"`
	PrincipalType string    `gorm:"column:principaltype"`
	Platform      string    ``
	Created       time.Time ``

	// AuthRoleMapping *AuthRoleMapping `json:"roleMapping,omitempty" gorm:"ForeignKey:ID;AssociationForeignKey:RoleID"`
}

// AuthAccessTokens is array of Authrole
type AuthAccessTokens []AuthAccessToken

// BeforeCreate hook
func (o *AuthAccessToken) BeforeCreate(scope *gorm.Scope) error {
	//scope.SetColumn("ID", util.TimeUUID().String())
	return nil
}

// AfterCreate hook
func (o *AuthAccessToken) AfterCreate(scope *gorm.Scope) error {
	return nil
}
