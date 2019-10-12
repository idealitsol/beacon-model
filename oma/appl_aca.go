package oma

import (
	"encoding/json"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// ApplAca database model
type ApplAca struct {
	ApplicantID    string         `json:"applicantId" gorm:"type:UUID;primary_key;size:36"`
	EducationLevel string         `json:"educationLevel" gorm:"type:varchar(20);not null;default:'None'"`
	LastSchool     string         `json:"lastSchool" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	SchStartDate   *time.Time     `json:"schStartDate"`
	SchEndDate     *time.Time     `json:"schEndDate"`
	StudyCampus    string         `json:"studyCampus" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	StudyCenter    string         `json:"studyCenter" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	ProgChoice1    string         `json:"progChoice1" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	ProgChoice2    string         `json:"progChoice2" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	ProgChoice3    string         `json:"progChoice3" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	ProgChoice4    string         `json:"progChoice4" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	ProgChoice5    string         `json:"progChoice5" gorm:"type:UUID;default:'00000000-0000-0000-0000-000000000000'"`
	FeePaying      bool           `json:"feePaying" gorm:"default:false"`
	GradResTopic   string         `json:"gradResTopic" gorm:"type:varchar(1000)"`
	Referee1       postgres.Jsonb `json:"referee1" gorm:"type:jsonb;not null;default:'{}'"`
	Referee2       postgres.Jsonb `json:"referee2" gorm:"type:jsonb;not null;default:'{}'"`
	Referee3       postgres.Jsonb `json:"referee3" gorm:"type:jsonb;not null;default:'{}'"`
	Referee4       postgres.Jsonb `json:"referee4" gorm:"type:jsonb;not null;default:'{}'"`
	IsComplete     bool           `json:"isComplete" gorm:"default:false"`
	InstitutionID  string         `json:"-" gorm:"type:UUID;"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplAcas is an array of ApplAca objects
type ApplAcas []ApplAca

// ApplAcaP2STransformer transforms ApplAca Protobuf to Struct
func ApplAcaP2STransformer(data *pbx.ApplAca) ApplAca {
	model := ApplAca{
		EducationLevel: data.GetEducationLevel(),
		LastSchool:     data.GetLastSchool(),
		SchStartDate:   util.GrpcTimeToGoTime(data.GetSchStartDate()),
		SchEndDate:     util.GrpcTimeToGoTime(data.GetSchEndDate()),
		StudyCampus:    data.GetStudyCampus(),
		StudyCenter:    data.GetStudyCenter(),
		ProgChoice1:    data.GetProgChoice_1(),
		ProgChoice2:    data.GetProgChoice_2(),
		ProgChoice3:    data.GetProgChoice_3(),
		ProgChoice4:    data.GetProgChoice_4(),
		ProgChoice5:    data.GetProgChoice_5(),
		FeePaying:      data.GetFeePaying(),
		GradResTopic:   data.GetGradResTopic(),
		Referee1:       postgres.Jsonb{json.RawMessage(data.GetReferee_1())},
		Referee2:       postgres.Jsonb{json.RawMessage(data.GetReferee_2())},
		Referee3:       postgres.Jsonb{json.RawMessage(data.GetReferee_3())},
		Referee4:       postgres.Jsonb{json.RawMessage(data.GetReferee_4())},
		IsComplete:     data.GetIsComplete(),
		InstitutionID:  data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetApplicantID has no value then it's a POST request (Create)
	if len(data.GetApplicantId()) != 0 {
		model.ApplicantID = data.GetApplicantId()
	}

	// Handle pointers after this

	return model
}

// ApplAcaS2PTransformer transforms ApplAca Struct to Protobuf
func ApplAcaS2PTransformer(data ApplAca) *pbx.ApplAca {
	model := &pbx.ApplAca{
		ApplicantId:    data.ApplicantID,
		EducationLevel: data.EducationLevel,
		LastSchool:     data.LastSchool,
		SchStartDate:   util.GoTimeToGrpcTime(data.SchStartDate),
		SchEndDate:     util.GoTimeToGrpcTime(data.SchEndDate),
		StudyCampus:    data.StudyCampus,
		StudyCenter:    data.StudyCenter,
		ProgChoice_1:   data.ProgChoice1,
		ProgChoice_2:   data.ProgChoice2,
		ProgChoice_3:   data.ProgChoice3,
		ProgChoice_4:   data.ProgChoice4,
		ProgChoice_5:   data.ProgChoice5,
		FeePaying:      data.FeePaying,
		GradResTopic:   data.GradResTopic,
		Referee_1:      string(data.Referee1.RawMessage),
		Referee_2:      string(data.Referee2.RawMessage),
		Referee_3:      string(data.Referee3.RawMessage),
		Referee_4:      string(data.Referee4.RawMessage),
		IsComplete:     data.IsComplete,
		InstitutionId:  data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
