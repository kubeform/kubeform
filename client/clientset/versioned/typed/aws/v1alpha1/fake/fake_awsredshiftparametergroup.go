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

// FakeAwsRedshiftParameterGroups implements AwsRedshiftParameterGroupInterface
type FakeAwsRedshiftParameterGroups struct {
	Fake *FakeAwsV1alpha1
}

var awsredshiftparametergroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsredshiftparametergroups"}

var awsredshiftparametergroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsRedshiftParameterGroup"}

// Get takes name of the awsRedshiftParameterGroup, and returns the corresponding awsRedshiftParameterGroup object, and an error if there is any.
func (c *FakeAwsRedshiftParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsRedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsredshiftparametergroupsResource, name), &v1alpha1.AwsRedshiftParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftParameterGroup), err
}

// List takes label and field selectors, and returns the list of AwsRedshiftParameterGroups that match those selectors.
func (c *FakeAwsRedshiftParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsRedshiftParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsredshiftparametergroupsResource, awsredshiftparametergroupsKind, opts), &v1alpha1.AwsRedshiftParameterGroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsRedshiftParameterGroupList{ListMeta: obj.(*v1alpha1.AwsRedshiftParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsRedshiftParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsRedshiftParameterGroups.
func (c *FakeAwsRedshiftParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsredshiftparametergroupsResource, opts))
}

// Create takes the representation of a awsRedshiftParameterGroup and creates it.  Returns the server's representation of the awsRedshiftParameterGroup, and an error, if there is any.
func (c *FakeAwsRedshiftParameterGroups) Create(awsRedshiftParameterGroup *v1alpha1.AwsRedshiftParameterGroup) (result *v1alpha1.AwsRedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsredshiftparametergroupsResource, awsRedshiftParameterGroup), &v1alpha1.AwsRedshiftParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftParameterGroup), err
}

// Update takes the representation of a awsRedshiftParameterGroup and updates it. Returns the server's representation of the awsRedshiftParameterGroup, and an error, if there is any.
func (c *FakeAwsRedshiftParameterGroups) Update(awsRedshiftParameterGroup *v1alpha1.AwsRedshiftParameterGroup) (result *v1alpha1.AwsRedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsredshiftparametergroupsResource, awsRedshiftParameterGroup), &v1alpha1.AwsRedshiftParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsRedshiftParameterGroups) UpdateStatus(awsRedshiftParameterGroup *v1alpha1.AwsRedshiftParameterGroup) (*v1alpha1.AwsRedshiftParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsredshiftparametergroupsResource, "status", awsRedshiftParameterGroup), &v1alpha1.AwsRedshiftParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftParameterGroup), err
}

// Delete takes name of the awsRedshiftParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeAwsRedshiftParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsredshiftparametergroupsResource, name), &v1alpha1.AwsRedshiftParameterGroup{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsRedshiftParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsredshiftparametergroupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsRedshiftParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched awsRedshiftParameterGroup.
func (c *FakeAwsRedshiftParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsRedshiftParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsredshiftparametergroupsResource, name, pt, data, subresources...), &v1alpha1.AwsRedshiftParameterGroup{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsRedshiftParameterGroup), err
}
