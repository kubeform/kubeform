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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeXraySamplingRules implements XraySamplingRuleInterface
type FakeXraySamplingRules struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var xraysamplingrulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "xraysamplingrules"}

var xraysamplingrulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "XraySamplingRule"}

// Get takes name of the xraySamplingRule, and returns the corresponding xraySamplingRule object, and an error if there is any.
func (c *FakeXraySamplingRules) Get(name string, options v1.GetOptions) (result *v1alpha1.XraySamplingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(xraysamplingrulesResource, c.ns, name), &v1alpha1.XraySamplingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XraySamplingRule), err
}

// List takes label and field selectors, and returns the list of XraySamplingRules that match those selectors.
func (c *FakeXraySamplingRules) List(opts v1.ListOptions) (result *v1alpha1.XraySamplingRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(xraysamplingrulesResource, xraysamplingrulesKind, c.ns, opts), &v1alpha1.XraySamplingRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.XraySamplingRuleList{ListMeta: obj.(*v1alpha1.XraySamplingRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.XraySamplingRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested xraySamplingRules.
func (c *FakeXraySamplingRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(xraysamplingrulesResource, c.ns, opts))

}

// Create takes the representation of a xraySamplingRule and creates it.  Returns the server's representation of the xraySamplingRule, and an error, if there is any.
func (c *FakeXraySamplingRules) Create(xraySamplingRule *v1alpha1.XraySamplingRule) (result *v1alpha1.XraySamplingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(xraysamplingrulesResource, c.ns, xraySamplingRule), &v1alpha1.XraySamplingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XraySamplingRule), err
}

// Update takes the representation of a xraySamplingRule and updates it. Returns the server's representation of the xraySamplingRule, and an error, if there is any.
func (c *FakeXraySamplingRules) Update(xraySamplingRule *v1alpha1.XraySamplingRule) (result *v1alpha1.XraySamplingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(xraysamplingrulesResource, c.ns, xraySamplingRule), &v1alpha1.XraySamplingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XraySamplingRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeXraySamplingRules) UpdateStatus(xraySamplingRule *v1alpha1.XraySamplingRule) (*v1alpha1.XraySamplingRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(xraysamplingrulesResource, "status", c.ns, xraySamplingRule), &v1alpha1.XraySamplingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XraySamplingRule), err
}

// Delete takes name of the xraySamplingRule and deletes it. Returns an error if one occurs.
func (c *FakeXraySamplingRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(xraysamplingrulesResource, c.ns, name), &v1alpha1.XraySamplingRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeXraySamplingRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(xraysamplingrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.XraySamplingRuleList{})
	return err
}

// Patch applies the patch and returns the patched xraySamplingRule.
func (c *FakeXraySamplingRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.XraySamplingRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(xraysamplingrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.XraySamplingRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XraySamplingRule), err
}
