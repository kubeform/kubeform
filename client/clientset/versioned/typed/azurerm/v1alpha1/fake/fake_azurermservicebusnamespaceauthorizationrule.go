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

// FakeAzurermServicebusNamespaceAuthorizationRules implements AzurermServicebusNamespaceAuthorizationRuleInterface
type FakeAzurermServicebusNamespaceAuthorizationRules struct {
	Fake *FakeAzurermV1alpha1
}

var azurermservicebusnamespaceauthorizationrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermservicebusnamespaceauthorizationrules"}

var azurermservicebusnamespaceauthorizationrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermServicebusNamespaceAuthorizationRule"}

// Get takes name of the azurermServicebusNamespaceAuthorizationRule, and returns the corresponding azurermServicebusNamespaceAuthorizationRule object, and an error if there is any.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermservicebusnamespaceauthorizationrulesResource, name), &v1alpha1.AzurermServicebusNamespaceAuthorizationRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServicebusNamespaceAuthorizationRule), err
}

// List takes label and field selectors, and returns the list of AzurermServicebusNamespaceAuthorizationRules that match those selectors.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermServicebusNamespaceAuthorizationRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermservicebusnamespaceauthorizationrulesResource, azurermservicebusnamespaceauthorizationrulesKind, opts), &v1alpha1.AzurermServicebusNamespaceAuthorizationRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermServicebusNamespaceAuthorizationRuleList{ListMeta: obj.(*v1alpha1.AzurermServicebusNamespaceAuthorizationRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermServicebusNamespaceAuthorizationRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermServicebusNamespaceAuthorizationRules.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermservicebusnamespaceauthorizationrulesResource, opts))
}

// Create takes the representation of a azurermServicebusNamespaceAuthorizationRule and creates it.  Returns the server's representation of the azurermServicebusNamespaceAuthorizationRule, and an error, if there is any.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) Create(azurermServicebusNamespaceAuthorizationRule *v1alpha1.AzurermServicebusNamespaceAuthorizationRule) (result *v1alpha1.AzurermServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermservicebusnamespaceauthorizationrulesResource, azurermServicebusNamespaceAuthorizationRule), &v1alpha1.AzurermServicebusNamespaceAuthorizationRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServicebusNamespaceAuthorizationRule), err
}

// Update takes the representation of a azurermServicebusNamespaceAuthorizationRule and updates it. Returns the server's representation of the azurermServicebusNamespaceAuthorizationRule, and an error, if there is any.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) Update(azurermServicebusNamespaceAuthorizationRule *v1alpha1.AzurermServicebusNamespaceAuthorizationRule) (result *v1alpha1.AzurermServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermservicebusnamespaceauthorizationrulesResource, azurermServicebusNamespaceAuthorizationRule), &v1alpha1.AzurermServicebusNamespaceAuthorizationRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServicebusNamespaceAuthorizationRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) UpdateStatus(azurermServicebusNamespaceAuthorizationRule *v1alpha1.AzurermServicebusNamespaceAuthorizationRule) (*v1alpha1.AzurermServicebusNamespaceAuthorizationRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermservicebusnamespaceauthorizationrulesResource, "status", azurermServicebusNamespaceAuthorizationRule), &v1alpha1.AzurermServicebusNamespaceAuthorizationRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServicebusNamespaceAuthorizationRule), err
}

// Delete takes name of the azurermServicebusNamespaceAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermservicebusnamespaceauthorizationrulesResource, name), &v1alpha1.AzurermServicebusNamespaceAuthorizationRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermservicebusnamespaceauthorizationrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermServicebusNamespaceAuthorizationRuleList{})
	return err
}

// Patch applies the patch and returns the patched azurermServicebusNamespaceAuthorizationRule.
func (c *FakeAzurermServicebusNamespaceAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermservicebusnamespaceauthorizationrulesResource, name, pt, data, subresources...), &v1alpha1.AzurermServicebusNamespaceAuthorizationRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermServicebusNamespaceAuthorizationRule), err
}
