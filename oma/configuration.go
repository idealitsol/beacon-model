package oma

import "github.com/idealitsol/beacon-proto/pbx"

// Configuration database model
type Configuration struct {
	Key           string `json:"key" gorm:"not null"`
	Value         string `json:"value" gorm:""`
	InstitutionID string `json:"institutionId" gorm:"type:UUID;"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Configurations is an array of Configuration objects
type Configurations []Configuration

// ConfigurationP2STransformer transforms Configuration Protobuf to Struct
func ConfigurationP2STransformer(data *pbx.Configuration) Configuration {
	model := Configuration{
		Value:         data.GetValue(),
		InstitutionID: data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetKey has no value then it's a POST request (Create)
	if len(data.GetKey()) != 0 {
		model.Key = data.GetKey()
	}

	// Handle pointers after this

	return model
}

// ConfigurationS2PTransformer transforms Configuration Struct to Protobuf
func ConfigurationS2PTransformer(data Configuration) *pbx.Configuration {
	model := &pbx.Configuration{
		Key:           data.Key,
		Value:         data.Value,
		InstitutionId: data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
