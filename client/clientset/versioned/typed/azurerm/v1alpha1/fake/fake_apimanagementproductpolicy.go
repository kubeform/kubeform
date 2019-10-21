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

// FakeApiManagementProductPolicies implements ApiManagementProductPolicyInterface
type FakeApiManagementProductPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementproductpoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementproductpolicies"}

var apimanagementproductpoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementProductPolicy"}

// Get takes name of the apiManagementProductPolicy, and returns the corresponding apiManagementProductPolicy object, and an error if there is any.
func (c *FakeApiManagementProductPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementproductpoliciesResource, c.ns, name), &v1alpha1.ApiManagementProductPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProductPolicy), err
}

// List takes label and field selectors, and returns the list of ApiManagementProductPolicies that match those selectors.
func (c *FakeApiManagementProductPolicies) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementProductPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementproductpoliciesResource, apimanagementproductpoliciesKind, c.ns, opts), &v1alpha1.ApiManagementProductPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementProductPolicyList{ListMeta: obj.(*v1alpha1.ApiManagementProductPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementProductPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementProductPolicies.
func (c *FakeApiManagementProductPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementproductpoliciesResource, c.ns, opts))

}

// Create takes the representation of a apiManagementProductPolicy and creates it.  Returns the server's representation of the apiManagementProductPolicy, and an error, if there is any.
func (c *FakeApiManagementProductPolicies) Create(apiManagementProductPolicy *v1alpha1.ApiManagementProductPolicy) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementproductpoliciesResource, c.ns, apiManagementProductPolicy), &v1alpha1.ApiManagementProductPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProductPolicy), err
}

// Update takes the representation of a apiManagementProductPolicy and updates it. Returns the server's representation of the apiManagementProductPolicy, and an error, if there is any.
func (c *FakeApiManagementProductPolicies) Update(apiManagementProductPolicy *v1alpha1.ApiManagementProductPolicy) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementproductpoliciesResource, c.ns, apiManagementProductPolicy), &v1alpha1.ApiManagementProductPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProductPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementProductPolicies) UpdateStatus(apiManagementProductPolicy *v1alpha1.ApiManagementProductPolicy) (*v1alpha1.ApiManagementProductPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementproductpoliciesResource, "status", c.ns, apiManagementProductPolicy), &v1alpha1.ApiManagementProductPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProductPolicy), err
}

// Delete takes name of the apiManagementProductPolicy and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementProductPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementproductpoliciesResource, c.ns, name), &v1alpha1.ApiManagementProductPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementProductPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementproductpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementProductPolicyList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementProductPolicy.
func (c *FakeApiManagementProductPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementProductPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementproductpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementProductPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementProductPolicy), err
}