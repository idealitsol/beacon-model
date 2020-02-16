package fss

import (
	"time"

	"github.com/idealitsol/beacon-proto/pbx"
	util "github.com/idealitsol/beacon-util"
)

// Container model
type Container struct {
	Name         string     `json:"name"`
	Size         string     `json:"size"`
	LastModified *time.Time `json:"lastModified"`
}

// Containers is an array of Container
type Containers []Container

// ContainerP2STransformer transforms Container Protobuf to Struct
func ContainerP2STransformer(data *pbx.Container) Container {
	model := Container{
		Name:         data.GetName(),
		Size:         data.GetSize(),
		LastModified: util.GrpcTimeToGoTime(data.GetLastModified()),
	}

	// Handle pointers after this

	return model
}

// ContainerS2PTransformer transforms Container Struct to Protobuf
func ContainerS2PTransformer(data Container) *pbx.Container {
	model := &pbx.Container{
		Name:         data.Name,
		Size:         data.Size,
		LastModified: util.GoTimeToGrpcTime(data.LastModified),
	}

	// Handle pointers after this

	return model
}
