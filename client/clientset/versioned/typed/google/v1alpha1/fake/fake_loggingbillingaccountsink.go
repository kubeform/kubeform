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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeLoggingBillingAccountSinks implements LoggingBillingAccountSinkInterface
type FakeLoggingBillingAccountSinks struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var loggingbillingaccountsinksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "loggingbillingaccountsinks"}

var loggingbillingaccountsinksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "LoggingBillingAccountSink"}

// Get takes name of the loggingBillingAccountSink, and returns the corresponding loggingBillingAccountSink object, and an error if there is any.
func (c *FakeLoggingBillingAccountSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.LoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loggingbillingaccountsinksResource, c.ns, name), &v1alpha1.LoggingBillingAccountSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingBillingAccountSink), err
}

// List takes label and field selectors, and returns the list of LoggingBillingAccountSinks that match those selectors.
func (c *FakeLoggingBillingAccountSinks) List(opts v1.ListOptions) (result *v1alpha1.LoggingBillingAccountSinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loggingbillingaccountsinksResource, loggingbillingaccountsinksKind, c.ns, opts), &v1alpha1.LoggingBillingAccountSinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoggingBillingAccountSinkList{ListMeta: obj.(*v1alpha1.LoggingBillingAccountSinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoggingBillingAccountSinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loggingBillingAccountSinks.
func (c *FakeLoggingBillingAccountSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loggingbillingaccountsinksResource, c.ns, opts))

}

// Create takes the representation of a loggingBillingAccountSink and creates it.  Returns the server's representation of the loggingBillingAccountSink, and an error, if there is any.
func (c *FakeLoggingBillingAccountSinks) Create(loggingBillingAccountSink *v1alpha1.LoggingBillingAccountSink) (result *v1alpha1.LoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loggingbillingaccountsinksResource, c.ns, loggingBillingAccountSink), &v1alpha1.LoggingBillingAccountSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingBillingAccountSink), err
}

// Update takes the representation of a loggingBillingAccountSink and updates it. Returns the server's representation of the loggingBillingAccountSink, and an error, if there is any.
func (c *FakeLoggingBillingAccountSinks) Update(loggingBillingAccountSink *v1alpha1.LoggingBillingAccountSink) (result *v1alpha1.LoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loggingbillingaccountsinksResource, c.ns, loggingBillingAccountSink), &v1alpha1.LoggingBillingAccountSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingBillingAccountSink), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoggingBillingAccountSinks) UpdateStatus(loggingBillingAccountSink *v1alpha1.LoggingBillingAccountSink) (*v1alpha1.LoggingBillingAccountSink, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loggingbillingaccountsinksResource, "status", c.ns, loggingBillingAccountSink), &v1alpha1.LoggingBillingAccountSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingBillingAccountSink), err
}

// Delete takes name of the loggingBillingAccountSink and deletes it. Returns an error if one occurs.
func (c *FakeLoggingBillingAccountSinks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loggingbillingaccountsinksResource, c.ns, name), &v1alpha1.LoggingBillingAccountSink{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoggingBillingAccountSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loggingbillingaccountsinksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoggingBillingAccountSinkList{})
	return err
}

// Patch applies the patch and returns the patched loggingBillingAccountSink.
func (c *FakeLoggingBillingAccountSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingBillingAccountSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loggingbillingaccountsinksResource, c.ns, name, pt, data, subresources...), &v1alpha1.LoggingBillingAccountSink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingBillingAccountSink), err
}
