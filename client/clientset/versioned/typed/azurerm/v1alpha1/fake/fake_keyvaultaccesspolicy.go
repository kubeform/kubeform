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

// FakeKeyVaultAccessPolicies implements KeyVaultAccessPolicyInterface
type FakeKeyVaultAccessPolicies struct {
	Fake *FakeAzurermV1alpha1
}

var keyvaultaccesspoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "keyvaultaccesspolicies"}

var keyvaultaccesspoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "KeyVaultAccessPolicy"}

// Get takes name of the keyVaultAccessPolicy, and returns the corresponding keyVaultAccessPolicy object, and an error if there is any.
func (c *FakeKeyVaultAccessPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.KeyVaultAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(keyvaultaccesspoliciesResource, name), &v1alpha1.KeyVaultAccessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultAccessPolicy), err
}

// List takes label and field selectors, and returns the list of KeyVaultAccessPolicies that match those selectors.
func (c *FakeKeyVaultAccessPolicies) List(opts v1.ListOptions) (result *v1alpha1.KeyVaultAccessPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(keyvaultaccesspoliciesResource, keyvaultaccesspoliciesKind, opts), &v1alpha1.KeyVaultAccessPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeyVaultAccessPolicyList{ListMeta: obj.(*v1alpha1.KeyVaultAccessPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeyVaultAccessPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keyVaultAccessPolicies.
func (c *FakeKeyVaultAccessPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(keyvaultaccesspoliciesResource, opts))
}

// Create takes the representation of a keyVaultAccessPolicy and creates it.  Returns the server's representation of the keyVaultAccessPolicy, and an error, if there is any.
func (c *FakeKeyVaultAccessPolicies) Create(keyVaultAccessPolicy *v1alpha1.KeyVaultAccessPolicy) (result *v1alpha1.KeyVaultAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(keyvaultaccesspoliciesResource, keyVaultAccessPolicy), &v1alpha1.KeyVaultAccessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultAccessPolicy), err
}

// Update takes the representation of a keyVaultAccessPolicy and updates it. Returns the server's representation of the keyVaultAccessPolicy, and an error, if there is any.
func (c *FakeKeyVaultAccessPolicies) Update(keyVaultAccessPolicy *v1alpha1.KeyVaultAccessPolicy) (result *v1alpha1.KeyVaultAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(keyvaultaccesspoliciesResource, keyVaultAccessPolicy), &v1alpha1.KeyVaultAccessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultAccessPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeyVaultAccessPolicies) UpdateStatus(keyVaultAccessPolicy *v1alpha1.KeyVaultAccessPolicy) (*v1alpha1.KeyVaultAccessPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(keyvaultaccesspoliciesResource, "status", keyVaultAccessPolicy), &v1alpha1.KeyVaultAccessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultAccessPolicy), err
}

// Delete takes name of the keyVaultAccessPolicy and deletes it. Returns an error if one occurs.
func (c *FakeKeyVaultAccessPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(keyvaultaccesspoliciesResource, name), &v1alpha1.KeyVaultAccessPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeyVaultAccessPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(keyvaultaccesspoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeyVaultAccessPolicyList{})
	return err
}

// Patch applies the patch and returns the patched keyVaultAccessPolicy.
func (c *FakeKeyVaultAccessPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KeyVaultAccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(keyvaultaccesspoliciesResource, name, pt, data, subresources...), &v1alpha1.KeyVaultAccessPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyVaultAccessPolicy), err
}
