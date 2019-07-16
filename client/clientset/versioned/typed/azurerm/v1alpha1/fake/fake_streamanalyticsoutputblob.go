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

// FakeStreamAnalyticsOutputBlobs implements StreamAnalyticsOutputBlobInterface
type FakeStreamAnalyticsOutputBlobs struct {
	Fake *FakeAzurermV1alpha1
}

var streamanalyticsoutputblobsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "streamanalyticsoutputblobs"}

var streamanalyticsoutputblobsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StreamAnalyticsOutputBlob"}

// Get takes name of the streamAnalyticsOutputBlob, and returns the corresponding streamAnalyticsOutputBlob object, and an error if there is any.
func (c *FakeStreamAnalyticsOutputBlobs) Get(name string, options v1.GetOptions) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(streamanalyticsoutputblobsResource, name), &v1alpha1.StreamAnalyticsOutputBlob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputBlob), err
}

// List takes label and field selectors, and returns the list of StreamAnalyticsOutputBlobs that match those selectors.
func (c *FakeStreamAnalyticsOutputBlobs) List(opts v1.ListOptions) (result *v1alpha1.StreamAnalyticsOutputBlobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(streamanalyticsoutputblobsResource, streamanalyticsoutputblobsKind, opts), &v1alpha1.StreamAnalyticsOutputBlobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StreamAnalyticsOutputBlobList{ListMeta: obj.(*v1alpha1.StreamAnalyticsOutputBlobList).ListMeta}
	for _, item := range obj.(*v1alpha1.StreamAnalyticsOutputBlobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested streamAnalyticsOutputBlobs.
func (c *FakeStreamAnalyticsOutputBlobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(streamanalyticsoutputblobsResource, opts))
}

// Create takes the representation of a streamAnalyticsOutputBlob and creates it.  Returns the server's representation of the streamAnalyticsOutputBlob, and an error, if there is any.
func (c *FakeStreamAnalyticsOutputBlobs) Create(streamAnalyticsOutputBlob *v1alpha1.StreamAnalyticsOutputBlob) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(streamanalyticsoutputblobsResource, streamAnalyticsOutputBlob), &v1alpha1.StreamAnalyticsOutputBlob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputBlob), err
}

// Update takes the representation of a streamAnalyticsOutputBlob and updates it. Returns the server's representation of the streamAnalyticsOutputBlob, and an error, if there is any.
func (c *FakeStreamAnalyticsOutputBlobs) Update(streamAnalyticsOutputBlob *v1alpha1.StreamAnalyticsOutputBlob) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(streamanalyticsoutputblobsResource, streamAnalyticsOutputBlob), &v1alpha1.StreamAnalyticsOutputBlob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputBlob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStreamAnalyticsOutputBlobs) UpdateStatus(streamAnalyticsOutputBlob *v1alpha1.StreamAnalyticsOutputBlob) (*v1alpha1.StreamAnalyticsOutputBlob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(streamanalyticsoutputblobsResource, "status", streamAnalyticsOutputBlob), &v1alpha1.StreamAnalyticsOutputBlob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputBlob), err
}

// Delete takes name of the streamAnalyticsOutputBlob and deletes it. Returns an error if one occurs.
func (c *FakeStreamAnalyticsOutputBlobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(streamanalyticsoutputblobsResource, name), &v1alpha1.StreamAnalyticsOutputBlob{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStreamAnalyticsOutputBlobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(streamanalyticsoutputblobsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StreamAnalyticsOutputBlobList{})
	return err
}

// Patch applies the patch and returns the patched streamAnalyticsOutputBlob.
func (c *FakeStreamAnalyticsOutputBlobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(streamanalyticsoutputblobsResource, name, pt, data, subresources...), &v1alpha1.StreamAnalyticsOutputBlob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputBlob), err
}
