package oma

import (
	"github.com/idealitsol/beacon-proto/pbx"
)

// EntryMode model
type EntryMode struct {
	ID            string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string `json:"name" gorm:"type:varchar(10);not null"`
	Ordering      int32  `json:"ordering" gorm:"not null"`
	Status        bool   `json:"status" gorm:"type:bool;default:false"`
	InstitutionID string `json:"-" gorm:"type:UUID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// EntryModes array
type EntryModes []EntryMode

// EntryModeP2STransformer transforms EntryMode Protobuf to Struct
func EntryModeP2STransformer(data *pbx.EntryMode) EntryMode {
	EntryMode := EntryMode{
		Name:          data.GetName(),
		Status:        data.GetStatus(),
		InstitutionID: data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}
	return EntryMode
}

// EntryModeS2PTransformer transforms EntryMode Struct to Protobuf
func EntryModeS2PTransformer(data EntryMode) *pbx.EntryMode {
	EntryMode := &pbx.EntryMode{
		Id:            data.ID,
		Name:          data.Name,
		Ordering:      data.Ordering,
		Status:        data.Status,
		InstitutionId: data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}
	return EntryMode
}
