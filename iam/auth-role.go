package iam

import (
	"fmt"

	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// AuthRole model

// Model Constants
const (
	AuthRoleUniqueConstraintName = "auth_role_name_key"
)

// AuthRole model
type AuthRole struct {
	Id          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"varchar(30);not null;unique_index"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Default     bool   `json:"default" gorm:"default:false"`
	Status      bool   `json:"status" gorm:"default:true"`

	util.ModelCUDAt

	AuthRoleMapping *AuthRoleMapping `json:"roleMapping,omitempty" gorm:"ForeignKey:ID;AssociationForeignKey:RoleID"`
}

// AuthRoles is array of Authrole
type AuthRoles []AuthRole

// BeforeCreate hook
func (o *AuthRole) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", util.TimeUUID().String())
	return nil
}

// AfterCreate hook
func (o *AuthRole) AfterCreate(scope *gorm.Scope) error {
	return nil
}

// ConstraintError handles all the database constrainst defined in a model
func (o *AuthRole) ConstraintError(err error) error {
	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Name '%s' already exists", o.Name), AuthRoleUniqueConstraintName); ok {
		return err
	}

	return nil
}
