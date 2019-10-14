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

// FakeAvailabilitySets implements AvailabilitySetInterface
type FakeAvailabilitySets struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var availabilitysetsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "availabilitysets"}

var availabilitysetsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AvailabilitySet"}

// Get takes name of the availabilitySet, and returns the corresponding availabilitySet object, and an error if there is any.
func (c *FakeAvailabilitySets) Get(name string, options v1.GetOptions) (result *v1alpha1.AvailabilitySet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(availabilitysetsResource, c.ns, name), &v1alpha1.AvailabilitySet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AvailabilitySet), err
}

// List takes label and field selectors, and returns the list of AvailabilitySets that match those selectors.
func (c *FakeAvailabilitySets) List(opts v1.ListOptions) (result *v1alpha1.AvailabilitySetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(availabilitysetsResource, availabilitysetsKind, c.ns, opts), &v1alpha1.AvailabilitySetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AvailabilitySetList{ListMeta: obj.(*v1alpha1.AvailabilitySetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AvailabilitySetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested availabilitySets.
func (c *FakeAvailabilitySets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(availabilitysetsResource, c.ns, opts))

}

// Create takes the representation of a availabilitySet and creates it.  Returns the server's representation of the availabilitySet, and an error, if there is any.
func (c *FakeAvailabilitySets) Create(availabilitySet *v1alpha1.AvailabilitySet) (result *v1alpha1.AvailabilitySet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(availabilitysetsResource, c.ns, availabilitySet), &v1alpha1.AvailabilitySet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AvailabilitySet), err
}

// Update takes the representation of a availabilitySet and updates it. Returns the server's representation of the availabilitySet, and an error, if there is any.
func (c *FakeAvailabilitySets) Update(availabilitySet *v1alpha1.AvailabilitySet) (result *v1alpha1.AvailabilitySet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(availabilitysetsResource, c.ns, availabilitySet), &v1alpha1.AvailabilitySet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AvailabilitySet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAvailabilitySets) UpdateStatus(availabilitySet *v1alpha1.AvailabilitySet) (*v1alpha1.AvailabilitySet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(availabilitysetsResource, "status", c.ns, availabilitySet), &v1alpha1.AvailabilitySet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AvailabilitySet), err
}

// Delete takes name of the availabilitySet and deletes it. Returns an error if one occurs.
func (c *FakeAvailabilitySets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(availabilitysetsResource, c.ns, name), &v1alpha1.AvailabilitySet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAvailabilitySets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(availabilitysetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AvailabilitySetList{})
	return err
}

// Patch applies the patch and returns the patched availabilitySet.
func (c *FakeAvailabilitySets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AvailabilitySet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(availabilitysetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AvailabilitySet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AvailabilitySet), err
}
