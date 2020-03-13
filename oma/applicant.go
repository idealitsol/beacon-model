package oma

import (
	"github.com/idealitsol/beacon-proto/pbx"
)

// Applicant ..
type Applicant struct {
	ID            string    `json:"id"`
	FormID        string    `json:"formId"`
	MainData      *ApplMain `json:"mainData"`
	BioData       *ApplBio  `json:"bioData"`
	AcademicInfo  *ApplAca  `json:"academicInfo"`
	FormData      *ApplForm `json:"formData"`
	InstitutionID string    `json:"institutionId" gorm:"type:UUID"`
}

// ApplicantAuthRequest is used to accept authentication request
type ApplicantAuthRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	LoginType string `json:"loginType"`

	Domain   string `json:"-"`
	Platform string `json:"-"`

	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

// ApplicantAuthResponse is response give back to authentication request
type ApplicantAuthResponse struct {
	Mesg           string      `json:"message"`
	Token          string      `json:"token"`
	UserID         string      `json:"userId"`
	Roles          interface{} `json:"roles,omitempty"`
	ForcePWDChange bool        `json:"forcePwdChange"`
	Err            error       `json:"err,omitempty"`
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

	if data.GetMainData() != nil {
		data.MainData.Id = model.ID

		mainData := ApplMainP2STransformer(data.GetMainData())
		model.MainData = &mainData
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

	if data.GetFormData() != nil {
		data.FormData.ApplicantId = model.ID
		data.FormData.InstitutionId = model.InstitutionID

		formData := ApplFormP2STransformer(data.GetFormData())
		model.FormData = &formData
	}

	return model
}

// ApplicantS2PTransformer transforms classification Struct to Protobuf
func ApplicantS2PTransformer(data Applicant) *pbx.Applicant {
	model := &pbx.Applicant{}

	model.Id = data.ID
	model.FormId = data.FormID
	if data.MainData != nil {
		model.MainData = ApplMainS2PTransformer(*data.MainData)
	}

	if data.BioData != nil {
		model.BioData = ApplBioS2PTransformer(*data.BioData)
	}

	if data.AcademicInfo != nil {
		model.AcademicInfo = ApplAcaS2PTransformer(*data.AcademicInfo)
	}

	if data.FormData != nil {
		model.FormData = ApplFormS2PTransformer(*data.FormData)
	}

	return model
}
