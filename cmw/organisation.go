package cmw

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
)

// Organisation database model
type Organisation struct {
	ID            string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string     `json:"name" gorm:"type:varchar(100);not null"`
	ShortName     string     `json:"shortName" gorm:"type:varchar(80)"`
	Parent        string     `json:"parent" gorm:"type:UUID;"`
	Type          string     `json:"type" gorm:"not null"` // Main, Campus, College, Faculty, Department, Division, Unit
	Category      []string   `json:"category" gorm:""`     // Institution, Academic, Administration
	InstitutionID string     `json:"institutionId" gorm:"type:UUID;"`
	Status        bool       `json:"status" gorm:"default:true"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	Ordering      int32      `json:"ordering" gorm:"default:'-1'"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Organisations is an array of Organisation objects
type Organisations []Organisation

// OrganisationP2STransformer transforms Organisation Protobuf to Struct
func OrganisationP2STransformer(data *pbx.Organisation) Organisation {
	model := Organisation{
		Name:          data.GetName(),
		ShortName:     data.GetShortName(),
		Parent:        data.GetParent(),
		Type:          data.GetType(),
		Category:      data.GetCategory(),
		InstitutionID: data.GetInstitutionId(),
		Status:        data.GetStatus(),
		CreatedAt:     util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:     util.GrpcTimeToGoTime(data.GetUpdatedAt()),
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

// OrganisationS2PTransformer transforms Organisation Struct to Protobuf
func OrganisationS2PTransformer(data Organisation) *pbx.Organisation {
	model := &pbx.Organisation{
		Id:            data.ID,
		Name:          data.Name,
		ShortName:     data.ShortName,
		Parent:        data.Parent,
		Type:          data.Type,
		Category:      data.Category,
		InstitutionId: data.InstitutionID,
		Status:        data.Status,
		CreatedAt:     util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:     util.GoTimeToGrpcTime(data.UpdatedAt),
		Ordering:      data.Ordering,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
