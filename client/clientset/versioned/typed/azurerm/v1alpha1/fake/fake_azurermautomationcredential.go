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

// FakeAzurermAutomationCredentials implements AzurermAutomationCredentialInterface
type FakeAzurermAutomationCredentials struct {
	Fake *FakeAzurermV1alpha1
}

var azurermautomationcredentialsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermautomationcredentials"}

var azurermautomationcredentialsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermAutomationCredential"}

// Get takes name of the azurermAutomationCredential, and returns the corresponding azurermAutomationCredential object, and an error if there is any.
func (c *FakeAzurermAutomationCredentials) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermAutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermautomationcredentialsResource, name), &v1alpha1.AzurermAutomationCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationCredential), err
}

// List takes label and field selectors, and returns the list of AzurermAutomationCredentials that match those selectors.
func (c *FakeAzurermAutomationCredentials) List(opts v1.ListOptions) (result *v1alpha1.AzurermAutomationCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermautomationcredentialsResource, azurermautomationcredentialsKind, opts), &v1alpha1.AzurermAutomationCredentialList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermAutomationCredentialList{ListMeta: obj.(*v1alpha1.AzurermAutomationCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermAutomationCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermAutomationCredentials.
func (c *FakeAzurermAutomationCredentials) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermautomationcredentialsResource, opts))
}

// Create takes the representation of a azurermAutomationCredential and creates it.  Returns the server's representation of the azurermAutomationCredential, and an error, if there is any.
func (c *FakeAzurermAutomationCredentials) Create(azurermAutomationCredential *v1alpha1.AzurermAutomationCredential) (result *v1alpha1.AzurermAutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermautomationcredentialsResource, azurermAutomationCredential), &v1alpha1.AzurermAutomationCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationCredential), err
}

// Update takes the representation of a azurermAutomationCredential and updates it. Returns the server's representation of the azurermAutomationCredential, and an error, if there is any.
func (c *FakeAzurermAutomationCredentials) Update(azurermAutomationCredential *v1alpha1.AzurermAutomationCredential) (result *v1alpha1.AzurermAutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermautomationcredentialsResource, azurermAutomationCredential), &v1alpha1.AzurermAutomationCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermAutomationCredentials) UpdateStatus(azurermAutomationCredential *v1alpha1.AzurermAutomationCredential) (*v1alpha1.AzurermAutomationCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermautomationcredentialsResource, "status", azurermAutomationCredential), &v1alpha1.AzurermAutomationCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationCredential), err
}

// Delete takes name of the azurermAutomationCredential and deletes it. Returns an error if one occurs.
func (c *FakeAzurermAutomationCredentials) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermautomationcredentialsResource, name), &v1alpha1.AzurermAutomationCredential{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermAutomationCredentials) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermautomationcredentialsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermAutomationCredentialList{})
	return err
}

// Patch applies the patch and returns the patched azurermAutomationCredential.
func (c *FakeAzurermAutomationCredentials) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermAutomationCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermautomationcredentialsResource, name, pt, data, subresources...), &v1alpha1.AzurermAutomationCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationCredential), err
}
