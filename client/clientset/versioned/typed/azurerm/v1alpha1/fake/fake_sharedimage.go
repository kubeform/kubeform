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

// FakeSharedImages implements SharedImageInterface
type FakeSharedImages struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var sharedimagesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "sharedimages"}

var sharedimagesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SharedImage"}

// Get takes name of the sharedImage, and returns the corresponding sharedImage object, and an error if there is any.
func (c *FakeSharedImages) Get(name string, options v1.GetOptions) (result *v1alpha1.SharedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharedimagesResource, c.ns, name), &v1alpha1.SharedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedImage), err
}

// List takes label and field selectors, and returns the list of SharedImages that match those selectors.
func (c *FakeSharedImages) List(opts v1.ListOptions) (result *v1alpha1.SharedImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharedimagesResource, sharedimagesKind, c.ns, opts), &v1alpha1.SharedImageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SharedImageList{ListMeta: obj.(*v1alpha1.SharedImageList).ListMeta}
	for _, item := range obj.(*v1alpha1.SharedImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharedImages.
func (c *FakeSharedImages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharedimagesResource, c.ns, opts))

}

// Create takes the representation of a sharedImage and creates it.  Returns the server's representation of the sharedImage, and an error, if there is any.
func (c *FakeSharedImages) Create(sharedImage *v1alpha1.SharedImage) (result *v1alpha1.SharedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharedimagesResource, c.ns, sharedImage), &v1alpha1.SharedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedImage), err
}

// Update takes the representation of a sharedImage and updates it. Returns the server's representation of the sharedImage, and an error, if there is any.
func (c *FakeSharedImages) Update(sharedImage *v1alpha1.SharedImage) (result *v1alpha1.SharedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharedimagesResource, c.ns, sharedImage), &v1alpha1.SharedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedImage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharedImages) UpdateStatus(sharedImage *v1alpha1.SharedImage) (*v1alpha1.SharedImage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharedimagesResource, "status", c.ns, sharedImage), &v1alpha1.SharedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedImage), err
}

// Delete takes name of the sharedImage and deletes it. Returns an error if one occurs.
func (c *FakeSharedImages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharedimagesResource, c.ns, name), &v1alpha1.SharedImage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharedImages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharedimagesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SharedImageList{})
	return err
}

// Patch applies the patch and returns the patched sharedImage.
func (c *FakeSharedImages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SharedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharedimagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SharedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SharedImage), err
}
