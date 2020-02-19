package fss

import (
	"fmt"
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
	"github.com/jinzhu/gorm"
)

// Model Constants
const (
	UniqueConstraintNameCreatedBy = "container_name_created_by_key"
)

// Container database model
type Container struct {
	ID        string     `json:"id" gorm:"type:UUID;primary_key;default:gen_random_uuid();size:36"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Size      string     `json:"size,omitempty" gorm:"type:varchar(15);not null"`
	Provider  string     `json:"-" gorm:"type:UUID;"`
	CreatedAt *time.Time `json:"createdAt"`
	CreatedBy string     `json:"createdBy" gorm:"type:UUID;"`

	BXXUpdatedFields []string `json:"-" gorm:"-"`
}

// Containers is an array of Container objects
type Containers []Container

// BeforeCreate hook   http://gorm.io/docs/hooks.html
func (o *Container) BeforeCreate(scope *gorm.Scope) error {
	if valid, err := o.validate(); !valid {
		return err
	}

	return nil
}

// ConstraintError handles all the database constrainst defined in a model
func (o *Container) ConstraintError(err error) error {
	if ok, err := util.IsConstraintError(err, fmt.Sprintf("Container name already exists"), UniqueConstraintNameCreatedBy); ok {
		return err
	}

	return nil
}

func (o *Container) validate() (bool, error) {
	if len(o.Name) == 0 {
		return false, fmt.Errorf("Container name is required")
	}
	return true, nil
}

// ContainerP2STransformer transforms Container Protobuf to Struct
func ContainerP2STransformer(data *pbx.Container) Container {
	model := Container{
		Name:      data.GetName(),
		Size:      data.GetSize(),
		Provider:  data.GetProvider(),
		CreatedAt: util.GrpcTimeToGoTime(data.GetCreatedAt()),
		CreatedBy: data.GetCreatedBy(),

		BXXUpdatedFields: data.GetBXX_UpdatedFields(),
	}

	// If GetID has no value then it's a POST request (Create)
	if len(data.GetId()) != 0 {
		model.ID = data.GetId()
	}

	// Handle pointers after this

	return model
}

// ContainerS2PTransformer transforms Container Struct to Protobuf
func ContainerS2PTransformer(data Container) *pbx.Container {
	model := &pbx.Container{
		Id:        data.ID,
		Name:      data.Name,
		Size:      data.Size,
		Provider:  data.Provider,
		CreatedAt: util.GoTimeToGrpcTime(data.CreatedAt),
		CreatedBy: data.CreatedBy,

		BXX_UpdatedFields: data.BXXUpdatedFields,
	}

	// Handle pointers after this

	return model
}
