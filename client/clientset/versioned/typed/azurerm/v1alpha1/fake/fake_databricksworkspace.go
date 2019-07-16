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

// FakeDatabricksWorkspaces implements DatabricksWorkspaceInterface
type FakeDatabricksWorkspaces struct {
	Fake *FakeAzurermV1alpha1
}

var databricksworkspacesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "databricksworkspaces"}

var databricksworkspacesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DatabricksWorkspace"}

// Get takes name of the databricksWorkspace, and returns the corresponding databricksWorkspace object, and an error if there is any.
func (c *FakeDatabricksWorkspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.DatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(databricksworkspacesResource, name), &v1alpha1.DatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabricksWorkspace), err
}

// List takes label and field selectors, and returns the list of DatabricksWorkspaces that match those selectors.
func (c *FakeDatabricksWorkspaces) List(opts v1.ListOptions) (result *v1alpha1.DatabricksWorkspaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(databricksworkspacesResource, databricksworkspacesKind, opts), &v1alpha1.DatabricksWorkspaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatabricksWorkspaceList{ListMeta: obj.(*v1alpha1.DatabricksWorkspaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatabricksWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested databricksWorkspaces.
func (c *FakeDatabricksWorkspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(databricksworkspacesResource, opts))
}

// Create takes the representation of a databricksWorkspace and creates it.  Returns the server's representation of the databricksWorkspace, and an error, if there is any.
func (c *FakeDatabricksWorkspaces) Create(databricksWorkspace *v1alpha1.DatabricksWorkspace) (result *v1alpha1.DatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(databricksworkspacesResource, databricksWorkspace), &v1alpha1.DatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabricksWorkspace), err
}

// Update takes the representation of a databricksWorkspace and updates it. Returns the server's representation of the databricksWorkspace, and an error, if there is any.
func (c *FakeDatabricksWorkspaces) Update(databricksWorkspace *v1alpha1.DatabricksWorkspace) (result *v1alpha1.DatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(databricksworkspacesResource, databricksWorkspace), &v1alpha1.DatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabricksWorkspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatabricksWorkspaces) UpdateStatus(databricksWorkspace *v1alpha1.DatabricksWorkspace) (*v1alpha1.DatabricksWorkspace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(databricksworkspacesResource, "status", databricksWorkspace), &v1alpha1.DatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabricksWorkspace), err
}

// Delete takes name of the databricksWorkspace and deletes it. Returns an error if one occurs.
func (c *FakeDatabricksWorkspaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(databricksworkspacesResource, name), &v1alpha1.DatabricksWorkspace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatabricksWorkspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(databricksworkspacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatabricksWorkspaceList{})
	return err
}

// Patch applies the patch and returns the patched databricksWorkspace.
func (c *FakeDatabricksWorkspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatabricksWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(databricksworkspacesResource, name, pt, data, subresources...), &v1alpha1.DatabricksWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatabricksWorkspace), err
}
