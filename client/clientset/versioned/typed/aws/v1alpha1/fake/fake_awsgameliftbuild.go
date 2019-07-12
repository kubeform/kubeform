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

// FakeAwsGameliftBuilds implements AwsGameliftBuildInterface
type FakeAwsGameliftBuilds struct {
	Fake *FakeAwsV1alpha1
}

var awsgameliftbuildsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsgameliftbuilds"}

var awsgameliftbuildsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsGameliftBuild"}

// Get takes name of the awsGameliftBuild, and returns the corresponding awsGameliftBuild object, and an error if there is any.
func (c *FakeAwsGameliftBuilds) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGameliftBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsgameliftbuildsResource, name), &v1alpha1.AwsGameliftBuild{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftBuild), err
}

// List takes label and field selectors, and returns the list of AwsGameliftBuilds that match those selectors.
func (c *FakeAwsGameliftBuilds) List(opts v1.ListOptions) (result *v1alpha1.AwsGameliftBuildList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsgameliftbuildsResource, awsgameliftbuildsKind, opts), &v1alpha1.AwsGameliftBuildList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGameliftBuildList{ListMeta: obj.(*v1alpha1.AwsGameliftBuildList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsGameliftBuildList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGameliftBuilds.
func (c *FakeAwsGameliftBuilds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsgameliftbuildsResource, opts))
}

// Create takes the representation of a awsGameliftBuild and creates it.  Returns the server's representation of the awsGameliftBuild, and an error, if there is any.
func (c *FakeAwsGameliftBuilds) Create(awsGameliftBuild *v1alpha1.AwsGameliftBuild) (result *v1alpha1.AwsGameliftBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsgameliftbuildsResource, awsGameliftBuild), &v1alpha1.AwsGameliftBuild{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftBuild), err
}

// Update takes the representation of a awsGameliftBuild and updates it. Returns the server's representation of the awsGameliftBuild, and an error, if there is any.
func (c *FakeAwsGameliftBuilds) Update(awsGameliftBuild *v1alpha1.AwsGameliftBuild) (result *v1alpha1.AwsGameliftBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsgameliftbuildsResource, awsGameliftBuild), &v1alpha1.AwsGameliftBuild{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftBuild), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsGameliftBuilds) UpdateStatus(awsGameliftBuild *v1alpha1.AwsGameliftBuild) (*v1alpha1.AwsGameliftBuild, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsgameliftbuildsResource, "status", awsGameliftBuild), &v1alpha1.AwsGameliftBuild{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftBuild), err
}

// Delete takes name of the awsGameliftBuild and deletes it. Returns an error if one occurs.
func (c *FakeAwsGameliftBuilds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsgameliftbuildsResource, name), &v1alpha1.AwsGameliftBuild{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGameliftBuilds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsgameliftbuildsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGameliftBuildList{})
	return err
}

// Patch applies the patch and returns the patched awsGameliftBuild.
func (c *FakeAwsGameliftBuilds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGameliftBuild, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsgameliftbuildsResource, name, pt, data, subresources...), &v1alpha1.AwsGameliftBuild{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGameliftBuild), err
}
