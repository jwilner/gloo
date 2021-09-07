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

func NewUpstream(namespace, name string) *Upstream {
	upstream := &Upstream{}
	upstream.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return upstream
}

func (r *Upstream) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

// Deprecated
func (r *Upstream) SetStatus(status *core.Status) {
	r.SetStatusForNamespace(status)
}

// Deprecated
func (r *Upstream) GetStatus() *core.Status {
	if r != nil {
		s, _ := r.GetStatusForNamespace()
		return s
	}
	return nil
}

func (r *Upstream) SetNamespacedStatuses(statuses *core.NamespacedStatuses) {
	r.NamespacedStatuses = statuses
}

// SetStatusForNamespace inserts the specified status into the NamespacedStatuses.Statuses map for
// the current namespace (as specified by POD_NAMESPACE env var).  If the resource does not yet
// have a NamespacedStatuses, one will be created.
// Note: POD_NAMESPACE environment variable must be set for this function to behave as expected.
// If unset, a podNamespaceErr is returned.
func (r *Upstream) SetStatusForNamespace(status *core.Status) error {
	return statusutils.SetStatusForPodNamespace(r, status)
}

// GetStatusForNamespace returns the status stored in the NamespacedStatuses.Statuses map for the
// controller specified by the POD_NAMESPACE env var, or nil if no status exists for that
// controller.
// Note: POD_NAMESPACE environment variable must be set for this function to behave as expected.
// If unset, a podNamespaceErr is returned.
func (r *Upstream) GetStatusForNamespace() (*core.Status, error) {
	return statusutils.GetStatusForPodNamespace(r)
}

func (r *Upstream) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *Upstream) GroupVersionKind() schema.GroupVersionKind {
	return UpstreamGVK
}

type UpstreamList []*Upstream

func (list UpstreamList) Find(namespace, name string) (*Upstream, error) {
	for _, upstream := range list {
		if upstream.GetMetadata().Name == name && upstream.GetMetadata().Namespace == namespace {
			return upstream, nil
		}
	}
	return nil, errors.Errorf("list did not find upstream %v.%v", namespace, name)
}

func (list UpstreamList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) Names() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.GetMetadata().Name)
	}
	return names
}

func (list UpstreamList) NamespacesDotNames() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.GetMetadata().Namespace+"."+upstream.GetMetadata().Name)
	}
	return names
}

func (list UpstreamList) Sort() UpstreamList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list UpstreamList) Clone() UpstreamList {
	var upstreamList UpstreamList
	for _, upstream := range list {
		upstreamList = append(upstreamList, resources.Clone(upstream).(*Upstream))
	}
	return upstreamList
}

func (list UpstreamList) Each(f func(element *Upstream)) {
	for _, upstream := range list {
		f(upstream)
	}
}

func (list UpstreamList) EachResource(f func(element resources.Resource)) {
	for _, upstream := range list {
		f(upstream)
	}
}

func (list UpstreamList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Upstream) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Upstream

func (o *Upstream) GetObjectKind() schema.ObjectKind {
	t := UpstreamCrd.TypeMeta()
	return &t
}

func (o *Upstream) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Upstream)
}

func (o *Upstream) DeepCopyInto(out *Upstream) {
	clone := resources.Clone(o).(*Upstream)
	*out = *clone
}

var (
	UpstreamCrd = crd.NewCrd(
		"upstreams",
		UpstreamGVK.Group,
		UpstreamGVK.Version,
		UpstreamGVK.Kind,
		"us",
		false,
		&Upstream{})
)

var (
	UpstreamGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gloo.solo.io",
		Kind:    "Upstream",
	}
)
