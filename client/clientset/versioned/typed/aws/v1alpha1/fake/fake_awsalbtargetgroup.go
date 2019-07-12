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

// FakeAwsAlbTargetGroups implements AwsAlbTargetGroupInterface
type FakeAwsAlbTargetGroups struct {
	Fake *FakeAwsV1alpha1
}

var awsalbtargetgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsalbtargetgroups"}

var awsalbtargetgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsAlbTargetGroup"}

// Get takes name of the awsAlbTargetGroup, and returns the corresponding awsAlbTargetGroup object, and an error if there is any.
func (c *FakeAwsAlbTargetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAlbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsalbtargetgroupsResource, name), &v1alpha1.AwsAlbTargetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlbTargetGroup), err
}

// List takes label and field selectors, and returns the list of AwsAlbTargetGroups that match those selectors.
func (c *FakeAwsAlbTargetGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsAlbTargetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsalbtargetgroupsResource, awsalbtargetgroupsKind, opts), &v1alpha1.AwsAlbTargetGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAlbTargetGroupList{ListMeta: obj.(*v1alpha1.AwsAlbTargetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsAlbTargetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAlbTargetGroups.
func (c *FakeAwsAlbTargetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsalbtargetgroupsResource, opts))
}

// Create takes the representation of a awsAlbTargetGroup and creates it.  Returns the server's representation of the awsAlbTargetGroup, and an error, if there is any.
func (c *FakeAwsAlbTargetGroups) Create(awsAlbTargetGroup *v1alpha1.AwsAlbTargetGroup) (result *v1alpha1.AwsAlbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsalbtargetgroupsResource, awsAlbTargetGroup), &v1alpha1.AwsAlbTargetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlbTargetGroup), err
}

// Update takes the representation of a awsAlbTargetGroup and updates it. Returns the server's representation of the awsAlbTargetGroup, and an error, if there is any.
func (c *FakeAwsAlbTargetGroups) Update(awsAlbTargetGroup *v1alpha1.AwsAlbTargetGroup) (result *v1alpha1.AwsAlbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsalbtargetgroupsResource, awsAlbTargetGroup), &v1alpha1.AwsAlbTargetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlbTargetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsAlbTargetGroups) UpdateStatus(awsAlbTargetGroup *v1alpha1.AwsAlbTargetGroup) (*v1alpha1.AwsAlbTargetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsalbtargetgroupsResource, "status", awsAlbTargetGroup), &v1alpha1.AwsAlbTargetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlbTargetGroup), err
}

// Delete takes name of the awsAlbTargetGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsAlbTargetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsalbtargetgroupsResource, name), &v1alpha1.AwsAlbTargetGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAlbTargetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsalbtargetgroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAlbTargetGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsAlbTargetGroup.
func (c *FakeAwsAlbTargetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAlbTargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsalbtargetgroupsResource, name, pt, data, subresources...), &v1alpha1.AwsAlbTargetGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAlbTargetGroup), err
}
