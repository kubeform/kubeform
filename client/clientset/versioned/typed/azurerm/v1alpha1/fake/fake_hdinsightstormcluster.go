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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHdinsightStormClusters implements HdinsightStormClusterInterface
type FakeHdinsightStormClusters struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var hdinsightstormclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "hdinsightstormclusters"}

var hdinsightstormclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "HdinsightStormCluster"}

// Get takes name of the hdinsightStormCluster, and returns the corresponding hdinsightStormCluster object, and an error if there is any.
func (c *FakeHdinsightStormClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.HdinsightStormCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hdinsightstormclustersResource, c.ns, name), &v1alpha1.HdinsightStormCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightStormCluster), err
}

// List takes label and field selectors, and returns the list of HdinsightStormClusters that match those selectors.
func (c *FakeHdinsightStormClusters) List(opts v1.ListOptions) (result *v1alpha1.HdinsightStormClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hdinsightstormclustersResource, hdinsightstormclustersKind, c.ns, opts), &v1alpha1.HdinsightStormClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HdinsightStormClusterList{ListMeta: obj.(*v1alpha1.HdinsightStormClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.HdinsightStormClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hdinsightStormClusters.
func (c *FakeHdinsightStormClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hdinsightstormclustersResource, c.ns, opts))

}

// Create takes the representation of a hdinsightStormCluster and creates it.  Returns the server's representation of the hdinsightStormCluster, and an error, if there is any.
func (c *FakeHdinsightStormClusters) Create(hdinsightStormCluster *v1alpha1.HdinsightStormCluster) (result *v1alpha1.HdinsightStormCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hdinsightstormclustersResource, c.ns, hdinsightStormCluster), &v1alpha1.HdinsightStormCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightStormCluster), err
}

// Update takes the representation of a hdinsightStormCluster and updates it. Returns the server's representation of the hdinsightStormCluster, and an error, if there is any.
func (c *FakeHdinsightStormClusters) Update(hdinsightStormCluster *v1alpha1.HdinsightStormCluster) (result *v1alpha1.HdinsightStormCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hdinsightstormclustersResource, c.ns, hdinsightStormCluster), &v1alpha1.HdinsightStormCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightStormCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHdinsightStormClusters) UpdateStatus(hdinsightStormCluster *v1alpha1.HdinsightStormCluster) (*v1alpha1.HdinsightStormCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hdinsightstormclustersResource, "status", c.ns, hdinsightStormCluster), &v1alpha1.HdinsightStormCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightStormCluster), err
}

// Delete takes name of the hdinsightStormCluster and deletes it. Returns an error if one occurs.
func (c *FakeHdinsightStormClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hdinsightstormclustersResource, c.ns, name), &v1alpha1.HdinsightStormCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHdinsightStormClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hdinsightstormclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HdinsightStormClusterList{})
	return err
}

// Patch applies the patch and returns the patched hdinsightStormCluster.
func (c *FakeHdinsightStormClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HdinsightStormCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hdinsightstormclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.HdinsightStormCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HdinsightStormCluster), err
}
