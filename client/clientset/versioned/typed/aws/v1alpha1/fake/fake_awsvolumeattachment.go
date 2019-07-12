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

// FakeAwsVolumeAttachments implements AwsVolumeAttachmentInterface
type FakeAwsVolumeAttachments struct {
	Fake *FakeAwsV1alpha1
}

var awsvolumeattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsvolumeattachments"}

var awsvolumeattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsVolumeAttachment"}

// Get takes name of the awsVolumeAttachment, and returns the corresponding awsVolumeAttachment object, and an error if there is any.
func (c *FakeAwsVolumeAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsvolumeattachmentsResource, name), &v1alpha1.AwsVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVolumeAttachment), err
}

// List takes label and field selectors, and returns the list of AwsVolumeAttachments that match those selectors.
func (c *FakeAwsVolumeAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsVolumeAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsvolumeattachmentsResource, awsvolumeattachmentsKind, opts), &v1alpha1.AwsVolumeAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVolumeAttachmentList{ListMeta: obj.(*v1alpha1.AwsVolumeAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsVolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVolumeAttachments.
func (c *FakeAwsVolumeAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsvolumeattachmentsResource, opts))
}

// Create takes the representation of a awsVolumeAttachment and creates it.  Returns the server's representation of the awsVolumeAttachment, and an error, if there is any.
func (c *FakeAwsVolumeAttachments) Create(awsVolumeAttachment *v1alpha1.AwsVolumeAttachment) (result *v1alpha1.AwsVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsvolumeattachmentsResource, awsVolumeAttachment), &v1alpha1.AwsVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVolumeAttachment), err
}

// Update takes the representation of a awsVolumeAttachment and updates it. Returns the server's representation of the awsVolumeAttachment, and an error, if there is any.
func (c *FakeAwsVolumeAttachments) Update(awsVolumeAttachment *v1alpha1.AwsVolumeAttachment) (result *v1alpha1.AwsVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsvolumeattachmentsResource, awsVolumeAttachment), &v1alpha1.AwsVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVolumeAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsVolumeAttachments) UpdateStatus(awsVolumeAttachment *v1alpha1.AwsVolumeAttachment) (*v1alpha1.AwsVolumeAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsvolumeattachmentsResource, "status", awsVolumeAttachment), &v1alpha1.AwsVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVolumeAttachment), err
}

// Delete takes name of the awsVolumeAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAwsVolumeAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsvolumeattachmentsResource, name), &v1alpha1.AwsVolumeAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVolumeAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsvolumeattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVolumeAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched awsVolumeAttachment.
func (c *FakeAwsVolumeAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVolumeAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsvolumeattachmentsResource, name, pt, data, subresources...), &v1alpha1.AwsVolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVolumeAttachment), err
}
