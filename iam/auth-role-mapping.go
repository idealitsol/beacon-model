package iam

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"

	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// AuthRoleMapping database model
type AuthRoleMapping struct {
	PrincipalID   string     `json:"principalId" gorm:"primary_key;not null"`
	RoleID        string     `json:"roleId" gorm:"primary_key;not null"`
	PrincipalType string     `json:"principalType" gorm:""`
	Expiry        *time.Time `json:"-"`
	Default       bool       `json:"default" gorm:"default:false"`
	Status        bool       `json:"status" gorm:"default:false"`
	CreatedBy     string     `json:"-" gorm:""`
	UpdatedBy     string     `json:"-" gorm:""`
	DeletedBy     string     `json:"-" gorm:""`
	CreatedAt     *time.Time `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
	UserID        string     `json:"userId" gorm:""`
	UserType      string     `json:"userType" gorm:""`
	Institution   string     `json:"institution" gorm:"type:UUID;"`

	AuthRole *AuthRole `json:"role,omitempty" gorm:"column:authRole;ForeignKey:ID;AssociationForeignKey:RoleID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// // AuthRoleMapping model
// type AuthRoleMapping struct {
// 	PrincipalId   string     `json:"principalId" gorm:"primary_key"`
// 	RoleId        string     `json:"roleId" gorm:"primary_key"`
// 	PrincipalType string     `json:"principalType" gorm:"varchar(30)"`
// 	Expiry        *time.Time `json:"-" gorm:"default:null"`
// 	Default       bool       `json:"default" gorm:"default:false"`
// 	Status        bool       `json:"status" gorm:"default:false"`

// 	AuthRole *AuthRole `json:"role,omitempty" gorm:"column:authRole;ForeignKey:ID;AssociationForeignKey:RoleID"`

// 	util.ModelCUD
// }

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

// AuthRoleMappingP2STransformer transforms AuthRoleMapping Protobuf to Struct
func AuthRoleMappingP2STransformer(data *pbx.AuthRoleMapping) AuthRoleMapping {
	model := AuthRoleMapping{
		RoleID:        data.GetRoleId(),
		PrincipalType: data.GetPrincipalType(),
		Expiry:        util.GrpcTimeToGoTime(data.GetExpiry()),
		Default:       data.GetDefault(),
		Status:        data.GetStatus(),
		CreatedBy:     data.GetCreatedBy(),
		UpdatedBy:     data.GetUpdatedBy(),
		DeletedBy:     data.GetDeletedBy(),
		CreatedAt:     util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:     util.GrpcTimeToGoTime(data.GetUpdatedAt()),
		DeletedAt:     util.GrpcTimeToGoTime(data.GetDeletedAt()),
		UserID:        data.GetUserId(),
		UserType:      data.GetUserType(),
		Institution:   data.GetInstitution(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetPrincipalID has no value then it's a POST request (Create)
	if len(data.GetPrincipalId()) != 0 {
		model.PrincipalID = data.GetPrincipalId()
	}

	// Handle pointers after this

	return model
}

// AuthRoleMappingS2PTransformer transforms AuthRoleMapping Struct to Protobuf
func AuthRoleMappingS2PTransformer(data AuthRoleMapping) *pbx.AuthRoleMapping {
	model := &pbx.AuthRoleMapping{
		PrincipalId:   data.PrincipalID,
		RoleId:        data.RoleID,
		PrincipalType: data.PrincipalType,
		Expiry:        util.GoTimeToGrpcTime(data.Expiry),
		Default:       data.Default,
		Status:        data.Status,
		CreatedBy:     data.CreatedBy,
		UpdatedBy:     data.UpdatedBy,
		DeletedBy:     data.DeletedBy,
		CreatedAt:     util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:     util.GoTimeToGrpcTime(data.UpdatedAt),
		DeletedAt:     util.GoTimeToGrpcTime(data.DeletedAt),
		UserId:        data.UserID,
		UserType:      data.UserType,
		Institution:   data.Institution,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
