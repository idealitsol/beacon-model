package fss

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
)

// File database model
type File struct {
	ID        string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	FileName  string     `json:"fileName" gorm:"type:varchar(500);not null"`
	FileMime  string     `json:"fileMime" gorm:"type:varchar(15)"`
	FileSize  string     `json:"fileSize" gorm:"type:varchar(20)"`
	Status    bool       `json:"status" gorm:"default:false"`
	Container string     `json:"container" gorm:"type:UUID;"`
	Service   string     `json:"service" gorm:"type:varchar(5);not null"`
	Module    string     `json:"module" gorm:"type:varchar(30);not null"`
	URI       string     `json:"uri" gorm:"not null"`
	CreatedBy string     `json:"createdBy" gorm:"type:varchar(255)"`
	UpdatedBy string     `json:"updatedBy" gorm:"type:varchar(255)"`
	DeletedBy string     `json:"deletedBy" gorm:"type:varchar(255)"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Tags      string     `json:"tags" gorm:""`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Files is an array of File objects
type Files []File

// FileP2STransformer transforms File Protobuf to Struct
func FileP2STransformer(data *pbx.File) File {
	model := File{
		FileName:  data.GetFileName(),
		FileMime:  data.GetFileMime(),
		FileSize:  data.GetFileSize(),
		Status:    data.GetStatus(),
		Container: data.GetContainer(),
		Service:   data.GetService(),
		Module:    data.GetModule(),
		URI:       data.GetUri(),
		CreatedBy: data.GetCreatedBy(),
		UpdatedBy: data.GetUpdatedBy(),
		DeletedBy: data.GetDeletedBy(),
		CreatedAt: util.GrpcTimeToGoTime(data.GetCreatedAt()),
		UpdatedAt: util.GrpcTimeToGoTime(data.GetUpdatedAt()),
		DeletedAt: util.GrpcTimeToGoTime(data.GetDeletedAt()),
		Tags:      data.GetTags(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// FileS2PTransformer transforms File Struct to Protobuf
func FileS2PTransformer(data File) *pbx.File {
	model := &pbx.File{
		Id:        data.ID,
		FileName:  data.FileName,
		FileMime:  data.FileMime,
		FileSize:  data.FileSize,
		Status:    data.Status,
		Container: data.Container,
		Service:   data.Service,
		Module:    data.Module,
		Uri:       data.URI,
		CreatedBy: data.CreatedBy,
		UpdatedBy: data.UpdatedBy,
		DeletedBy: data.DeletedBy,
		CreatedAt: util.GoTimeToGrpcTime(data.CreatedAt),
		UpdatedAt: util.GoTimeToGrpcTime(data.UpdatedAt),
		DeletedAt: util.GoTimeToGrpcTime(data.DeletedAt),
		Tags:      data.Tags,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
