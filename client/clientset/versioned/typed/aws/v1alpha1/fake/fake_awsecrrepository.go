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

// FakeAwsEcrRepositories implements AwsEcrRepositoryInterface
type FakeAwsEcrRepositories struct {
	Fake *FakeAwsV1alpha1
}

var awsecrrepositoriesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsecrrepositories"}

var awsecrrepositoriesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsEcrRepository"}

// Get takes name of the awsEcrRepository, and returns the corresponding awsEcrRepository object, and an error if there is any.
func (c *FakeAwsEcrRepositories) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEcrRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsecrrepositoriesResource, name), &v1alpha1.AwsEcrRepository{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcrRepository), err
}

// List takes label and field selectors, and returns the list of AwsEcrRepositories that match those selectors.
func (c *FakeAwsEcrRepositories) List(opts v1.ListOptions) (result *v1alpha1.AwsEcrRepositoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsecrrepositoriesResource, awsecrrepositoriesKind, opts), &v1alpha1.AwsEcrRepositoryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsEcrRepositoryList{ListMeta: obj.(*v1alpha1.AwsEcrRepositoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsEcrRepositoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsEcrRepositories.
func (c *FakeAwsEcrRepositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsecrrepositoriesResource, opts))
}

// Create takes the representation of a awsEcrRepository and creates it.  Returns the server's representation of the awsEcrRepository, and an error, if there is any.
func (c *FakeAwsEcrRepositories) Create(awsEcrRepository *v1alpha1.AwsEcrRepository) (result *v1alpha1.AwsEcrRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsecrrepositoriesResource, awsEcrRepository), &v1alpha1.AwsEcrRepository{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcrRepository), err
}

// Update takes the representation of a awsEcrRepository and updates it. Returns the server's representation of the awsEcrRepository, and an error, if there is any.
func (c *FakeAwsEcrRepositories) Update(awsEcrRepository *v1alpha1.AwsEcrRepository) (result *v1alpha1.AwsEcrRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsecrrepositoriesResource, awsEcrRepository), &v1alpha1.AwsEcrRepository{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcrRepository), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsEcrRepositories) UpdateStatus(awsEcrRepository *v1alpha1.AwsEcrRepository) (*v1alpha1.AwsEcrRepository, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsecrrepositoriesResource, "status", awsEcrRepository), &v1alpha1.AwsEcrRepository{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcrRepository), err
}

// Delete takes name of the awsEcrRepository and deletes it. Returns an error if one occurs.
func (c *FakeAwsEcrRepositories) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsecrrepositoriesResource, name), &v1alpha1.AwsEcrRepository{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsEcrRepositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsecrrepositoriesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsEcrRepositoryList{})
	return err
}

// Patch applies the patch and returns the patched awsEcrRepository.
func (c *FakeAwsEcrRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcrRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsecrrepositoriesResource, name, pt, data, subresources...), &v1alpha1.AwsEcrRepository{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsEcrRepository), err
}
