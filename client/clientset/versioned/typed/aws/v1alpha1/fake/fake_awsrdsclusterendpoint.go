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

// FakeAwsRdsClusterEndpoints implements AwsRdsClusterEndpointInterface
type FakeAwsRdsClusterEndpoints struct {
	Fake *FakeAwsV1alpha1
}

var awsrdsclusterendpointsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsrdsclusterendpoints"}

var awsrdsclusterendpointsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsRdsClusterEndpoint"}

// Get takes name of the awsRdsClusterEndpoint, and returns the corresponding awsRdsClusterEndpoint object, and an error if there is any.
func (c *FakeAwsRdsClusterEndpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRdsClusterEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsrdsclusterendpointsResource, name), &v1alpha1.AwsRdsClusterEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRdsClusterEndpoint), err
}

// List takes label and field selectors, and returns the list of AwsRdsClusterEndpoints that match those selectors.
func (c *FakeAwsRdsClusterEndpoints) List(opts v1.ListOptions) (result *v1alpha1.AwsRdsClusterEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsrdsclusterendpointsResource, awsrdsclusterendpointsKind, opts), &v1alpha1.AwsRdsClusterEndpointList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsRdsClusterEndpointList{ListMeta: obj.(*v1alpha1.AwsRdsClusterEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsRdsClusterEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRdsClusterEndpoints.
func (c *FakeAwsRdsClusterEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsrdsclusterendpointsResource, opts))
}

// Create takes the representation of a awsRdsClusterEndpoint and creates it.  Returns the server's representation of the awsRdsClusterEndpoint, and an error, if there is any.
func (c *FakeAwsRdsClusterEndpoints) Create(awsRdsClusterEndpoint *v1alpha1.AwsRdsClusterEndpoint) (result *v1alpha1.AwsRdsClusterEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsrdsclusterendpointsResource, awsRdsClusterEndpoint), &v1alpha1.AwsRdsClusterEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRdsClusterEndpoint), err
}

// Update takes the representation of a awsRdsClusterEndpoint and updates it. Returns the server's representation of the awsRdsClusterEndpoint, and an error, if there is any.
func (c *FakeAwsRdsClusterEndpoints) Update(awsRdsClusterEndpoint *v1alpha1.AwsRdsClusterEndpoint) (result *v1alpha1.AwsRdsClusterEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsrdsclusterendpointsResource, awsRdsClusterEndpoint), &v1alpha1.AwsRdsClusterEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRdsClusterEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsRdsClusterEndpoints) UpdateStatus(awsRdsClusterEndpoint *v1alpha1.AwsRdsClusterEndpoint) (*v1alpha1.AwsRdsClusterEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsrdsclusterendpointsResource, "status", awsRdsClusterEndpoint), &v1alpha1.AwsRdsClusterEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRdsClusterEndpoint), err
}

// Delete takes name of the awsRdsClusterEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeAwsRdsClusterEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsrdsclusterendpointsResource, name), &v1alpha1.AwsRdsClusterEndpoint{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRdsClusterEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsrdsclusterendpointsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsRdsClusterEndpointList{})
	return err
}

// Patch applies the patch and returns the patched awsRdsClusterEndpoint.
func (c *FakeAwsRdsClusterEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRdsClusterEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsrdsclusterendpointsResource, name, pt, data, subresources...), &v1alpha1.AwsRdsClusterEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRdsClusterEndpoint), err
}
