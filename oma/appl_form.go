package oma

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
)

// ApplForm database model
type ApplForm struct {
	ApplicantID string     `json:"applicantId" gorm:"type:UUID;primary_key;size:36"`
	FormID      string     `json:"formId" gorm:"type:UUID;primary_key;size:36"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplForms is an array of ApplForm objects
type ApplForms []ApplForm

// ApplFormP2STransformer transforms ApplForm Protobuf to Struct
func ApplFormP2STransformer(data *pbx.ApplForm) ApplForm {
	model := ApplForm{
		CreatedAt: util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt: util.GrpcTimeToGoTime(data.GetUpdatedAt()),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetApplicantID has no value then it's a POST request (Create)
	if len(data.GetApplicantId()) != 0 {
		model.ApplicantID = data.GetApplicantId()
	}

	if len(data.GetFormId()) != 0 {
		model.FormID = data.GetFormId()
	}

	// Handle pointers after this

	return model
}

// ApplFormS2PTransformer transforms ApplForm Struct to Protobuf
func ApplFormS2PTransformer(data ApplForm) *pbx.ApplForm {
	model := &pbx.ApplForm{
		ApplicantId: data.ApplicantID,
		FormId:      data.FormID,
		CreatedAt:   util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:   util.GoTimeToGrpcTime(data.UpdatedAt),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
