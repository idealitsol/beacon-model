package oma

import (
	"encoding/json"

	"github.com/idealitsol/beacon-proto/pbx"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Form database model
type Form struct {
	ID            string         `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string         `json:"name" gorm:"not null"`
	MachineName   string         `json:"machineName" gorm:"default:'100'"`
	Description   string         `json:"description" gorm:"default:'255'"`
	Status        bool           `json:"status" gorm:""`
	Fields        postgres.Jsonb `json:"fields" gorm:"type:jsonb;default:'{}'"`
	Display       postgres.Jsonb `json:"display" gorm:"type:jsonb;default:'{}'"`
	IsSystem      bool           `json:"isSystem" gorm:""`
	InstitutionID string         `json:"institutionId" gorm:"type:UUID;"`
	Groups        postgres.Jsonb `json:"groups" gorm:"type:jsonb;default:'[]'"`
	Settings      postgres.Jsonb `json:"settings" gorm:"type:jsonb;default:'{}'"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Forms is an array of Form objects
type Forms []Form

// FormP2STransformer transforms Form Protobuf to Struct
func FormP2STransformer(data *pbx.Form) Form {
	model := Form{
		Name:          data.GetName(),
		MachineName:   data.GetMachineName(),
		Description:   data.GetDescription(),
		Status:        data.GetStatus(),
		Fields:        postgres.Jsonb{json.RawMessage(data.GetFields())},
		Display:       postgres.Jsonb{json.RawMessage(data.GetDisplay())},
		IsSystem:      data.GetIsSystem(),
		InstitutionID: data.GetInstitutionId(),
		Groups:        postgres.Jsonb{json.RawMessage(data.GetGroups())},
		Settings:      postgres.Jsonb{json.RawMessage(data.GetSettings())},

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// FormS2PTransformer transforms Form Struct to Protobuf
func FormS2PTransformer(data Form) *pbx.Form {
	model := &pbx.Form{
		Id:            data.ID,
		Name:          data.Name,
		MachineName:   data.MachineName,
		Description:   data.Description,
		Status:        data.Status,
		Fields:        string(data.Fields.RawMessage),
		Display:       string(data.Display.RawMessage),
		IsSystem:      data.IsSystem,
		InstitutionId: data.InstitutionID,
		Groups:        string(data.Groups.RawMessage),
		Settings:      string(data.Settings.RawMessage),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
