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

// FakeGoogleLoggingBillingAccountSinks implements GoogleLoggingBillingAccountSinkInterface
type FakeGoogleLoggingBillingAccountSinks struct {
	Fake *FakeGoogleV1alpha1
}

var googleloggingbillingaccountsinksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googleloggingbillingaccountsinks"}

var googleloggingbillingaccountsinksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleLoggingBillingAccountSink"}

// Get takes name of the googleLoggingBillingAccountSink, and returns the corresponding googleLoggingBillingAccountSink object, and an error if there is any.
func (c *FakeGoogleLoggingBillingAccountSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleLoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googleloggingbillingaccountsinksResource, name), &v1alpha1.GoogleLoggingBillingAccountSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingBillingAccountSink), err
}

// List takes label and field selectors, and returns the list of GoogleLoggingBillingAccountSinks that match those selectors.
func (c *FakeGoogleLoggingBillingAccountSinks) List(opts v1.ListOptions) (result *v1alpha1.GoogleLoggingBillingAccountSinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googleloggingbillingaccountsinksResource, googleloggingbillingaccountsinksKind, opts), &v1alpha1.GoogleLoggingBillingAccountSinkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleLoggingBillingAccountSinkList{ListMeta: obj.(*v1alpha1.GoogleLoggingBillingAccountSinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleLoggingBillingAccountSinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleLoggingBillingAccountSinks.
func (c *FakeGoogleLoggingBillingAccountSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googleloggingbillingaccountsinksResource, opts))
}

// Create takes the representation of a googleLoggingBillingAccountSink and creates it.  Returns the server's representation of the googleLoggingBillingAccountSink, and an error, if there is any.
func (c *FakeGoogleLoggingBillingAccountSinks) Create(googleLoggingBillingAccountSink *v1alpha1.GoogleLoggingBillingAccountSink) (result *v1alpha1.GoogleLoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googleloggingbillingaccountsinksResource, googleLoggingBillingAccountSink), &v1alpha1.GoogleLoggingBillingAccountSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingBillingAccountSink), err
}

// Update takes the representation of a googleLoggingBillingAccountSink and updates it. Returns the server's representation of the googleLoggingBillingAccountSink, and an error, if there is any.
func (c *FakeGoogleLoggingBillingAccountSinks) Update(googleLoggingBillingAccountSink *v1alpha1.GoogleLoggingBillingAccountSink) (result *v1alpha1.GoogleLoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googleloggingbillingaccountsinksResource, googleLoggingBillingAccountSink), &v1alpha1.GoogleLoggingBillingAccountSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingBillingAccountSink), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleLoggingBillingAccountSinks) UpdateStatus(googleLoggingBillingAccountSink *v1alpha1.GoogleLoggingBillingAccountSink) (*v1alpha1.GoogleLoggingBillingAccountSink, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googleloggingbillingaccountsinksResource, "status", googleLoggingBillingAccountSink), &v1alpha1.GoogleLoggingBillingAccountSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingBillingAccountSink), err
}

// Delete takes name of the googleLoggingBillingAccountSink and deletes it. Returns an error if one occurs.
func (c *FakeGoogleLoggingBillingAccountSinks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googleloggingbillingaccountsinksResource, name), &v1alpha1.GoogleLoggingBillingAccountSink{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleLoggingBillingAccountSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googleloggingbillingaccountsinksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleLoggingBillingAccountSinkList{})
	return err
}

// Patch applies the patch and returns the patched googleLoggingBillingAccountSink.
func (c *FakeGoogleLoggingBillingAccountSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleLoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googleloggingbillingaccountsinksResource, name, pt, data, subresources...), &v1alpha1.GoogleLoggingBillingAccountSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingBillingAccountSink), err
}
