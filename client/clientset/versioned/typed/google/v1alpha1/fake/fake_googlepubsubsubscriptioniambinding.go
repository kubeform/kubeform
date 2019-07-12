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

// FakeGooglePubsubSubscriptionIamBindings implements GooglePubsubSubscriptionIamBindingInterface
type FakeGooglePubsubSubscriptionIamBindings struct {
	Fake *FakeGoogleV1alpha1
}

var googlepubsubsubscriptioniambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlepubsubsubscriptioniambindings"}

var googlepubsubsubscriptioniambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GooglePubsubSubscriptionIamBinding"}

// Get takes name of the googlePubsubSubscriptionIamBinding, and returns the corresponding googlePubsubSubscriptionIamBinding object, and an error if there is any.
func (c *FakeGooglePubsubSubscriptionIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GooglePubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlepubsubsubscriptioniambindingsResource, name), &v1alpha1.GooglePubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamBinding), err
}

// List takes label and field selectors, and returns the list of GooglePubsubSubscriptionIamBindings that match those selectors.
func (c *FakeGooglePubsubSubscriptionIamBindings) List(opts v1.ListOptions) (result *v1alpha1.GooglePubsubSubscriptionIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlepubsubsubscriptioniambindingsResource, googlepubsubsubscriptioniambindingsKind, opts), &v1alpha1.GooglePubsubSubscriptionIamBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GooglePubsubSubscriptionIamBindingList{ListMeta: obj.(*v1alpha1.GooglePubsubSubscriptionIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.GooglePubsubSubscriptionIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googlePubsubSubscriptionIamBindings.
func (c *FakeGooglePubsubSubscriptionIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlepubsubsubscriptioniambindingsResource, opts))
}

// Create takes the representation of a googlePubsubSubscriptionIamBinding and creates it.  Returns the server's representation of the googlePubsubSubscriptionIamBinding, and an error, if there is any.
func (c *FakeGooglePubsubSubscriptionIamBindings) Create(googlePubsubSubscriptionIamBinding *v1alpha1.GooglePubsubSubscriptionIamBinding) (result *v1alpha1.GooglePubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlepubsubsubscriptioniambindingsResource, googlePubsubSubscriptionIamBinding), &v1alpha1.GooglePubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamBinding), err
}

// Update takes the representation of a googlePubsubSubscriptionIamBinding and updates it. Returns the server's representation of the googlePubsubSubscriptionIamBinding, and an error, if there is any.
func (c *FakeGooglePubsubSubscriptionIamBindings) Update(googlePubsubSubscriptionIamBinding *v1alpha1.GooglePubsubSubscriptionIamBinding) (result *v1alpha1.GooglePubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlepubsubsubscriptioniambindingsResource, googlePubsubSubscriptionIamBinding), &v1alpha1.GooglePubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGooglePubsubSubscriptionIamBindings) UpdateStatus(googlePubsubSubscriptionIamBinding *v1alpha1.GooglePubsubSubscriptionIamBinding) (*v1alpha1.GooglePubsubSubscriptionIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlepubsubsubscriptioniambindingsResource, "status", googlePubsubSubscriptionIamBinding), &v1alpha1.GooglePubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamBinding), err
}

// Delete takes name of the googlePubsubSubscriptionIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeGooglePubsubSubscriptionIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlepubsubsubscriptioniambindingsResource, name), &v1alpha1.GooglePubsubSubscriptionIamBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGooglePubsubSubscriptionIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlepubsubsubscriptioniambindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GooglePubsubSubscriptionIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched googlePubsubSubscriptionIamBinding.
func (c *FakeGooglePubsubSubscriptionIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GooglePubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlepubsubsubscriptioniambindingsResource, name, pt, data, subresources...), &v1alpha1.GooglePubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GooglePubsubSubscriptionIamBinding), err
}
