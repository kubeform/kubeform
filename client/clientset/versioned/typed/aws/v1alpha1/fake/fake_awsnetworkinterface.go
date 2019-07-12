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

// FakeAwsNetworkInterfaces implements AwsNetworkInterfaceInterface
type FakeAwsNetworkInterfaces struct {
	Fake *FakeAwsV1alpha1
}

var awsnetworkinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsnetworkinterfaces"}

var awsnetworkinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsNetworkInterface"}

// Get takes name of the awsNetworkInterface, and returns the corresponding awsNetworkInterface object, and an error if there is any.
func (c *FakeAwsNetworkInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsnetworkinterfacesResource, name), &v1alpha1.AwsNetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterface), err
}

// List takes label and field selectors, and returns the list of AwsNetworkInterfaces that match those selectors.
func (c *FakeAwsNetworkInterfaces) List(opts v1.ListOptions) (result *v1alpha1.AwsNetworkInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsnetworkinterfacesResource, awsnetworkinterfacesKind, opts), &v1alpha1.AwsNetworkInterfaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsNetworkInterfaceList{ListMeta: obj.(*v1alpha1.AwsNetworkInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsNetworkInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsNetworkInterfaces.
func (c *FakeAwsNetworkInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsnetworkinterfacesResource, opts))
}

// Create takes the representation of a awsNetworkInterface and creates it.  Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *FakeAwsNetworkInterfaces) Create(awsNetworkInterface *v1alpha1.AwsNetworkInterface) (result *v1alpha1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsnetworkinterfacesResource, awsNetworkInterface), &v1alpha1.AwsNetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterface), err
}

// Update takes the representation of a awsNetworkInterface and updates it. Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *FakeAwsNetworkInterfaces) Update(awsNetworkInterface *v1alpha1.AwsNetworkInterface) (result *v1alpha1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsnetworkinterfacesResource, awsNetworkInterface), &v1alpha1.AwsNetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsNetworkInterfaces) UpdateStatus(awsNetworkInterface *v1alpha1.AwsNetworkInterface) (*v1alpha1.AwsNetworkInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsnetworkinterfacesResource, "status", awsNetworkInterface), &v1alpha1.AwsNetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterface), err
}

// Delete takes name of the awsNetworkInterface and deletes it. Returns an error if one occurs.
func (c *FakeAwsNetworkInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsnetworkinterfacesResource, name), &v1alpha1.AwsNetworkInterface{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsNetworkInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsnetworkinterfacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsNetworkInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched awsNetworkInterface.
func (c *FakeAwsNetworkInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsnetworkinterfacesResource, name, pt, data, subresources...), &v1alpha1.AwsNetworkInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsNetworkInterface), err
}
