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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsSecurityGroupRules implements AwsSecurityGroupRuleInterface
type FakeAwsSecurityGroupRules struct {
	Fake *FakeAwsV1alpha1
}

var awssecuritygrouprulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awssecuritygrouprules"}

var awssecuritygrouprulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSecurityGroupRule"}

// Get takes name of the awsSecurityGroupRule, and returns the corresponding awsSecurityGroupRule object, and an error if there is any.
func (c *FakeAwsSecurityGroupRules) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awssecuritygrouprulesResource, name), &v1alpha1.AwsSecurityGroupRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}

// List takes label and field selectors, and returns the list of AwsSecurityGroupRules that match those selectors.
func (c *FakeAwsSecurityGroupRules) List(opts v1.ListOptions) (result *v1alpha1.AwsSecurityGroupRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awssecuritygrouprulesResource, awssecuritygrouprulesKind, opts), &v1alpha1.AwsSecurityGroupRuleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSecurityGroupRuleList{ListMeta: obj.(*v1alpha1.AwsSecurityGroupRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSecurityGroupRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSecurityGroupRules.
func (c *FakeAwsSecurityGroupRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awssecuritygrouprulesResource, opts))
}

// Create takes the representation of a awsSecurityGroupRule and creates it.  Returns the server's representation of the awsSecurityGroupRule, and an error, if there is any.
func (c *FakeAwsSecurityGroupRules) Create(awsSecurityGroupRule *v1alpha1.AwsSecurityGroupRule) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awssecuritygrouprulesResource, awsSecurityGroupRule), &v1alpha1.AwsSecurityGroupRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}

// Update takes the representation of a awsSecurityGroupRule and updates it. Returns the server's representation of the awsSecurityGroupRule, and an error, if there is any.
func (c *FakeAwsSecurityGroupRules) Update(awsSecurityGroupRule *v1alpha1.AwsSecurityGroupRule) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awssecuritygrouprulesResource, awsSecurityGroupRule), &v1alpha1.AwsSecurityGroupRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSecurityGroupRules) UpdateStatus(awsSecurityGroupRule *v1alpha1.AwsSecurityGroupRule) (*v1alpha1.AwsSecurityGroupRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awssecuritygrouprulesResource, "status", awsSecurityGroupRule), &v1alpha1.AwsSecurityGroupRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}

// Delete takes name of the awsSecurityGroupRule and deletes it. Returns an error if one occurs.
func (c *FakeAwsSecurityGroupRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awssecuritygrouprulesResource, name), &v1alpha1.AwsSecurityGroupRule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSecurityGroupRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awssecuritygrouprulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSecurityGroupRuleList{})
	return err
}

// Patch applies the patch and returns the patched awsSecurityGroupRule.
func (c *FakeAwsSecurityGroupRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSecurityGroupRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awssecuritygrouprulesResource, name, pt, data, subresources...), &v1alpha1.AwsSecurityGroupRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroupRule), err
}
