package oma

import (
	"github.com/idealitsol/beacon-proto/pbx"
)

// Classification model
type Classification struct {
	ID            string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string `json:"name" gorm:"type:varchar(10);not null"`
	Ordering      int32  `json:"ordering" gorm:"not null"`
	Status        bool   `json:"status" gorm:"type:bool;default:false"`
	InstitutionID string `json:"-" gorm:"type:UUID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Classifications array
type Classifications []Classification

// ClassificationP2STransformer transforms classification Protobuf to Struct
func ClassificationP2STransformer(data *pbx.Classification) Classification {
	classification := Classification{
		Name:          data.GetName(),
		Status:        data.GetStatus(),
		InstitutionID: data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}
	return classification
}

// ClassificationS2PTransformer transforms classification Struct to Protobuf
func ClassificationS2PTransformer(data Classification) *pbx.Classification {
	classification := &pbx.Classification{
		Id:            data.ID,
		Name:          data.Name,
		Ordering:      data.Ordering,
		Status:        data.Status,
		InstitutionId: data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}
	return classification
}
