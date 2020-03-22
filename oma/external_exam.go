package oma

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

var validate *validator.Validate

// ExternalExam database model
type ExternalExam struct {
	ID            string         `json:"id" gorm:"type:UUID;primary_key;size:36"`
	ApplicantID   string         `json:"applicantId" gorm:"type:UUID;" validate:"required"`
	WaecExam      postgres.Jsonb `json:"waecExams" gorm:"type:jsonb;default:'{}'"`
	NonWaecExam   postgres.Jsonb `json:"nonWaecExams" gorm:"type:jsonb;default:'{}'"`
	Verified      bool           `json:"verified" gorm:"default:false"`
	VerifiedBy    string         `json:"verifiedBy" gorm:"type:varchar(100);not null;default:''''''"`
	CreatedAt     *time.Time     `json:"createdAt"`
	UpdatedAt     *time.Time     `json:"updatedAt"`
	InstitutionID string         `json:"-" gorm:"type:UUID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// WaecExam for validation
type WaecExam struct {
	Level          string           `json:"level" validate:"required"`
	IndexNo        string           `json:"indexNo" validate:"required"`
	MonthYear      int32            `json:"monthYear" validate:"required"`
	Sitting        int32            `json:"sitting" validate:"required"`
	AwaitingResult bool             `json:"awaitingResult"`
	Results        []WaecExamResult `json:"results" validate:"required,dive,required"`
}

// WaecExamResult for validation
type WaecExamResult struct {
	SubjectCode int32  `json:"subjectCode" validate:"required"`
	Grade       string `json:"grade" validate:"required"`
}

// NonWaecExam for validation
type NonWaecExam struct {
	QualificationType      string     `json:"qualType" validate:"required"`
	QualificationTypeOther string     `json:"qualTypeOther"`
	Qualification          string     `json:"qualification" validate:"required"`
	IndexNo                string     `json:"indexNo"`
	GPA                    string     `json:"gpa"`
	Institution            string     `json:"institution" validate:"required"`
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
	validate = validator.New()
	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.Struct(o); err != nil {
		verrs := err.(validator.ValidationErrors)
		fmt.Println(verrs)
		return false, fmt.Errorf("%s is %s", verrs[0].Field(), verrs[0].Tag())
	}

	// Validate WaecExam JSON by unmarshalling into struct to see if it's valid
	var wes []WaecExam
	if err := json.Unmarshal([]byte(string(o.WaecExam.RawMessage)), &wes); err != nil {
		fmt.Printf("Json Unmarshal Error %v", err)
		return false, fmt.Errorf("Invalid Waec Exam JSON property")
	}
	for _, we := range wes {
		if err := validate.Struct(we); err != nil {
			verrs := err.(validator.ValidationErrors)
			fmt.Println(verrs)
			return false, fmt.Errorf("%s is %s", verrs[0].Namespace(), verrs[0].Tag())
		}
	}

	var nwes []NonWaecExam
	if err := json.Unmarshal([]byte(string(o.NonWaecExam.RawMessage)), &nwes); err != nil {
		fmt.Printf("Json Unmarshal Error %v", err)
		return false, fmt.Errorf("Invalid Non Waec Exam JSON property")
	}

	for _, nwe := range nwes {
		if err := validate.Struct(nwe); err != nil {
			verrs := err.(validator.ValidationErrors)
			fmt.Println(verrs)
			return false, fmt.Errorf("%s is %s", verrs[0].Namespace(), verrs[0].Tag())
		}
	}

	if len(wes) == 0 && len(nwes) == 0 {
		return false, fmt.Errorf("Enter your WAEC or Non-WAEC Exams")
	}

	return true, nil
}
