package oma

import (
	"encoding/json"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// FormSetup database model
type FormSetup struct {
	ID             string         `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name           string         `json:"name" gorm:"not null"`
	AcademicYear   string         `json:"academicYear" gorm:"not null"`
	Zone           string         `json:"zone" gorm:"not null"`
	Classification string         `json:"classification" gorm:"not null"`
	FormMode       string         `json:"formMode" gorm:"not null"`
	OpenDate       *time.Time     `json:"openDate"`
	CloseDate      *time.Time     `json:"closeDate"`
	InstitutionID  string         `json:"institutionId" gorm:"type:UUID"`
	StepperID      string         `json:"stepperId" gorm:"type:UUID"`
	Tags           postgres.Jsonb `json:"tags" gorm:"type:jsonb"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// FormSetups is an array of FormSetup objects
type FormSetups []FormSetup

// FormSetupP2STransformer transforms FormSetup Protobuf to Struct
func FormSetupP2STransformer(data *pbx.FormSetup) FormSetup {
	model := FormSetup{
		Name:           data.GetName(),
		AcademicYear:   data.GetAcademicYear(),
		Zone:           data.GetZone(),
		Classification: data.GetClassification(),
		FormMode:       data.GetFormMode(),
		OpenDate:       util.GrpcTimeToGoTime(data.GetOpenDate()),
		CloseDate:      util.GrpcTimeToGoTime(data.GetCloseDate()),
		InstitutionID:  data.GetInstitutionId(),
		StepperID:      data.GetStepperId(),
		Tags:           postgres.Jsonb{json.RawMessage(data.GetTags())},

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// FormSetupS2PTransformer transforms FormSetup Struct to Protobuf
func FormSetupS2PTransformer(data FormSetup) *pbx.FormSetup {
	model := &pbx.FormSetup{
		Id:             data.ID,
		Name:           data.Name,
		AcademicYear:   data.AcademicYear,
		Zone:           data.Zone,
		Classification: data.Classification,
		FormMode:       data.FormMode,
		OpenDate:       util.GoTimeToGrpcTime(data.OpenDate),
		CloseDate:      util.GoTimeToGrpcTime(data.CloseDate),
		InstitutionId:  data.InstitutionID,
		StepperId:      data.StepperID,
		Tags:           string(data.Tags.RawMessage),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
