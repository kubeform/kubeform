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

// FakeAwsOpsworksPermissions implements AwsOpsworksPermissionInterface
type FakeAwsOpsworksPermissions struct {
	Fake *FakeAwsV1alpha1
}

var awsopsworkspermissionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsopsworkspermissions"}

var awsopsworkspermissionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsOpsworksPermission"}

// Get takes name of the awsOpsworksPermission, and returns the corresponding awsOpsworksPermission object, and an error if there is any.
func (c *FakeAwsOpsworksPermissions) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsopsworkspermissionsResource, name), &v1alpha1.AwsOpsworksPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksPermission), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksPermissions that match those selectors.
func (c *FakeAwsOpsworksPermissions) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksPermissionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsopsworkspermissionsResource, awsopsworkspermissionsKind, opts), &v1alpha1.AwsOpsworksPermissionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOpsworksPermissionList{ListMeta: obj.(*v1alpha1.AwsOpsworksPermissionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsOpsworksPermissionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksPermissions.
func (c *FakeAwsOpsworksPermissions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsopsworkspermissionsResource, opts))
}

// Create takes the representation of a awsOpsworksPermission and creates it.  Returns the server's representation of the awsOpsworksPermission, and an error, if there is any.
func (c *FakeAwsOpsworksPermissions) Create(awsOpsworksPermission *v1alpha1.AwsOpsworksPermission) (result *v1alpha1.AwsOpsworksPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsopsworkspermissionsResource, awsOpsworksPermission), &v1alpha1.AwsOpsworksPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksPermission), err
}

// Update takes the representation of a awsOpsworksPermission and updates it. Returns the server's representation of the awsOpsworksPermission, and an error, if there is any.
func (c *FakeAwsOpsworksPermissions) Update(awsOpsworksPermission *v1alpha1.AwsOpsworksPermission) (result *v1alpha1.AwsOpsworksPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsopsworkspermissionsResource, awsOpsworksPermission), &v1alpha1.AwsOpsworksPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksPermission), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsOpsworksPermissions) UpdateStatus(awsOpsworksPermission *v1alpha1.AwsOpsworksPermission) (*v1alpha1.AwsOpsworksPermission, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsopsworkspermissionsResource, "status", awsOpsworksPermission), &v1alpha1.AwsOpsworksPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksPermission), err
}

// Delete takes name of the awsOpsworksPermission and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksPermissions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsopsworkspermissionsResource, name), &v1alpha1.AwsOpsworksPermission{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksPermissions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsopsworkspermissionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOpsworksPermissionList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksPermission.
func (c *FakeAwsOpsworksPermissions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsopsworkspermissionsResource, name, pt, data, subresources...), &v1alpha1.AwsOpsworksPermission{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksPermission), err
}
