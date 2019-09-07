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

// FakeStreamAnalyticsStreamInputIothubs implements StreamAnalyticsStreamInputIothubInterface
type FakeStreamAnalyticsStreamInputIothubs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var streamanalyticsstreaminputiothubsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "streamanalyticsstreaminputiothubs"}

var streamanalyticsstreaminputiothubsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StreamAnalyticsStreamInputIothub"}

// Get takes name of the streamAnalyticsStreamInputIothub, and returns the corresponding streamAnalyticsStreamInputIothub object, and an error if there is any.
func (c *FakeStreamAnalyticsStreamInputIothubs) Get(name string, options v1.GetOptions) (result *v1alpha1.StreamAnalyticsStreamInputIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(streamanalyticsstreaminputiothubsResource, c.ns, name), &v1alpha1.StreamAnalyticsStreamInputIothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputIothub), err
}

// List takes label and field selectors, and returns the list of StreamAnalyticsStreamInputIothubs that match those selectors.
func (c *FakeStreamAnalyticsStreamInputIothubs) List(opts v1.ListOptions) (result *v1alpha1.StreamAnalyticsStreamInputIothubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(streamanalyticsstreaminputiothubsResource, streamanalyticsstreaminputiothubsKind, c.ns, opts), &v1alpha1.StreamAnalyticsStreamInputIothubList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StreamAnalyticsStreamInputIothubList{ListMeta: obj.(*v1alpha1.StreamAnalyticsStreamInputIothubList).ListMeta}
	for _, item := range obj.(*v1alpha1.StreamAnalyticsStreamInputIothubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested streamAnalyticsStreamInputIothubs.
func (c *FakeStreamAnalyticsStreamInputIothubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(streamanalyticsstreaminputiothubsResource, c.ns, opts))

}

// Create takes the representation of a streamAnalyticsStreamInputIothub and creates it.  Returns the server's representation of the streamAnalyticsStreamInputIothub, and an error, if there is any.
func (c *FakeStreamAnalyticsStreamInputIothubs) Create(streamAnalyticsStreamInputIothub *v1alpha1.StreamAnalyticsStreamInputIothub) (result *v1alpha1.StreamAnalyticsStreamInputIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(streamanalyticsstreaminputiothubsResource, c.ns, streamAnalyticsStreamInputIothub), &v1alpha1.StreamAnalyticsStreamInputIothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputIothub), err
}

// Update takes the representation of a streamAnalyticsStreamInputIothub and updates it. Returns the server's representation of the streamAnalyticsStreamInputIothub, and an error, if there is any.
func (c *FakeStreamAnalyticsStreamInputIothubs) Update(streamAnalyticsStreamInputIothub *v1alpha1.StreamAnalyticsStreamInputIothub) (result *v1alpha1.StreamAnalyticsStreamInputIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(streamanalyticsstreaminputiothubsResource, c.ns, streamAnalyticsStreamInputIothub), &v1alpha1.StreamAnalyticsStreamInputIothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputIothub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStreamAnalyticsStreamInputIothubs) UpdateStatus(streamAnalyticsStreamInputIothub *v1alpha1.StreamAnalyticsStreamInputIothub) (*v1alpha1.StreamAnalyticsStreamInputIothub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(streamanalyticsstreaminputiothubsResource, "status", c.ns, streamAnalyticsStreamInputIothub), &v1alpha1.StreamAnalyticsStreamInputIothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputIothub), err
}

// Delete takes name of the streamAnalyticsStreamInputIothub and deletes it. Returns an error if one occurs.
func (c *FakeStreamAnalyticsStreamInputIothubs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(streamanalyticsstreaminputiothubsResource, c.ns, name), &v1alpha1.StreamAnalyticsStreamInputIothub{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStreamAnalyticsStreamInputIothubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(streamanalyticsstreaminputiothubsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StreamAnalyticsStreamInputIothubList{})
	return err
}

// Patch applies the patch and returns the patched streamAnalyticsStreamInputIothub.
func (c *FakeStreamAnalyticsStreamInputIothubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StreamAnalyticsStreamInputIothub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(streamanalyticsstreaminputiothubsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StreamAnalyticsStreamInputIothub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputIothub), err
}