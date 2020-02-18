package fss

import "github.com/idealitsol/beacon-proto/pbx"

// Provider database model
type Provider struct {
	ID     string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name   string `json:"name" gorm:"type:varchar(50);not null"`
	Status bool   `json:"status" gorm:"default:true"`
	Abbr   string `json:"abbr" gorm:"type:varchar(10)"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Providers is an array of Provider objects
type Providers []Provider

// ProviderP2STransformer transforms Provider Protobuf to Struct
func ProviderP2STransformer(data *pbx.Provider) Provider {
	model := Provider{
		Name:   data.GetName(),
		Status: data.GetStatus(),
		Abbr:   data.GetAbbr(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// ProviderS2PTransformer transforms Provider Struct to Protobuf
func ProviderS2PTransformer(data Provider) *pbx.Provider {
	model := &pbx.Provider{
		Id:     data.ID,
		Name:   data.Name,
		Status: data.Status,
		Abbr:   data.Abbr,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
