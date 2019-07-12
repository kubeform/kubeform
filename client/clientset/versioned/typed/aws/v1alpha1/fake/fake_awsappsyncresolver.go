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

// FakeAwsAppsyncResolvers implements AwsAppsyncResolverInterface
type FakeAwsAppsyncResolvers struct {
	Fake *FakeAwsV1alpha1
}

var awsappsyncresolversResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsappsyncresolvers"}

var awsappsyncresolversKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsAppsyncResolver"}

// Get takes name of the awsAppsyncResolver, and returns the corresponding awsAppsyncResolver object, and an error if there is any.
func (c *FakeAwsAppsyncResolvers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppsyncResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsappsyncresolversResource, name), &v1alpha1.AwsAppsyncResolver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncResolver), err
}

// List takes label and field selectors, and returns the list of AwsAppsyncResolvers that match those selectors.
func (c *FakeAwsAppsyncResolvers) List(opts v1.ListOptions) (result *v1alpha1.AwsAppsyncResolverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsappsyncresolversResource, awsappsyncresolversKind, opts), &v1alpha1.AwsAppsyncResolverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsAppsyncResolverList{ListMeta: obj.(*v1alpha1.AwsAppsyncResolverList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsAppsyncResolverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAppsyncResolvers.
func (c *FakeAwsAppsyncResolvers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsappsyncresolversResource, opts))
}

// Create takes the representation of a awsAppsyncResolver and creates it.  Returns the server's representation of the awsAppsyncResolver, and an error, if there is any.
func (c *FakeAwsAppsyncResolvers) Create(awsAppsyncResolver *v1alpha1.AwsAppsyncResolver) (result *v1alpha1.AwsAppsyncResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsappsyncresolversResource, awsAppsyncResolver), &v1alpha1.AwsAppsyncResolver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncResolver), err
}

// Update takes the representation of a awsAppsyncResolver and updates it. Returns the server's representation of the awsAppsyncResolver, and an error, if there is any.
func (c *FakeAwsAppsyncResolvers) Update(awsAppsyncResolver *v1alpha1.AwsAppsyncResolver) (result *v1alpha1.AwsAppsyncResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsappsyncresolversResource, awsAppsyncResolver), &v1alpha1.AwsAppsyncResolver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncResolver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsAppsyncResolvers) UpdateStatus(awsAppsyncResolver *v1alpha1.AwsAppsyncResolver) (*v1alpha1.AwsAppsyncResolver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsappsyncresolversResource, "status", awsAppsyncResolver), &v1alpha1.AwsAppsyncResolver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncResolver), err
}

// Delete takes name of the awsAppsyncResolver and deletes it. Returns an error if one occurs.
func (c *FakeAwsAppsyncResolvers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsappsyncresolversResource, name), &v1alpha1.AwsAppsyncResolver{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAppsyncResolvers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsappsyncresolversResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsAppsyncResolverList{})
	return err
}

// Patch applies the patch and returns the patched awsAppsyncResolver.
func (c *FakeAwsAppsyncResolvers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppsyncResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsappsyncresolversResource, name, pt, data, subresources...), &v1alpha1.AwsAppsyncResolver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsAppsyncResolver), err
}
