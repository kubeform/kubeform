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

// FakeDxPublicVirtualInterfaces implements DxPublicVirtualInterfaceInterface
type FakeDxPublicVirtualInterfaces struct {
	Fake *FakeAwsV1alpha1
}

var dxpublicvirtualinterfacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxpublicvirtualinterfaces"}

var dxpublicvirtualinterfacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxPublicVirtualInterface"}

// Get takes name of the dxPublicVirtualInterface, and returns the corresponding dxPublicVirtualInterface object, and an error if there is any.
func (c *FakeDxPublicVirtualInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.DxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dxpublicvirtualinterfacesResource, name), &v1alpha1.DxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPublicVirtualInterface), err
}

// List takes label and field selectors, and returns the list of DxPublicVirtualInterfaces that match those selectors.
func (c *FakeDxPublicVirtualInterfaces) List(opts v1.ListOptions) (result *v1alpha1.DxPublicVirtualInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dxpublicvirtualinterfacesResource, dxpublicvirtualinterfacesKind, opts), &v1alpha1.DxPublicVirtualInterfaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxPublicVirtualInterfaceList{ListMeta: obj.(*v1alpha1.DxPublicVirtualInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxPublicVirtualInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxPublicVirtualInterfaces.
func (c *FakeDxPublicVirtualInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dxpublicvirtualinterfacesResource, opts))
}

// Create takes the representation of a dxPublicVirtualInterface and creates it.  Returns the server's representation of the dxPublicVirtualInterface, and an error, if there is any.
func (c *FakeDxPublicVirtualInterfaces) Create(dxPublicVirtualInterface *v1alpha1.DxPublicVirtualInterface) (result *v1alpha1.DxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dxpublicvirtualinterfacesResource, dxPublicVirtualInterface), &v1alpha1.DxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPublicVirtualInterface), err
}

// Update takes the representation of a dxPublicVirtualInterface and updates it. Returns the server's representation of the dxPublicVirtualInterface, and an error, if there is any.
func (c *FakeDxPublicVirtualInterfaces) Update(dxPublicVirtualInterface *v1alpha1.DxPublicVirtualInterface) (result *v1alpha1.DxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dxpublicvirtualinterfacesResource, dxPublicVirtualInterface), &v1alpha1.DxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPublicVirtualInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxPublicVirtualInterfaces) UpdateStatus(dxPublicVirtualInterface *v1alpha1.DxPublicVirtualInterface) (*v1alpha1.DxPublicVirtualInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dxpublicvirtualinterfacesResource, "status", dxPublicVirtualInterface), &v1alpha1.DxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPublicVirtualInterface), err
}

// Delete takes name of the dxPublicVirtualInterface and deletes it. Returns an error if one occurs.
func (c *FakeDxPublicVirtualInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dxpublicvirtualinterfacesResource, name), &v1alpha1.DxPublicVirtualInterface{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxPublicVirtualInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dxpublicvirtualinterfacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxPublicVirtualInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched dxPublicVirtualInterface.
func (c *FakeDxPublicVirtualInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxPublicVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dxpublicvirtualinterfacesResource, name, pt, data, subresources...), &v1alpha1.DxPublicVirtualInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxPublicVirtualInterface), err
}
