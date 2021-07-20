package admission

import (
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
	routeinformers "github.com/openshift/client-go/route/informers/externalversions"
)

type WantsOpenShiftConfigInformers interface {
	SetOpenShiftConfigInformers(informers configinformers.SharedInformerFactory)
}

type WantsOpenShiftRouteInformers interface {
	SetOpenShiftRouteInformers(informers routeinformers.SharedInformerFactory)
}
