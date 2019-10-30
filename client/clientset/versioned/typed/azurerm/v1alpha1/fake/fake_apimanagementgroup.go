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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiManagementGroups implements ApiManagementGroupInterface
type FakeApiManagementGroups struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementgroupsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementgroups"}

var apimanagementgroupsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementGroup"}

// Get takes name of the apiManagementGroup, and returns the corresponding apiManagementGroup object, and an error if there is any.
func (c *FakeApiManagementGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementgroupsResource, c.ns, name), &v1alpha1.ApiManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroup), err
}

// List takes label and field selectors, and returns the list of ApiManagementGroups that match those selectors.
func (c *FakeApiManagementGroups) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementgroupsResource, apimanagementgroupsKind, c.ns, opts), &v1alpha1.ApiManagementGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementGroupList{ListMeta: obj.(*v1alpha1.ApiManagementGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementGroups.
func (c *FakeApiManagementGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementgroupsResource, c.ns, opts))

}

// Create takes the representation of a apiManagementGroup and creates it.  Returns the server's representation of the apiManagementGroup, and an error, if there is any.
func (c *FakeApiManagementGroups) Create(apiManagementGroup *v1alpha1.ApiManagementGroup) (result *v1alpha1.ApiManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementgroupsResource, c.ns, apiManagementGroup), &v1alpha1.ApiManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroup), err
}

// Update takes the representation of a apiManagementGroup and updates it. Returns the server's representation of the apiManagementGroup, and an error, if there is any.
func (c *FakeApiManagementGroups) Update(apiManagementGroup *v1alpha1.ApiManagementGroup) (result *v1alpha1.ApiManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementgroupsResource, c.ns, apiManagementGroup), &v1alpha1.ApiManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementGroups) UpdateStatus(apiManagementGroup *v1alpha1.ApiManagementGroup) (*v1alpha1.ApiManagementGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementgroupsResource, "status", c.ns, apiManagementGroup), &v1alpha1.ApiManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroup), err
}

// Delete takes name of the apiManagementGroup and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementgroupsResource, c.ns, name), &v1alpha1.ApiManagementGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementGroupList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementGroup.
func (c *FakeApiManagementGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementGroup), err
}
