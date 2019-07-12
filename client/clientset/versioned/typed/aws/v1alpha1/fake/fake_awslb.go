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

// FakeAwsLbs implements AwsLbInterface
type FakeAwsLbs struct {
	Fake *FakeAwsV1alpha1
}

var awslbsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awslbs"}

var awslbsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLb"}

// Get takes name of the awsLb, and returns the corresponding awsLb object, and an error if there is any.
func (c *FakeAwsLbs) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awslbsResource, name), &v1alpha1.AwsLb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLb), err
}

// List takes label and field selectors, and returns the list of AwsLbs that match those selectors.
func (c *FakeAwsLbs) List(opts v1.ListOptions) (result *v1alpha1.AwsLbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awslbsResource, awslbsKind, opts), &v1alpha1.AwsLbList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLbList{ListMeta: obj.(*v1alpha1.AwsLbList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLbs.
func (c *FakeAwsLbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awslbsResource, opts))
}

// Create takes the representation of a awsLb and creates it.  Returns the server's representation of the awsLb, and an error, if there is any.
func (c *FakeAwsLbs) Create(awsLb *v1alpha1.AwsLb) (result *v1alpha1.AwsLb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awslbsResource, awsLb), &v1alpha1.AwsLb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLb), err
}

// Update takes the representation of a awsLb and updates it. Returns the server's representation of the awsLb, and an error, if there is any.
func (c *FakeAwsLbs) Update(awsLb *v1alpha1.AwsLb) (result *v1alpha1.AwsLb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awslbsResource, awsLb), &v1alpha1.AwsLb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLbs) UpdateStatus(awsLb *v1alpha1.AwsLb) (*v1alpha1.AwsLb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awslbsResource, "status", awsLb), &v1alpha1.AwsLb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLb), err
}

// Delete takes name of the awsLb and deletes it. Returns an error if one occurs.
func (c *FakeAwsLbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awslbsResource, name), &v1alpha1.AwsLb{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awslbsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLbList{})
	return err
}

// Patch applies the patch and returns the patched awsLb.
func (c *FakeAwsLbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awslbsResource, name, pt, data, subresources...), &v1alpha1.AwsLb{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLb), err
}
