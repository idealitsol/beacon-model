package oma

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
)

// Session database model
type Session struct {
	ID            string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name          string     `json:"name" gorm:"type:varchar(30);not null"`
	AcademicYear  string     `json:"academicYear" gorm:"type:UUID;"`
	StartDate     *time.Time `json:"startDate"`
	EndDate       *time.Time `json:"endDate"`
	PayType       string     `json:"payType" gorm:"type:varchar(15);not null;default:'free'"`
	InstitutionID string     `json:"institutionId" gorm:"type:UUID;"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Sessions is an array of Session objects
type Sessions []Session

// SessionP2STransformer transforms Session Protobuf to Struct
func SessionP2STransformer(data *pbx.Session) Session {
	model := Session{
		Name:          data.GetName(),
		AcademicYear:  data.GetAcademicYear(),
		StartDate:     util.GrpcTimeToGoTime(data.GetStartDate()),
		EndDate:       util.GrpcTimeToGoTime(data.GetEndDate()),
		PayType:       data.GetPayType(),
		InstitutionID: data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// SessionS2PTransformer transforms Session Struct to Protobuf
func SessionS2PTransformer(data Session) *pbx.Session {
	model := &pbx.Session{
		Id:            data.ID,
		Name:          data.Name,
		AcademicYear:  data.AcademicYear,
		StartDate:     util.GoTimeToGrpcTime(data.StartDate),
		EndDate:       util.GoTimeToGrpcTime(data.EndDate),
		PayType:       data.PayType,
		InstitutionId: data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
