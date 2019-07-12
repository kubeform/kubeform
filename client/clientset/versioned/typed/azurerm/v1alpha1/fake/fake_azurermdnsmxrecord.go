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

// FakeAzurermDnsMxRecords implements AzurermDnsMxRecordInterface
type FakeAzurermDnsMxRecords struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdnsmxrecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdnsmxrecords"}

var azurermdnsmxrecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDnsMxRecord"}

// Get takes name of the azurermDnsMxRecord, and returns the corresponding azurermDnsMxRecord object, and an error if there is any.
func (c *FakeAzurermDnsMxRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdnsmxrecordsResource, name), &v1alpha1.AzurermDnsMxRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsMxRecord), err
}

// List takes label and field selectors, and returns the list of AzurermDnsMxRecords that match those selectors.
func (c *FakeAzurermDnsMxRecords) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsMxRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdnsmxrecordsResource, azurermdnsmxrecordsKind, opts), &v1alpha1.AzurermDnsMxRecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDnsMxRecordList{ListMeta: obj.(*v1alpha1.AzurermDnsMxRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDnsMxRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDnsMxRecords.
func (c *FakeAzurermDnsMxRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdnsmxrecordsResource, opts))
}

// Create takes the representation of a azurermDnsMxRecord and creates it.  Returns the server's representation of the azurermDnsMxRecord, and an error, if there is any.
func (c *FakeAzurermDnsMxRecords) Create(azurermDnsMxRecord *v1alpha1.AzurermDnsMxRecord) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdnsmxrecordsResource, azurermDnsMxRecord), &v1alpha1.AzurermDnsMxRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsMxRecord), err
}

// Update takes the representation of a azurermDnsMxRecord and updates it. Returns the server's representation of the azurermDnsMxRecord, and an error, if there is any.
func (c *FakeAzurermDnsMxRecords) Update(azurermDnsMxRecord *v1alpha1.AzurermDnsMxRecord) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdnsmxrecordsResource, azurermDnsMxRecord), &v1alpha1.AzurermDnsMxRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsMxRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDnsMxRecords) UpdateStatus(azurermDnsMxRecord *v1alpha1.AzurermDnsMxRecord) (*v1alpha1.AzurermDnsMxRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdnsmxrecordsResource, "status", azurermDnsMxRecord), &v1alpha1.AzurermDnsMxRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsMxRecord), err
}

// Delete takes name of the azurermDnsMxRecord and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDnsMxRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdnsmxrecordsResource, name), &v1alpha1.AzurermDnsMxRecord{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDnsMxRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdnsmxrecordsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDnsMxRecordList{})
	return err
}

// Patch applies the patch and returns the patched azurermDnsMxRecord.
func (c *FakeAzurermDnsMxRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdnsmxrecordsResource, name, pt, data, subresources...), &v1alpha1.AzurermDnsMxRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsMxRecord), err
}
