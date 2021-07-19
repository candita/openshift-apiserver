package admission

import (
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
	routeinformers "github.com/openshift/client-go/route/informers/externalversions"
	"k8s.io/apiserver/pkg/admission"
)

type openshiftInformersInitializer struct {
	configInformers configinformers.SharedInformerFactory
	routeInformers  routeinformers.SharedInformerFactory
}

func NewOpenShiftInformersInitializer(
	configInformers configinformers.SharedInformerFactory,
	routeInformers routeinformers.SharedInformerFactory,
) *openshiftInformersInitializer {
	return &openshiftInformersInitializer{
		configInformers: configInformers,
		routeInformers:  routeInformers,
	}
}

func (i *openshiftInformersInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsOpenShiftConfigInformers); ok {
		wants.SetOpenShiftConfigInformers(i.configInformers)
	}
	if wants, ok := plugin.(WantsOpenShiftRouteInformers); ok {
		wants.SetOpenShiftRouteInformers(i.routeInformers)
	}
}
