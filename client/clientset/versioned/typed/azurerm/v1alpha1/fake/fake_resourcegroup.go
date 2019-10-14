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

// FakeResourceGroups implements ResourceGroupInterface
type FakeResourceGroups struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var resourcegroupsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "resourcegroups"}

var resourcegroupsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ResourceGroup"}

// Get takes name of the resourceGroup, and returns the corresponding resourceGroup object, and an error if there is any.
func (c *FakeResourceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcegroupsResource, c.ns, name), &v1alpha1.ResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceGroup), err
}

// List takes label and field selectors, and returns the list of ResourceGroups that match those selectors.
func (c *FakeResourceGroups) List(opts v1.ListOptions) (result *v1alpha1.ResourceGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcegroupsResource, resourcegroupsKind, c.ns, opts), &v1alpha1.ResourceGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceGroupList{ListMeta: obj.(*v1alpha1.ResourceGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceGroups.
func (c *FakeResourceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcegroupsResource, c.ns, opts))

}

// Create takes the representation of a resourceGroup and creates it.  Returns the server's representation of the resourceGroup, and an error, if there is any.
func (c *FakeResourceGroups) Create(resourceGroup *v1alpha1.ResourceGroup) (result *v1alpha1.ResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcegroupsResource, c.ns, resourceGroup), &v1alpha1.ResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceGroup), err
}

// Update takes the representation of a resourceGroup and updates it. Returns the server's representation of the resourceGroup, and an error, if there is any.
func (c *FakeResourceGroups) Update(resourceGroup *v1alpha1.ResourceGroup) (result *v1alpha1.ResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcegroupsResource, c.ns, resourceGroup), &v1alpha1.ResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceGroups) UpdateStatus(resourceGroup *v1alpha1.ResourceGroup) (*v1alpha1.ResourceGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcegroupsResource, "status", c.ns, resourceGroup), &v1alpha1.ResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceGroup), err
}

// Delete takes name of the resourceGroup and deletes it. Returns an error if one occurs.
func (c *FakeResourceGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourcegroupsResource, c.ns, name), &v1alpha1.ResourceGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceGroupList{})
	return err
}

// Patch applies the patch and returns the patched resourceGroup.
func (c *FakeResourceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceGroup), err
}
