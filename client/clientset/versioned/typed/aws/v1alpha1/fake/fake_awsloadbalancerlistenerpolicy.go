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

// FakeAwsLoadBalancerListenerPolicies implements AwsLoadBalancerListenerPolicyInterface
type FakeAwsLoadBalancerListenerPolicies struct {
	Fake *FakeAwsV1alpha1
}

var awsloadbalancerlistenerpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsloadbalancerlistenerpolicies"}

var awsloadbalancerlistenerpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLoadBalancerListenerPolicy"}

// Get takes name of the awsLoadBalancerListenerPolicy, and returns the corresponding awsLoadBalancerListenerPolicy object, and an error if there is any.
func (c *FakeAwsLoadBalancerListenerPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsloadbalancerlistenerpoliciesResource, name), &v1alpha1.AwsLoadBalancerListenerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerListenerPolicy), err
}

// List takes label and field selectors, and returns the list of AwsLoadBalancerListenerPolicies that match those selectors.
func (c *FakeAwsLoadBalancerListenerPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsLoadBalancerListenerPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsloadbalancerlistenerpoliciesResource, awsloadbalancerlistenerpoliciesKind, opts), &v1alpha1.AwsLoadBalancerListenerPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLoadBalancerListenerPolicyList{ListMeta: obj.(*v1alpha1.AwsLoadBalancerListenerPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLoadBalancerListenerPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLoadBalancerListenerPolicies.
func (c *FakeAwsLoadBalancerListenerPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsloadbalancerlistenerpoliciesResource, opts))
}

// Create takes the representation of a awsLoadBalancerListenerPolicy and creates it.  Returns the server's representation of the awsLoadBalancerListenerPolicy, and an error, if there is any.
func (c *FakeAwsLoadBalancerListenerPolicies) Create(awsLoadBalancerListenerPolicy *v1alpha1.AwsLoadBalancerListenerPolicy) (result *v1alpha1.AwsLoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsloadbalancerlistenerpoliciesResource, awsLoadBalancerListenerPolicy), &v1alpha1.AwsLoadBalancerListenerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerListenerPolicy), err
}

// Update takes the representation of a awsLoadBalancerListenerPolicy and updates it. Returns the server's representation of the awsLoadBalancerListenerPolicy, and an error, if there is any.
func (c *FakeAwsLoadBalancerListenerPolicies) Update(awsLoadBalancerListenerPolicy *v1alpha1.AwsLoadBalancerListenerPolicy) (result *v1alpha1.AwsLoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsloadbalancerlistenerpoliciesResource, awsLoadBalancerListenerPolicy), &v1alpha1.AwsLoadBalancerListenerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerListenerPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLoadBalancerListenerPolicies) UpdateStatus(awsLoadBalancerListenerPolicy *v1alpha1.AwsLoadBalancerListenerPolicy) (*v1alpha1.AwsLoadBalancerListenerPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsloadbalancerlistenerpoliciesResource, "status", awsLoadBalancerListenerPolicy), &v1alpha1.AwsLoadBalancerListenerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerListenerPolicy), err
}

// Delete takes name of the awsLoadBalancerListenerPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsLoadBalancerListenerPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsloadbalancerlistenerpoliciesResource, name), &v1alpha1.AwsLoadBalancerListenerPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLoadBalancerListenerPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsloadbalancerlistenerpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLoadBalancerListenerPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsLoadBalancerListenerPolicy.
func (c *FakeAwsLoadBalancerListenerPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsloadbalancerlistenerpoliciesResource, name, pt, data, subresources...), &v1alpha1.AwsLoadBalancerListenerPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLoadBalancerListenerPolicy), err
}
