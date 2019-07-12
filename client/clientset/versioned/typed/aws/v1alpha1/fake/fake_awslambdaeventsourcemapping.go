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

// FakeAwsLambdaEventSourceMappings implements AwsLambdaEventSourceMappingInterface
type FakeAwsLambdaEventSourceMappings struct {
	Fake *FakeAwsV1alpha1
}

var awslambdaeventsourcemappingsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awslambdaeventsourcemappings"}

var awslambdaeventsourcemappingsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLambdaEventSourceMapping"}

// Get takes name of the awsLambdaEventSourceMapping, and returns the corresponding awsLambdaEventSourceMapping object, and an error if there is any.
func (c *FakeAwsLambdaEventSourceMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awslambdaeventsourcemappingsResource, name), &v1alpha1.AwsLambdaEventSourceMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaEventSourceMapping), err
}

// List takes label and field selectors, and returns the list of AwsLambdaEventSourceMappings that match those selectors.
func (c *FakeAwsLambdaEventSourceMappings) List(opts v1.ListOptions) (result *v1alpha1.AwsLambdaEventSourceMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awslambdaeventsourcemappingsResource, awslambdaeventsourcemappingsKind, opts), &v1alpha1.AwsLambdaEventSourceMappingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLambdaEventSourceMappingList{ListMeta: obj.(*v1alpha1.AwsLambdaEventSourceMappingList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLambdaEventSourceMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLambdaEventSourceMappings.
func (c *FakeAwsLambdaEventSourceMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awslambdaeventsourcemappingsResource, opts))
}

// Create takes the representation of a awsLambdaEventSourceMapping and creates it.  Returns the server's representation of the awsLambdaEventSourceMapping, and an error, if there is any.
func (c *FakeAwsLambdaEventSourceMappings) Create(awsLambdaEventSourceMapping *v1alpha1.AwsLambdaEventSourceMapping) (result *v1alpha1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awslambdaeventsourcemappingsResource, awsLambdaEventSourceMapping), &v1alpha1.AwsLambdaEventSourceMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaEventSourceMapping), err
}

// Update takes the representation of a awsLambdaEventSourceMapping and updates it. Returns the server's representation of the awsLambdaEventSourceMapping, and an error, if there is any.
func (c *FakeAwsLambdaEventSourceMappings) Update(awsLambdaEventSourceMapping *v1alpha1.AwsLambdaEventSourceMapping) (result *v1alpha1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awslambdaeventsourcemappingsResource, awsLambdaEventSourceMapping), &v1alpha1.AwsLambdaEventSourceMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaEventSourceMapping), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLambdaEventSourceMappings) UpdateStatus(awsLambdaEventSourceMapping *v1alpha1.AwsLambdaEventSourceMapping) (*v1alpha1.AwsLambdaEventSourceMapping, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awslambdaeventsourcemappingsResource, "status", awsLambdaEventSourceMapping), &v1alpha1.AwsLambdaEventSourceMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaEventSourceMapping), err
}

// Delete takes name of the awsLambdaEventSourceMapping and deletes it. Returns an error if one occurs.
func (c *FakeAwsLambdaEventSourceMappings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awslambdaeventsourcemappingsResource, name), &v1alpha1.AwsLambdaEventSourceMapping{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLambdaEventSourceMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awslambdaeventsourcemappingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLambdaEventSourceMappingList{})
	return err
}

// Patch applies the patch and returns the patched awsLambdaEventSourceMapping.
func (c *FakeAwsLambdaEventSourceMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLambdaEventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awslambdaeventsourcemappingsResource, name, pt, data, subresources...), &v1alpha1.AwsLambdaEventSourceMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLambdaEventSourceMapping), err
}
