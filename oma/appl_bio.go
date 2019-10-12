package oma

import (
	"encoding/json"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// ApplBio database model
type ApplBio struct {
	ApplicantID       string         `json:"applicantId" gorm:"type:UUID;primary_key;size:36"`
	Source            string         `json:"source" gorm:"type:varchar(6);not null;default:'MANUAL'"`
	Title             string         `json:"title" gorm:"type:varchar(5);not null"`
	Fname             string         `json:"fname" gorm:"type:varchar(50);not null"`
	Mname             string         `json:"mname" gorm:"type:varchar(50)"`
	Sname             string         `json:"sname" gorm:"type:varchar(50);not null"`
	Gender            string         `json:"gender" gorm:"type:varchar(1);not null;default:'U'"`
	Dob               *time.Time     `json:"dob"`
	Email             string         `json:"email" gorm:"type:varchar(80)"`
	Cellphone         string         `json:"cellphone" gorm:"type:varchar(13);not null"`
	HomeAddress       string         `json:"homeAddress" gorm:"type:varchar(200)"`
	HomeAddressRegion string         `json:"homeAddressRegion" gorm:"type:varchar(50);default:'None'"`
	PostAddress       string         `json:"postAddress" gorm:"type:varchar(200)"`
	PostAddressRegion string         `json:"postAddressRegion" gorm:"type:varchar(50);default:'None'"`
	Disability        string         `json:"disability" gorm:"type:varchar(2);not null;default:'NO'"`
	BirthPlace        string         `json:"birthPlace" gorm:"type:varchar(150)"`
	BirthRegion       string         `json:"birthRegion" gorm:"type:varchar(50);default:'None'"`
	HomeTown          string         `json:"homeTown" gorm:"type:varchar(150)"`
	HomeTownRegion    string         `json:"homeTownRegion" gorm:"type:varchar(50);default:'None'"`
	Religion          string         `json:"religion" gorm:"type:varchar(15)"`
	Denomination      string         `json:"denomination" gorm:"type:varchar(80)"`
	MaritalStatus     string         `json:"maritalStatus" gorm:"type:varchar(10);not null;default:'None'"`
	NoChildren        int32          `json:"noChildren" gorm:"default:0"`
	NationalID        string         `json:"nationalId" gorm:"type:varchar(20)"`
	Country           string         `json:"country" gorm:"type:varchar(20);not null"`
	IsComplete        bool           `json:"isComplete" gorm:"default:false"`
	EmergencyContact  postgres.Jsonb `json:"emergencyContact" gorm:"type:jsonb;default:'{}'"`
	InstitutionID     string         `json:"institutionId" gorm:"type:UUID"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// ApplBios is an array of ApplBio objects
type ApplBios []ApplBio

// BeforeCreate hook   http://gorm.io/docs/hooks.html
func (o *ApplBio) BeforeCreate(scope *gorm.Scope) error {
	// if valid, err := o.validate(); !valid {
	// 	return err
	// }

	return nil
}

// ApplBioP2STransformer transforms ApplBio Protobuf to Struct
func ApplBioP2STransformer(data *pbx.ApplBio) ApplBio {
	model := ApplBio{
		Source:            data.GetSource(),
		Title:             data.GetTitle(),
		Fname:             data.GetFname(),
		Mname:             data.GetMname(),
		Sname:             data.GetSname(),
		Gender:            data.GetGender(),
		Dob:               util.GrpcTimeToGoTime(data.GetDob()),
		Email:             data.GetEmail(),
		Cellphone:         data.GetCellphone(),
		HomeAddress:       data.GetHomeAddress(),
		HomeAddressRegion: data.GetHomeAddressRegion(),
		PostAddress:       data.GetPostAddress(),
		PostAddressRegion: data.GetPostAddressRegion(),
		Disability:        data.GetDisability(),
		BirthPlace:        data.GetBirthPlace(),
		BirthRegion:       data.GetBirthRegion(),
		HomeTown:          data.GetHomeTown(),
		HomeTownRegion:    data.GetHomeTownRegion(),
		Religion:          data.GetReligion(),
		Denomination:      data.GetDenomination(),
		MaritalStatus:     data.GetMaritalStatus(),
		NoChildren:        data.GetNoChildren(),
		NationalID:        data.GetNationalId(),
		Country:           data.GetCountry(),
		IsComplete:        data.GetIsComplete(),
		EmergencyContact:  postgres.Jsonb{json.RawMessage(data.GetEmergencyContact())},
		InstitutionID:     data.GetInstitutionId(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetApplicantID has no value then it's a POST request (Create)
	if len(data.GetApplicantId()) != 0 {
		model.ApplicantID = data.GetApplicantId()
	}

	// Handle pointers after this

	return model
}

// ApplBioS2PTransformer transforms ApplBio Struct to Protobuf
func ApplBioS2PTransformer(data ApplBio) *pbx.ApplBio {
	model := &pbx.ApplBio{
		ApplicantId:       data.ApplicantID,
		Source:            data.Source,
		Title:             data.Title,
		Fname:             data.Fname,
		Mname:             data.Mname,
		Sname:             data.Sname,
		Gender:            data.Gender,
		Dob:               util.GoTimeToGrpcTime(data.Dob),
		Email:             data.Email,
		Cellphone:         data.Cellphone,
		HomeAddress:       data.HomeAddress,
		HomeAddressRegion: data.HomeAddressRegion,
		PostAddress:       data.PostAddress,
		PostAddressRegion: data.PostAddressRegion,
		Disability:        data.Disability,
		BirthPlace:        data.BirthPlace,
		BirthRegion:       data.BirthRegion,
		HomeTown:          data.HomeTown,
		HomeTownRegion:    data.HomeTownRegion,
		Religion:          data.Religion,
		Denomination:      data.Denomination,
		MaritalStatus:     data.MaritalStatus,
		NoChildren:        data.NoChildren,
		NationalId:        data.NationalID,
		Country:           data.Country,
		IsComplete:        data.IsComplete,
		EmergencyContact:  string(data.EmergencyContact.RawMessage),
		InstitutionId:     data.InstitutionID,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
