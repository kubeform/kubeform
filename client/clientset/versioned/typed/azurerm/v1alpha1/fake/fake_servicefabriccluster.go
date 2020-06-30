/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceFabricClusters implements ServiceFabricClusterInterface
type FakeServiceFabricClusters struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var servicefabricclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicefabricclusters"}

var servicefabricclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServiceFabricCluster"}

// Get takes name of the serviceFabricCluster, and returns the corresponding serviceFabricCluster object, and an error if there is any.
func (c *FakeServiceFabricClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicefabricclustersResource, c.ns, name), &v1alpha1.ServiceFabricCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceFabricCluster), err
}

// List takes label and field selectors, and returns the list of ServiceFabricClusters that match those selectors.
func (c *FakeServiceFabricClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceFabricClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicefabricclustersResource, servicefabricclustersKind, c.ns, opts), &v1alpha1.ServiceFabricClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceFabricClusterList{ListMeta: obj.(*v1alpha1.ServiceFabricClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceFabricClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceFabricClusters.
func (c *FakeServiceFabricClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicefabricclustersResource, c.ns, opts))

}

// Create takes the representation of a serviceFabricCluster and creates it.  Returns the server's representation of the serviceFabricCluster, and an error, if there is any.
func (c *FakeServiceFabricClusters) Create(ctx context.Context, serviceFabricCluster *v1alpha1.ServiceFabricCluster, opts v1.CreateOptions) (result *v1alpha1.ServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicefabricclustersResource, c.ns, serviceFabricCluster), &v1alpha1.ServiceFabricCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceFabricCluster), err
}

// Update takes the representation of a serviceFabricCluster and updates it. Returns the server's representation of the serviceFabricCluster, and an error, if there is any.
func (c *FakeServiceFabricClusters) Update(ctx context.Context, serviceFabricCluster *v1alpha1.ServiceFabricCluster, opts v1.UpdateOptions) (result *v1alpha1.ServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicefabricclustersResource, c.ns, serviceFabricCluster), &v1alpha1.ServiceFabricCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceFabricCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceFabricClusters) UpdateStatus(ctx context.Context, serviceFabricCluster *v1alpha1.ServiceFabricCluster, opts v1.UpdateOptions) (*v1alpha1.ServiceFabricCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicefabricclustersResource, "status", c.ns, serviceFabricCluster), &v1alpha1.ServiceFabricCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceFabricCluster), err
}

// Delete takes name of the serviceFabricCluster and deletes it. Returns an error if one occurs.
func (c *FakeServiceFabricClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicefabricclustersResource, c.ns, name), &v1alpha1.ServiceFabricCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceFabricClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicefabricclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceFabricClusterList{})
	return err
}

// Patch applies the patch and returns the patched serviceFabricCluster.
func (c *FakeServiceFabricClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceFabricCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicefabricclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceFabricCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceFabricCluster), err
}
