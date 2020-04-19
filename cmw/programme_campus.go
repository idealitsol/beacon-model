package cmw

import "github.com/idealitsol/beacon-proto/pbx"

// ProgrammeCampus database model
type ProgrammeCampus struct {
	ID           string `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	ProgrammeID  string `json:"programmeId" gorm:"type:UUID;"`
	CampusID     string `json:"campusId" gorm:"type:UUID;"`
	FacultyID    string `json:"facultyId" gorm:"type:UUID;"`
	DepartmentID string `json:"departmentId" gorm:"type:UUID;"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ProgrammeCampuses is an array of ProgrammeCampus objects
type ProgrammeCampuses []ProgrammeCampus

// ProgrammeCampusP2STransformer transforms ProgrammeCampus Protobuf to Struct
func ProgrammeCampusP2STransformer(data *pbx.ProgrammeCampus) ProgrammeCampus {
	model := ProgrammeCampus{
		ProgrammeID:  data.GetProgrammeId(),
		CampusID:     data.GetCampusId(),
		FacultyID:    data.GetFacultyId(),
		DepartmentID: data.GetDepartmentId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// ProgrammeCampusS2PTransformer transforms ProgrammeCampus Struct to Protobuf
func ProgrammeCampusS2PTransformer(data ProgrammeCampus) *pbx.ProgrammeCampus {
	model := &pbx.ProgrammeCampus{
		Id:           data.ID,
		ProgrammeId:  data.ProgrammeID,
		CampusId:     data.CampusID,
		FacultyId:    data.FacultyID,
		DepartmentId: data.DepartmentID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
