package srm

import (
	"github.com/idealitsol/beacon-proto/pbx"
)

// Course model
type Course struct {
	ID          string  `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Code        string  `json:"code" gorm:"type:varchar(10);not null;unique_index"`
	Title       string  `json:"title" gorm:"type:varchar(555); not null;index"`
	Description string  `json:"description" gorm:"type:varchar(255)"`
	Credits     float32 `json:"credits" gorm:"type:decimal; default:0.00"`
	Owner       int32   `json:"owner" gorm:"size:3;not null"`                                             // this is deptID or unitId in legacy system
	Type        string  `json:"type" gorm:"type:varchar(15);default:'Regular'"`                           // Regular or General Course
	SchemeID    string  `json:"schemeId" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"` // Default Grading Scheme for the Course
	PreCourse   string  `json:"preCourse" gorm:"type:varchar(10)"`                                        // If a course is a prerequisite then we set it here (This is pcrel_preq)
	Status      bool    `json:"status" gorm:"type:bool;default:false"`
	Institution string  `json:"-" gorm:"type:UUID"`

	Scheme *Scheme `json:"scheme,omitempty" gorm:"ForeignKey:ID;AssociationForeignKey:SchemeID"`
}

// Courses array
type Courses []Course

// CourseP2STransformer transforms course Protobuf to Struct
func CourseP2STransformer(data *pbx.Course) Course {
	course := Course{
		Code:        data.GetCode(),
		Title:       data.GetTitle(),
		Description: data.GetDescription(),
		Credits:     data.GetCredits(),
		Type:        data.GetType(),
		Owner:       data.GetOwner(),
		SchemeID:    data.GetSchemeId(),
		PreCourse:   data.GetPreCourse(),
		Status:      data.GetStatus(),
		Institution: data.GetInstitution(),
	}

	// If GetId has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		course.ID = data.GetId()
	}

	// includes scheme
	if data.GetScheme() != nil {
		course.Scheme = &Scheme{
			ID:     data.GetScheme().Id,
			Scheme: data.GetScheme().Scheme,
			Status: data.GetScheme().Status,
		}
	}

	return course
}

// CourseS2PTransformer transforms course Struct to Protobuf
func CourseS2PTransformer(data Course) *pbx.Course {
	course := &pbx.Course{
		Id:          data.ID,
		Code:        data.Code,
		Title:       data.Title,
		Description: data.Description,
		Credits:     data.Credits,
		Owner:       data.Owner,
		Type:        data.Type,
		SchemeId:    data.SchemeID,
		PreCourse:   data.PreCourse,
		Status:      data.Status,
	}

	// includes scheme
	if data.Scheme != nil {
		course.Scheme = &pbx.Scheme{
			Id:     data.Scheme.ID,
			Scheme: data.Scheme.Scheme,
			Status: data.Scheme.Status,
		}
	}

	return course
}
