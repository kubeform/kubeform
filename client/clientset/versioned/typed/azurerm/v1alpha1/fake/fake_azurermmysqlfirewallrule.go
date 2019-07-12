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

// FakeAzurermMysqlFirewallRules implements AzurermMysqlFirewallRuleInterface
type FakeAzurermMysqlFirewallRules struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmysqlfirewallrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmysqlfirewallrules"}

var azurermmysqlfirewallrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMysqlFirewallRule"}

// Get takes name of the azurermMysqlFirewallRule, and returns the corresponding azurermMysqlFirewallRule object, and an error if there is any.
func (c *FakeAzurermMysqlFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMysqlFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmysqlfirewallrulesResource, name), &v1alpha1.AzurermMysqlFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlFirewallRule), err
}

// List takes label and field selectors, and returns the list of AzurermMysqlFirewallRules that match those selectors.
func (c *FakeAzurermMysqlFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.AzurermMysqlFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmysqlfirewallrulesResource, azurermmysqlfirewallrulesKind, opts), &v1alpha1.AzurermMysqlFirewallRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMysqlFirewallRuleList{ListMeta: obj.(*v1alpha1.AzurermMysqlFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMysqlFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMysqlFirewallRules.
func (c *FakeAzurermMysqlFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmysqlfirewallrulesResource, opts))
}

// Create takes the representation of a azurermMysqlFirewallRule and creates it.  Returns the server's representation of the azurermMysqlFirewallRule, and an error, if there is any.
func (c *FakeAzurermMysqlFirewallRules) Create(azurermMysqlFirewallRule *v1alpha1.AzurermMysqlFirewallRule) (result *v1alpha1.AzurermMysqlFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmysqlfirewallrulesResource, azurermMysqlFirewallRule), &v1alpha1.AzurermMysqlFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlFirewallRule), err
}

// Update takes the representation of a azurermMysqlFirewallRule and updates it. Returns the server's representation of the azurermMysqlFirewallRule, and an error, if there is any.
func (c *FakeAzurermMysqlFirewallRules) Update(azurermMysqlFirewallRule *v1alpha1.AzurermMysqlFirewallRule) (result *v1alpha1.AzurermMysqlFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmysqlfirewallrulesResource, azurermMysqlFirewallRule), &v1alpha1.AzurermMysqlFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMysqlFirewallRules) UpdateStatus(azurermMysqlFirewallRule *v1alpha1.AzurermMysqlFirewallRule) (*v1alpha1.AzurermMysqlFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmysqlfirewallrulesResource, "status", azurermMysqlFirewallRule), &v1alpha1.AzurermMysqlFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlFirewallRule), err
}

// Delete takes name of the azurermMysqlFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMysqlFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmysqlfirewallrulesResource, name), &v1alpha1.AzurermMysqlFirewallRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMysqlFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmysqlfirewallrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMysqlFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched azurermMysqlFirewallRule.
func (c *FakeAzurermMysqlFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMysqlFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmysqlfirewallrulesResource, name, pt, data, subresources...), &v1alpha1.AzurermMysqlFirewallRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMysqlFirewallRule), err
}
