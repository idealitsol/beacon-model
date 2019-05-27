package srm

// Scheme model
type Scheme struct {
	Id     string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Scheme string `json:"scheme" gorm:"type:varchar(255);not null;unique_index"`
	Status bool   `json:"status" gorm:"type:bool;default:false"`
	// Course *Course `gorm:"ForeignKey:SchemeID;AssociationForeignKey:ID"`
}

// Schemes array
type Schemes []Scheme
