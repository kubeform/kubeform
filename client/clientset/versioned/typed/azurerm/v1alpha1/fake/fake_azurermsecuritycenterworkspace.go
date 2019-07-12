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

// FakeAzurermSecurityCenterWorkspaces implements AzurermSecurityCenterWorkspaceInterface
type FakeAzurermSecurityCenterWorkspaces struct {
	Fake *FakeAzurermV1alpha1
}

var azurermsecuritycenterworkspacesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermsecuritycenterworkspaces"}

var azurermsecuritycenterworkspacesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermSecurityCenterWorkspace"}

// Get takes name of the azurermSecurityCenterWorkspace, and returns the corresponding azurermSecurityCenterWorkspace object, and an error if there is any.
func (c *FakeAzurermSecurityCenterWorkspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermSecurityCenterWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermsecuritycenterworkspacesResource, name), &v1alpha1.AzurermSecurityCenterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSecurityCenterWorkspace), err
}

// List takes label and field selectors, and returns the list of AzurermSecurityCenterWorkspaces that match those selectors.
func (c *FakeAzurermSecurityCenterWorkspaces) List(opts v1.ListOptions) (result *v1alpha1.AzurermSecurityCenterWorkspaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermsecuritycenterworkspacesResource, azurermsecuritycenterworkspacesKind, opts), &v1alpha1.AzurermSecurityCenterWorkspaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermSecurityCenterWorkspaceList{ListMeta: obj.(*v1alpha1.AzurermSecurityCenterWorkspaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermSecurityCenterWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermSecurityCenterWorkspaces.
func (c *FakeAzurermSecurityCenterWorkspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermsecuritycenterworkspacesResource, opts))
}

// Create takes the representation of a azurermSecurityCenterWorkspace and creates it.  Returns the server's representation of the azurermSecurityCenterWorkspace, and an error, if there is any.
func (c *FakeAzurermSecurityCenterWorkspaces) Create(azurermSecurityCenterWorkspace *v1alpha1.AzurermSecurityCenterWorkspace) (result *v1alpha1.AzurermSecurityCenterWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermsecuritycenterworkspacesResource, azurermSecurityCenterWorkspace), &v1alpha1.AzurermSecurityCenterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSecurityCenterWorkspace), err
}

// Update takes the representation of a azurermSecurityCenterWorkspace and updates it. Returns the server's representation of the azurermSecurityCenterWorkspace, and an error, if there is any.
func (c *FakeAzurermSecurityCenterWorkspaces) Update(azurermSecurityCenterWorkspace *v1alpha1.AzurermSecurityCenterWorkspace) (result *v1alpha1.AzurermSecurityCenterWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermsecuritycenterworkspacesResource, azurermSecurityCenterWorkspace), &v1alpha1.AzurermSecurityCenterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSecurityCenterWorkspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermSecurityCenterWorkspaces) UpdateStatus(azurermSecurityCenterWorkspace *v1alpha1.AzurermSecurityCenterWorkspace) (*v1alpha1.AzurermSecurityCenterWorkspace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermsecuritycenterworkspacesResource, "status", azurermSecurityCenterWorkspace), &v1alpha1.AzurermSecurityCenterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSecurityCenterWorkspace), err
}

// Delete takes name of the azurermSecurityCenterWorkspace and deletes it. Returns an error if one occurs.
func (c *FakeAzurermSecurityCenterWorkspaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermsecuritycenterworkspacesResource, name), &v1alpha1.AzurermSecurityCenterWorkspace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermSecurityCenterWorkspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermsecuritycenterworkspacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermSecurityCenterWorkspaceList{})
	return err
}

// Patch applies the patch and returns the patched azurermSecurityCenterWorkspace.
func (c *FakeAzurermSecurityCenterWorkspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermSecurityCenterWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermsecuritycenterworkspacesResource, name, pt, data, subresources...), &v1alpha1.AzurermSecurityCenterWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermSecurityCenterWorkspace), err
}
