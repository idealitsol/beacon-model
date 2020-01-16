package oma

import "github.com/idealitsol/beacon-proto/pbx"

// Configuration database model
type Configuration struct {
	Key           string `json:"key" gorm:"primary_key"`
	Value         string `json:"value" gorm:""`
	InstitutionID string `json:"institutionId" gorm:"type:UUID;"`
	Visibility    string `json:"visibility" gorm:"type:varchar(10);not null;default:'private'"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Configurations is an array of Configuration objects
type Configurations []Configuration

// ConfigurationP2STransformer transforms Configuration Protobuf to Struct
func ConfigurationP2STransformer(data *pbx.Configuration) Configuration {
	model := Configuration{
		Value:         data.GetValue(),
		InstitutionID: data.GetInstitutionId(),
		Visibility:    data.GetVisibility(),

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
		Visibility:    data.Visibility,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}

// ConfigurationsS2PTransformer transforms Configurations Struct which is array to Protobuf
func ConfigurationsS2PTransformer(datas Configurations) []*pbx.Configuration {
	models := []*pbx.Configuration{}

	for _, data := range datas {
		model := &pbx.Configuration{
			Key:           data.Key,
			Value:         data.Value,
			Visibility:    data.Visibility,
			InstitutionId: data.InstitutionID,

			BXX_UpdatedFields: data.BXXUpdatedFields,
		}
		models = append(models, model)
	}

	// Handle pointers after this

	return models
}

// ConfigurationsP2STransformer transforms Configurations which is array Protobuf to Struct
func ConfigurationsP2STransformer(datas []*pbx.Configuration) Configurations {
	models := Configurations{}
	for _, data := range datas {
		model := Configuration{
			Value:         data.GetValue(),
			InstitutionID: data.GetInstitutionId(),
			Visibility:    data.GetVisibility(),

			BXXUpdatedFields: data.GetBXX_UpdatedFields(),
		}

		// If GetKey has no value then it's a POST request (Create)
		if len(data.GetKey()) != 0 {
			model.Key = data.GetKey()
		}

		models = append(models, model)
	}

	// Handle pointers after this

	return models
}
