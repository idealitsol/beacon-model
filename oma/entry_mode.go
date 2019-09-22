package oma

import "github.com/idealitsol/beacon-proto/pbx"

// EntryMode database model
type EntryMode struct {
	ID            string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string `json:"name" gorm:"type:varchar(30)not null"`
	InstitutionID string `json:"institutionId" gorm:"type:UUID"`
	Status        bool   `json:"status" gorm:"default:false"`
	Ordering      int32  `json:"ordering" gorm:""`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// EntryModes is an array of EntryMode objects
type EntryModes []EntryMode

// EntryModeP2STransformer transforms EntryMode Protobuf to Struct
func EntryModeP2STransformer(data *pbx.EntryMode) EntryMode {
	model := EntryMode{
		Name:          data.GetName(),
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

// EntryModeS2PTransformer transforms EntryMode Struct to Protobuf
func EntryModeS2PTransformer(data EntryMode) *pbx.EntryMode {
	model := &pbx.EntryMode{
		Id:            data.ID,
		Name:          data.Name,
		InstitutionId: data.InstitutionID,
		Status:        data.Status,
		Ordering:      data.Ordering,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
