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

// FakePublicIPPrefixes implements PublicIPPrefixInterface
type FakePublicIPPrefixes struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var publicipprefixesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "publicipprefixes"}

var publicipprefixesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PublicIPPrefix"}

// Get takes name of the publicIPPrefix, and returns the corresponding publicIPPrefix object, and an error if there is any.
func (c *FakePublicIPPrefixes) Get(name string, options v1.GetOptions) (result *v1alpha1.PublicIPPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(publicipprefixesResource, c.ns, name), &v1alpha1.PublicIPPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIPPrefix), err
}

// List takes label and field selectors, and returns the list of PublicIPPrefixes that match those selectors.
func (c *FakePublicIPPrefixes) List(opts v1.ListOptions) (result *v1alpha1.PublicIPPrefixList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(publicipprefixesResource, publicipprefixesKind, c.ns, opts), &v1alpha1.PublicIPPrefixList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PublicIPPrefixList{ListMeta: obj.(*v1alpha1.PublicIPPrefixList).ListMeta}
	for _, item := range obj.(*v1alpha1.PublicIPPrefixList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested publicIPPrefixes.
func (c *FakePublicIPPrefixes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(publicipprefixesResource, c.ns, opts))

}

// Create takes the representation of a publicIPPrefix and creates it.  Returns the server's representation of the publicIPPrefix, and an error, if there is any.
func (c *FakePublicIPPrefixes) Create(publicIPPrefix *v1alpha1.PublicIPPrefix) (result *v1alpha1.PublicIPPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(publicipprefixesResource, c.ns, publicIPPrefix), &v1alpha1.PublicIPPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIPPrefix), err
}

// Update takes the representation of a publicIPPrefix and updates it. Returns the server's representation of the publicIPPrefix, and an error, if there is any.
func (c *FakePublicIPPrefixes) Update(publicIPPrefix *v1alpha1.PublicIPPrefix) (result *v1alpha1.PublicIPPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(publicipprefixesResource, c.ns, publicIPPrefix), &v1alpha1.PublicIPPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIPPrefix), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePublicIPPrefixes) UpdateStatus(publicIPPrefix *v1alpha1.PublicIPPrefix) (*v1alpha1.PublicIPPrefix, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(publicipprefixesResource, "status", c.ns, publicIPPrefix), &v1alpha1.PublicIPPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIPPrefix), err
}

// Delete takes name of the publicIPPrefix and deletes it. Returns an error if one occurs.
func (c *FakePublicIPPrefixes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(publicipprefixesResource, c.ns, name), &v1alpha1.PublicIPPrefix{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePublicIPPrefixes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(publicipprefixesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PublicIPPrefixList{})
	return err
}

// Patch applies the patch and returns the patched publicIPPrefix.
func (c *FakePublicIPPrefixes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PublicIPPrefix, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(publicipprefixesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PublicIPPrefix{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIPPrefix), err
}
