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

// FakeRedisFirewallRules implements RedisFirewallRuleInterface
type FakeRedisFirewallRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var redisfirewallrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "redisfirewallrules"}

var redisfirewallrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "RedisFirewallRule"}

// Get takes name of the redisFirewallRule, and returns the corresponding redisFirewallRule object, and an error if there is any.
func (c *FakeRedisFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisfirewallrulesResource, c.ns, name), &v1alpha1.RedisFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisFirewallRule), err
}

// List takes label and field selectors, and returns the list of RedisFirewallRules that match those selectors.
func (c *FakeRedisFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.RedisFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisfirewallrulesResource, redisfirewallrulesKind, c.ns, opts), &v1alpha1.RedisFirewallRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedisFirewallRuleList{ListMeta: obj.(*v1alpha1.RedisFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.RedisFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisFirewallRules.
func (c *FakeRedisFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisfirewallrulesResource, c.ns, opts))

}

// Create takes the representation of a redisFirewallRule and creates it.  Returns the server's representation of the redisFirewallRule, and an error, if there is any.
func (c *FakeRedisFirewallRules) Create(redisFirewallRule *v1alpha1.RedisFirewallRule) (result *v1alpha1.RedisFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisfirewallrulesResource, c.ns, redisFirewallRule), &v1alpha1.RedisFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisFirewallRule), err
}

// Update takes the representation of a redisFirewallRule and updates it. Returns the server's representation of the redisFirewallRule, and an error, if there is any.
func (c *FakeRedisFirewallRules) Update(redisFirewallRule *v1alpha1.RedisFirewallRule) (result *v1alpha1.RedisFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisfirewallrulesResource, c.ns, redisFirewallRule), &v1alpha1.RedisFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedisFirewallRules) UpdateStatus(redisFirewallRule *v1alpha1.RedisFirewallRule) (*v1alpha1.RedisFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redisfirewallrulesResource, "status", c.ns, redisFirewallRule), &v1alpha1.RedisFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisFirewallRule), err
}

// Delete takes name of the redisFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeRedisFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisfirewallrulesResource, c.ns, name), &v1alpha1.RedisFirewallRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisfirewallrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedisFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched redisFirewallRule.
func (c *FakeRedisFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisfirewallrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RedisFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisFirewallRule), err
}
