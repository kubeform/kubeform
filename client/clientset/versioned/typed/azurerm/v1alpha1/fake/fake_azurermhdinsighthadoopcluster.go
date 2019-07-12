/*
Copyright 2019 The KubeDB Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermHdinsightHadoopClusters implements AzurermHdinsightHadoopClusterInterface
type FakeAzurermHdinsightHadoopClusters struct {
	Fake *FakeAzurermV1alpha1
}

var azurermhdinsighthadoopclustersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermhdinsighthadoopclusters"}

var azurermhdinsighthadoopclustersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermHdinsightHadoopCluster"}

// Get takes name of the azurermHdinsightHadoopCluster, and returns the corresponding azurermHdinsightHadoopCluster object, and an error if there is any.
func (c *FakeAzurermHdinsightHadoopClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermHdinsightHadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermhdinsighthadoopclustersResource, name), &v1alpha1.AzurermHdinsightHadoopCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightHadoopCluster), err
}

// List takes label and field selectors, and returns the list of AzurermHdinsightHadoopClusters that match those selectors.
func (c *FakeAzurermHdinsightHadoopClusters) List(opts v1.ListOptions) (result *v1alpha1.AzurermHdinsightHadoopClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermhdinsighthadoopclustersResource, azurermhdinsighthadoopclustersKind, opts), &v1alpha1.AzurermHdinsightHadoopClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermHdinsightHadoopClusterList{ListMeta: obj.(*v1alpha1.AzurermHdinsightHadoopClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermHdinsightHadoopClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermHdinsightHadoopClusters.
func (c *FakeAzurermHdinsightHadoopClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermhdinsighthadoopclustersResource, opts))
}

// Create takes the representation of a azurermHdinsightHadoopCluster and creates it.  Returns the server's representation of the azurermHdinsightHadoopCluster, and an error, if there is any.
func (c *FakeAzurermHdinsightHadoopClusters) Create(azurermHdinsightHadoopCluster *v1alpha1.AzurermHdinsightHadoopCluster) (result *v1alpha1.AzurermHdinsightHadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermhdinsighthadoopclustersResource, azurermHdinsightHadoopCluster), &v1alpha1.AzurermHdinsightHadoopCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightHadoopCluster), err
}

// Update takes the representation of a azurermHdinsightHadoopCluster and updates it. Returns the server's representation of the azurermHdinsightHadoopCluster, and an error, if there is any.
func (c *FakeAzurermHdinsightHadoopClusters) Update(azurermHdinsightHadoopCluster *v1alpha1.AzurermHdinsightHadoopCluster) (result *v1alpha1.AzurermHdinsightHadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermhdinsighthadoopclustersResource, azurermHdinsightHadoopCluster), &v1alpha1.AzurermHdinsightHadoopCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightHadoopCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermHdinsightHadoopClusters) UpdateStatus(azurermHdinsightHadoopCluster *v1alpha1.AzurermHdinsightHadoopCluster) (*v1alpha1.AzurermHdinsightHadoopCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermhdinsighthadoopclustersResource, "status", azurermHdinsightHadoopCluster), &v1alpha1.AzurermHdinsightHadoopCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightHadoopCluster), err
}

// Delete takes name of the azurermHdinsightHadoopCluster and deletes it. Returns an error if one occurs.
func (c *FakeAzurermHdinsightHadoopClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermhdinsighthadoopclustersResource, name), &v1alpha1.AzurermHdinsightHadoopCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermHdinsightHadoopClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermhdinsighthadoopclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermHdinsightHadoopClusterList{})
	return err
}

// Patch applies the patch and returns the patched azurermHdinsightHadoopCluster.
func (c *FakeAzurermHdinsightHadoopClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermHdinsightHadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermhdinsighthadoopclustersResource, name, pt, data, subresources...), &v1alpha1.AzurermHdinsightHadoopCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermHdinsightHadoopCluster), err
}
