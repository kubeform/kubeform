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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermRoleAssignments implements AzurermRoleAssignmentInterface
type FakeAzurermRoleAssignments struct {
	Fake *FakeAzurermV1alpha1
}

var azurermroleassignmentsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermroleassignments"}

var azurermroleassignmentsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermRoleAssignment"}

// Get takes name of the azurermRoleAssignment, and returns the corresponding azurermRoleAssignment object, and an error if there is any.
func (c *FakeAzurermRoleAssignments) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermRoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermroleassignmentsResource, name), &v1alpha1.AzurermRoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermRoleAssignment), err
}

// List takes label and field selectors, and returns the list of AzurermRoleAssignments that match those selectors.
func (c *FakeAzurermRoleAssignments) List(opts v1.ListOptions) (result *v1alpha1.AzurermRoleAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermroleassignmentsResource, azurermroleassignmentsKind, opts), &v1alpha1.AzurermRoleAssignmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermRoleAssignmentList{ListMeta: obj.(*v1alpha1.AzurermRoleAssignmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermRoleAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermRoleAssignments.
func (c *FakeAzurermRoleAssignments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermroleassignmentsResource, opts))
}

// Create takes the representation of a azurermRoleAssignment and creates it.  Returns the server's representation of the azurermRoleAssignment, and an error, if there is any.
func (c *FakeAzurermRoleAssignments) Create(azurermRoleAssignment *v1alpha1.AzurermRoleAssignment) (result *v1alpha1.AzurermRoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermroleassignmentsResource, azurermRoleAssignment), &v1alpha1.AzurermRoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermRoleAssignment), err
}

// Update takes the representation of a azurermRoleAssignment and updates it. Returns the server's representation of the azurermRoleAssignment, and an error, if there is any.
func (c *FakeAzurermRoleAssignments) Update(azurermRoleAssignment *v1alpha1.AzurermRoleAssignment) (result *v1alpha1.AzurermRoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermroleassignmentsResource, azurermRoleAssignment), &v1alpha1.AzurermRoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermRoleAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermRoleAssignments) UpdateStatus(azurermRoleAssignment *v1alpha1.AzurermRoleAssignment) (*v1alpha1.AzurermRoleAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermroleassignmentsResource, "status", azurermRoleAssignment), &v1alpha1.AzurermRoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermRoleAssignment), err
}

// Delete takes name of the azurermRoleAssignment and deletes it. Returns an error if one occurs.
func (c *FakeAzurermRoleAssignments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermroleassignmentsResource, name), &v1alpha1.AzurermRoleAssignment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermRoleAssignments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermroleassignmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermRoleAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched azurermRoleAssignment.
func (c *FakeAzurermRoleAssignments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermRoleAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermroleassignmentsResource, name, pt, data, subresources...), &v1alpha1.AzurermRoleAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermRoleAssignment), err
}
