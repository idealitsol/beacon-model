package oma

import (
	"encoding/json"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// ApplAca database model
type ApplAca struct {
	ApplicantID    string         `json:"applicantId" gorm:"type:UUID;primary_key;size:36"`
	EducationLevel string         `json:"educationLevel" gorm:"type:varchar(20);not null;default:'None'"`
	LastSchool     string         `json:"lastSchool" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	SchStartDate   *time.Time     `json:"schStartDate"`
	SchEndDate     *time.Time     `json:"schEndDate"`
	Referee        postgres.Jsonb `json:"referee" gorm:"type:jsonb;not null;default:'[]'"`
	IsComplete     bool           `json:"isComplete" gorm:"default:false"`
	InstitutionID  string         `json:"institutionId" gorm:"type:UUID;"`
	Attachment     postgres.Jsonb `json:"attachment" gorm:"type:jsonb;not null;default:'{}'"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplAcas is an array of ApplAca objects
type ApplAcas []ApplAca

// BeforeCreate hook   http://gorm.io/docs/hooks.html
func (o *ApplAca) BeforeCreate(scope *gorm.Scope) error {
	// if valid, err := o.validate(); !valid {
	// 	return err
	// }

	if len(o.LastSchool) == 0 {
		o.LastSchool = util.DefaultUUID
	}

	return nil
}

// ApplAcaP2STransformer transforms ApplAca Protobuf to Struct
func ApplAcaP2STransformer(data *pbx.ApplAca) ApplAca {
	model := ApplAca{
		EducationLevel: data.GetEducationLevel(),
		LastSchool:     data.GetLastSchool(),
		SchStartDate:   util.GrpcTimeToGoTime(data.GetSchStartDate()),
		SchEndDate:     util.GrpcTimeToGoTime(data.GetSchEndDate()),
		Referee:        postgres.Jsonb{json.RawMessage(data.GetReferee())},
		IsComplete:     data.GetIsComplete(),
		InstitutionID:  data.GetInstitutionId(),
		Attachment:     postgres.Jsonb{json.RawMessage(data.GetAttachment())},

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetApplicantID has no value then it's a POST request (Create)
	if len(data.GetApplicantId()) != 0 {
		model.ApplicantID = data.GetApplicantId()
	}

	// Handle pointers after this

	return model
}

// ApplAcaS2PTransformer transforms ApplAca Struct to Protobuf
func ApplAcaS2PTransformer(data ApplAca) *pbx.ApplAca {
	model := &pbx.ApplAca{
		ApplicantId:    data.ApplicantID,
		EducationLevel: data.EducationLevel,
		LastSchool:     data.LastSchool,
		SchStartDate:   util.GoTimeToGrpcTime(data.SchStartDate),
		SchEndDate:     util.GoTimeToGrpcTime(data.SchEndDate),
		Referee:        string(data.Referee.RawMessage),
		IsComplete:     data.IsComplete,
		InstitutionId:  data.InstitutionID,
		Attachment:     string(data.Attachment.RawMessage),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
