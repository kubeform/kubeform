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

// FakeGoogleSpannerDatabaseIamBindings implements GoogleSpannerDatabaseIamBindingInterface
type FakeGoogleSpannerDatabaseIamBindings struct {
	Fake *FakeGoogleV1alpha1
}

var googlespannerdatabaseiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlespannerdatabaseiambindings"}

var googlespannerdatabaseiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleSpannerDatabaseIamBinding"}

// Get takes name of the googleSpannerDatabaseIamBinding, and returns the corresponding googleSpannerDatabaseIamBinding object, and an error if there is any.
func (c *FakeGoogleSpannerDatabaseIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleSpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlespannerdatabaseiambindingsResource, name), &v1alpha1.GoogleSpannerDatabaseIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamBinding), err
}

// List takes label and field selectors, and returns the list of GoogleSpannerDatabaseIamBindings that match those selectors.
func (c *FakeGoogleSpannerDatabaseIamBindings) List(opts v1.ListOptions) (result *v1alpha1.GoogleSpannerDatabaseIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlespannerdatabaseiambindingsResource, googlespannerdatabaseiambindingsKind, opts), &v1alpha1.GoogleSpannerDatabaseIamBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleSpannerDatabaseIamBindingList{ListMeta: obj.(*v1alpha1.GoogleSpannerDatabaseIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleSpannerDatabaseIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleSpannerDatabaseIamBindings.
func (c *FakeGoogleSpannerDatabaseIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlespannerdatabaseiambindingsResource, opts))
}

// Create takes the representation of a googleSpannerDatabaseIamBinding and creates it.  Returns the server's representation of the googleSpannerDatabaseIamBinding, and an error, if there is any.
func (c *FakeGoogleSpannerDatabaseIamBindings) Create(googleSpannerDatabaseIamBinding *v1alpha1.GoogleSpannerDatabaseIamBinding) (result *v1alpha1.GoogleSpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlespannerdatabaseiambindingsResource, googleSpannerDatabaseIamBinding), &v1alpha1.GoogleSpannerDatabaseIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamBinding), err
}

// Update takes the representation of a googleSpannerDatabaseIamBinding and updates it. Returns the server's representation of the googleSpannerDatabaseIamBinding, and an error, if there is any.
func (c *FakeGoogleSpannerDatabaseIamBindings) Update(googleSpannerDatabaseIamBinding *v1alpha1.GoogleSpannerDatabaseIamBinding) (result *v1alpha1.GoogleSpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlespannerdatabaseiambindingsResource, googleSpannerDatabaseIamBinding), &v1alpha1.GoogleSpannerDatabaseIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleSpannerDatabaseIamBindings) UpdateStatus(googleSpannerDatabaseIamBinding *v1alpha1.GoogleSpannerDatabaseIamBinding) (*v1alpha1.GoogleSpannerDatabaseIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlespannerdatabaseiambindingsResource, "status", googleSpannerDatabaseIamBinding), &v1alpha1.GoogleSpannerDatabaseIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamBinding), err
}

// Delete takes name of the googleSpannerDatabaseIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeGoogleSpannerDatabaseIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlespannerdatabaseiambindingsResource, name), &v1alpha1.GoogleSpannerDatabaseIamBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleSpannerDatabaseIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlespannerdatabaseiambindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleSpannerDatabaseIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched googleSpannerDatabaseIamBinding.
func (c *FakeGoogleSpannerDatabaseIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleSpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlespannerdatabaseiambindingsResource, name, pt, data, subresources...), &v1alpha1.GoogleSpannerDatabaseIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleSpannerDatabaseIamBinding), err
}
