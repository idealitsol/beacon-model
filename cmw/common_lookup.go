package cmw

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
)

// CommonLookup database model
type CommonLookup struct {
	ID            string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	AttrGroup     string     `json:"attrGroup" gorm:"not null"`
	AttrValue     string     `json:"attrValue" gorm:"type:varchar(20)"`
	AttrName      string     `json:"attrName" gorm:"type:varchar(100)"`
	AttrOther     string     `json:"attrOther" gorm:"type:varchar(100)"`
	Ordering      int32      `json:"ordering" gorm:";default:0"`
	Status        bool       `json:"status" gorm:"default:true"`
	InstitutionID string     `json:"institutionId" gorm:"type:UUID;"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// CommonLookups is an array of CommonLookup objects
type CommonLookups []CommonLookup

// CommonLookupP2STransformer transforms CommonLookup Protobuf to Struct
func CommonLookupP2STransformer(data *pbx.CommonLookup) CommonLookup {
	model := CommonLookup{
		AttrGroup:     data.GetAttrGroup(),
		AttrValue:     data.GetAttrValue(),
		AttrName:      data.GetAttrName(),
		AttrOther:     data.GetAttrOther(),
		Ordering:      int32(data.GetOrdering()),
		Status:        data.GetStatus(),
		InstitutionID: data.GetInstitutionId(),
		CreatedAt:     util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:     util.GrpcTimeToGoTime(data.GetUpdatedAt()),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// CommonLookupS2PTransformer transforms CommonLookup Struct to Protobuf
func CommonLookupS2PTransformer(data CommonLookup) *pbx.CommonLookup {
	model := &pbx.CommonLookup{
		Id:            data.ID,
		AttrGroup:     data.AttrGroup,
		AttrValue:     data.AttrValue,
		AttrName:      data.AttrName,
		AttrOther:     data.AttrOther,
		Ordering:      int32(data.Ordering),
		Status:        data.Status,
		InstitutionId: data.InstitutionID,
		CreatedAt:     util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:     util.GoTimeToGrpcTime(data.UpdatedAt),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
