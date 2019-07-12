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

// FakeGoogleFolderIamMembers implements GoogleFolderIamMemberInterface
type FakeGoogleFolderIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var googlefolderiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlefolderiammembers"}

var googlefolderiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleFolderIamMember"}

// Get takes name of the googleFolderIamMember, and returns the corresponding googleFolderIamMember object, and an error if there is any.
func (c *FakeGoogleFolderIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleFolderIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlefolderiammembersResource, name), &v1alpha1.GoogleFolderIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderIamMember), err
}

// List takes label and field selectors, and returns the list of GoogleFolderIamMembers that match those selectors.
func (c *FakeGoogleFolderIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleFolderIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlefolderiammembersResource, googlefolderiammembersKind, opts), &v1alpha1.GoogleFolderIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleFolderIamMemberList{ListMeta: obj.(*v1alpha1.GoogleFolderIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleFolderIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleFolderIamMembers.
func (c *FakeGoogleFolderIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlefolderiammembersResource, opts))
}

// Create takes the representation of a googleFolderIamMember and creates it.  Returns the server's representation of the googleFolderIamMember, and an error, if there is any.
func (c *FakeGoogleFolderIamMembers) Create(googleFolderIamMember *v1alpha1.GoogleFolderIamMember) (result *v1alpha1.GoogleFolderIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlefolderiammembersResource, googleFolderIamMember), &v1alpha1.GoogleFolderIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderIamMember), err
}

// Update takes the representation of a googleFolderIamMember and updates it. Returns the server's representation of the googleFolderIamMember, and an error, if there is any.
func (c *FakeGoogleFolderIamMembers) Update(googleFolderIamMember *v1alpha1.GoogleFolderIamMember) (result *v1alpha1.GoogleFolderIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlefolderiammembersResource, googleFolderIamMember), &v1alpha1.GoogleFolderIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleFolderIamMembers) UpdateStatus(googleFolderIamMember *v1alpha1.GoogleFolderIamMember) (*v1alpha1.GoogleFolderIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlefolderiammembersResource, "status", googleFolderIamMember), &v1alpha1.GoogleFolderIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderIamMember), err
}

// Delete takes name of the googleFolderIamMember and deletes it. Returns an error if one occurs.
func (c *FakeGoogleFolderIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlefolderiammembersResource, name), &v1alpha1.GoogleFolderIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleFolderIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlefolderiammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleFolderIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched googleFolderIamMember.
func (c *FakeGoogleFolderIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleFolderIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlefolderiammembersResource, name, pt, data, subresources...), &v1alpha1.GoogleFolderIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFolderIamMember), err
}
