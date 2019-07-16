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

// FakeLightsailStaticIpAttachments implements LightsailStaticIpAttachmentInterface
type FakeLightsailStaticIpAttachments struct {
	Fake *FakeAwsV1alpha1
}

var lightsailstaticipattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lightsailstaticipattachments"}

var lightsailstaticipattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LightsailStaticIpAttachment"}

// Get takes name of the lightsailStaticIpAttachment, and returns the corresponding lightsailStaticIpAttachment object, and an error if there is any.
func (c *FakeLightsailStaticIpAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.LightsailStaticIpAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(lightsailstaticipattachmentsResource, name), &v1alpha1.LightsailStaticIpAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIpAttachment), err
}

// List takes label and field selectors, and returns the list of LightsailStaticIpAttachments that match those selectors.
func (c *FakeLightsailStaticIpAttachments) List(opts v1.ListOptions) (result *v1alpha1.LightsailStaticIpAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(lightsailstaticipattachmentsResource, lightsailstaticipattachmentsKind, opts), &v1alpha1.LightsailStaticIpAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LightsailStaticIpAttachmentList{ListMeta: obj.(*v1alpha1.LightsailStaticIpAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.LightsailStaticIpAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lightsailStaticIpAttachments.
func (c *FakeLightsailStaticIpAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(lightsailstaticipattachmentsResource, opts))
}

// Create takes the representation of a lightsailStaticIpAttachment and creates it.  Returns the server's representation of the lightsailStaticIpAttachment, and an error, if there is any.
func (c *FakeLightsailStaticIpAttachments) Create(lightsailStaticIpAttachment *v1alpha1.LightsailStaticIpAttachment) (result *v1alpha1.LightsailStaticIpAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(lightsailstaticipattachmentsResource, lightsailStaticIpAttachment), &v1alpha1.LightsailStaticIpAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIpAttachment), err
}

// Update takes the representation of a lightsailStaticIpAttachment and updates it. Returns the server's representation of the lightsailStaticIpAttachment, and an error, if there is any.
func (c *FakeLightsailStaticIpAttachments) Update(lightsailStaticIpAttachment *v1alpha1.LightsailStaticIpAttachment) (result *v1alpha1.LightsailStaticIpAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(lightsailstaticipattachmentsResource, lightsailStaticIpAttachment), &v1alpha1.LightsailStaticIpAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIpAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLightsailStaticIpAttachments) UpdateStatus(lightsailStaticIpAttachment *v1alpha1.LightsailStaticIpAttachment) (*v1alpha1.LightsailStaticIpAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(lightsailstaticipattachmentsResource, "status", lightsailStaticIpAttachment), &v1alpha1.LightsailStaticIpAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIpAttachment), err
}

// Delete takes name of the lightsailStaticIpAttachment and deletes it. Returns an error if one occurs.
func (c *FakeLightsailStaticIpAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(lightsailstaticipattachmentsResource, name), &v1alpha1.LightsailStaticIpAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLightsailStaticIpAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(lightsailstaticipattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LightsailStaticIpAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched lightsailStaticIpAttachment.
func (c *FakeLightsailStaticIpAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailStaticIpAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(lightsailstaticipattachmentsResource, name, pt, data, subresources...), &v1alpha1.LightsailStaticIpAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIpAttachment), err
}
