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

// FakeBatchApplications implements BatchApplicationInterface
type FakeBatchApplications struct {
	Fake *FakeAzurermV1alpha1
}

var batchapplicationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "batchapplications"}

var batchapplicationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BatchApplication"}

// Get takes name of the batchApplication, and returns the corresponding batchApplication object, and an error if there is any.
func (c *FakeBatchApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.BatchApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(batchapplicationsResource, name), &v1alpha1.BatchApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchApplication), err
}

// List takes label and field selectors, and returns the list of BatchApplications that match those selectors.
func (c *FakeBatchApplications) List(opts v1.ListOptions) (result *v1alpha1.BatchApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(batchapplicationsResource, batchapplicationsKind, opts), &v1alpha1.BatchApplicationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BatchApplicationList{ListMeta: obj.(*v1alpha1.BatchApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.BatchApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested batchApplications.
func (c *FakeBatchApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(batchapplicationsResource, opts))
}

// Create takes the representation of a batchApplication and creates it.  Returns the server's representation of the batchApplication, and an error, if there is any.
func (c *FakeBatchApplications) Create(batchApplication *v1alpha1.BatchApplication) (result *v1alpha1.BatchApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(batchapplicationsResource, batchApplication), &v1alpha1.BatchApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchApplication), err
}

// Update takes the representation of a batchApplication and updates it. Returns the server's representation of the batchApplication, and an error, if there is any.
func (c *FakeBatchApplications) Update(batchApplication *v1alpha1.BatchApplication) (result *v1alpha1.BatchApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(batchapplicationsResource, batchApplication), &v1alpha1.BatchApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBatchApplications) UpdateStatus(batchApplication *v1alpha1.BatchApplication) (*v1alpha1.BatchApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(batchapplicationsResource, "status", batchApplication), &v1alpha1.BatchApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchApplication), err
}

// Delete takes name of the batchApplication and deletes it. Returns an error if one occurs.
func (c *FakeBatchApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(batchapplicationsResource, name), &v1alpha1.BatchApplication{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBatchApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(batchapplicationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BatchApplicationList{})
	return err
}

// Patch applies the patch and returns the patched batchApplication.
func (c *FakeBatchApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(batchapplicationsResource, name, pt, data, subresources...), &v1alpha1.BatchApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchApplication), err
}
