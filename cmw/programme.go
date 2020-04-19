package cmw

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/lib/pq"
)

// Programme database model
type Programme struct {
	ID             string        `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	LongName       string        `json:"longName" gorm:"type:varchar(200);not null"`
	ShortName      string        `json:"shortName" gorm:"type:varchar(100)"`
	Duration       int32         `json:"duration" gorm:"default:0"`
	Levels         pq.Int64Array `json:"levels" gorm:""` // Lookup {100,200,300,400}
	Qualification  string        `json:"qualification" gorm:"not null;default:'None'"`
	ProgIDx        int32         `json:"progIdx" gorm:""`                      // Programme Identifier Index , prog_idx is a 3 digit to indentify a programme in an index number
	Classification pq.Int64Array `json:"classification" gorm:"default:'{-1}'"` // Classification Eg. Fulltime, Parttime, Sandwich, Sandwich, Evening
	ProgType       int32         `json:"progType" gorm:""`                     // 1=Certificate; 2=Diploma; 3=Bachelors; 4=Post Diploma; 5=Post Graduate Certificate; 6=Post Graduate Diploma; 7=Masters; 8=MPhil; 9=PhD
	StatusAdm      bool          `json:"statusAdm" gorm:"default:true"`        // Controls programme visibility at admissions
	StatusApp      bool          `json:"statusApp" gorm:"default:true"`        // Controls programme visibility at applications
	Status         bool          `json:"status" gorm:"default:true"`
	CreatedAt      *time.Time    `json:"createdAt"`
	UpdatedAt      *time.Time    `json:"updatedAt"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Programmes is an array of Programme objects
type Programmes []Programme

// ProgrammeP2STransformer transforms Programme Protobuf to Struct
func ProgrammeP2STransformer(data *pbx.Programme) Programme {
	model := Programme{
		LongName:       data.GetLongName(),
		ShortName:      data.GetShortName(),
		Duration:       data.GetDuration(),
		Levels:         data.GetLevels(),
		Qualification:  data.GetQualification(),
		ProgIDx:        data.GetProgIdx(),
		Classification: data.GetClassification(),
		ProgType:       data.GetProgType(),
		StatusAdm:      data.GetStatusAdm(),
		StatusApp:      data.GetStatusApp(),
		Status:         data.GetStatus(),
		CreatedAt:      util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt:      util.GrpcTimeToGoTime(data.GetUpdatedAt()),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// ProgrammeS2PTransformer transforms Programme Struct to Protobuf
func ProgrammeS2PTransformer(data Programme) *pbx.Programme {
	model := &pbx.Programme{
		Id:             data.ID,
		LongName:       data.LongName,
		ShortName:      data.ShortName,
		Duration:       data.Duration,
		Levels:         data.Levels,
		Qualification:  data.Qualification,
		ProgIdx:        data.ProgIDx,
		Classification: data.Classification,
		ProgType:       data.ProgType,
		StatusAdm:      data.StatusAdm,
		StatusApp:      data.StatusApp,
		Status:         data.Status,
		CreatedAt:      util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt:      util.GoTimeToGrpcTime(data.UpdatedAt),

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
