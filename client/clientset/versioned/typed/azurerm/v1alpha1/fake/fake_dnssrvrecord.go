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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDnsSrvRecords implements DnsSrvRecordInterface
type FakeDnsSrvRecords struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var dnssrvrecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "dnssrvrecords"}

var dnssrvrecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DnsSrvRecord"}

// Get takes name of the dnsSrvRecord, and returns the corresponding dnsSrvRecord object, and an error if there is any.
func (c *FakeDnsSrvRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnssrvrecordsResource, c.ns, name), &v1alpha1.DnsSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsSrvRecord), err
}

// List takes label and field selectors, and returns the list of DnsSrvRecords that match those selectors.
func (c *FakeDnsSrvRecords) List(opts v1.ListOptions) (result *v1alpha1.DnsSrvRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnssrvrecordsResource, dnssrvrecordsKind, c.ns, opts), &v1alpha1.DnsSrvRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DnsSrvRecordList{ListMeta: obj.(*v1alpha1.DnsSrvRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.DnsSrvRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnsSrvRecords.
func (c *FakeDnsSrvRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnssrvrecordsResource, c.ns, opts))

}

// Create takes the representation of a dnsSrvRecord and creates it.  Returns the server's representation of the dnsSrvRecord, and an error, if there is any.
func (c *FakeDnsSrvRecords) Create(dnsSrvRecord *v1alpha1.DnsSrvRecord) (result *v1alpha1.DnsSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnssrvrecordsResource, c.ns, dnsSrvRecord), &v1alpha1.DnsSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsSrvRecord), err
}

// Update takes the representation of a dnsSrvRecord and updates it. Returns the server's representation of the dnsSrvRecord, and an error, if there is any.
func (c *FakeDnsSrvRecords) Update(dnsSrvRecord *v1alpha1.DnsSrvRecord) (result *v1alpha1.DnsSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnssrvrecordsResource, c.ns, dnsSrvRecord), &v1alpha1.DnsSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsSrvRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnsSrvRecords) UpdateStatus(dnsSrvRecord *v1alpha1.DnsSrvRecord) (*v1alpha1.DnsSrvRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnssrvrecordsResource, "status", c.ns, dnsSrvRecord), &v1alpha1.DnsSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsSrvRecord), err
}

// Delete takes name of the dnsSrvRecord and deletes it. Returns an error if one occurs.
func (c *FakeDnsSrvRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnssrvrecordsResource, c.ns, name), &v1alpha1.DnsSrvRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnsSrvRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnssrvrecordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DnsSrvRecordList{})
	return err
}

// Patch applies the patch and returns the patched dnsSrvRecord.
func (c *FakeDnsSrvRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnssrvrecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DnsSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsSrvRecord), err
}
