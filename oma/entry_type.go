package oma

import (
	"github.com/idealitsol/beacon-proto/pbx"
)

// EntryType model
type EntryType struct {
	ID            string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string `json:"name" gorm:"type:varchar(10);not null"`
	Description   string `json:"description" gorm:"type:varchar(100);not null"`
	Ordering      int32  `json:"ordering" gorm:"not null"`
	Status        bool   `json:"status" gorm:"type:bool;default:false"`
	InstitutionID string `json:"-" gorm:"type:UUID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// EntryTypes array
type EntryTypes []EntryType

// EntryTypeP2STransformer transforms EntryType Protobuf to Struct
func EntryTypeP2STransformer(data *pbx.EntryType) EntryType {
	EntryType := EntryType{
		Name:          data.GetName(),
		Description:   data.GetDescription(),
		Status:        data.GetStatus(),
		InstitutionID: data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}
	return EntryType
}

// EntryTypeS2PTransformer transforms EntryType Struct to Protobuf
func EntryTypeS2PTransformer(data EntryType) *pbx.EntryType {
	EntryType := &pbx.EntryType{
		Id:            data.ID,
		Name:          data.Name,
		Description:   data.Description,
		Ordering:      data.Ordering,
		Status:        data.Status,
		InstitutionId: data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}
	return EntryType
}
