package oma

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// ApplMain database model
type ApplMain struct {
	ID            string     `json:"id" gorm:"type:UUID;primary_key;size:36"`
	ApplYear      int32      `json:"applYear" gorm:""`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	InstitutionID string     `json:"institutionId" gorm:"type:UUID;"`
	IsComplete    bool       `json:"isComplete" gorm:"default:false"`
	ApplStatus    int32      `json:"applStatus" gorm:";default:'0'"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplMains is an array of ApplMain objects
type ApplMains []ApplMain

// BeforeCreate hook
func (o *ApplMain) BeforeCreate(scope *gorm.Scope) error {
	if valid, err := o.validate(); !valid {
		return err
	}

	return nil
}

// ApplMainP2STransformer transforms ApplMain Protobuf to Struct
func ApplMainP2STransformer(data *pbx.ApplMain) ApplMain {
	model := ApplMain{
		ApplYear:      data.GetApplYear(),
		CreatedAt:     util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:     util.GrpcTimeToGoTime(data.GetUpdatedAt()),
		InstitutionID: data.GetInstitutionId(),
		IsComplete:    data.GetIsComplete(),
		ApplStatus:    int32(data.GetApplStatus()),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// ApplMainS2PTransformer transforms ApplMain Struct to Protobuf
func ApplMainS2PTransformer(data ApplMain) *pbx.ApplMain {
	model := &pbx.ApplMain{
		Id:            data.ID,
		ApplYear:      data.ApplYear,
		CreatedAt:     util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:     util.GoTimeToGrpcTime(data.UpdatedAt),
		InstitutionId: data.InstitutionID,
		IsComplete:    data.IsComplete,
		ApplStatus:    int32(data.ApplStatus),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}

func (o *ApplMain) validate() (bool, error) {

	return true, nil
}
