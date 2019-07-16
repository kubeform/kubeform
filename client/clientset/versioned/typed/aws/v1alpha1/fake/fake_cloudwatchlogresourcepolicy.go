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

// FakeCloudwatchLogResourcePolicies implements CloudwatchLogResourcePolicyInterface
type FakeCloudwatchLogResourcePolicies struct {
	Fake *FakeAwsV1alpha1
}

var cloudwatchlogresourcepoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudwatchlogresourcepolicies"}

var cloudwatchlogresourcepoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudwatchLogResourcePolicy"}

// Get takes name of the cloudwatchLogResourcePolicy, and returns the corresponding cloudwatchLogResourcePolicy object, and an error if there is any.
func (c *FakeCloudwatchLogResourcePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudwatchlogresourcepoliciesResource, name), &v1alpha1.CloudwatchLogResourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogResourcePolicy), err
}

// List takes label and field selectors, and returns the list of CloudwatchLogResourcePolicies that match those selectors.
func (c *FakeCloudwatchLogResourcePolicies) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogResourcePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudwatchlogresourcepoliciesResource, cloudwatchlogresourcepoliciesKind, opts), &v1alpha1.CloudwatchLogResourcePolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudwatchLogResourcePolicyList{ListMeta: obj.(*v1alpha1.CloudwatchLogResourcePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudwatchLogResourcePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogResourcePolicies.
func (c *FakeCloudwatchLogResourcePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudwatchlogresourcepoliciesResource, opts))
}

// Create takes the representation of a cloudwatchLogResourcePolicy and creates it.  Returns the server's representation of the cloudwatchLogResourcePolicy, and an error, if there is any.
func (c *FakeCloudwatchLogResourcePolicies) Create(cloudwatchLogResourcePolicy *v1alpha1.CloudwatchLogResourcePolicy) (result *v1alpha1.CloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudwatchlogresourcepoliciesResource, cloudwatchLogResourcePolicy), &v1alpha1.CloudwatchLogResourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogResourcePolicy), err
}

// Update takes the representation of a cloudwatchLogResourcePolicy and updates it. Returns the server's representation of the cloudwatchLogResourcePolicy, and an error, if there is any.
func (c *FakeCloudwatchLogResourcePolicies) Update(cloudwatchLogResourcePolicy *v1alpha1.CloudwatchLogResourcePolicy) (result *v1alpha1.CloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudwatchlogresourcepoliciesResource, cloudwatchLogResourcePolicy), &v1alpha1.CloudwatchLogResourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogResourcePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudwatchLogResourcePolicies) UpdateStatus(cloudwatchLogResourcePolicy *v1alpha1.CloudwatchLogResourcePolicy) (*v1alpha1.CloudwatchLogResourcePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cloudwatchlogresourcepoliciesResource, "status", cloudwatchLogResourcePolicy), &v1alpha1.CloudwatchLogResourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogResourcePolicy), err
}

// Delete takes name of the cloudwatchLogResourcePolicy and deletes it. Returns an error if one occurs.
func (c *FakeCloudwatchLogResourcePolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cloudwatchlogresourcepoliciesResource, name), &v1alpha1.CloudwatchLogResourcePolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudwatchLogResourcePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudwatchlogresourcepoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudwatchLogResourcePolicyList{})
	return err
}

// Patch applies the patch and returns the patched cloudwatchLogResourcePolicy.
func (c *FakeCloudwatchLogResourcePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudwatchlogresourcepoliciesResource, name, pt, data, subresources...), &v1alpha1.CloudwatchLogResourcePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogResourcePolicy), err
}
