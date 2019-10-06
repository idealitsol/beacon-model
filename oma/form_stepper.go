package oma

import (
	"encoding/json"

	"github.com/idealitsol/beacon-proto/pbx"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// FormStepper database model
type FormStepper struct {
	ID            string         `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Label         string         `json:"label" gorm:"type:varchar(100)not null"`
	MachineName   string         `json:"machineName" gorm:"type:varchar(30)not null"`
	Type          string         `json:"type" gorm:"not null"`
	Linear        bool           `json:"linear" gorm:""`
	Items         postgres.Jsonb `json:"items" gorm:"type:jsonb"`
	InstitutionID string         `json:"institutionId" gorm:"type:UUID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// FormSteppers is an array of FormStepper objects
type FormSteppers []FormStepper

// FormStepperP2STransformer transforms FormStepper Protobuf to Struct
func FormStepperP2STransformer(data *pbx.FormStepper) FormStepper {
	model := FormStepper{
		Label:         data.GetLabel(),
		MachineName:   data.GetMachineName(),
		Type:          data.GetType(),
		Linear:        data.GetLinear(),
		Items:         postgres.Jsonb{json.RawMessage(data.GetItems())},
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

// FormStepperS2PTransformer transforms FormStepper Struct to Protobuf
func FormStepperS2PTransformer(data FormStepper) *pbx.FormStepper {
	model := &pbx.FormStepper{
		Id:            data.ID,
		Label:         data.Label,
		MachineName:   data.MachineName,
		Type:          data.Type,
		Linear:        data.Linear,
		Items:         string(data.Items.RawMessage),
		InstitutionId: data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
