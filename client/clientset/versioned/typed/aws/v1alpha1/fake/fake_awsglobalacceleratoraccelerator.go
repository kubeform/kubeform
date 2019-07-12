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

// FakeAwsGlobalacceleratorAccelerators implements AwsGlobalacceleratorAcceleratorInterface
type FakeAwsGlobalacceleratorAccelerators struct {
	Fake *FakeAwsV1alpha1
}

var awsglobalacceleratoracceleratorsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsglobalacceleratoraccelerators"}

var awsglobalacceleratoracceleratorsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsGlobalacceleratorAccelerator"}

// Get takes name of the awsGlobalacceleratorAccelerator, and returns the corresponding awsGlobalacceleratorAccelerator object, and an error if there is any.
func (c *FakeAwsGlobalacceleratorAccelerators) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGlobalacceleratorAccelerator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsglobalacceleratoracceleratorsResource, name), &v1alpha1.AwsGlobalacceleratorAccelerator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorAccelerator), err
}

// List takes label and field selectors, and returns the list of AwsGlobalacceleratorAccelerators that match those selectors.
func (c *FakeAwsGlobalacceleratorAccelerators) List(opts v1.ListOptions) (result *v1alpha1.AwsGlobalacceleratorAcceleratorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsglobalacceleratoracceleratorsResource, awsglobalacceleratoracceleratorsKind, opts), &v1alpha1.AwsGlobalacceleratorAcceleratorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGlobalacceleratorAcceleratorList{ListMeta: obj.(*v1alpha1.AwsGlobalacceleratorAcceleratorList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsGlobalacceleratorAcceleratorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGlobalacceleratorAccelerators.
func (c *FakeAwsGlobalacceleratorAccelerators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsglobalacceleratoracceleratorsResource, opts))
}

// Create takes the representation of a awsGlobalacceleratorAccelerator and creates it.  Returns the server's representation of the awsGlobalacceleratorAccelerator, and an error, if there is any.
func (c *FakeAwsGlobalacceleratorAccelerators) Create(awsGlobalacceleratorAccelerator *v1alpha1.AwsGlobalacceleratorAccelerator) (result *v1alpha1.AwsGlobalacceleratorAccelerator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsglobalacceleratoracceleratorsResource, awsGlobalacceleratorAccelerator), &v1alpha1.AwsGlobalacceleratorAccelerator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorAccelerator), err
}

// Update takes the representation of a awsGlobalacceleratorAccelerator and updates it. Returns the server's representation of the awsGlobalacceleratorAccelerator, and an error, if there is any.
func (c *FakeAwsGlobalacceleratorAccelerators) Update(awsGlobalacceleratorAccelerator *v1alpha1.AwsGlobalacceleratorAccelerator) (result *v1alpha1.AwsGlobalacceleratorAccelerator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsglobalacceleratoracceleratorsResource, awsGlobalacceleratorAccelerator), &v1alpha1.AwsGlobalacceleratorAccelerator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorAccelerator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsGlobalacceleratorAccelerators) UpdateStatus(awsGlobalacceleratorAccelerator *v1alpha1.AwsGlobalacceleratorAccelerator) (*v1alpha1.AwsGlobalacceleratorAccelerator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsglobalacceleratoracceleratorsResource, "status", awsGlobalacceleratorAccelerator), &v1alpha1.AwsGlobalacceleratorAccelerator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorAccelerator), err
}

// Delete takes name of the awsGlobalacceleratorAccelerator and deletes it. Returns an error if one occurs.
func (c *FakeAwsGlobalacceleratorAccelerators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsglobalacceleratoracceleratorsResource, name), &v1alpha1.AwsGlobalacceleratorAccelerator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGlobalacceleratorAccelerators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsglobalacceleratoracceleratorsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGlobalacceleratorAcceleratorList{})
	return err
}

// Patch applies the patch and returns the patched awsGlobalacceleratorAccelerator.
func (c *FakeAwsGlobalacceleratorAccelerators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlobalacceleratorAccelerator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsglobalacceleratoracceleratorsResource, name, pt, data, subresources...), &v1alpha1.AwsGlobalacceleratorAccelerator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlobalacceleratorAccelerator), err
}
