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

// FakeAzurermDnsCaaRecords implements AzurermDnsCaaRecordInterface
type FakeAzurermDnsCaaRecords struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdnscaarecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdnscaarecords"}

var azurermdnscaarecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDnsCaaRecord"}

// Get takes name of the azurermDnsCaaRecord, and returns the corresponding azurermDnsCaaRecord object, and an error if there is any.
func (c *FakeAzurermDnsCaaRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDnsCaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdnscaarecordsResource, name), &v1alpha1.AzurermDnsCaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsCaaRecord), err
}

// List takes label and field selectors, and returns the list of AzurermDnsCaaRecords that match those selectors.
func (c *FakeAzurermDnsCaaRecords) List(opts v1.ListOptions) (result *v1alpha1.AzurermDnsCaaRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdnscaarecordsResource, azurermdnscaarecordsKind, opts), &v1alpha1.AzurermDnsCaaRecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDnsCaaRecordList{ListMeta: obj.(*v1alpha1.AzurermDnsCaaRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDnsCaaRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDnsCaaRecords.
func (c *FakeAzurermDnsCaaRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdnscaarecordsResource, opts))
}

// Create takes the representation of a azurermDnsCaaRecord and creates it.  Returns the server's representation of the azurermDnsCaaRecord, and an error, if there is any.
func (c *FakeAzurermDnsCaaRecords) Create(azurermDnsCaaRecord *v1alpha1.AzurermDnsCaaRecord) (result *v1alpha1.AzurermDnsCaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdnscaarecordsResource, azurermDnsCaaRecord), &v1alpha1.AzurermDnsCaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsCaaRecord), err
}

// Update takes the representation of a azurermDnsCaaRecord and updates it. Returns the server's representation of the azurermDnsCaaRecord, and an error, if there is any.
func (c *FakeAzurermDnsCaaRecords) Update(azurermDnsCaaRecord *v1alpha1.AzurermDnsCaaRecord) (result *v1alpha1.AzurermDnsCaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdnscaarecordsResource, azurermDnsCaaRecord), &v1alpha1.AzurermDnsCaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsCaaRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDnsCaaRecords) UpdateStatus(azurermDnsCaaRecord *v1alpha1.AzurermDnsCaaRecord) (*v1alpha1.AzurermDnsCaaRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdnscaarecordsResource, "status", azurermDnsCaaRecord), &v1alpha1.AzurermDnsCaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsCaaRecord), err
}

// Delete takes name of the azurermDnsCaaRecord and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDnsCaaRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdnscaarecordsResource, name), &v1alpha1.AzurermDnsCaaRecord{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDnsCaaRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdnscaarecordsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDnsCaaRecordList{})
	return err
}

// Patch applies the patch and returns the patched azurermDnsCaaRecord.
func (c *FakeAzurermDnsCaaRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDnsCaaRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdnscaarecordsResource, name, pt, data, subresources...), &v1alpha1.AzurermDnsCaaRecord{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDnsCaaRecord), err
}
