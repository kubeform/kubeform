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

// FakeAzurermApiManagementGroupUsers implements AzurermApiManagementGroupUserInterface
type FakeAzurermApiManagementGroupUsers struct {
	Fake *FakeAzurermV1alpha1
}

var azurermapimanagementgroupusersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermapimanagementgroupusers"}

var azurermapimanagementgroupusersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermApiManagementGroupUser"}

// Get takes name of the azurermApiManagementGroupUser, and returns the corresponding azurermApiManagementGroupUser object, and an error if there is any.
func (c *FakeAzurermApiManagementGroupUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermapimanagementgroupusersResource, name), &v1alpha1.AzurermApiManagementGroupUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementGroupUser), err
}

// List takes label and field selectors, and returns the list of AzurermApiManagementGroupUsers that match those selectors.
func (c *FakeAzurermApiManagementGroupUsers) List(opts v1.ListOptions) (result *v1alpha1.AzurermApiManagementGroupUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermapimanagementgroupusersResource, azurermapimanagementgroupusersKind, opts), &v1alpha1.AzurermApiManagementGroupUserList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermApiManagementGroupUserList{ListMeta: obj.(*v1alpha1.AzurermApiManagementGroupUserList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermApiManagementGroupUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermApiManagementGroupUsers.
func (c *FakeAzurermApiManagementGroupUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermapimanagementgroupusersResource, opts))
}

// Create takes the representation of a azurermApiManagementGroupUser and creates it.  Returns the server's representation of the azurermApiManagementGroupUser, and an error, if there is any.
func (c *FakeAzurermApiManagementGroupUsers) Create(azurermApiManagementGroupUser *v1alpha1.AzurermApiManagementGroupUser) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermapimanagementgroupusersResource, azurermApiManagementGroupUser), &v1alpha1.AzurermApiManagementGroupUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementGroupUser), err
}

// Update takes the representation of a azurermApiManagementGroupUser and updates it. Returns the server's representation of the azurermApiManagementGroupUser, and an error, if there is any.
func (c *FakeAzurermApiManagementGroupUsers) Update(azurermApiManagementGroupUser *v1alpha1.AzurermApiManagementGroupUser) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermapimanagementgroupusersResource, azurermApiManagementGroupUser), &v1alpha1.AzurermApiManagementGroupUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementGroupUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermApiManagementGroupUsers) UpdateStatus(azurermApiManagementGroupUser *v1alpha1.AzurermApiManagementGroupUser) (*v1alpha1.AzurermApiManagementGroupUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermapimanagementgroupusersResource, "status", azurermApiManagementGroupUser), &v1alpha1.AzurermApiManagementGroupUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementGroupUser), err
}

// Delete takes name of the azurermApiManagementGroupUser and deletes it. Returns an error if one occurs.
func (c *FakeAzurermApiManagementGroupUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermapimanagementgroupusersResource, name), &v1alpha1.AzurermApiManagementGroupUser{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermApiManagementGroupUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermapimanagementgroupusersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermApiManagementGroupUserList{})
	return err
}

// Patch applies the patch and returns the patched azurermApiManagementGroupUser.
func (c *FakeAzurermApiManagementGroupUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApiManagementGroupUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermapimanagementgroupusersResource, name, pt, data, subresources...), &v1alpha1.AzurermApiManagementGroupUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApiManagementGroupUser), err
}
