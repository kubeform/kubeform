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

// FakeAzurermDnsPtrRecords implements AzurermDnsPtrRecordInterface
type FakeAzurermDnsPtrRecords struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdnsptrrecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdnsptrrecords"}

var azurermdnsptrrecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDnsPtrRecord"}

// Get takes name of the azurermDnsPtrRecord, and returns the corresponding azurermDnsPtrRecord object, and an error if there is any.
func (c *FakeAzurermDnsPtrRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdnsptrrecordsResource, name), &v1alpha1.AzurermDnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsPtrRecord), err
}

// List takes label and field selectors, and returns the list of AzurermDnsPtrRecords that match those selectors.
func (c *FakeAzurermDnsPtrRecords) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsPtrRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdnsptrrecordsResource, azurermdnsptrrecordsKind, opts), &v1alpha1.AzurermDnsPtrRecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDnsPtrRecordList{ListMeta: obj.(*v1alpha1.AzurermDnsPtrRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDnsPtrRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDnsPtrRecords.
func (c *FakeAzurermDnsPtrRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdnsptrrecordsResource, opts))
}

// Create takes the representation of a azurermDnsPtrRecord and creates it.  Returns the server's representation of the azurermDnsPtrRecord, and an error, if there is any.
func (c *FakeAzurermDnsPtrRecords) Create(azurermDnsPtrRecord *v1alpha1.AzurermDnsPtrRecord) (result *v1alpha1.AzurermDnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdnsptrrecordsResource, azurermDnsPtrRecord), &v1alpha1.AzurermDnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsPtrRecord), err
}

// Update takes the representation of a azurermDnsPtrRecord and updates it. Returns the server's representation of the azurermDnsPtrRecord, and an error, if there is any.
func (c *FakeAzurermDnsPtrRecords) Update(azurermDnsPtrRecord *v1alpha1.AzurermDnsPtrRecord) (result *v1alpha1.AzurermDnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdnsptrrecordsResource, azurermDnsPtrRecord), &v1alpha1.AzurermDnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsPtrRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDnsPtrRecords) UpdateStatus(azurermDnsPtrRecord *v1alpha1.AzurermDnsPtrRecord) (*v1alpha1.AzurermDnsPtrRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdnsptrrecordsResource, "status", azurermDnsPtrRecord), &v1alpha1.AzurermDnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsPtrRecord), err
}

// Delete takes name of the azurermDnsPtrRecord and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDnsPtrRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdnsptrrecordsResource, name), &v1alpha1.AzurermDnsPtrRecord{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDnsPtrRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdnsptrrecordsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDnsPtrRecordList{})
	return err
}

// Patch applies the patch and returns the patched azurermDnsPtrRecord.
func (c *FakeAzurermDnsPtrRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsPtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdnsptrrecordsResource, name, pt, data, subresources...), &v1alpha1.AzurermDnsPtrRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsPtrRecord), err
}
