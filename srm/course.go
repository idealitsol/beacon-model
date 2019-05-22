package srm

// Course model
type Course struct {
	ID          string  `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Code        string  `json:"code" gorm:"type:varchar(10);not null;unique_index"`
	Title       string  `json:"title" gorm:"type:varchar(555); not null;index"`
	Description string  `json:"description" gorm:"type:varchar(255)"`
	Credits     float32 `json:"credits" gorm:"type:decimal; default:0.00"`
	Owner       int32   `json:"owner" gorm:"size:3;not null"`                                          // this is deptID or unitId in legacy system
	Type        string  `json:"type" gorm:"type:varchar(15);default:'Regular'"`                       // Regular or General Course
	SchemeID    string  `json:"schemeId" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"` // Default Grading Scheme for the Course
	PreCourse   string  `json:"preCourse" gorm:"type:varchar(10)"`                                         // If a course is a prerequisite then we set it here (This is pcrel_preq)
	Status      bool    `json:"status" gorm:"type:bool;default:false"`
	Scheme      *Scheme `json:"scheme" gorm:"ForeignKey:ID;AssociationForeignKey:SchemeID"`
}

// Courses array
type Courses []Course
