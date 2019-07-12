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

// FakeAwsIamGroups implements AwsIamGroupInterface
type FakeAwsIamGroups struct {
	Fake *FakeAwsV1alpha1
}

var awsiamgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsiamgroups"}

var awsiamgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsIamGroup"}

// Get takes name of the awsIamGroup, and returns the corresponding awsIamGroup object, and an error if there is any.
func (c *FakeAwsIamGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsiamgroupsResource, name), &v1alpha1.AwsIamGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroup), err
}

// List takes label and field selectors, and returns the list of AwsIamGroups that match those selectors.
func (c *FakeAwsIamGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsIamGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsiamgroupsResource, awsiamgroupsKind, opts), &v1alpha1.AwsIamGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsIamGroupList{ListMeta: obj.(*v1alpha1.AwsIamGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsIamGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIamGroups.
func (c *FakeAwsIamGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsiamgroupsResource, opts))
}

// Create takes the representation of a awsIamGroup and creates it.  Returns the server's representation of the awsIamGroup, and an error, if there is any.
func (c *FakeAwsIamGroups) Create(awsIamGroup *v1alpha1.AwsIamGroup) (result *v1alpha1.AwsIamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsiamgroupsResource, awsIamGroup), &v1alpha1.AwsIamGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroup), err
}

// Update takes the representation of a awsIamGroup and updates it. Returns the server's representation of the awsIamGroup, and an error, if there is any.
func (c *FakeAwsIamGroups) Update(awsIamGroup *v1alpha1.AwsIamGroup) (result *v1alpha1.AwsIamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsiamgroupsResource, awsIamGroup), &v1alpha1.AwsIamGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsIamGroups) UpdateStatus(awsIamGroup *v1alpha1.AwsIamGroup) (*v1alpha1.AwsIamGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsiamgroupsResource, "status", awsIamGroup), &v1alpha1.AwsIamGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroup), err
}

// Delete takes name of the awsIamGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsIamGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsiamgroupsResource, name), &v1alpha1.AwsIamGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIamGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsiamgroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsIamGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsIamGroup.
func (c *FakeAwsIamGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIamGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsiamgroupsResource, name, pt, data, subresources...), &v1alpha1.AwsIamGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsIamGroup), err
}
