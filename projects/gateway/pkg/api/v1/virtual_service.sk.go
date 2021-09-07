// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewVirtualService(namespace, name string) *VirtualService {
	virtualservice := &VirtualService{}
	virtualservice.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return virtualservice
}

func (r *VirtualService) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

// Deprecated
func (r *VirtualService) SetStatus(status *core.Status) {
	r.SetStatusForNamespace(status)
}

// Deprecated
func (r *VirtualService) GetStatus() *core.Status {
	if r != nil {
		s, _ := r.GetStatusForNamespace()
		return s
	}
	return nil
}

func (r *VirtualService) SetNamespacedStatuses(statuses *core.NamespacedStatuses) {
	r.NamespacedStatuses = statuses
}

// SetStatusForNamespace inserts the specified status into the NamespacedStatuses.Statuses map for
// the current namespace (as specified by POD_NAMESPACE env var).  If the resource does not yet
// have a NamespacedStatuses, one will be created.
// Note: POD_NAMESPACE environment variable must be set for this function to behave as expected.
// If unset, a podNamespaceErr is returned.
func (r *VirtualService) SetStatusForNamespace(status *core.Status) error {
	return statusutils.SetStatusForPodNamespace(r, status)
}

// GetStatusForNamespace returns the status stored in the NamespacedStatuses.Statuses map for the
// controller specified by the POD_NAMESPACE env var, or nil if no status exists for that
// controller.
// Note: POD_NAMESPACE environment variable must be set for this function to behave as expected.
// If unset, a podNamespaceErr is returned.
func (r *VirtualService) GetStatusForNamespace() (*core.Status, error) {
	return statusutils.GetStatusForPodNamespace(r)
}

func (r *VirtualService) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *VirtualService) GroupVersionKind() schema.GroupVersionKind {
	return VirtualServiceGVK
}

type VirtualServiceList []*VirtualService

func (list VirtualServiceList) Find(namespace, name string) (*VirtualService, error) {
	for _, virtualService := range list {
		if virtualService.GetMetadata().Name == name && virtualService.GetMetadata().Namespace == namespace {
			return virtualService, nil
		}
	}
	return nil, errors.Errorf("list did not find virtualService %v.%v", namespace, name)
}

func (list VirtualServiceList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, virtualService := range list {
		ress = append(ress, virtualService)
	}
	return ress
}

func (list VirtualServiceList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, virtualService := range list {
		ress = append(ress, virtualService)
	}
	return ress
}

func (list VirtualServiceList) Names() []string {
	var names []string
	for _, virtualService := range list {
		names = append(names, virtualService.GetMetadata().Name)
	}
	return names
}

func (list VirtualServiceList) NamespacesDotNames() []string {
	var names []string
	for _, virtualService := range list {
		names = append(names, virtualService.GetMetadata().Namespace+"."+virtualService.GetMetadata().Name)
	}
	return names
}

func (list VirtualServiceList) Sort() VirtualServiceList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list VirtualServiceList) Clone() VirtualServiceList {
	var virtualServiceList VirtualServiceList
	for _, virtualService := range list {
		virtualServiceList = append(virtualServiceList, resources.Clone(virtualService).(*VirtualService))
	}
	return virtualServiceList
}

func (list VirtualServiceList) Each(f func(element *VirtualService)) {
	for _, virtualService := range list {
		f(virtualService)
	}
}

func (list VirtualServiceList) EachResource(f func(element resources.Resource)) {
	for _, virtualService := range list {
		f(virtualService)
	}
}

func (list VirtualServiceList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *VirtualService) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for VirtualService

func (o *VirtualService) GetObjectKind() schema.ObjectKind {
	t := VirtualServiceCrd.TypeMeta()
	return &t
}

func (o *VirtualService) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*VirtualService)
}

func (o *VirtualService) DeepCopyInto(out *VirtualService) {
	clone := resources.Clone(o).(*VirtualService)
	*out = *clone
}

var (
	VirtualServiceCrd = crd.NewCrd(
		"virtualservices",
		VirtualServiceGVK.Group,
		VirtualServiceGVK.Version,
		VirtualServiceGVK.Kind,
		"vs",
		false,
		&VirtualService{})
)

var (
	VirtualServiceGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gateway.solo.io",
		Kind:    "VirtualService",
	}
)
