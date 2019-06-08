package srm

import "github.com/idealitsol/beacon-proto/pbx"

// Scheme model
type Scheme struct {
	ID            string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Scheme        string `json:"scheme" gorm:"type:varchar(255);not null;unique_index"`
	Status        bool   `json:"status" gorm:"type:bool;default:false"`
	InstitutionID string `json:"-" gorm:"type:UUID"`
	// Scheme *Scheme `gorm:"ForeignKey:SchemeID;AssociationForeignKey:ID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Schemes array
type Schemes []Scheme

// StructTransformer transforms Scheme Protobuf to Struct
func (o *Scheme) StructTransformer(data *pbx.Scheme) Scheme {
	o.Scheme = data.GetScheme()
	o.Status = data.GetStatus()
	o.InstitutionID = data.GetInstitutionId()

	o.BXXUpdatedFields = data.GetBXX_UpdatedFields()
	return *o
}

// ProtoTransformer transforms Scheme Struct to Protobuf
func (o *Scheme) ProtoTransformer(data Scheme) *pbx.Scheme {
	return &pbx.Scheme{
		Id:            o.ID,
		Scheme:        o.Scheme,
		Status:        o.Status,
		InstitutionId: o.InstitutionID,

		BXX_UpdatedFields: o.BXXUpdatedFields,
	}
}
