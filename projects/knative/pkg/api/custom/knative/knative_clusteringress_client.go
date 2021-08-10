package knative

import (
	"context"
	"sort"

	"github.com/solo-io/gloo/projects/knative/api/external/knative"
	v1alpha1 "github.com/solo-io/gloo/projects/knative/pkg/api/external/knative"
	knativev1alpha1 "knative.dev/networking/pkg/apis/networking/v1alpha1"
	knativeclient "knative.dev/networking/pkg/client/clientset/versioned"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type ResourceClient struct {
	knativeClient knativeclient.Interface
	cache         Cache
}

func NewResourceClient(knativeClient knativeclient.Interface, cache Cache) *ResourceClient {
	return &ResourceClient{
		knativeClient: knativeClient,
		cache:         cache,
	}
}

func FromKube(ci *knativev1alpha1.Ingress) *v1alpha1.Ingress {
	deepCopy := ci.DeepCopy()
	baseType := knative.Ingress(*deepCopy)
	resource := &v1alpha1.Ingress{
		Ingress: baseType,
	}

	return resource
}

func ToKube(resource resources.Resource) (*knativev1alpha1.Ingress, error) {
	ingressResource, ok := resource.(*v1alpha1.Ingress)
	if !ok {
		return nil, errors.Errorf("internal error: invalid resource %v passed to ingress client", resources.Kind(resource))
	}

	ingress := knativev1alpha1.Ingress(ingressResource.Ingress)

	return &ingress, nil
}

var _ clients.ResourceClient = &ResourceClient{}

func (rc *ResourceClient) Kind() string {
	return resources.Kind(&v1alpha1.Ingress{})
}

func (rc *ResourceClient) NewResource() resources.Resource {
	return resources.Clone(&v1alpha1.Ingress{})
}

func (rc *ResourceClient) Register() error {
	return nil
}

func (rc *ResourceClient) Read(namespace, name string, opts clients.ReadOpts) (resources.Resource, error) {
	panic("this client does not support read operations")
}

func (rc *ResourceClient) Write(resource resources.Resource, opts clients.WriteOpts) (resources.Resource, error) {
	panic("this client does not support write operations")
}

func (rc *ResourceClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	panic("this client does not support delete operations")
}

func (rc *ResourceClient) List(namespace string, opts clients.ListOpts) (resources.ResourceList, error) {
	opts = opts.WithDefaults()

	ingressObjList, err := rc.cache.IngressLister().Ingresses(namespace).List(labels.SelectorFromSet(opts.Selector))
	if err != nil {
		return nil, errors.Wrapf(err, "listing Ingresses")
	}
	var resourceList resources.ResourceList
	for _, IngressObj := range ingressObjList {
		resource := FromKube(IngressObj)

		if resource == nil {
			continue
		}
		resourceList = append(resourceList, resource)
	}

	sort.SliceStable(resourceList, func(i, j int) bool {
		return resourceList[i].GetMetadata().GetName() < resourceList[j].GetMetadata().GetName()
	})

	return resourceList, nil
}

func (rc *ResourceClient) Watch(namespace string, opts clients.WatchOpts) (<-chan resources.ResourceList, <-chan error, error) {
	opts = opts.WithDefaults()
	watch := rc.cache.Subscribe()

	resourcesChan := make(chan resources.ResourceList)
	errs := make(chan error)
	// prevent flooding the channel with duplicates
	var previous *resources.ResourceList
	updateResourceList := func() {
		list, err := rc.List(namespace, clients.ListOpts{
			Ctx:      opts.Ctx,
			Selector: opts.Selector,
		})
		if err != nil {
			errs <- err
			return
		}
		if previous != nil {
			if list.Equal(*previous) {
				return
			}
		}
		previous = &list
		resourcesChan <- list
	}

	go func() {
		defer rc.cache.Unsubscribe(watch)
		defer close(resourcesChan)
		defer close(errs)

		// watch should open up with an initial read
		updateResourceList()
		for {
			select {
			case _, ok := <-watch:
				if !ok {
					return
				}
				updateResourceList()
			case <-opts.Ctx.Done():
				return
			}
		}
	}()

	return resourcesChan, errs, nil
}

func (rc *ResourceClient) exist(ctx context.Context, namespace, name string) bool {
	_, err := rc.knativeClient.NetworkingV1alpha1().Ingresses(namespace).Get(ctx, name, metav1.GetOptions{})
	return err == nil
}
