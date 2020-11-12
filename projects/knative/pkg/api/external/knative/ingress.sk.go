// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
	"log"
	"sort"

	github_com_solo_io_gloo_edge_projects_knative_api_external_knative "github.com/solo-io/gloo-edge/projects/knative/api/external/knative"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewIngress(namespace, name string) *Ingress {
	ingress := &Ingress{}
	ingress.Ingress.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return ingress
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneableIngress interface {
	resources.Resource
	Clone() *github_com_solo_io_gloo_edge_projects_knative_api_external_knative.Ingress
}

var _ CloneableIngress = &github_com_solo_io_gloo_edge_projects_knative_api_external_knative.Ingress{}

type Ingress struct {
	github_com_solo_io_gloo_edge_projects_knative_api_external_knative.Ingress
}

func (r *Ingress) Clone() resources.Resource {
	return &Ingress{Ingress: *r.Ingress.Clone()}
}

func (r *Ingress) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	clone := r.Ingress.Clone()
	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})
	err := binary.Write(hasher, binary.LittleEndian, hashutils.HashAll(clone))
	if err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (r *Ingress) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *Ingress) GroupVersionKind() schema.GroupVersionKind {
	return IngressGVK
}

type IngressList []*Ingress

func (list IngressList) Find(namespace, name string) (*Ingress, error) {
	for _, ingress := range list {
		if ingress.GetMetadata().Name == name && ingress.GetMetadata().Namespace == namespace {
			return ingress, nil
		}
	}
	return nil, errors.Errorf("list did not find ingress %v.%v", namespace, name)
}

func (list IngressList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, ingress := range list {
		ress = append(ress, ingress)
	}
	return ress
}

func (list IngressList) Names() []string {
	var names []string
	for _, ingress := range list {
		names = append(names, ingress.GetMetadata().Name)
	}
	return names
}

func (list IngressList) NamespacesDotNames() []string {
	var names []string
	for _, ingress := range list {
		names = append(names, ingress.GetMetadata().Namespace+"."+ingress.GetMetadata().Name)
	}
	return names
}

func (list IngressList) Sort() IngressList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list IngressList) Clone() IngressList {
	var ingressList IngressList
	for _, ingress := range list {
		ingressList = append(ingressList, resources.Clone(ingress).(*Ingress))
	}
	return ingressList
}

func (list IngressList) Each(f func(element *Ingress)) {
	for _, ingress := range list {
		f(ingress)
	}
}

func (list IngressList) EachResource(f func(element resources.Resource)) {
	for _, ingress := range list {
		f(ingress)
	}
}

func (list IngressList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Ingress) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var (
	IngressGVK = schema.GroupVersionKind{
		Version: "v1alpha1",
		Group:   "networking.internal.knative.dev",
		Kind:    "Ingress",
	}
)
