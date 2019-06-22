package iam

import (
	"fmt"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"

	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// Model Constants
const (
	UniqueConstraintUsername = "uix_auth_user_username"
	UniqueConstraintEmail    = "uix_auth_user_email"
)

// AdminUser model
type AdminUser struct {
	Id             string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Username       string     `json:"username" gorm:"unique;not null"`
	Password       string     `json:"-"`
	Fullname       string     `json:"fullname" gorm:"not null"`
	Email          string     `json:"email" gorm:"type:varchar(100);unique_index"`
	AccountAccess  bool       `json:"accountAccess" gorm:"default:true"`
	LoginCounter   int        `json:"loginCounter" gorm:"default:0"`
	LastLogin      *time.Time `json:"lastLogin" gorm:"default:null"`
	AccountExpiry  *time.Time `json:"accountExpiry"`
	Photo          *string    `json:"photo"`
	PwdExpiry      bool       `json:"pwdExpiry" gorm:"default:false"`
	PwdExpiryTime  *time.Time `json:"pwdExpiryTime"`
	PwdLifeInDays  int        `json:"pwdLifeInDays" gorm:"default:0"`
	ForcePwdChange bool       `json:"forcePwdChange" gorm:"default:false"`
	InstitutionID  string     `json:"-" gorm:"type:UUID"`

	util.ModelCUD

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// AdminUsers is an array of AdminUser objects
type AdminUsers []AdminUser

// BeforeCreate hook   http://gorm.io/docs/hooks.html
func (o *AdminUser) BeforeCreate(scope *gorm.Scope) error {
	if valid, err := o.validate(); !valid {
		return err
	}

	scope.SetColumn("ID", util.TimeUUID().String())
	scope.SetColumn("Password", util.HashAndSalt([]byte(o.Password)))
	now := time.Now()
	o.CreatedAt = &now
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

func (o *AdminUser) validate() (bool, error) {
	if len(o.Username) == 0 {
		return false, fmt.Errorf("Username is required")
	}

	if len(o.Password) == 0 {
		return false, fmt.Errorf("Password is required")
	}

	if len(o.Email) == 0 {
		return false, fmt.Errorf("Email is required")
	}

	if len(o.Fullname) == 0 {
		return false, fmt.Errorf("Fullname is required")
	}

	return true, nil
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
	return o.Id + " " + o.Fullname
}

// AdminUserP2STransformer transforms Protobuf to Struct
func AdminUserP2STransformer(data *pbx.AdminUser) AdminUser {
	model := AdminUser{
		Username:       data.GetUsername(),
		Password:       data.GetPassword(),
		Fullname:       data.GetFullname(),
		Email:          data.GetEmail(),
		AccountAccess:  data.GetAccountAccess(),
		LoginCounter:   int(data.GetLoginCounter()),
		LastLogin:      util.GrpcTimeToGoTime(data.GetLastLogin()),
		AccountExpiry:  util.GrpcTimeToGoTime(data.GetAccountExpiry()),
		PwdExpiry:      data.GetPwdExpiry(),
		PwdExpiryTime:  util.GrpcTimeToGoTime(data.GetPwdExpiryTime()),
		PwdLifeInDays:  int(data.GetPwdLifeInDays()),
		ForcePwdChange: data.GetForcePwdChange(),
		InstitutionID:  data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetId has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.Id = data.GetId()
	}

	// Handling pointer string
	if len(data.GetPhoto()) != 0 {
		pix := data.GetPhoto()
		model.Photo = &pix
	}

	return model
}

// AdminUserS2PTransformer transforms Struct to Protobuf
func AdminUserS2PTransformer(data AdminUser) *pbx.AdminUser {
	model := &pbx.AdminUser{
		Id:             data.Id,
		Username:       data.Username,
		Password:       data.Password,
		Fullname:       data.Fullname,
		Email:          data.Email,
		AccountAccess:  data.AccountAccess,
		LoginCounter:   int32(data.LoginCounter),
		LastLogin:      util.GoTimeToGrpcTime(data.LastLogin),
		AccountExpiry:  util.GoTimeToGrpcTime(data.AccountExpiry),
		PwdExpiry:      data.PwdExpiry,
		PwdExpiryTime:  util.GoTimeToGrpcTime(data.PwdExpiryTime),
		PwdLifeInDays:  int32(data.PwdLifeInDays),
		ForcePwdChange: data.ForcePwdChange,
		InstitutionId:  data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handling pointer string
	if data.Photo != nil {
		model.Photo = *data.Photo
	}

	return model
}
