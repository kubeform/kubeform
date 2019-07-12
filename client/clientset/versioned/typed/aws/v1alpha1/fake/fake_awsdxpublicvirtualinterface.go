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

// FakeAwsDxPublicVirtualInterfaces implements AwsDxPublicVirtualInterfaceInterface
type FakeAwsDxPublicVirtualInterfaces struct {
	Fake *FakeAwsV1alpha1
}

var awsdxpublicvirtualinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdxpublicvirtualinterfaces"}

var awsdxpublicvirtualinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDxPublicVirtualInterface"}

// Get takes name of the awsDxPublicVirtualInterface, and returns the corresponding awsDxPublicVirtualInterface object, and an error if there is any.
func (c *FakeAwsDxPublicVirtualInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdxpublicvirtualinterfacesResource, name), &v1alpha1.AwsDxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxPublicVirtualInterface), err
}

// List takes label and field selectors, and returns the list of AwsDxPublicVirtualInterfaces that match those selectors.
func (c *FakeAwsDxPublicVirtualInterfaces) List(opts v1.ListOptions) (result *v1alpha1.AwsDxPublicVirtualInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdxpublicvirtualinterfacesResource, awsdxpublicvirtualinterfacesKind, opts), &v1alpha1.AwsDxPublicVirtualInterfaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDxPublicVirtualInterfaceList{ListMeta: obj.(*v1alpha1.AwsDxPublicVirtualInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDxPublicVirtualInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDxPublicVirtualInterfaces.
func (c *FakeAwsDxPublicVirtualInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdxpublicvirtualinterfacesResource, opts))
}

// Create takes the representation of a awsDxPublicVirtualInterface and creates it.  Returns the server's representation of the awsDxPublicVirtualInterface, and an error, if there is any.
func (c *FakeAwsDxPublicVirtualInterfaces) Create(awsDxPublicVirtualInterface *v1alpha1.AwsDxPublicVirtualInterface) (result *v1alpha1.AwsDxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdxpublicvirtualinterfacesResource, awsDxPublicVirtualInterface), &v1alpha1.AwsDxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxPublicVirtualInterface), err
}

// Update takes the representation of a awsDxPublicVirtualInterface and updates it. Returns the server's representation of the awsDxPublicVirtualInterface, and an error, if there is any.
func (c *FakeAwsDxPublicVirtualInterfaces) Update(awsDxPublicVirtualInterface *v1alpha1.AwsDxPublicVirtualInterface) (result *v1alpha1.AwsDxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdxpublicvirtualinterfacesResource, awsDxPublicVirtualInterface), &v1alpha1.AwsDxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxPublicVirtualInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDxPublicVirtualInterfaces) UpdateStatus(awsDxPublicVirtualInterface *v1alpha1.AwsDxPublicVirtualInterface) (*v1alpha1.AwsDxPublicVirtualInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdxpublicvirtualinterfacesResource, "status", awsDxPublicVirtualInterface), &v1alpha1.AwsDxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxPublicVirtualInterface), err
}

// Delete takes name of the awsDxPublicVirtualInterface and deletes it. Returns an error if one occurs.
func (c *FakeAwsDxPublicVirtualInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdxpublicvirtualinterfacesResource, name), &v1alpha1.AwsDxPublicVirtualInterface{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDxPublicVirtualInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdxpublicvirtualinterfacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDxPublicVirtualInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched awsDxPublicVirtualInterface.
func (c *FakeAwsDxPublicVirtualInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdxpublicvirtualinterfacesResource, name, pt, data, subresources...), &v1alpha1.AwsDxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDxPublicVirtualInterface), err
}
