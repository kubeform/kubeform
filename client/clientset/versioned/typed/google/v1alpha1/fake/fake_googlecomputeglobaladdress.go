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

// FakeGoogleComputeGlobalAddresses implements GoogleComputeGlobalAddressInterface
type FakeGoogleComputeGlobalAddresses struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputeglobaladdressesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputeglobaladdresses"}

var googlecomputeglobaladdressesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeGlobalAddress"}

// Get takes name of the googleComputeGlobalAddress, and returns the corresponding googleComputeGlobalAddress object, and an error if there is any.
func (c *FakeGoogleComputeGlobalAddresses) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputeglobaladdressesResource, name), &v1alpha1.GoogleComputeGlobalAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeGlobalAddress), err
}

// List takes label and field selectors, and returns the list of GoogleComputeGlobalAddresses that match those selectors.
func (c *FakeGoogleComputeGlobalAddresses) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeGlobalAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputeglobaladdressesResource, googlecomputeglobaladdressesKind, opts), &v1alpha1.GoogleComputeGlobalAddressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeGlobalAddressList{ListMeta: obj.(*v1alpha1.GoogleComputeGlobalAddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeGlobalAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeGlobalAddresses.
func (c *FakeGoogleComputeGlobalAddresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputeglobaladdressesResource, opts))
}

// Create takes the representation of a googleComputeGlobalAddress and creates it.  Returns the server's representation of the googleComputeGlobalAddress, and an error, if there is any.
func (c *FakeGoogleComputeGlobalAddresses) Create(googleComputeGlobalAddress *v1alpha1.GoogleComputeGlobalAddress) (result *v1alpha1.GoogleComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputeglobaladdressesResource, googleComputeGlobalAddress), &v1alpha1.GoogleComputeGlobalAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeGlobalAddress), err
}

// Update takes the representation of a googleComputeGlobalAddress and updates it. Returns the server's representation of the googleComputeGlobalAddress, and an error, if there is any.
func (c *FakeGoogleComputeGlobalAddresses) Update(googleComputeGlobalAddress *v1alpha1.GoogleComputeGlobalAddress) (result *v1alpha1.GoogleComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputeglobaladdressesResource, googleComputeGlobalAddress), &v1alpha1.GoogleComputeGlobalAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeGlobalAddress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeGlobalAddresses) UpdateStatus(googleComputeGlobalAddress *v1alpha1.GoogleComputeGlobalAddress) (*v1alpha1.GoogleComputeGlobalAddress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputeglobaladdressesResource, "status", googleComputeGlobalAddress), &v1alpha1.GoogleComputeGlobalAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeGlobalAddress), err
}

// Delete takes name of the googleComputeGlobalAddress and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeGlobalAddresses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputeglobaladdressesResource, name), &v1alpha1.GoogleComputeGlobalAddress{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeGlobalAddresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputeglobaladdressesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeGlobalAddressList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeGlobalAddress.
func (c *FakeGoogleComputeGlobalAddresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeGlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputeglobaladdressesResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeGlobalAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeGlobalAddress), err
}
