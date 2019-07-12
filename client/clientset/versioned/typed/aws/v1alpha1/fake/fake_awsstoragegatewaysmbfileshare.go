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

// FakeAwsStoragegatewaySmbFileShares implements AwsStoragegatewaySmbFileShareInterface
type FakeAwsStoragegatewaySmbFileShares struct {
	Fake *FakeAwsV1alpha1
}

var awsstoragegatewaysmbfilesharesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsstoragegatewaysmbfileshares"}

var awsstoragegatewaysmbfilesharesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsStoragegatewaySmbFileShare"}

// Get takes name of the awsStoragegatewaySmbFileShare, and returns the corresponding awsStoragegatewaySmbFileShare object, and an error if there is any.
func (c *FakeAwsStoragegatewaySmbFileShares) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsstoragegatewaysmbfilesharesResource, name), &v1alpha1.AwsStoragegatewaySmbFileShare{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewaySmbFileShare), err
}

// List takes label and field selectors, and returns the list of AwsStoragegatewaySmbFileShares that match those selectors.
func (c *FakeAwsStoragegatewaySmbFileShares) List(opts v1.ListOptions) (result *v1alpha1.AwsStoragegatewaySmbFileShareList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsstoragegatewaysmbfilesharesResource, awsstoragegatewaysmbfilesharesKind, opts), &v1alpha1.AwsStoragegatewaySmbFileShareList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsStoragegatewaySmbFileShareList{ListMeta: obj.(*v1alpha1.AwsStoragegatewaySmbFileShareList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsStoragegatewaySmbFileShareList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsStoragegatewaySmbFileShares.
func (c *FakeAwsStoragegatewaySmbFileShares) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsstoragegatewaysmbfilesharesResource, opts))
}

// Create takes the representation of a awsStoragegatewaySmbFileShare and creates it.  Returns the server's representation of the awsStoragegatewaySmbFileShare, and an error, if there is any.
func (c *FakeAwsStoragegatewaySmbFileShares) Create(awsStoragegatewaySmbFileShare *v1alpha1.AwsStoragegatewaySmbFileShare) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsstoragegatewaysmbfilesharesResource, awsStoragegatewaySmbFileShare), &v1alpha1.AwsStoragegatewaySmbFileShare{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewaySmbFileShare), err
}

// Update takes the representation of a awsStoragegatewaySmbFileShare and updates it. Returns the server's representation of the awsStoragegatewaySmbFileShare, and an error, if there is any.
func (c *FakeAwsStoragegatewaySmbFileShares) Update(awsStoragegatewaySmbFileShare *v1alpha1.AwsStoragegatewaySmbFileShare) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsstoragegatewaysmbfilesharesResource, awsStoragegatewaySmbFileShare), &v1alpha1.AwsStoragegatewaySmbFileShare{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewaySmbFileShare), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsStoragegatewaySmbFileShares) UpdateStatus(awsStoragegatewaySmbFileShare *v1alpha1.AwsStoragegatewaySmbFileShare) (*v1alpha1.AwsStoragegatewaySmbFileShare, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsstoragegatewaysmbfilesharesResource, "status", awsStoragegatewaySmbFileShare), &v1alpha1.AwsStoragegatewaySmbFileShare{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewaySmbFileShare), err
}

// Delete takes name of the awsStoragegatewaySmbFileShare and deletes it. Returns an error if one occurs.
func (c *FakeAwsStoragegatewaySmbFileShares) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsstoragegatewaysmbfilesharesResource, name), &v1alpha1.AwsStoragegatewaySmbFileShare{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsStoragegatewaySmbFileShares) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsstoragegatewaysmbfilesharesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsStoragegatewaySmbFileShareList{})
	return err
}

// Patch applies the patch and returns the patched awsStoragegatewaySmbFileShare.
func (c *FakeAwsStoragegatewaySmbFileShares) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsstoragegatewaysmbfilesharesResource, name, pt, data, subresources...), &v1alpha1.AwsStoragegatewaySmbFileShare{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsStoragegatewaySmbFileShare), err
}
