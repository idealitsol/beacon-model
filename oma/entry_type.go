package oma

import "github.com/idealitsol/beacon-proto/pbx"

// EntryType database model
type EntryType struct {
	ID            string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string `json:"name" gorm:"type:varchar(30)not null"`
	Description   string `json:"description" gorm:"type:varchar(255)"`
	InstitutionID string `json:"institutionId" gorm:"type:UUID"`
	Status        bool   `json:"status" gorm:""`
	Ordering      int32  `json:"ordering" gorm:""`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// EntryTypes is an array of EntryType objects
type EntryTypes []EntryType

// EntryTypeP2STransformer transforms EntryType Protobuf to Struct
func EntryTypeP2STransformer(data *pbx.EntryType) EntryType {
	model := EntryType{
		Name:          data.GetName(),
		Description:   data.GetDescription(),
		InstitutionID: data.GetInstitutionId(),
		Status:        data.GetStatus(),
		Ordering:      data.GetOrdering(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// EntryTypeS2PTransformer transforms EntryType Struct to Protobuf
func EntryTypeS2PTransformer(data EntryType) *pbx.EntryType {
	model := &pbx.EntryType{
		Id:            data.ID,
		Name:          data.Name,
		Description:   data.Description,
		InstitutionId: data.InstitutionID,
		Status:        data.Status,
		Ordering:      data.Ordering,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
