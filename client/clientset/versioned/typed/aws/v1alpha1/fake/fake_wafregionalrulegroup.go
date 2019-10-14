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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeWafregionalRuleGroups implements WafregionalRuleGroupInterface
type FakeWafregionalRuleGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var wafregionalrulegroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "wafregionalrulegroups"}

var wafregionalrulegroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WafregionalRuleGroup"}

// Get takes name of the wafregionalRuleGroup, and returns the corresponding wafregionalRuleGroup object, and an error if there is any.
func (c *FakeWafregionalRuleGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(wafregionalrulegroupsResource, c.ns, name), &v1alpha1.WafregionalRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRuleGroup), err
}

// List takes label and field selectors, and returns the list of WafregionalRuleGroups that match those selectors.
func (c *FakeWafregionalRuleGroups) List(opts v1.ListOptions) (result *v1alpha1.WafregionalRuleGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(wafregionalrulegroupsResource, wafregionalrulegroupsKind, c.ns, opts), &v1alpha1.WafregionalRuleGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafregionalRuleGroupList{ListMeta: obj.(*v1alpha1.WafregionalRuleGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafregionalRuleGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafregionalRuleGroups.
func (c *FakeWafregionalRuleGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(wafregionalrulegroupsResource, c.ns, opts))

}

// Create takes the representation of a wafregionalRuleGroup and creates it.  Returns the server's representation of the wafregionalRuleGroup, and an error, if there is any.
func (c *FakeWafregionalRuleGroups) Create(wafregionalRuleGroup *v1alpha1.WafregionalRuleGroup) (result *v1alpha1.WafregionalRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(wafregionalrulegroupsResource, c.ns, wafregionalRuleGroup), &v1alpha1.WafregionalRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRuleGroup), err
}

// Update takes the representation of a wafregionalRuleGroup and updates it. Returns the server's representation of the wafregionalRuleGroup, and an error, if there is any.
func (c *FakeWafregionalRuleGroups) Update(wafregionalRuleGroup *v1alpha1.WafregionalRuleGroup) (result *v1alpha1.WafregionalRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(wafregionalrulegroupsResource, c.ns, wafregionalRuleGroup), &v1alpha1.WafregionalRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRuleGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafregionalRuleGroups) UpdateStatus(wafregionalRuleGroup *v1alpha1.WafregionalRuleGroup) (*v1alpha1.WafregionalRuleGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(wafregionalrulegroupsResource, "status", c.ns, wafregionalRuleGroup), &v1alpha1.WafregionalRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRuleGroup), err
}

// Delete takes name of the wafregionalRuleGroup and deletes it. Returns an error if one occurs.
func (c *FakeWafregionalRuleGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(wafregionalrulegroupsResource, c.ns, name), &v1alpha1.WafregionalRuleGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafregionalRuleGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(wafregionalrulegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafregionalRuleGroupList{})
	return err
}

// Patch applies the patch and returns the patched wafregionalRuleGroup.
func (c *FakeWafregionalRuleGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalRuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(wafregionalrulegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WafregionalRuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRuleGroup), err
}
