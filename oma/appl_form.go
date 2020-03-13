package oma

import (
	"encoding/json"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// ApplForm database model
type ApplForm struct {
	ApplicantID    string         `json:"applicantId" gorm:"type:UUID;primary_key;size:36"`
	FormID         string         `json:"formId" gorm:"type:UUID;"`
	InstitutionID  string         `json:"institutionId" gorm:"type:UUID;"`
	CreatedAt      *time.Time     `json:"createdAt"`
	UpdatedAt      *time.Time     `json:"updatedAt"`
	ProgChoice     postgres.Jsonb `json:"progChoice" gorm:"type:jsonb;default:'[]'"`
	GradResTopic   string         `json:"gradResTopic" gorm:"type:varchar(8000)"`
	StudyCampus    string         `json:"studyCampus" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	StudyCenter    string         `json:"studyCenter" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	FeePaying      bool           `json:"feePaying" gorm:"default:false"`
	EntryMode      string         `json:"entryMode" gorm:"type:UUID;"`
	EntryType      string         `json:"entryType" gorm:"type:UUID;"`
	Classification string         `json:"classification" gorm:"type:UUID;"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplForms is an array of ApplForm objects
type ApplForms []ApplForm

// ApplFormP2STransformer transforms ApplForm Protobuf to Struct
func ApplFormP2STransformer(data *pbx.ApplForm) ApplForm {
	model := ApplForm{
		CreatedAt:      util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:      util.GrpcTimeToGoTime(data.GetUpdatedAt()),
		ProgChoice:     postgres.Jsonb{json.RawMessage(data.GetProgChoice())},
		GradResTopic:   data.GetGradResTopic(),
		StudyCampus:    data.GetStudyCampus(),
		StudyCenter:    data.GetStudyCenter(),
		FeePaying:      data.GetFeePaying(),
		EntryMode:      data.GetEntryMode(),
		EntryType:      data.GetEntryType(),
		Classification: data.GetClassification(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetApplicantID has no value then it's a POST request (Create)
	if len(data.GetApplicantId()) != 0 {
		model.ApplicantID = data.GetApplicantId()
	}

	// If GetFormId has no value then it's a POST request (Create)
	if len(data.GetFormId()) != 0 {
		model.FormID = data.GetFormId()
	}

	// If GetInstitutionId has no value then it's a POST request (Create)
	if len(data.GetInstitutionId()) != 0 {
		model.InstitutionID = data.GetInstitutionId()
	}

	// Handle pointers after this

	return model
}

// ApplFormS2PTransformer transforms ApplForm Struct to Protobuf
func ApplFormS2PTransformer(data ApplForm) *pbx.ApplForm {
	model := &pbx.ApplForm{
		ApplicantId:    data.ApplicantID,
		FormId:         data.FormID,
		InstitutionId:  data.InstitutionID,
		CreatedAt:      util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:      util.GoTimeToGrpcTime(data.UpdatedAt),
		ProgChoice:     string(data.ProgChoice.RawMessage),
		GradResTopic:   data.GradResTopic,
		StudyCampus:    data.StudyCampus,
		StudyCenter:    data.StudyCenter,
		FeePaying:      data.FeePaying,
		EntryMode:      data.EntryMode,
		EntryType:      data.EntryType,
		Classification: data.Classification,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
