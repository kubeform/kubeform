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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBillingAccountIamBindings implements BillingAccountIamBindingInterface
type FakeBillingAccountIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var billingaccountiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "billingaccountiambindings"}

var billingaccountiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BillingAccountIamBinding"}

// Get takes name of the billingAccountIamBinding, and returns the corresponding billingAccountIamBinding object, and an error if there is any.
func (c *FakeBillingAccountIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.BillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(billingaccountiambindingsResource, c.ns, name), &v1alpha1.BillingAccountIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamBinding), err
}

// List takes label and field selectors, and returns the list of BillingAccountIamBindings that match those selectors.
func (c *FakeBillingAccountIamBindings) List(opts v1.ListOptions) (result *v1alpha1.BillingAccountIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(billingaccountiambindingsResource, billingaccountiambindingsKind, c.ns, opts), &v1alpha1.BillingAccountIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BillingAccountIamBindingList{ListMeta: obj.(*v1alpha1.BillingAccountIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.BillingAccountIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested billingAccountIamBindings.
func (c *FakeBillingAccountIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(billingaccountiambindingsResource, c.ns, opts))

}

// Create takes the representation of a billingAccountIamBinding and creates it.  Returns the server's representation of the billingAccountIamBinding, and an error, if there is any.
func (c *FakeBillingAccountIamBindings) Create(billingAccountIamBinding *v1alpha1.BillingAccountIamBinding) (result *v1alpha1.BillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(billingaccountiambindingsResource, c.ns, billingAccountIamBinding), &v1alpha1.BillingAccountIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamBinding), err
}

// Update takes the representation of a billingAccountIamBinding and updates it. Returns the server's representation of the billingAccountIamBinding, and an error, if there is any.
func (c *FakeBillingAccountIamBindings) Update(billingAccountIamBinding *v1alpha1.BillingAccountIamBinding) (result *v1alpha1.BillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(billingaccountiambindingsResource, c.ns, billingAccountIamBinding), &v1alpha1.BillingAccountIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBillingAccountIamBindings) UpdateStatus(billingAccountIamBinding *v1alpha1.BillingAccountIamBinding) (*v1alpha1.BillingAccountIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(billingaccountiambindingsResource, "status", c.ns, billingAccountIamBinding), &v1alpha1.BillingAccountIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamBinding), err
}

// Delete takes name of the billingAccountIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeBillingAccountIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(billingaccountiambindingsResource, c.ns, name), &v1alpha1.BillingAccountIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBillingAccountIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(billingaccountiambindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BillingAccountIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched billingAccountIamBinding.
func (c *FakeBillingAccountIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BillingAccountIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(billingaccountiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BillingAccountIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingAccountIamBinding), err
}
