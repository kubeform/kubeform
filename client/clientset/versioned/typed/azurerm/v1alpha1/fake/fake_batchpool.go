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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeBatchPools implements BatchPoolInterface
type FakeBatchPools struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var batchpoolsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "batchpools"}

var batchpoolsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BatchPool"}

// Get takes name of the batchPool, and returns the corresponding batchPool object, and an error if there is any.
func (c *FakeBatchPools) Get(name string, options v1.GetOptions) (result *v1alpha1.BatchPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(batchpoolsResource, c.ns, name), &v1alpha1.BatchPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchPool), err
}

// List takes label and field selectors, and returns the list of BatchPools that match those selectors.
func (c *FakeBatchPools) List(opts v1.ListOptions) (result *v1alpha1.BatchPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(batchpoolsResource, batchpoolsKind, c.ns, opts), &v1alpha1.BatchPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BatchPoolList{ListMeta: obj.(*v1alpha1.BatchPoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.BatchPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested batchPools.
func (c *FakeBatchPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(batchpoolsResource, c.ns, opts))

}

// Create takes the representation of a batchPool and creates it.  Returns the server's representation of the batchPool, and an error, if there is any.
func (c *FakeBatchPools) Create(batchPool *v1alpha1.BatchPool) (result *v1alpha1.BatchPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(batchpoolsResource, c.ns, batchPool), &v1alpha1.BatchPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchPool), err
}

// Update takes the representation of a batchPool and updates it. Returns the server's representation of the batchPool, and an error, if there is any.
func (c *FakeBatchPools) Update(batchPool *v1alpha1.BatchPool) (result *v1alpha1.BatchPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(batchpoolsResource, c.ns, batchPool), &v1alpha1.BatchPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBatchPools) UpdateStatus(batchPool *v1alpha1.BatchPool) (*v1alpha1.BatchPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(batchpoolsResource, "status", c.ns, batchPool), &v1alpha1.BatchPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchPool), err
}

// Delete takes name of the batchPool and deletes it. Returns an error if one occurs.
func (c *FakeBatchPools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(batchpoolsResource, c.ns, name), &v1alpha1.BatchPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBatchPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(batchpoolsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BatchPoolList{})
	return err
}

// Patch applies the patch and returns the patched batchPool.
func (c *FakeBatchPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(batchpoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BatchPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchPool), err
}
