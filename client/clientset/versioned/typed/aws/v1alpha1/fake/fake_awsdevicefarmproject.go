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

// FakeAwsDevicefarmProjects implements AwsDevicefarmProjectInterface
type FakeAwsDevicefarmProjects struct {
	Fake *FakeAwsV1alpha1
}

var awsdevicefarmprojectsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdevicefarmprojects"}

var awsdevicefarmprojectsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDevicefarmProject"}

// Get takes name of the awsDevicefarmProject, and returns the corresponding awsDevicefarmProject object, and an error if there is any.
func (c *FakeAwsDevicefarmProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdevicefarmprojectsResource, name), &v1alpha1.AwsDevicefarmProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDevicefarmProject), err
}

// List takes label and field selectors, and returns the list of AwsDevicefarmProjects that match those selectors.
func (c *FakeAwsDevicefarmProjects) List(opts v1.ListOptions) (result *v1alpha1.AwsDevicefarmProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdevicefarmprojectsResource, awsdevicefarmprojectsKind, opts), &v1alpha1.AwsDevicefarmProjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDevicefarmProjectList{ListMeta: obj.(*v1alpha1.AwsDevicefarmProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDevicefarmProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDevicefarmProjects.
func (c *FakeAwsDevicefarmProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdevicefarmprojectsResource, opts))
}

// Create takes the representation of a awsDevicefarmProject and creates it.  Returns the server's representation of the awsDevicefarmProject, and an error, if there is any.
func (c *FakeAwsDevicefarmProjects) Create(awsDevicefarmProject *v1alpha1.AwsDevicefarmProject) (result *v1alpha1.AwsDevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdevicefarmprojectsResource, awsDevicefarmProject), &v1alpha1.AwsDevicefarmProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDevicefarmProject), err
}

// Update takes the representation of a awsDevicefarmProject and updates it. Returns the server's representation of the awsDevicefarmProject, and an error, if there is any.
func (c *FakeAwsDevicefarmProjects) Update(awsDevicefarmProject *v1alpha1.AwsDevicefarmProject) (result *v1alpha1.AwsDevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdevicefarmprojectsResource, awsDevicefarmProject), &v1alpha1.AwsDevicefarmProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDevicefarmProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDevicefarmProjects) UpdateStatus(awsDevicefarmProject *v1alpha1.AwsDevicefarmProject) (*v1alpha1.AwsDevicefarmProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdevicefarmprojectsResource, "status", awsDevicefarmProject), &v1alpha1.AwsDevicefarmProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDevicefarmProject), err
}

// Delete takes name of the awsDevicefarmProject and deletes it. Returns an error if one occurs.
func (c *FakeAwsDevicefarmProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdevicefarmprojectsResource, name), &v1alpha1.AwsDevicefarmProject{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDevicefarmProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdevicefarmprojectsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDevicefarmProjectList{})
	return err
}

// Patch applies the patch and returns the patched awsDevicefarmProject.
func (c *FakeAwsDevicefarmProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDevicefarmProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdevicefarmprojectsResource, name, pt, data, subresources...), &v1alpha1.AwsDevicefarmProject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDevicefarmProject), err
}
