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

// FakeDataLakeAnalyticsFirewallRules implements DataLakeAnalyticsFirewallRuleInterface
type FakeDataLakeAnalyticsFirewallRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var datalakeanalyticsfirewallrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datalakeanalyticsfirewallrules"}

var datalakeanalyticsfirewallrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataLakeAnalyticsFirewallRule"}

// Get takes name of the dataLakeAnalyticsFirewallRule, and returns the corresponding dataLakeAnalyticsFirewallRule object, and an error if there is any.
func (c *FakeDataLakeAnalyticsFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datalakeanalyticsfirewallrulesResource, c.ns, name), &v1alpha1.DataLakeAnalyticsFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsFirewallRule), err
}

// List takes label and field selectors, and returns the list of DataLakeAnalyticsFirewallRules that match those selectors.
func (c *FakeDataLakeAnalyticsFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.DataLakeAnalyticsFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datalakeanalyticsfirewallrulesResource, datalakeanalyticsfirewallrulesKind, c.ns, opts), &v1alpha1.DataLakeAnalyticsFirewallRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataLakeAnalyticsFirewallRuleList{ListMeta: obj.(*v1alpha1.DataLakeAnalyticsFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataLakeAnalyticsFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataLakeAnalyticsFirewallRules.
func (c *FakeDataLakeAnalyticsFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datalakeanalyticsfirewallrulesResource, c.ns, opts))

}

// Create takes the representation of a dataLakeAnalyticsFirewallRule and creates it.  Returns the server's representation of the dataLakeAnalyticsFirewallRule, and an error, if there is any.
func (c *FakeDataLakeAnalyticsFirewallRules) Create(dataLakeAnalyticsFirewallRule *v1alpha1.DataLakeAnalyticsFirewallRule) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datalakeanalyticsfirewallrulesResource, c.ns, dataLakeAnalyticsFirewallRule), &v1alpha1.DataLakeAnalyticsFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsFirewallRule), err
}

// Update takes the representation of a dataLakeAnalyticsFirewallRule and updates it. Returns the server's representation of the dataLakeAnalyticsFirewallRule, and an error, if there is any.
func (c *FakeDataLakeAnalyticsFirewallRules) Update(dataLakeAnalyticsFirewallRule *v1alpha1.DataLakeAnalyticsFirewallRule) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datalakeanalyticsfirewallrulesResource, c.ns, dataLakeAnalyticsFirewallRule), &v1alpha1.DataLakeAnalyticsFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataLakeAnalyticsFirewallRules) UpdateStatus(dataLakeAnalyticsFirewallRule *v1alpha1.DataLakeAnalyticsFirewallRule) (*v1alpha1.DataLakeAnalyticsFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datalakeanalyticsfirewallrulesResource, "status", c.ns, dataLakeAnalyticsFirewallRule), &v1alpha1.DataLakeAnalyticsFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsFirewallRule), err
}

// Delete takes name of the dataLakeAnalyticsFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeDataLakeAnalyticsFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datalakeanalyticsfirewallrulesResource, c.ns, name), &v1alpha1.DataLakeAnalyticsFirewallRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataLakeAnalyticsFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datalakeanalyticsfirewallrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataLakeAnalyticsFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched dataLakeAnalyticsFirewallRule.
func (c *FakeDataLakeAnalyticsFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datalakeanalyticsfirewallrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataLakeAnalyticsFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeAnalyticsFirewallRule), err
}
