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

// FakeAwsSecurityGroups implements AwsSecurityGroupInterface
type FakeAwsSecurityGroups struct {
	Fake *FakeAwsV1alpha1
}

var awssecuritygroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awssecuritygroups"}

var awssecuritygroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSecurityGroup"}

// Get takes name of the awsSecurityGroup, and returns the corresponding awsSecurityGroup object, and an error if there is any.
func (c *FakeAwsSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awssecuritygroupsResource, name), &v1alpha1.AwsSecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroup), err
}

// List takes label and field selectors, and returns the list of AwsSecurityGroups that match those selectors.
func (c *FakeAwsSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsSecurityGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awssecuritygroupsResource, awssecuritygroupsKind, opts), &v1alpha1.AwsSecurityGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSecurityGroupList{ListMeta: obj.(*v1alpha1.AwsSecurityGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSecurityGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSecurityGroups.
func (c *FakeAwsSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awssecuritygroupsResource, opts))
}

// Create takes the representation of a awsSecurityGroup and creates it.  Returns the server's representation of the awsSecurityGroup, and an error, if there is any.
func (c *FakeAwsSecurityGroups) Create(awsSecurityGroup *v1alpha1.AwsSecurityGroup) (result *v1alpha1.AwsSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awssecuritygroupsResource, awsSecurityGroup), &v1alpha1.AwsSecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroup), err
}

// Update takes the representation of a awsSecurityGroup and updates it. Returns the server's representation of the awsSecurityGroup, and an error, if there is any.
func (c *FakeAwsSecurityGroups) Update(awsSecurityGroup *v1alpha1.AwsSecurityGroup) (result *v1alpha1.AwsSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awssecuritygroupsResource, awsSecurityGroup), &v1alpha1.AwsSecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSecurityGroups) UpdateStatus(awsSecurityGroup *v1alpha1.AwsSecurityGroup) (*v1alpha1.AwsSecurityGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awssecuritygroupsResource, "status", awsSecurityGroup), &v1alpha1.AwsSecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroup), err
}

// Delete takes name of the awsSecurityGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awssecuritygroupsResource, name), &v1alpha1.AwsSecurityGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awssecuritygroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSecurityGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsSecurityGroup.
func (c *FakeAwsSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awssecuritygroupsResource, name, pt, data, subresources...), &v1alpha1.AwsSecurityGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecurityGroup), err
}
