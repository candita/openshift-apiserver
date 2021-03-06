package imagestreammapping

import (
	"context"

	imageapi "github.com/openshift/openshift-apiserver/pkg/image/apis/image"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Registry is an interface for things that know how to store ImageStreamMapping objects.
type Registry interface {
	// CreateImageStreamMapping creates a new image stream mapping.
	CreateImageStreamMapping(ctx context.Context, mapping *imageapi.ImageStreamMapping) (*metav1.Status, error)
}

// Storage is an interface for a standard REST Storage backend
type Storage interface {
	Create(ctx context.Context, obj runtime.Object) (runtime.Object, error)
}

// storage puts strong typing around storage calls
type storage struct {
	Storage
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
func NewRegistry(s Storage) Registry {
	return &storage{s}
}

// CreateImageStreamMapping will create an image stream mapping.
func (s *storage) CreateImageStreamMapping(ctx context.Context, mapping *imageapi.ImageStreamMapping) (*metav1.Status, error) {
	obj, err := s.Create(ctx, mapping)
	if err != nil {
		return nil, err
	}
	return obj.(*metav1.Status), nil
}
