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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiManagementAPIOperations implements ApiManagementAPIOperationInterface
type FakeApiManagementAPIOperations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementapioperationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementapioperations"}

var apimanagementapioperationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementAPIOperation"}

// Get takes name of the apiManagementAPIOperation, and returns the corresponding apiManagementAPIOperation object, and an error if there is any.
func (c *FakeApiManagementAPIOperations) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementAPIOperation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementapioperationsResource, c.ns, name), &v1alpha1.ApiManagementAPIOperation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperation), err
}

// List takes label and field selectors, and returns the list of ApiManagementAPIOperations that match those selectors.
func (c *FakeApiManagementAPIOperations) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementAPIOperationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementapioperationsResource, apimanagementapioperationsKind, c.ns, opts), &v1alpha1.ApiManagementAPIOperationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementAPIOperationList{ListMeta: obj.(*v1alpha1.ApiManagementAPIOperationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementAPIOperationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementAPIOperations.
func (c *FakeApiManagementAPIOperations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementapioperationsResource, c.ns, opts))

}

// Create takes the representation of a apiManagementAPIOperation and creates it.  Returns the server's representation of the apiManagementAPIOperation, and an error, if there is any.
func (c *FakeApiManagementAPIOperations) Create(apiManagementAPIOperation *v1alpha1.ApiManagementAPIOperation) (result *v1alpha1.ApiManagementAPIOperation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementapioperationsResource, c.ns, apiManagementAPIOperation), &v1alpha1.ApiManagementAPIOperation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperation), err
}

// Update takes the representation of a apiManagementAPIOperation and updates it. Returns the server's representation of the apiManagementAPIOperation, and an error, if there is any.
func (c *FakeApiManagementAPIOperations) Update(apiManagementAPIOperation *v1alpha1.ApiManagementAPIOperation) (result *v1alpha1.ApiManagementAPIOperation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementapioperationsResource, c.ns, apiManagementAPIOperation), &v1alpha1.ApiManagementAPIOperation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementAPIOperations) UpdateStatus(apiManagementAPIOperation *v1alpha1.ApiManagementAPIOperation) (*v1alpha1.ApiManagementAPIOperation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementapioperationsResource, "status", c.ns, apiManagementAPIOperation), &v1alpha1.ApiManagementAPIOperation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperation), err
}

// Delete takes name of the apiManagementAPIOperation and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementAPIOperations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementapioperationsResource, c.ns, name), &v1alpha1.ApiManagementAPIOperation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementAPIOperations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementapioperationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementAPIOperationList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementAPIOperation.
func (c *FakeApiManagementAPIOperations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAPIOperation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementapioperationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementAPIOperation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperation), err
}
