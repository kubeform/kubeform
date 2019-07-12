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

// FakeGoogleOrganizationIamMembers implements GoogleOrganizationIamMemberInterface
type FakeGoogleOrganizationIamMembers struct {
	Fake *FakeGoogleV1alpha1
}

var googleorganizationiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googleorganizationiammembers"}

var googleorganizationiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleOrganizationIamMember"}

// Get takes name of the googleOrganizationIamMember, and returns the corresponding googleOrganizationIamMember object, and an error if there is any.
func (c *FakeGoogleOrganizationIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleOrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googleorganizationiammembersResource, name), &v1alpha1.GoogleOrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleOrganizationIamMember), err
}

// List takes label and field selectors, and returns the list of GoogleOrganizationIamMembers that match those selectors.
func (c *FakeGoogleOrganizationIamMembers) List(opts v1.ListOptions) (result *v1alpha1.GoogleOrganizationIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googleorganizationiammembersResource, googleorganizationiammembersKind, opts), &v1alpha1.GoogleOrganizationIamMemberList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleOrganizationIamMemberList{ListMeta: obj.(*v1alpha1.GoogleOrganizationIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleOrganizationIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleOrganizationIamMembers.
func (c *FakeGoogleOrganizationIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googleorganizationiammembersResource, opts))
}

// Create takes the representation of a googleOrganizationIamMember and creates it.  Returns the server's representation of the googleOrganizationIamMember, and an error, if there is any.
func (c *FakeGoogleOrganizationIamMembers) Create(googleOrganizationIamMember *v1alpha1.GoogleOrganizationIamMember) (result *v1alpha1.GoogleOrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googleorganizationiammembersResource, googleOrganizationIamMember), &v1alpha1.GoogleOrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleOrganizationIamMember), err
}

// Update takes the representation of a googleOrganizationIamMember and updates it. Returns the server's representation of the googleOrganizationIamMember, and an error, if there is any.
func (c *FakeGoogleOrganizationIamMembers) Update(googleOrganizationIamMember *v1alpha1.GoogleOrganizationIamMember) (result *v1alpha1.GoogleOrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googleorganizationiammembersResource, googleOrganizationIamMember), &v1alpha1.GoogleOrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleOrganizationIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleOrganizationIamMembers) UpdateStatus(googleOrganizationIamMember *v1alpha1.GoogleOrganizationIamMember) (*v1alpha1.GoogleOrganizationIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googleorganizationiammembersResource, "status", googleOrganizationIamMember), &v1alpha1.GoogleOrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleOrganizationIamMember), err
}

// Delete takes name of the googleOrganizationIamMember and deletes it. Returns an error if one occurs.
func (c *FakeGoogleOrganizationIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googleorganizationiammembersResource, name), &v1alpha1.GoogleOrganizationIamMember{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleOrganizationIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googleorganizationiammembersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleOrganizationIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched googleOrganizationIamMember.
func (c *FakeGoogleOrganizationIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleOrganizationIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googleorganizationiammembersResource, name, pt, data, subresources...), &v1alpha1.GoogleOrganizationIamMember{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleOrganizationIamMember), err
}
