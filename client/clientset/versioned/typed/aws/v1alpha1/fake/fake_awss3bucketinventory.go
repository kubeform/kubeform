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

// FakeAwsS3BucketInventories implements AwsS3BucketInventoryInterface
type FakeAwsS3BucketInventories struct {
	Fake *FakeAwsV1alpha1
}

var awss3bucketinventoriesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awss3bucketinventories"}

var awss3bucketinventoriesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsS3BucketInventory"}

// Get takes name of the awsS3BucketInventory, and returns the corresponding awsS3BucketInventory object, and an error if there is any.
func (c *FakeAwsS3BucketInventories) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsS3BucketInventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awss3bucketinventoriesResource, name), &v1alpha1.AwsS3BucketInventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsS3BucketInventory), err
}

// List takes label and field selectors, and returns the list of AwsS3BucketInventories that match those selectors.
func (c *FakeAwsS3BucketInventories) List(opts v1.ListOptions) (result *v1alpha1.AwsS3BucketInventoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awss3bucketinventoriesResource, awss3bucketinventoriesKind, opts), &v1alpha1.AwsS3BucketInventoryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsS3BucketInventoryList{ListMeta: obj.(*v1alpha1.AwsS3BucketInventoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsS3BucketInventoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsS3BucketInventories.
func (c *FakeAwsS3BucketInventories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awss3bucketinventoriesResource, opts))
}

// Create takes the representation of a awsS3BucketInventory and creates it.  Returns the server's representation of the awsS3BucketInventory, and an error, if there is any.
func (c *FakeAwsS3BucketInventories) Create(awsS3BucketInventory *v1alpha1.AwsS3BucketInventory) (result *v1alpha1.AwsS3BucketInventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awss3bucketinventoriesResource, awsS3BucketInventory), &v1alpha1.AwsS3BucketInventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsS3BucketInventory), err
}

// Update takes the representation of a awsS3BucketInventory and updates it. Returns the server's representation of the awsS3BucketInventory, and an error, if there is any.
func (c *FakeAwsS3BucketInventories) Update(awsS3BucketInventory *v1alpha1.AwsS3BucketInventory) (result *v1alpha1.AwsS3BucketInventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awss3bucketinventoriesResource, awsS3BucketInventory), &v1alpha1.AwsS3BucketInventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsS3BucketInventory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsS3BucketInventories) UpdateStatus(awsS3BucketInventory *v1alpha1.AwsS3BucketInventory) (*v1alpha1.AwsS3BucketInventory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awss3bucketinventoriesResource, "status", awsS3BucketInventory), &v1alpha1.AwsS3BucketInventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsS3BucketInventory), err
}

// Delete takes name of the awsS3BucketInventory and deletes it. Returns an error if one occurs.
func (c *FakeAwsS3BucketInventories) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awss3bucketinventoriesResource, name), &v1alpha1.AwsS3BucketInventory{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsS3BucketInventories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awss3bucketinventoriesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsS3BucketInventoryList{})
	return err
}

// Patch applies the patch and returns the patched awsS3BucketInventory.
func (c *FakeAwsS3BucketInventories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsS3BucketInventory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awss3bucketinventoriesResource, name, pt, data, subresources...), &v1alpha1.AwsS3BucketInventory{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsS3BucketInventory), err
}
