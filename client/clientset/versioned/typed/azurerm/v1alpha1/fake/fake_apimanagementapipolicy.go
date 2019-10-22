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

// FakeApiManagementAPIPolicies implements ApiManagementAPIPolicyInterface
type FakeApiManagementAPIPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementapipoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementapipolicies"}

var apimanagementapipoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementAPIPolicy"}

// Get takes name of the apiManagementAPIPolicy, and returns the corresponding apiManagementAPIPolicy object, and an error if there is any.
func (c *FakeApiManagementAPIPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementAPIPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementapipoliciesResource, c.ns, name), &v1alpha1.ApiManagementAPIPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIPolicy), err
}

// List takes label and field selectors, and returns the list of ApiManagementAPIPolicies that match those selectors.
func (c *FakeApiManagementAPIPolicies) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementAPIPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementapipoliciesResource, apimanagementapipoliciesKind, c.ns, opts), &v1alpha1.ApiManagementAPIPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementAPIPolicyList{ListMeta: obj.(*v1alpha1.ApiManagementAPIPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementAPIPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementAPIPolicies.
func (c *FakeApiManagementAPIPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementapipoliciesResource, c.ns, opts))

}

// Create takes the representation of a apiManagementAPIPolicy and creates it.  Returns the server's representation of the apiManagementAPIPolicy, and an error, if there is any.
func (c *FakeApiManagementAPIPolicies) Create(apiManagementAPIPolicy *v1alpha1.ApiManagementAPIPolicy) (result *v1alpha1.ApiManagementAPIPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementapipoliciesResource, c.ns, apiManagementAPIPolicy), &v1alpha1.ApiManagementAPIPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIPolicy), err
}

// Update takes the representation of a apiManagementAPIPolicy and updates it. Returns the server's representation of the apiManagementAPIPolicy, and an error, if there is any.
func (c *FakeApiManagementAPIPolicies) Update(apiManagementAPIPolicy *v1alpha1.ApiManagementAPIPolicy) (result *v1alpha1.ApiManagementAPIPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementapipoliciesResource, c.ns, apiManagementAPIPolicy), &v1alpha1.ApiManagementAPIPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementAPIPolicies) UpdateStatus(apiManagementAPIPolicy *v1alpha1.ApiManagementAPIPolicy) (*v1alpha1.ApiManagementAPIPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementapipoliciesResource, "status", c.ns, apiManagementAPIPolicy), &v1alpha1.ApiManagementAPIPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIPolicy), err
}

// Delete takes name of the apiManagementAPIPolicy and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementAPIPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementapipoliciesResource, c.ns, name), &v1alpha1.ApiManagementAPIPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementAPIPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementapipoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementAPIPolicyList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementAPIPolicy.
func (c *FakeApiManagementAPIPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAPIPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementapipoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementAPIPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIPolicy), err
}
