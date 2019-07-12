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

// FakeAwsStoragegatewayWorkingStorages implements AwsStoragegatewayWorkingStorageInterface
type FakeAwsStoragegatewayWorkingStorages struct {
	Fake *FakeAwsV1alpha1
}

var awsstoragegatewayworkingstoragesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsstoragegatewayworkingstorages"}

var awsstoragegatewayworkingstoragesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsStoragegatewayWorkingStorage"}

// Get takes name of the awsStoragegatewayWorkingStorage, and returns the corresponding awsStoragegatewayWorkingStorage object, and an error if there is any.
func (c *FakeAwsStoragegatewayWorkingStorages) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsStoragegatewayWorkingStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsstoragegatewayworkingstoragesResource, name), &v1alpha1.AwsStoragegatewayWorkingStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayWorkingStorage), err
}

// List takes label and field selectors, and returns the list of AwsStoragegatewayWorkingStorages that match those selectors.
func (c *FakeAwsStoragegatewayWorkingStorages) List(opts v1.ListOptions) (result *v1alpha1.AwsStoragegatewayWorkingStorageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsstoragegatewayworkingstoragesResource, awsstoragegatewayworkingstoragesKind, opts), &v1alpha1.AwsStoragegatewayWorkingStorageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsStoragegatewayWorkingStorageList{ListMeta: obj.(*v1alpha1.AwsStoragegatewayWorkingStorageList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsStoragegatewayWorkingStorageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsStoragegatewayWorkingStorages.
func (c *FakeAwsStoragegatewayWorkingStorages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsstoragegatewayworkingstoragesResource, opts))
}

// Create takes the representation of a awsStoragegatewayWorkingStorage and creates it.  Returns the server's representation of the awsStoragegatewayWorkingStorage, and an error, if there is any.
func (c *FakeAwsStoragegatewayWorkingStorages) Create(awsStoragegatewayWorkingStorage *v1alpha1.AwsStoragegatewayWorkingStorage) (result *v1alpha1.AwsStoragegatewayWorkingStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsstoragegatewayworkingstoragesResource, awsStoragegatewayWorkingStorage), &v1alpha1.AwsStoragegatewayWorkingStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayWorkingStorage), err
}

// Update takes the representation of a awsStoragegatewayWorkingStorage and updates it. Returns the server's representation of the awsStoragegatewayWorkingStorage, and an error, if there is any.
func (c *FakeAwsStoragegatewayWorkingStorages) Update(awsStoragegatewayWorkingStorage *v1alpha1.AwsStoragegatewayWorkingStorage) (result *v1alpha1.AwsStoragegatewayWorkingStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsstoragegatewayworkingstoragesResource, awsStoragegatewayWorkingStorage), &v1alpha1.AwsStoragegatewayWorkingStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayWorkingStorage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsStoragegatewayWorkingStorages) UpdateStatus(awsStoragegatewayWorkingStorage *v1alpha1.AwsStoragegatewayWorkingStorage) (*v1alpha1.AwsStoragegatewayWorkingStorage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsstoragegatewayworkingstoragesResource, "status", awsStoragegatewayWorkingStorage), &v1alpha1.AwsStoragegatewayWorkingStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayWorkingStorage), err
}

// Delete takes name of the awsStoragegatewayWorkingStorage and deletes it. Returns an error if one occurs.
func (c *FakeAwsStoragegatewayWorkingStorages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsstoragegatewayworkingstoragesResource, name), &v1alpha1.AwsStoragegatewayWorkingStorage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsStoragegatewayWorkingStorages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsstoragegatewayworkingstoragesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsStoragegatewayWorkingStorageList{})
	return err
}

// Patch applies the patch and returns the patched awsStoragegatewayWorkingStorage.
func (c *FakeAwsStoragegatewayWorkingStorages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewayWorkingStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsstoragegatewayworkingstoragesResource, name, pt, data, subresources...), &v1alpha1.AwsStoragegatewayWorkingStorage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewayWorkingStorage), err
}
