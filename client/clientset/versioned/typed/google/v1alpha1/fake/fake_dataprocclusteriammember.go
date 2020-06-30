/*
Copyright The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataprocClusterIamMembers implements DataprocClusterIamMemberInterface
type FakeDataprocClusterIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var dataprocclusteriammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "dataprocclusteriammembers"}

var dataprocclusteriammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "DataprocClusterIamMember"}

// Get takes name of the dataprocClusterIamMember, and returns the corresponding dataprocClusterIamMember object, and an error if there is any.
func (c *FakeDataprocClusterIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.DataprocClusterIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataprocclusteriammembersResource, c.ns, name), &v1alpha1.DataprocClusterIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamMember), err
}

// List takes label and field selectors, and returns the list of DataprocClusterIamMembers that match those selectors.
func (c *FakeDataprocClusterIamMembers) List(opts v1.ListOptions) (result *v1alpha1.DataprocClusterIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataprocclusteriammembersResource, dataprocclusteriammembersKind, c.ns, opts), &v1alpha1.DataprocClusterIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataprocClusterIamMemberList{ListMeta: obj.(*v1alpha1.DataprocClusterIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataprocClusterIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataprocClusterIamMembers.
func (c *FakeDataprocClusterIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataprocclusteriammembersResource, c.ns, opts))

}

// Create takes the representation of a dataprocClusterIamMember and creates it.  Returns the server's representation of the dataprocClusterIamMember, and an error, if there is any.
func (c *FakeDataprocClusterIamMembers) Create(dataprocClusterIamMember *v1alpha1.DataprocClusterIamMember) (result *v1alpha1.DataprocClusterIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataprocclusteriammembersResource, c.ns, dataprocClusterIamMember), &v1alpha1.DataprocClusterIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamMember), err
}

// Update takes the representation of a dataprocClusterIamMember and updates it. Returns the server's representation of the dataprocClusterIamMember, and an error, if there is any.
func (c *FakeDataprocClusterIamMembers) Update(dataprocClusterIamMember *v1alpha1.DataprocClusterIamMember) (result *v1alpha1.DataprocClusterIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataprocclusteriammembersResource, c.ns, dataprocClusterIamMember), &v1alpha1.DataprocClusterIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataprocClusterIamMembers) UpdateStatus(dataprocClusterIamMember *v1alpha1.DataprocClusterIamMember) (*v1alpha1.DataprocClusterIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataprocclusteriammembersResource, "status", c.ns, dataprocClusterIamMember), &v1alpha1.DataprocClusterIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamMember), err
}

// Delete takes name of the dataprocClusterIamMember and deletes it. Returns an error if one occurs.
func (c *FakeDataprocClusterIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataprocclusteriammembersResource, c.ns, name), &v1alpha1.DataprocClusterIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataprocClusterIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataprocclusteriammembersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataprocClusterIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched dataprocClusterIamMember.
func (c *FakeDataprocClusterIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataprocClusterIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataprocclusteriammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataprocClusterIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataprocClusterIamMember), err
}
