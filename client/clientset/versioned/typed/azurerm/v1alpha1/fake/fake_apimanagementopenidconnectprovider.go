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

// FakeApiManagementOpenidConnectProviders implements ApiManagementOpenidConnectProviderInterface
type FakeApiManagementOpenidConnectProviders struct {
	Fake *FakeAzurermV1alpha1
}

var apimanagementopenidconnectprovidersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementopenidconnectproviders"}

var apimanagementopenidconnectprovidersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementOpenidConnectProvider"}

// Get takes name of the apiManagementOpenidConnectProvider, and returns the corresponding apiManagementOpenidConnectProvider object, and an error if there is any.
func (c *FakeApiManagementOpenidConnectProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apimanagementopenidconnectprovidersResource, name), &v1alpha1.ApiManagementOpenidConnectProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementOpenidConnectProvider), err
}

// List takes label and field selectors, and returns the list of ApiManagementOpenidConnectProviders that match those selectors.
func (c *FakeApiManagementOpenidConnectProviders) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementOpenidConnectProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apimanagementopenidconnectprovidersResource, apimanagementopenidconnectprovidersKind, opts), &v1alpha1.ApiManagementOpenidConnectProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementOpenidConnectProviderList{ListMeta: obj.(*v1alpha1.ApiManagementOpenidConnectProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementOpenidConnectProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementOpenidConnectProviders.
func (c *FakeApiManagementOpenidConnectProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apimanagementopenidconnectprovidersResource, opts))
}

// Create takes the representation of a apiManagementOpenidConnectProvider and creates it.  Returns the server's representation of the apiManagementOpenidConnectProvider, and an error, if there is any.
func (c *FakeApiManagementOpenidConnectProviders) Create(apiManagementOpenidConnectProvider *v1alpha1.ApiManagementOpenidConnectProvider) (result *v1alpha1.ApiManagementOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apimanagementopenidconnectprovidersResource, apiManagementOpenidConnectProvider), &v1alpha1.ApiManagementOpenidConnectProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementOpenidConnectProvider), err
}

// Update takes the representation of a apiManagementOpenidConnectProvider and updates it. Returns the server's representation of the apiManagementOpenidConnectProvider, and an error, if there is any.
func (c *FakeApiManagementOpenidConnectProviders) Update(apiManagementOpenidConnectProvider *v1alpha1.ApiManagementOpenidConnectProvider) (result *v1alpha1.ApiManagementOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apimanagementopenidconnectprovidersResource, apiManagementOpenidConnectProvider), &v1alpha1.ApiManagementOpenidConnectProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementOpenidConnectProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementOpenidConnectProviders) UpdateStatus(apiManagementOpenidConnectProvider *v1alpha1.ApiManagementOpenidConnectProvider) (*v1alpha1.ApiManagementOpenidConnectProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apimanagementopenidconnectprovidersResource, "status", apiManagementOpenidConnectProvider), &v1alpha1.ApiManagementOpenidConnectProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementOpenidConnectProvider), err
}

// Delete takes name of the apiManagementOpenidConnectProvider and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementOpenidConnectProviders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apimanagementopenidconnectprovidersResource, name), &v1alpha1.ApiManagementOpenidConnectProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementOpenidConnectProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apimanagementopenidconnectprovidersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementOpenidConnectProviderList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementOpenidConnectProvider.
func (c *FakeApiManagementOpenidConnectProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apimanagementopenidconnectprovidersResource, name, pt, data, subresources...), &v1alpha1.ApiManagementOpenidConnectProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementOpenidConnectProvider), err
}
