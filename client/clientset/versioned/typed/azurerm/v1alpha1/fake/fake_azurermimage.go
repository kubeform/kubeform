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

// FakeAzurermImages implements AzurermImageInterface
type FakeAzurermImages struct {
	Fake *FakeAzurermV1alpha1
}

var azurermimagesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermimages"}

var azurermimagesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermImage"}

// Get takes name of the azurermImage, and returns the corresponding azurermImage object, and an error if there is any.
func (c *FakeAzurermImages) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermimagesResource, name), &v1alpha1.AzurermImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermImage), err
}

// List takes label and field selectors, and returns the list of AzurermImages that match those selectors.
func (c *FakeAzurermImages) List(opts v1.ListOptions) (result *v1alpha1.AzurermImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermimagesResource, azurermimagesKind, opts), &v1alpha1.AzurermImageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermImageList{ListMeta: obj.(*v1alpha1.AzurermImageList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermImages.
func (c *FakeAzurermImages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermimagesResource, opts))
}

// Create takes the representation of a azurermImage and creates it.  Returns the server's representation of the azurermImage, and an error, if there is any.
func (c *FakeAzurermImages) Create(azurermImage *v1alpha1.AzurermImage) (result *v1alpha1.AzurermImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermimagesResource, azurermImage), &v1alpha1.AzurermImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermImage), err
}

// Update takes the representation of a azurermImage and updates it. Returns the server's representation of the azurermImage, and an error, if there is any.
func (c *FakeAzurermImages) Update(azurermImage *v1alpha1.AzurermImage) (result *v1alpha1.AzurermImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermimagesResource, azurermImage), &v1alpha1.AzurermImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermImage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermImages) UpdateStatus(azurermImage *v1alpha1.AzurermImage) (*v1alpha1.AzurermImage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermimagesResource, "status", azurermImage), &v1alpha1.AzurermImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermImage), err
}

// Delete takes name of the azurermImage and deletes it. Returns an error if one occurs.
func (c *FakeAzurermImages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermimagesResource, name), &v1alpha1.AzurermImage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermImages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermimagesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermImageList{})
	return err
}

// Patch applies the patch and returns the patched azurermImage.
func (c *FakeAzurermImages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermimagesResource, name, pt, data, subresources...), &v1alpha1.AzurermImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermImage), err
}
