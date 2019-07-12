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

// FakeAwsMqConfigurations implements AwsMqConfigurationInterface
type FakeAwsMqConfigurations struct {
	Fake *FakeAwsV1alpha1
}

var awsmqconfigurationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsmqconfigurations"}

var awsmqconfigurationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsMqConfiguration"}

// Get takes name of the awsMqConfiguration, and returns the corresponding awsMqConfiguration object, and an error if there is any.
func (c *FakeAwsMqConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsMqConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsmqconfigurationsResource, name), &v1alpha1.AwsMqConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMqConfiguration), err
}

// List takes label and field selectors, and returns the list of AwsMqConfigurations that match those selectors.
func (c *FakeAwsMqConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AwsMqConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsmqconfigurationsResource, awsmqconfigurationsKind, opts), &v1alpha1.AwsMqConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsMqConfigurationList{ListMeta: obj.(*v1alpha1.AwsMqConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsMqConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsMqConfigurations.
func (c *FakeAwsMqConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsmqconfigurationsResource, opts))
}

// Create takes the representation of a awsMqConfiguration and creates it.  Returns the server's representation of the awsMqConfiguration, and an error, if there is any.
func (c *FakeAwsMqConfigurations) Create(awsMqConfiguration *v1alpha1.AwsMqConfiguration) (result *v1alpha1.AwsMqConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsmqconfigurationsResource, awsMqConfiguration), &v1alpha1.AwsMqConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMqConfiguration), err
}

// Update takes the representation of a awsMqConfiguration and updates it. Returns the server's representation of the awsMqConfiguration, and an error, if there is any.
func (c *FakeAwsMqConfigurations) Update(awsMqConfiguration *v1alpha1.AwsMqConfiguration) (result *v1alpha1.AwsMqConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsmqconfigurationsResource, awsMqConfiguration), &v1alpha1.AwsMqConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMqConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsMqConfigurations) UpdateStatus(awsMqConfiguration *v1alpha1.AwsMqConfiguration) (*v1alpha1.AwsMqConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsmqconfigurationsResource, "status", awsMqConfiguration), &v1alpha1.AwsMqConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMqConfiguration), err
}

// Delete takes name of the awsMqConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeAwsMqConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsmqconfigurationsResource, name), &v1alpha1.AwsMqConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsMqConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsmqconfigurationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsMqConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched awsMqConfiguration.
func (c *FakeAwsMqConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsMqConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsmqconfigurationsResource, name, pt, data, subresources...), &v1alpha1.AwsMqConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMqConfiguration), err
}
