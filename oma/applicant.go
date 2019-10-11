package oma

import (
	"github.com/idealitsol/beacon-proto/pbx"
)

// Applicant ..
type Applicant struct {
	ID            string   `json:"id"`
	AccountInfo   *ApplAcc `json:"accountInfo"`
	BioData       *ApplBio `json:"bioData"`
	AcademicInfo  *ApplAca `json:"academicInfo"`
	InstitutionID string   `json:"institutionId" gorm:"type:UUID"`
}

// Applicants ..
type Applicants []Applicant

// ApplicantP2STransformer transforms applicant Protobuf to Struct
func ApplicantP2STransformer(data *pbx.Applicant) Applicant {
	model := Applicant{}

	// If GetId has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	if data.GetAccountInfo() != nil {
		data.AccountInfo.Id = model.ID

		accountInfo := ApplAccP2STransformer(data.GetAccountInfo())
		model.AccountInfo = &accountInfo
	}

	if data.GetBioData() != nil {
		data.BioData.ApplicantId = model.ID

		bioData := ApplBioP2STransformer(data.GetBioData())
		model.BioData = &bioData
	}

	if data.GetAcademicInfo() != nil {
		data.AcademicInfo.ApplicantId = model.ID

		academicInfo := ApplAcaP2STransformer(data.GetAcademicInfo())
		model.AcademicInfo = &academicInfo
	}

	return model
}

// ApplicantS2PTransformer transforms classification Struct to Protobuf
func ApplicantS2PTransformer(data Applicant) *pbx.Applicant {
	model := &pbx.Applicant{}

	model.Id = data.ID
	if data.AccountInfo != nil {
		model.AccountInfo = ApplAccS2PTransformer(*data.AccountInfo)
	}

	return model
}
