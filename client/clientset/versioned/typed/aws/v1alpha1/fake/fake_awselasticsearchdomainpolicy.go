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

// FakeAwsElasticsearchDomainPolicies implements AwsElasticsearchDomainPolicyInterface
type FakeAwsElasticsearchDomainPolicies struct {
	Fake *FakeAwsV1alpha1
}

var awselasticsearchdomainpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awselasticsearchdomainpolicies"}

var awselasticsearchdomainpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsElasticsearchDomainPolicy"}

// Get takes name of the awsElasticsearchDomainPolicy, and returns the corresponding awsElasticsearchDomainPolicy object, and an error if there is any.
func (c *FakeAwsElasticsearchDomainPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticsearchDomainPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awselasticsearchdomainpoliciesResource, name), &v1alpha1.AwsElasticsearchDomainPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomainPolicy), err
}

// List takes label and field selectors, and returns the list of AwsElasticsearchDomainPolicies that match those selectors.
func (c *FakeAwsElasticsearchDomainPolicies) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticsearchDomainPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awselasticsearchdomainpoliciesResource, awselasticsearchdomainpoliciesKind, opts), &v1alpha1.AwsElasticsearchDomainPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsElasticsearchDomainPolicyList{ListMeta: obj.(*v1alpha1.AwsElasticsearchDomainPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsElasticsearchDomainPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsElasticsearchDomainPolicies.
func (c *FakeAwsElasticsearchDomainPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awselasticsearchdomainpoliciesResource, opts))
}

// Create takes the representation of a awsElasticsearchDomainPolicy and creates it.  Returns the server's representation of the awsElasticsearchDomainPolicy, and an error, if there is any.
func (c *FakeAwsElasticsearchDomainPolicies) Create(awsElasticsearchDomainPolicy *v1alpha1.AwsElasticsearchDomainPolicy) (result *v1alpha1.AwsElasticsearchDomainPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awselasticsearchdomainpoliciesResource, awsElasticsearchDomainPolicy), &v1alpha1.AwsElasticsearchDomainPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomainPolicy), err
}

// Update takes the representation of a awsElasticsearchDomainPolicy and updates it. Returns the server's representation of the awsElasticsearchDomainPolicy, and an error, if there is any.
func (c *FakeAwsElasticsearchDomainPolicies) Update(awsElasticsearchDomainPolicy *v1alpha1.AwsElasticsearchDomainPolicy) (result *v1alpha1.AwsElasticsearchDomainPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awselasticsearchdomainpoliciesResource, awsElasticsearchDomainPolicy), &v1alpha1.AwsElasticsearchDomainPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomainPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsElasticsearchDomainPolicies) UpdateStatus(awsElasticsearchDomainPolicy *v1alpha1.AwsElasticsearchDomainPolicy) (*v1alpha1.AwsElasticsearchDomainPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awselasticsearchdomainpoliciesResource, "status", awsElasticsearchDomainPolicy), &v1alpha1.AwsElasticsearchDomainPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomainPolicy), err
}

// Delete takes name of the awsElasticsearchDomainPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsElasticsearchDomainPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awselasticsearchdomainpoliciesResource, name), &v1alpha1.AwsElasticsearchDomainPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsElasticsearchDomainPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awselasticsearchdomainpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsElasticsearchDomainPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsElasticsearchDomainPolicy.
func (c *FakeAwsElasticsearchDomainPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticsearchDomainPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awselasticsearchdomainpoliciesResource, name, pt, data, subresources...), &v1alpha1.AwsElasticsearchDomainPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsElasticsearchDomainPolicy), err
}
