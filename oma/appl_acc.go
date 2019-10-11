package oma

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"

	util "github.com/idealitsol/beacon-util"
)

// ApplAcc database model
type ApplAcc struct {
	ID            string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	ApplYear      int32      `json:"applYear" gorm:""`
	Username      string     `json:"username" gorm:"type:varchar(20);not null"`
	Password      string     `json:"password" gorm:"type:varchar(50);not null"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	LastLogin     *time.Time `json:"lastLogin"`
	SelectedForm  string     `json:"selectedForm" gorm:"type:UUID"`
	InstitutionID string     `json:"institutionId" gorm:"type:UUID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplAccs is an array of ApplAcc objects
type ApplAccs []ApplAcc

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

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
