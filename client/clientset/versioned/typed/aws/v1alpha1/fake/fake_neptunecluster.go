/*
Copyright 2019 The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeNeptuneClusters implements NeptuneClusterInterface
type FakeNeptuneClusters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var neptuneclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "neptuneclusters"}

var neptuneclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "NeptuneCluster"}

// Get takes name of the neptuneCluster, and returns the corresponding neptuneCluster object, and an error if there is any.
func (c *FakeNeptuneClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.NeptuneCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(neptuneclustersResource, c.ns, name), &v1alpha1.NeptuneCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneCluster), err
}

// List takes label and field selectors, and returns the list of NeptuneClusters that match those selectors.
func (c *FakeNeptuneClusters) List(opts v1.ListOptions) (result *v1alpha1.NeptuneClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(neptuneclustersResource, neptuneclustersKind, c.ns, opts), &v1alpha1.NeptuneClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NeptuneClusterList{ListMeta: obj.(*v1alpha1.NeptuneClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.NeptuneClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested neptuneClusters.
func (c *FakeNeptuneClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(neptuneclustersResource, c.ns, opts))

}

// Create takes the representation of a neptuneCluster and creates it.  Returns the server's representation of the neptuneCluster, and an error, if there is any.
func (c *FakeNeptuneClusters) Create(neptuneCluster *v1alpha1.NeptuneCluster) (result *v1alpha1.NeptuneCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(neptuneclustersResource, c.ns, neptuneCluster), &v1alpha1.NeptuneCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneCluster), err
}

// Update takes the representation of a neptuneCluster and updates it. Returns the server's representation of the neptuneCluster, and an error, if there is any.
func (c *FakeNeptuneClusters) Update(neptuneCluster *v1alpha1.NeptuneCluster) (result *v1alpha1.NeptuneCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(neptuneclustersResource, c.ns, neptuneCluster), &v1alpha1.NeptuneCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNeptuneClusters) UpdateStatus(neptuneCluster *v1alpha1.NeptuneCluster) (*v1alpha1.NeptuneCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(neptuneclustersResource, "status", c.ns, neptuneCluster), &v1alpha1.NeptuneCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneCluster), err
}

// Delete takes name of the neptuneCluster and deletes it. Returns an error if one occurs.
func (c *FakeNeptuneClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(neptuneclustersResource, c.ns, name), &v1alpha1.NeptuneCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNeptuneClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(neptuneclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NeptuneClusterList{})
	return err
}

// Patch applies the patch and returns the patched neptuneCluster.
func (c *FakeNeptuneClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NeptuneCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(neptuneclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.NeptuneCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneCluster), err
}