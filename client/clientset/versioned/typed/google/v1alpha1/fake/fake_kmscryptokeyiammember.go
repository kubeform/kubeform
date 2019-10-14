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

// FakeKmsCryptoKeyIamMembers implements KmsCryptoKeyIamMemberInterface
type FakeKmsCryptoKeyIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var kmscryptokeyiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "kmscryptokeyiammembers"}

var kmscryptokeyiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "KmsCryptoKeyIamMember"}

// Get takes name of the kmsCryptoKeyIamMember, and returns the corresponding kmsCryptoKeyIamMember object, and an error if there is any.
func (c *FakeKmsCryptoKeyIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kmscryptokeyiammembersResource, c.ns, name), &v1alpha1.KmsCryptoKeyIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCryptoKeyIamMember), err
}

// List takes label and field selectors, and returns the list of KmsCryptoKeyIamMembers that match those selectors.
func (c *FakeKmsCryptoKeyIamMembers) List(opts v1.ListOptions) (result *v1alpha1.KmsCryptoKeyIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kmscryptokeyiammembersResource, kmscryptokeyiammembersKind, c.ns, opts), &v1alpha1.KmsCryptoKeyIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KmsCryptoKeyIamMemberList{ListMeta: obj.(*v1alpha1.KmsCryptoKeyIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.KmsCryptoKeyIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kmsCryptoKeyIamMembers.
func (c *FakeKmsCryptoKeyIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kmscryptokeyiammembersResource, c.ns, opts))

}

// Create takes the representation of a kmsCryptoKeyIamMember and creates it.  Returns the server's representation of the kmsCryptoKeyIamMember, and an error, if there is any.
func (c *FakeKmsCryptoKeyIamMembers) Create(kmsCryptoKeyIamMember *v1alpha1.KmsCryptoKeyIamMember) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kmscryptokeyiammembersResource, c.ns, kmsCryptoKeyIamMember), &v1alpha1.KmsCryptoKeyIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCryptoKeyIamMember), err
}

// Update takes the representation of a kmsCryptoKeyIamMember and updates it. Returns the server's representation of the kmsCryptoKeyIamMember, and an error, if there is any.
func (c *FakeKmsCryptoKeyIamMembers) Update(kmsCryptoKeyIamMember *v1alpha1.KmsCryptoKeyIamMember) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kmscryptokeyiammembersResource, c.ns, kmsCryptoKeyIamMember), &v1alpha1.KmsCryptoKeyIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCryptoKeyIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKmsCryptoKeyIamMembers) UpdateStatus(kmsCryptoKeyIamMember *v1alpha1.KmsCryptoKeyIamMember) (*v1alpha1.KmsCryptoKeyIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kmscryptokeyiammembersResource, "status", c.ns, kmsCryptoKeyIamMember), &v1alpha1.KmsCryptoKeyIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCryptoKeyIamMember), err
}

// Delete takes name of the kmsCryptoKeyIamMember and deletes it. Returns an error if one occurs.
func (c *FakeKmsCryptoKeyIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kmscryptokeyiammembersResource, c.ns, name), &v1alpha1.KmsCryptoKeyIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKmsCryptoKeyIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kmscryptokeyiammembersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KmsCryptoKeyIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched kmsCryptoKeyIamMember.
func (c *FakeKmsCryptoKeyIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCryptoKeyIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kmscryptokeyiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.KmsCryptoKeyIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsCryptoKeyIamMember), err
}
