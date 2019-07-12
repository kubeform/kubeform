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

// FakeAwsOpsworksUserProfiles implements AwsOpsworksUserProfileInterface
type FakeAwsOpsworksUserProfiles struct {
	Fake *FakeAwsV1alpha1
}

var awsopsworksuserprofilesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsopsworksuserprofiles"}

var awsopsworksuserprofilesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsOpsworksUserProfile"}

// Get takes name of the awsOpsworksUserProfile, and returns the corresponding awsOpsworksUserProfile object, and an error if there is any.
func (c *FakeAwsOpsworksUserProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsopsworksuserprofilesResource, name), &v1alpha1.AwsOpsworksUserProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksUserProfile), err
}

// List takes label and field selectors, and returns the list of AwsOpsworksUserProfiles that match those selectors.
func (c *FakeAwsOpsworksUserProfiles) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksUserProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsopsworksuserprofilesResource, awsopsworksuserprofilesKind, opts), &v1alpha1.AwsOpsworksUserProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsOpsworksUserProfileList{ListMeta: obj.(*v1alpha1.AwsOpsworksUserProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsOpsworksUserProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOpsworksUserProfiles.
func (c *FakeAwsOpsworksUserProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsopsworksuserprofilesResource, opts))
}

// Create takes the representation of a awsOpsworksUserProfile and creates it.  Returns the server's representation of the awsOpsworksUserProfile, and an error, if there is any.
func (c *FakeAwsOpsworksUserProfiles) Create(awsOpsworksUserProfile *v1alpha1.AwsOpsworksUserProfile) (result *v1alpha1.AwsOpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsopsworksuserprofilesResource, awsOpsworksUserProfile), &v1alpha1.AwsOpsworksUserProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksUserProfile), err
}

// Update takes the representation of a awsOpsworksUserProfile and updates it. Returns the server's representation of the awsOpsworksUserProfile, and an error, if there is any.
func (c *FakeAwsOpsworksUserProfiles) Update(awsOpsworksUserProfile *v1alpha1.AwsOpsworksUserProfile) (result *v1alpha1.AwsOpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsopsworksuserprofilesResource, awsOpsworksUserProfile), &v1alpha1.AwsOpsworksUserProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksUserProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsOpsworksUserProfiles) UpdateStatus(awsOpsworksUserProfile *v1alpha1.AwsOpsworksUserProfile) (*v1alpha1.AwsOpsworksUserProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsopsworksuserprofilesResource, "status", awsOpsworksUserProfile), &v1alpha1.AwsOpsworksUserProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksUserProfile), err
}

// Delete takes name of the awsOpsworksUserProfile and deletes it. Returns an error if one occurs.
func (c *FakeAwsOpsworksUserProfiles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsopsworksuserprofilesResource, name), &v1alpha1.AwsOpsworksUserProfile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOpsworksUserProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsopsworksuserprofilesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsOpsworksUserProfileList{})
	return err
}

// Patch applies the patch and returns the patched awsOpsworksUserProfile.
func (c *FakeAwsOpsworksUserProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksUserProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsopsworksuserprofilesResource, name, pt, data, subresources...), &v1alpha1.AwsOpsworksUserProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsOpsworksUserProfile), err
}
