package oma

import (
	"fmt"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

const (
	PayAfter  = "payAfter"
	PayBefore = "payBefore"
	Open      = "open"
)

// ApplAcc database model
// ApplAcc database model
type ApplAcc struct {
	ID            string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	ApplYear      int32      `json:"applYear" gorm:""`
	Username      string     `json:"username" gorm:"type:varchar(30);not null"`
	Password      string     `json:"password" gorm:"not null"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	LastLogin     *time.Time `json:"lastLogin"`
	SelectedForm  string     `json:"selectedForm" gorm:"type:UUID;"`
	InstitutionID string     `json:"institutionId" gorm:"type:UUID;"`
	IsComplete    bool       `json:"isComplete" gorm:"default:false"`
	Sname         string     `json:"sname" gorm:"type:varchar(40)"`
	Fname         string     `json:"fname" gorm:"type:varchar(60)"`
	Oname         string     `json:"oname" gorm:"type:varchar(30)"`
	LoginType     string     `json:"loginType" gorm:"type:varchar(10)"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplAccs is an array of ApplAcc objects
type ApplAccs []ApplAcc

// BeforeCreate hook
func (o *ApplAcc) BeforeCreate(scope *gorm.Scope) error {
	if valid, err := o.validate(); !valid {
		return err
	}

	scope.SetColumn("Password", util.HashAndSalt([]byte(o.Password)))
	return nil
}

// ApplAccP2STransformer transforms ApplAcc Protobuf to Struct
func ApplAccP2STransformer(data *pbx.ApplAcc) ApplAcc {
	model := ApplAcc{
		ApplYear:      data.GetApplYear(),
		Username:      data.GetUsername(),
		Password:      data.GetPassword(),
		CreatedAt:     util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:     util.GrpcTimeToGoTime(data.GetUpdatedAt()),
		LastLogin:     util.GrpcTimeToGoTime(data.GetLastLogin()),
		SelectedForm:  data.GetSelectedForm(),
		InstitutionID: data.GetInstitutionId(),
		IsComplete:    data.GetIsComplete(),
		Sname:         data.GetSname(),
		Fname:         data.GetFname(),
		Oname:         data.GetOname(),
		LoginType:     data.GetLoginType(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// ApplAccS2PTransformer transforms ApplAcc Struct to Protobuf
func ApplAccS2PTransformer(data ApplAcc) *pbx.ApplAcc {
	model := &pbx.ApplAcc{
		Id:            data.ID,
		ApplYear:      data.ApplYear,
		Username:      data.Username,
		Password:      data.Password,
		CreatedAt:     util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:     util.GoTimeToGrpcTime(data.UpdatedAt),
		LastLogin:     util.GoTimeToGrpcTime(data.LastLogin),
		SelectedForm:  data.SelectedForm,
		InstitutionId: data.InstitutionID,
		IsComplete:    data.IsComplete,
		Sname:         data.Sname,
		Fname:         data.Fname,
		Oname:         data.Oname,
		LoginType:     data.LoginType,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}

func (o *ApplAcc) validate() (bool, error) {
	if len(o.Username) == 0 {
		return false, fmt.Errorf("Username is required")
	}

	if len(o.Password) == 0 {
		return false, fmt.Errorf("Password is required")
	}

	/* Validate login Type */
	if len(o.LoginType) == 0 {
		return false, fmt.Errorf("Login type required")
	}
	if o.LoginType != PayAfter && o.LoginType != PayBefore && o.LoginType != Open {
		return false, fmt.Errorf("Invalid login type provided")
	}

	return true, nil
}
