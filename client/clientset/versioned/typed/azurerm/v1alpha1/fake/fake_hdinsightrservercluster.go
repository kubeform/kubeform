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

// FakeHdinsightRserverClusters implements HdinsightRserverClusterInterface
type FakeHdinsightRserverClusters struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var hdinsightrserverclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "hdinsightrserverclusters"}

var hdinsightrserverclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "HdinsightRserverCluster"}

// Get takes name of the hdinsightRserverCluster, and returns the corresponding hdinsightRserverCluster object, and an error if there is any.
func (c *FakeHdinsightRserverClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hdinsightrserverclustersResource, c.ns, name), &v1alpha1.HdinsightRserverCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightRserverCluster), err
}

// List takes label and field selectors, and returns the list of HdinsightRserverClusters that match those selectors.
func (c *FakeHdinsightRserverClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HdinsightRserverClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hdinsightrserverclustersResource, hdinsightrserverclustersKind, c.ns, opts), &v1alpha1.HdinsightRserverClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HdinsightRserverClusterList{ListMeta: obj.(*v1alpha1.HdinsightRserverClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.HdinsightRserverClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hdinsightRserverClusters.
func (c *FakeHdinsightRserverClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hdinsightrserverclustersResource, c.ns, opts))

}

// Create takes the representation of a hdinsightRserverCluster and creates it.  Returns the server's representation of the hdinsightRserverCluster, and an error, if there is any.
func (c *FakeHdinsightRserverClusters) Create(ctx context.Context, hdinsightRserverCluster *v1alpha1.HdinsightRserverCluster, opts v1.CreateOptions) (result *v1alpha1.HdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hdinsightrserverclustersResource, c.ns, hdinsightRserverCluster), &v1alpha1.HdinsightRserverCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightRserverCluster), err
}

// Update takes the representation of a hdinsightRserverCluster and updates it. Returns the server's representation of the hdinsightRserverCluster, and an error, if there is any.
func (c *FakeHdinsightRserverClusters) Update(ctx context.Context, hdinsightRserverCluster *v1alpha1.HdinsightRserverCluster, opts v1.UpdateOptions) (result *v1alpha1.HdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hdinsightrserverclustersResource, c.ns, hdinsightRserverCluster), &v1alpha1.HdinsightRserverCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightRserverCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHdinsightRserverClusters) UpdateStatus(ctx context.Context, hdinsightRserverCluster *v1alpha1.HdinsightRserverCluster, opts v1.UpdateOptions) (*v1alpha1.HdinsightRserverCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hdinsightrserverclustersResource, "status", c.ns, hdinsightRserverCluster), &v1alpha1.HdinsightRserverCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightRserverCluster), err
}

// Delete takes name of the hdinsightRserverCluster and deletes it. Returns an error if one occurs.
func (c *FakeHdinsightRserverClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hdinsightrserverclustersResource, c.ns, name), &v1alpha1.HdinsightRserverCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHdinsightRserverClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hdinsightrserverclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HdinsightRserverClusterList{})
	return err
}

// Patch applies the patch and returns the patched hdinsightRserverCluster.
func (c *FakeHdinsightRserverClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HdinsightRserverCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hdinsightrserverclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.HdinsightRserverCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightRserverCluster), err
}
