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

// FakeAwsIamGroupPolicies implements AwsIamGroupPolicyInterface
type FakeAwsIamGroupPolicies struct {
	Fake *FakeAwsV1alpha1
}

var awsiamgrouppoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsiamgrouppolicies"}

var awsiamgrouppoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsIamGroupPolicy"}

// Get takes name of the awsIamGroupPolicy, and returns the corresponding awsIamGroupPolicy object, and an error if there is any.
func (c *FakeAwsIamGroupPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsiamgrouppoliciesResource, name), &v1alpha1.AwsIamGroupPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicy), err
}

// List takes label and field selectors, and returns the list of AwsIamGroupPolicies that match those selectors.
func (c *FakeAwsIamGroupPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsIamGroupPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsiamgrouppoliciesResource, awsiamgrouppoliciesKind, opts), &v1alpha1.AwsIamGroupPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamGroupPolicyList{ListMeta: obj.(*v1alpha1.AwsIamGroupPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsIamGroupPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamGroupPolicies.
func (c *FakeAwsIamGroupPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsiamgrouppoliciesResource, opts))
}

// Create takes the representation of a awsIamGroupPolicy and creates it.  Returns the server's representation of the awsIamGroupPolicy, and an error, if there is any.
func (c *FakeAwsIamGroupPolicies) Create(awsIamGroupPolicy *v1alpha1.AwsIamGroupPolicy) (result *v1alpha1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsiamgrouppoliciesResource, awsIamGroupPolicy), &v1alpha1.AwsIamGroupPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicy), err
}

// Update takes the representation of a awsIamGroupPolicy and updates it. Returns the server's representation of the awsIamGroupPolicy, and an error, if there is any.
func (c *FakeAwsIamGroupPolicies) Update(awsIamGroupPolicy *v1alpha1.AwsIamGroupPolicy) (result *v1alpha1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsiamgrouppoliciesResource, awsIamGroupPolicy), &v1alpha1.AwsIamGroupPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsIamGroupPolicies) UpdateStatus(awsIamGroupPolicy *v1alpha1.AwsIamGroupPolicy) (*v1alpha1.AwsIamGroupPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsiamgrouppoliciesResource, "status", awsIamGroupPolicy), &v1alpha1.AwsIamGroupPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicy), err
}

// Delete takes name of the awsIamGroupPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamGroupPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsiamgrouppoliciesResource, name), &v1alpha1.AwsIamGroupPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamGroupPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsiamgrouppoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamGroupPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsIamGroupPolicy.
func (c *FakeAwsIamGroupPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsiamgrouppoliciesResource, name, pt, data, subresources...), &v1alpha1.AwsIamGroupPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroupPolicy), err
}
