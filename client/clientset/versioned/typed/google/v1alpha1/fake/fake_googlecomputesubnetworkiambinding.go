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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleComputeSubnetworkIamBindings implements GoogleComputeSubnetworkIamBindingInterface
type FakeGoogleComputeSubnetworkIamBindings struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputesubnetworkiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputesubnetworkiambindings"}

var googlecomputesubnetworkiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeSubnetworkIamBinding"}

// Get takes name of the googleComputeSubnetworkIamBinding, and returns the corresponding googleComputeSubnetworkIamBinding object, and an error if there is any.
func (c *FakeGoogleComputeSubnetworkIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeSubnetworkIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputesubnetworkiambindingsResource, name), &v1alpha1.GoogleComputeSubnetworkIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeSubnetworkIamBinding), err
}

// List takes label and field selectors, and returns the list of GoogleComputeSubnetworkIamBindings that match those selectors.
func (c *FakeGoogleComputeSubnetworkIamBindings) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeSubnetworkIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputesubnetworkiambindingsResource, googlecomputesubnetworkiambindingsKind, opts), &v1alpha1.GoogleComputeSubnetworkIamBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeSubnetworkIamBindingList{ListMeta: obj.(*v1alpha1.GoogleComputeSubnetworkIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeSubnetworkIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeSubnetworkIamBindings.
func (c *FakeGoogleComputeSubnetworkIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputesubnetworkiambindingsResource, opts))
}

// Create takes the representation of a googleComputeSubnetworkIamBinding and creates it.  Returns the server's representation of the googleComputeSubnetworkIamBinding, and an error, if there is any.
func (c *FakeGoogleComputeSubnetworkIamBindings) Create(googleComputeSubnetworkIamBinding *v1alpha1.GoogleComputeSubnetworkIamBinding) (result *v1alpha1.GoogleComputeSubnetworkIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputesubnetworkiambindingsResource, googleComputeSubnetworkIamBinding), &v1alpha1.GoogleComputeSubnetworkIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeSubnetworkIamBinding), err
}

// Update takes the representation of a googleComputeSubnetworkIamBinding and updates it. Returns the server's representation of the googleComputeSubnetworkIamBinding, and an error, if there is any.
func (c *FakeGoogleComputeSubnetworkIamBindings) Update(googleComputeSubnetworkIamBinding *v1alpha1.GoogleComputeSubnetworkIamBinding) (result *v1alpha1.GoogleComputeSubnetworkIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputesubnetworkiambindingsResource, googleComputeSubnetworkIamBinding), &v1alpha1.GoogleComputeSubnetworkIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeSubnetworkIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeSubnetworkIamBindings) UpdateStatus(googleComputeSubnetworkIamBinding *v1alpha1.GoogleComputeSubnetworkIamBinding) (*v1alpha1.GoogleComputeSubnetworkIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputesubnetworkiambindingsResource, "status", googleComputeSubnetworkIamBinding), &v1alpha1.GoogleComputeSubnetworkIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeSubnetworkIamBinding), err
}

// Delete takes name of the googleComputeSubnetworkIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeSubnetworkIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputesubnetworkiambindingsResource, name), &v1alpha1.GoogleComputeSubnetworkIamBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeSubnetworkIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputesubnetworkiambindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeSubnetworkIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeSubnetworkIamBinding.
func (c *FakeGoogleComputeSubnetworkIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeSubnetworkIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputesubnetworkiambindingsResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeSubnetworkIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeSubnetworkIamBinding), err
}
