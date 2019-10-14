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

// FakeStreamAnalyticsStreamInputBlobs implements StreamAnalyticsStreamInputBlobInterface
type FakeStreamAnalyticsStreamInputBlobs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var streamanalyticsstreaminputblobsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "streamanalyticsstreaminputblobs"}

var streamanalyticsstreaminputblobsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StreamAnalyticsStreamInputBlob"}

// Get takes name of the streamAnalyticsStreamInputBlob, and returns the corresponding streamAnalyticsStreamInputBlob object, and an error if there is any.
func (c *FakeStreamAnalyticsStreamInputBlobs) Get(name string, options v1.GetOptions) (result *v1alpha1.StreamAnalyticsStreamInputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(streamanalyticsstreaminputblobsResource, c.ns, name), &v1alpha1.StreamAnalyticsStreamInputBlob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputBlob), err
}

// List takes label and field selectors, and returns the list of StreamAnalyticsStreamInputBlobs that match those selectors.
func (c *FakeStreamAnalyticsStreamInputBlobs) List(opts v1.ListOptions) (result *v1alpha1.StreamAnalyticsStreamInputBlobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(streamanalyticsstreaminputblobsResource, streamanalyticsstreaminputblobsKind, c.ns, opts), &v1alpha1.StreamAnalyticsStreamInputBlobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StreamAnalyticsStreamInputBlobList{ListMeta: obj.(*v1alpha1.StreamAnalyticsStreamInputBlobList).ListMeta}
	for _, item := range obj.(*v1alpha1.StreamAnalyticsStreamInputBlobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested streamAnalyticsStreamInputBlobs.
func (c *FakeStreamAnalyticsStreamInputBlobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(streamanalyticsstreaminputblobsResource, c.ns, opts))

}

// Create takes the representation of a streamAnalyticsStreamInputBlob and creates it.  Returns the server's representation of the streamAnalyticsStreamInputBlob, and an error, if there is any.
func (c *FakeStreamAnalyticsStreamInputBlobs) Create(streamAnalyticsStreamInputBlob *v1alpha1.StreamAnalyticsStreamInputBlob) (result *v1alpha1.StreamAnalyticsStreamInputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(streamanalyticsstreaminputblobsResource, c.ns, streamAnalyticsStreamInputBlob), &v1alpha1.StreamAnalyticsStreamInputBlob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputBlob), err
}

// Update takes the representation of a streamAnalyticsStreamInputBlob and updates it. Returns the server's representation of the streamAnalyticsStreamInputBlob, and an error, if there is any.
func (c *FakeStreamAnalyticsStreamInputBlobs) Update(streamAnalyticsStreamInputBlob *v1alpha1.StreamAnalyticsStreamInputBlob) (result *v1alpha1.StreamAnalyticsStreamInputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(streamanalyticsstreaminputblobsResource, c.ns, streamAnalyticsStreamInputBlob), &v1alpha1.StreamAnalyticsStreamInputBlob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputBlob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStreamAnalyticsStreamInputBlobs) UpdateStatus(streamAnalyticsStreamInputBlob *v1alpha1.StreamAnalyticsStreamInputBlob) (*v1alpha1.StreamAnalyticsStreamInputBlob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(streamanalyticsstreaminputblobsResource, "status", c.ns, streamAnalyticsStreamInputBlob), &v1alpha1.StreamAnalyticsStreamInputBlob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputBlob), err
}

// Delete takes name of the streamAnalyticsStreamInputBlob and deletes it. Returns an error if one occurs.
func (c *FakeStreamAnalyticsStreamInputBlobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(streamanalyticsstreaminputblobsResource, c.ns, name), &v1alpha1.StreamAnalyticsStreamInputBlob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStreamAnalyticsStreamInputBlobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(streamanalyticsstreaminputblobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StreamAnalyticsStreamInputBlobList{})
	return err
}

// Patch applies the patch and returns the patched streamAnalyticsStreamInputBlob.
func (c *FakeStreamAnalyticsStreamInputBlobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StreamAnalyticsStreamInputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(streamanalyticsstreaminputblobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StreamAnalyticsStreamInputBlob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputBlob), err
}
