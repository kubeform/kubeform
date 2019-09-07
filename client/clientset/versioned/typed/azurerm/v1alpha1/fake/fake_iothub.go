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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeIothubs implements IothubInterface
type FakeIothubs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var iothubsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "iothubs"}

var iothubsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "Iothub"}

// Get takes name of the iothub, and returns the corresponding iothub object, and an error if there is any.
func (c *FakeIothubs) Get(name string, options v1.GetOptions) (result *v1alpha1.Iothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iothubsResource, c.ns, name), &v1alpha1.Iothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iothub), err
}

// List takes label and field selectors, and returns the list of Iothubs that match those selectors.
func (c *FakeIothubs) List(opts v1.ListOptions) (result *v1alpha1.IothubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iothubsResource, iothubsKind, c.ns, opts), &v1alpha1.IothubList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IothubList{ListMeta: obj.(*v1alpha1.IothubList).ListMeta}
	for _, item := range obj.(*v1alpha1.IothubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iothubs.
func (c *FakeIothubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iothubsResource, c.ns, opts))

}

// Create takes the representation of a iothub and creates it.  Returns the server's representation of the iothub, and an error, if there is any.
func (c *FakeIothubs) Create(iothub *v1alpha1.Iothub) (result *v1alpha1.Iothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iothubsResource, c.ns, iothub), &v1alpha1.Iothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iothub), err
}

// Update takes the representation of a iothub and updates it. Returns the server's representation of the iothub, and an error, if there is any.
func (c *FakeIothubs) Update(iothub *v1alpha1.Iothub) (result *v1alpha1.Iothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iothubsResource, c.ns, iothub), &v1alpha1.Iothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iothub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIothubs) UpdateStatus(iothub *v1alpha1.Iothub) (*v1alpha1.Iothub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iothubsResource, "status", c.ns, iothub), &v1alpha1.Iothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iothub), err
}

// Delete takes name of the iothub and deletes it. Returns an error if one occurs.
func (c *FakeIothubs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iothubsResource, c.ns, name), &v1alpha1.Iothub{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIothubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iothubsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IothubList{})
	return err
}

// Patch applies the patch and returns the patched iothub.
func (c *FakeIothubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Iothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iothubsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Iothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iothub), err
}