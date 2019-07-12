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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeDigitaloceanVolumeAttachments implements DigitaloceanVolumeAttachmentInterface
type FakeDigitaloceanVolumeAttachments struct {
	Fake *FakeDigitaloceanV1alpha1
}

var digitaloceanvolumeattachmentsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "digitaloceanvolumeattachments"}

var digitaloceanvolumeattachmentsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DigitaloceanVolumeAttachment"}

// Get takes name of the digitaloceanVolumeAttachment, and returns the corresponding digitaloceanVolumeAttachment object, and an error if there is any.
func (c *FakeDigitaloceanVolumeAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(digitaloceanvolumeattachmentsResource, name), &v1alpha1.DigitaloceanVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeAttachment), err
}

// List takes label and field selectors, and returns the list of DigitaloceanVolumeAttachments that match those selectors.
func (c *FakeDigitaloceanVolumeAttachments) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanVolumeAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(digitaloceanvolumeattachmentsResource, digitaloceanvolumeattachmentsKind, opts), &v1alpha1.DigitaloceanVolumeAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DigitaloceanVolumeAttachmentList{ListMeta: obj.(*v1alpha1.DigitaloceanVolumeAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.DigitaloceanVolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested digitaloceanVolumeAttachments.
func (c *FakeDigitaloceanVolumeAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(digitaloceanvolumeattachmentsResource, opts))
}

// Create takes the representation of a digitaloceanVolumeAttachment and creates it.  Returns the server's representation of the digitaloceanVolumeAttachment, and an error, if there is any.
func (c *FakeDigitaloceanVolumeAttachments) Create(digitaloceanVolumeAttachment *v1alpha1.DigitaloceanVolumeAttachment) (result *v1alpha1.DigitaloceanVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(digitaloceanvolumeattachmentsResource, digitaloceanVolumeAttachment), &v1alpha1.DigitaloceanVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeAttachment), err
}

// Update takes the representation of a digitaloceanVolumeAttachment and updates it. Returns the server's representation of the digitaloceanVolumeAttachment, and an error, if there is any.
func (c *FakeDigitaloceanVolumeAttachments) Update(digitaloceanVolumeAttachment *v1alpha1.DigitaloceanVolumeAttachment) (result *v1alpha1.DigitaloceanVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(digitaloceanvolumeattachmentsResource, digitaloceanVolumeAttachment), &v1alpha1.DigitaloceanVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDigitaloceanVolumeAttachments) UpdateStatus(digitaloceanVolumeAttachment *v1alpha1.DigitaloceanVolumeAttachment) (*v1alpha1.DigitaloceanVolumeAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(digitaloceanvolumeattachmentsResource, "status", digitaloceanVolumeAttachment), &v1alpha1.DigitaloceanVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeAttachment), err
}

// Delete takes name of the digitaloceanVolumeAttachment and deletes it. Returns an error if one occurs.
func (c *FakeDigitaloceanVolumeAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(digitaloceanvolumeattachmentsResource, name), &v1alpha1.DigitaloceanVolumeAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDigitaloceanVolumeAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(digitaloceanvolumeattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DigitaloceanVolumeAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched digitaloceanVolumeAttachment.
func (c *FakeDigitaloceanVolumeAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(digitaloceanvolumeattachmentsResource, name, pt, data, subresources...), &v1alpha1.DigitaloceanVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolumeAttachment), err
}
