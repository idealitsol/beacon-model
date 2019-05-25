package iam

import (
	"time"

	util "github.com/bekinsoft/beacon-util"
	"github.com/jinzhu/gorm"
)

// AuthRoleMapping model
type AuthRoleMapping struct {
	PrincipalId   string     `json:"principalId" gorm:"primary_key"`
	RoleId        string     `json:"roleId" gorm:"primary_key"`
	PrincipalType string     `json:"principalType" gorm:"varchar(30)"`
	Expiry        *time.Time `json:"-" gorm:"default:null"`
	Default       bool       `json:"default" gorm:"default:false"`
	Status        bool       `json:"status" gorm:"default:false"`

	AuthRole *AuthRole `json:"role,omitempty" gorm:"column:authRole;ForeignKey:ID;AssociationForeignKey:RoleID"`

	util.ModelCUD
}

// AuthRoleMappings is an array of AuthRoleMapping
type AuthRoleMappings []AuthRoleMapping

// BeforeCreate hook
func (o *AuthRoleMapping) BeforeCreate(scope *gorm.Scope) error {
	// if len(strconv.Itoa(o.ID)) < 10 {
	// 	return fmt.Errorf("Invalid Id length: Id must be a 10 digit number")
	// }

	return nil
}

// ConstraintError handles all the database constrainst defined in a model
func (o *AuthRoleMapping) ConstraintError(err error) error {
	if ok, err := util.IsConstraintError(err, "", ""); ok {
		return err
	}

	return nil
}
