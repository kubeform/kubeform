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

// FakePrivateDNSCnameRecords implements PrivateDNSCnameRecordInterface
type FakePrivateDNSCnameRecords struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var privatednscnamerecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "privatednscnamerecords"}

var privatednscnamerecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PrivateDNSCnameRecord"}

// Get takes name of the privateDNSCnameRecord, and returns the corresponding privateDNSCnameRecord object, and an error if there is any.
func (c *FakePrivateDNSCnameRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrivateDNSCnameRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatednscnamerecordsResource, c.ns, name), &v1alpha1.PrivateDNSCnameRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSCnameRecord), err
}

// List takes label and field selectors, and returns the list of PrivateDNSCnameRecords that match those selectors.
func (c *FakePrivateDNSCnameRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrivateDNSCnameRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatednscnamerecordsResource, privatednscnamerecordsKind, c.ns, opts), &v1alpha1.PrivateDNSCnameRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrivateDNSCnameRecordList{ListMeta: obj.(*v1alpha1.PrivateDNSCnameRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrivateDNSCnameRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateDNSCnameRecords.
func (c *FakePrivateDNSCnameRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatednscnamerecordsResource, c.ns, opts))

}

// Create takes the representation of a privateDNSCnameRecord and creates it.  Returns the server's representation of the privateDNSCnameRecord, and an error, if there is any.
func (c *FakePrivateDNSCnameRecords) Create(ctx context.Context, privateDNSCnameRecord *v1alpha1.PrivateDNSCnameRecord, opts v1.CreateOptions) (result *v1alpha1.PrivateDNSCnameRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatednscnamerecordsResource, c.ns, privateDNSCnameRecord), &v1alpha1.PrivateDNSCnameRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSCnameRecord), err
}

// Update takes the representation of a privateDNSCnameRecord and updates it. Returns the server's representation of the privateDNSCnameRecord, and an error, if there is any.
func (c *FakePrivateDNSCnameRecords) Update(ctx context.Context, privateDNSCnameRecord *v1alpha1.PrivateDNSCnameRecord, opts v1.UpdateOptions) (result *v1alpha1.PrivateDNSCnameRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatednscnamerecordsResource, c.ns, privateDNSCnameRecord), &v1alpha1.PrivateDNSCnameRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSCnameRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateDNSCnameRecords) UpdateStatus(ctx context.Context, privateDNSCnameRecord *v1alpha1.PrivateDNSCnameRecord, opts v1.UpdateOptions) (*v1alpha1.PrivateDNSCnameRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatednscnamerecordsResource, "status", c.ns, privateDNSCnameRecord), &v1alpha1.PrivateDNSCnameRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSCnameRecord), err
}

// Delete takes name of the privateDNSCnameRecord and deletes it. Returns an error if one occurs.
func (c *FakePrivateDNSCnameRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(privatednscnamerecordsResource, c.ns, name), &v1alpha1.PrivateDNSCnameRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateDNSCnameRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatednscnamerecordsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrivateDNSCnameRecordList{})
	return err
}

// Patch applies the patch and returns the patched privateDNSCnameRecord.
func (c *FakePrivateDNSCnameRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrivateDNSCnameRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatednscnamerecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrivateDNSCnameRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSCnameRecord), err
}
