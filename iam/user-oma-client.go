package iam

import (
	"fmt"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	"github.com/jinzhu/gorm"

	util "github.com/idealitsol/beacon-util"
)

// Model Constants
const (
	uniqueConstraintUsername = "uix_auth_user_username"
)

// OmaClientUser database model
type OmaClientUser struct {
	ID                string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Username          string     `json:"username" gorm:"type:varchar(30);not null"`
	Password          string     `json:"password" gorm:"not null"`
	Fullname          string     `json:"fullname" gorm:"type:varchar(200);not null"`
	Email             string     `json:"email" gorm:"type:varchar(100)"`
	AccountAccess     bool       `json:"accountAccess" gorm:"default:true"`
	LoginPassCount    int32      `json:"loginPassCount" gorm:"default:(0)"`
	LastLogin         *time.Time `json:"lastLogin"`
	AccountExpiry     *time.Time `json:"accountExpiry"`
	Photo             string     `json:"photo" gorm:""`
	PwdExpiry         bool       `json:"pwdExpiry" gorm:"default:false"`
	PwdExpiryTime     *time.Time `json:"pwdExpiryTime"`
	PwdLifeInDays     int32      `json:"pwdLifeInDays" gorm:";default:(0)"`
	CreatedBy         string     `json:"createdBy" gorm:""`
	UpdatedBy         string     `json:"updatedBy" gorm:""`
	DeletedBy         string     `json:"deletedBy" gorm:""`
	CreatedAt         *time.Time `json:"createdAt"`
	UpdatedAt         *time.Time `json:"updatedAt"`
	DeletedAt         *time.Time `json:"deletedAt"`
	Pin               string     `json:"pin" gorm:"type:varchar(4)"`
	ForcePwdChange    bool       `json:"forcePwdChange" gorm:"default:false"`
	LoginFailCount    int32      `json:"loginFailCount" gorm:"default:0"`
	LoginCounter      int32      `json:"loginCounter" gorm:"default:0"`
	InstitutionID     string     `json:"institutionId" gorm:"type:UUID;"`
	Verified          int32      `json:"verified" gorm:";default:-1"`
	VerificationToken string     `json:"verificationToken" gorm:""`
	VerificationType  string     `json:"verificationType" gorm:"type:varchar(4)"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// OmaClientUsers is an array of OmaClientUser objects
type OmaClientUsers []OmaClientUser

// BeforeCreate hook   http://gorm.io/docs/hooks.html
func (o *OmaClientUser) BeforeCreate(scope *gorm.Scope) error {
	if valid, err := o.validate(); !valid {
		return err
	}

	scope.SetColumn("ID", util.TimeUUID().String())
	scope.SetColumn("Password", util.HashAndSalt([]byte(o.Password)))
	now := time.Now()
	o.CreatedAt = &now
	return nil
}

func (o *OmaClientUser) validate() (bool, error) {
	if len(o.Username) == 0 {
		return false, fmt.Errorf("Username is required")
	}

	if len(o.Password) == 0 {
		return false, fmt.Errorf("Password is required")
	}

	// if len(o.Email) == 0 {
	// 	return false, fmt.Errorf("Email is required")
	// }

	if len(o.Fullname) == 0 {
		return false, fmt.Errorf("Fullname is required")
	}

	return true, nil
}

// ConstraintError handles all the database constrainst defined in a model
// func (o *OmaClientUser) ConstraintError(err error) error {
// 	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Username '%s' already exists", o.Username), UniqueConstraintUsername); ok {
// 		return err
// 	}

// 	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Email '%s' already exists", o.Email), UniqueConstraintEmail); ok {
// 		return err
// 	}

// 	return nil
// }

// AllowLogin Checks whether user should be allowed to login
func (o *OmaClientUser) AllowLogin() error {
	if o.AccountAccess == false {
		// return util.util.ErrLoginAccessDenied
		// return utils.NewError("Your account does not have access to this system", "LOGIN_ERROR")
		return fmt.Errorf("Sorry! Account does not have access to this system")
	}

	return nil
}

// OmaClientUserP2STransformer transforms OmaClientUser Protobuf to Struct
func OmaClientUserP2STransformer(data *pbx.OmaClientUser) OmaClientUser {
	model := OmaClientUser{
		Username:          data.GetUsername(),
		Password:          data.GetPassword(),
		Fullname:          data.GetFullname(),
		Email:             data.GetEmail(),
		AccountAccess:     data.GetAccountAccess(),
		LoginPassCount:    data.GetLoginPassCount(),
		LastLogin:         util.GrpcTimeToGoTime(data.GetLastLogin()),
		AccountExpiry:     util.GrpcTimeToGoTime(data.GetAccountExpiry()),
		Photo:             data.GetPhoto(),
		PwdExpiry:         data.GetPwdExpiry(),
		PwdExpiryTime:     util.GrpcTimeToGoTime(data.GetPwdExpiryTime()),
		PwdLifeInDays:     int32(data.GetPwdLifeInDays()),
		CreatedBy:         data.GetCreatedBy(),
		UpdatedBy:         data.GetUpdatedBy(),
		DeletedBy:         data.GetDeletedBy(),
		CreatedAt:         util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:         util.GrpcTimeToGoTime(data.GetUpdatedAt()),
		DeletedAt:         util.GrpcTimeToGoTime(data.GetDeletedAt()),
		Pin:               data.GetPin(),
		ForcePwdChange:    data.GetForcePwdChange(),
		LoginFailCount:    int32(data.GetLoginFailCount()),
		LoginCounter:      int32(data.GetLoginCounter()),
		InstitutionID:     data.GetInstitutionId(),
		Verified:          int32(data.GetVerified()),
		VerificationToken: data.GetVerificationToken(),
		VerificationType:  data.GetVerificationType(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// OmaClientUserS2PTransformer transforms OmaClientUser Struct to Protobuf
func OmaClientUserS2PTransformer(data OmaClientUser) *pbx.OmaClientUser {
	model := &pbx.OmaClientUser{
		Id:                data.ID,
		Username:          data.Username,
		Password:          data.Password,
		Fullname:          data.Fullname,
		Email:             data.Email,
		AccountAccess:     data.AccountAccess,
		LoginPassCount:    data.LoginPassCount,
		LastLogin:         util.GoTimeToGrpcTime(data.LastLogin),
		AccountExpiry:     util.GoTimeToGrpcTime(data.AccountExpiry),
		Photo:             data.Photo,
		PwdExpiry:         data.PwdExpiry,
		PwdExpiryTime:     util.GoTimeToGrpcTime(data.PwdExpiryTime),
		PwdLifeInDays:     int32(data.PwdLifeInDays),
		CreatedBy:         data.CreatedBy,
		UpdatedBy:         data.UpdatedBy,
		DeletedBy:         data.DeletedBy,
		CreatedAt:         util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:         util.GoTimeToGrpcTime(data.UpdatedAt),
		DeletedAt:         util.GoTimeToGrpcTime(data.DeletedAt),
		Pin:               data.Pin,
		ForcePwdChange:    data.ForcePwdChange,
		LoginFailCount:    int32(data.LoginFailCount),
		LoginCounter:      int32(data.LoginCounter),
		InstitutionId:     data.InstitutionID,
		Verified:          int32(data.Verified),
		VerificationToken: data.VerificationToken,
		VerificationType:  data.VerificationType,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
