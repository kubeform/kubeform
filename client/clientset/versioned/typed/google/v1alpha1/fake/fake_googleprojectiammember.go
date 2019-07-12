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

// FakeGoogleProjectIamMembers implements GoogleProjectIamMemberInterface
type FakeGoogleProjectIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var googleprojectiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googleprojectiammembers"}

var googleprojectiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleProjectIamMember"}

// Get takes name of the googleProjectIamMember, and returns the corresponding googleProjectIamMember object, and an error if there is any.
func (c *FakeGoogleProjectIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleProjectIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googleprojectiammembersResource, name), &v1alpha1.GoogleProjectIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleProjectIamMember), err
}

// List takes label and field selectors, and returns the list of GoogleProjectIamMembers that match those selectors.
func (c *FakeGoogleProjectIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleProjectIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googleprojectiammembersResource, googleprojectiammembersKind, opts), &v1alpha1.GoogleProjectIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleProjectIamMemberList{ListMeta: obj.(*v1alpha1.GoogleProjectIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleProjectIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleProjectIamMembers.
func (c *FakeGoogleProjectIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googleprojectiammembersResource, opts))
}

// Create takes the representation of a googleProjectIamMember and creates it.  Returns the server's representation of the googleProjectIamMember, and an error, if there is any.
func (c *FakeGoogleProjectIamMembers) Create(googleProjectIamMember *v1alpha1.GoogleProjectIamMember) (result *v1alpha1.GoogleProjectIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googleprojectiammembersResource, googleProjectIamMember), &v1alpha1.GoogleProjectIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleProjectIamMember), err
}

// Update takes the representation of a googleProjectIamMember and updates it. Returns the server's representation of the googleProjectIamMember, and an error, if there is any.
func (c *FakeGoogleProjectIamMembers) Update(googleProjectIamMember *v1alpha1.GoogleProjectIamMember) (result *v1alpha1.GoogleProjectIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googleprojectiammembersResource, googleProjectIamMember), &v1alpha1.GoogleProjectIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleProjectIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleProjectIamMembers) UpdateStatus(googleProjectIamMember *v1alpha1.GoogleProjectIamMember) (*v1alpha1.GoogleProjectIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googleprojectiammembersResource, "status", googleProjectIamMember), &v1alpha1.GoogleProjectIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleProjectIamMember), err
}

// Delete takes name of the googleProjectIamMember and deletes it. Returns an error if one occurs.
func (c *FakeGoogleProjectIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googleprojectiammembersResource, name), &v1alpha1.GoogleProjectIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleProjectIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googleprojectiammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleProjectIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched googleProjectIamMember.
func (c *FakeGoogleProjectIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleProjectIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googleprojectiammembersResource, name, pt, data, subresources...), &v1alpha1.GoogleProjectIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleProjectIamMember), err
}
