/*
Copyright 2019 The Kubeform Authors.

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

// FakeSagemakerEndpointConfigurations implements SagemakerEndpointConfigurationInterface
type FakeSagemakerEndpointConfigurations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var sagemakerendpointconfigurationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "sagemakerendpointconfigurations"}

var sagemakerendpointconfigurationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SagemakerEndpointConfiguration"}

// Get takes name of the sagemakerEndpointConfiguration, and returns the corresponding sagemakerEndpointConfiguration object, and an error if there is any.
func (c *FakeSagemakerEndpointConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.SagemakerEndpointConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sagemakerendpointconfigurationsResource, c.ns, name), &v1alpha1.SagemakerEndpointConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerEndpointConfiguration), err
}

// List takes label and field selectors, and returns the list of SagemakerEndpointConfigurations that match those selectors.
func (c *FakeSagemakerEndpointConfigurations) List(opts v1.ListOptions) (result *v1alpha1.SagemakerEndpointConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sagemakerendpointconfigurationsResource, sagemakerendpointconfigurationsKind, c.ns, opts), &v1alpha1.SagemakerEndpointConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SagemakerEndpointConfigurationList{ListMeta: obj.(*v1alpha1.SagemakerEndpointConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SagemakerEndpointConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sagemakerEndpointConfigurations.
func (c *FakeSagemakerEndpointConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sagemakerendpointconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a sagemakerEndpointConfiguration and creates it.  Returns the server's representation of the sagemakerEndpointConfiguration, and an error, if there is any.
func (c *FakeSagemakerEndpointConfigurations) Create(sagemakerEndpointConfiguration *v1alpha1.SagemakerEndpointConfiguration) (result *v1alpha1.SagemakerEndpointConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sagemakerendpointconfigurationsResource, c.ns, sagemakerEndpointConfiguration), &v1alpha1.SagemakerEndpointConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerEndpointConfiguration), err
}

// Update takes the representation of a sagemakerEndpointConfiguration and updates it. Returns the server's representation of the sagemakerEndpointConfiguration, and an error, if there is any.
func (c *FakeSagemakerEndpointConfigurations) Update(sagemakerEndpointConfiguration *v1alpha1.SagemakerEndpointConfiguration) (result *v1alpha1.SagemakerEndpointConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sagemakerendpointconfigurationsResource, c.ns, sagemakerEndpointConfiguration), &v1alpha1.SagemakerEndpointConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerEndpointConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSagemakerEndpointConfigurations) UpdateStatus(sagemakerEndpointConfiguration *v1alpha1.SagemakerEndpointConfiguration) (*v1alpha1.SagemakerEndpointConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sagemakerendpointconfigurationsResource, "status", c.ns, sagemakerEndpointConfiguration), &v1alpha1.SagemakerEndpointConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerEndpointConfiguration), err
}

// Delete takes name of the sagemakerEndpointConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeSagemakerEndpointConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sagemakerendpointconfigurationsResource, c.ns, name), &v1alpha1.SagemakerEndpointConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSagemakerEndpointConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sagemakerendpointconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SagemakerEndpointConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched sagemakerEndpointConfiguration.
func (c *FakeSagemakerEndpointConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SagemakerEndpointConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sagemakerendpointconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SagemakerEndpointConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerEndpointConfiguration), err
}