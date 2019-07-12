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

// FakeAzurermDatabricksWorkspaces implements AzurermDatabricksWorkspaceInterface
type FakeAzurermDatabricksWorkspaces struct {
	Fake *FakeAzurermV1alpha1
}

var azurermdatabricksworkspacesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermdatabricksworkspaces"}

var azurermdatabricksworkspacesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermDatabricksWorkspace"}

// Get takes name of the azurermDatabricksWorkspace, and returns the corresponding azurermDatabricksWorkspace object, and an error if there is any.
func (c *FakeAzurermDatabricksWorkspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermdatabricksworkspacesResource, name), &v1alpha1.AzurermDatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDatabricksWorkspace), err
}

// List takes label and field selectors, and returns the list of AzurermDatabricksWorkspaces that match those selectors.
func (c *FakeAzurermDatabricksWorkspaces) List(opts v1.ListOptions) (result *v1alpha1.AzurermDatabricksWorkspaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermdatabricksworkspacesResource, azurermdatabricksworkspacesKind, opts), &v1alpha1.AzurermDatabricksWorkspaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermDatabricksWorkspaceList{ListMeta: obj.(*v1alpha1.AzurermDatabricksWorkspaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermDatabricksWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermDatabricksWorkspaces.
func (c *FakeAzurermDatabricksWorkspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermdatabricksworkspacesResource, opts))
}

// Create takes the representation of a azurermDatabricksWorkspace and creates it.  Returns the server's representation of the azurermDatabricksWorkspace, and an error, if there is any.
func (c *FakeAzurermDatabricksWorkspaces) Create(azurermDatabricksWorkspace *v1alpha1.AzurermDatabricksWorkspace) (result *v1alpha1.AzurermDatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermdatabricksworkspacesResource, azurermDatabricksWorkspace), &v1alpha1.AzurermDatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDatabricksWorkspace), err
}

// Update takes the representation of a azurermDatabricksWorkspace and updates it. Returns the server's representation of the azurermDatabricksWorkspace, and an error, if there is any.
func (c *FakeAzurermDatabricksWorkspaces) Update(azurermDatabricksWorkspace *v1alpha1.AzurermDatabricksWorkspace) (result *v1alpha1.AzurermDatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermdatabricksworkspacesResource, azurermDatabricksWorkspace), &v1alpha1.AzurermDatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDatabricksWorkspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermDatabricksWorkspaces) UpdateStatus(azurermDatabricksWorkspace *v1alpha1.AzurermDatabricksWorkspace) (*v1alpha1.AzurermDatabricksWorkspace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermdatabricksworkspacesResource, "status", azurermDatabricksWorkspace), &v1alpha1.AzurermDatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDatabricksWorkspace), err
}

// Delete takes name of the azurermDatabricksWorkspace and deletes it. Returns an error if one occurs.
func (c *FakeAzurermDatabricksWorkspaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermdatabricksworkspacesResource, name), &v1alpha1.AzurermDatabricksWorkspace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermDatabricksWorkspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermdatabricksworkspacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermDatabricksWorkspaceList{})
	return err
}

// Patch applies the patch and returns the patched azurermDatabricksWorkspace.
func (c *FakeAzurermDatabricksWorkspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermdatabricksworkspacesResource, name, pt, data, subresources...), &v1alpha1.AzurermDatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermDatabricksWorkspace), err
}
