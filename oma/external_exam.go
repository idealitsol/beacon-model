package oma

import (
	"encoding/json"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// ExternalExam database model
type ExternalExam struct {
	ID          string         `json:"id" gorm:"type:UUID;primary_key;size:36"`
	ApplicantID string         `json:"applicantId" gorm:"type:UUID;"`
	WaecExam    postgres.Jsonb `json:"waecExam" gorm:"type:jsonb;default:'{}'"`
	NonWaecExam postgres.Jsonb `json:"nonWaecExam" gorm:"type:jsonb;default:'{}'"`
	Verified    bool           `json:"verified" gorm:"default:false"`
	VerifiedBy  string         `json:"verifiedBy" gorm:"type:varchar(100);not null;default:''''''"`
	CreatedAt   *time.Time     `json:"createdAt"`
	UpdatedAt   *time.Time     `json:"updatedAt"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

type WaecExam struct {
	Level          string           `json:"level"`
	IndexNo        string           `json:"indexNo"`
	MonthYear      int32            `json:"monthYear"`
	Sitting        int32            `json:"sitting"`
	AwaitingResult bool             `json:"awaitingResult"`
	Results        []WaecExamResult `json:"results"`
}

type WaecExamResult struct {
	SubjectCode int32  `json:"subjectCode"`
	Grade       string `json:"grade"`
}

type NonWaecExam struct {
	QualificationType      string     `json:"qualType"`
	QualificationTypeOther string     `json:"qualTypeOther"`
	Qualification          string     `json:"qualification"`
	IndexNo                string     `json:"indexNo"`
	GPA                    string     `json:"gpa"`
	Institution            string     `json:"institution"`
	ClassHonour            string     `json:"classHonour"`
	AttendFrom             *time.Time `json:"attendFrom"`
	AttendTo               *time.Time `json:"attendTo"`
}

// ExternalExams is an array of ExternalExam objects
type ExternalExams []ExternalExam

// BeforeCreate hook
func (o *ExternalExam) BeforeCreate(scope *gorm.Scope) error {
	if valid, err := o.validate(); !valid {
		return err
	}

	return nil
}

// ExternalExamP2STransformer transforms ExternalExam Protobuf to Struct
func ExternalExamP2STransformer(data *pbx.ExternalExam) ExternalExam {
	model := ExternalExam{
		ApplicantID: data.GetApplicantId(),
		WaecExam:    postgres.Jsonb{json.RawMessage(data.GetWaecExam())},
		NonWaecExam: postgres.Jsonb{json.RawMessage(data.GetNonWaecExam())},
		Verified:    data.GetVerified(),
		VerifiedBy:  data.GetVerifiedBy(),
		CreatedAt:   util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:   util.GrpcTimeToGoTime(data.GetUpdatedAt()),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// ExternalExamS2PTransformer transforms ExternalExam Struct to Protobuf
func ExternalExamS2PTransformer(data ExternalExam) *pbx.ExternalExam {
	model := &pbx.ExternalExam{
		Id:          data.ID,
		ApplicantId: data.ApplicantID,
		WaecExam:    string(data.WaecExam.RawMessage),
		NonWaecExam: string(data.NonWaecExam.RawMessage),
		Verified:    data.Verified,
		VerifiedBy:  data.VerifiedBy,
		CreatedAt:   util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:   util.GoTimeToGrpcTime(data.UpdatedAt),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}

func (o *ExternalExam) validate() (bool, error) {

	return true, nil
}
