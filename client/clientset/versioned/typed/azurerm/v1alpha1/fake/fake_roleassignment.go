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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeRoleAssignments implements RoleAssignmentInterface
type FakeRoleAssignments struct {
	Fake *FakeAzurermV1alpha1
}

var roleassignmentsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "roleassignments"}

var roleassignmentsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "RoleAssignment"}

// Get takes name of the roleAssignment, and returns the corresponding roleAssignment object, and an error if there is any.
func (c *FakeRoleAssignments) Get(name string, options v1.GetOptions) (result *v1alpha1.RoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(roleassignmentsResource, name), &v1alpha1.RoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAssignment), err
}

// List takes label and field selectors, and returns the list of RoleAssignments that match those selectors.
func (c *FakeRoleAssignments) List(opts v1.ListOptions) (result *v1alpha1.RoleAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(roleassignmentsResource, roleassignmentsKind, opts), &v1alpha1.RoleAssignmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RoleAssignmentList{ListMeta: obj.(*v1alpha1.RoleAssignmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.RoleAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleAssignments.
func (c *FakeRoleAssignments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(roleassignmentsResource, opts))
}

// Create takes the representation of a roleAssignment and creates it.  Returns the server's representation of the roleAssignment, and an error, if there is any.
func (c *FakeRoleAssignments) Create(roleAssignment *v1alpha1.RoleAssignment) (result *v1alpha1.RoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(roleassignmentsResource, roleAssignment), &v1alpha1.RoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAssignment), err
}

// Update takes the representation of a roleAssignment and updates it. Returns the server's representation of the roleAssignment, and an error, if there is any.
func (c *FakeRoleAssignments) Update(roleAssignment *v1alpha1.RoleAssignment) (result *v1alpha1.RoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(roleassignmentsResource, roleAssignment), &v1alpha1.RoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoleAssignments) UpdateStatus(roleAssignment *v1alpha1.RoleAssignment) (*v1alpha1.RoleAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(roleassignmentsResource, "status", roleAssignment), &v1alpha1.RoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAssignment), err
}

// Delete takes name of the roleAssignment and deletes it. Returns an error if one occurs.
func (c *FakeRoleAssignments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(roleassignmentsResource, name), &v1alpha1.RoleAssignment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoleAssignments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(roleassignmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RoleAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched roleAssignment.
func (c *FakeRoleAssignments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(roleassignmentsResource, name, pt, data, subresources...), &v1alpha1.RoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAssignment), err
}
