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

// FakeAutoscalingAttachments implements AutoscalingAttachmentInterface
type FakeAutoscalingAttachments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var autoscalingattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "autoscalingattachments"}

var autoscalingattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AutoscalingAttachment"}

// Get takes name of the autoscalingAttachment, and returns the corresponding autoscalingAttachment object, and an error if there is any.
func (c *FakeAutoscalingAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscalingAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(autoscalingattachmentsResource, c.ns, name), &v1alpha1.AutoscalingAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingAttachment), err
}

// List takes label and field selectors, and returns the list of AutoscalingAttachments that match those selectors.
func (c *FakeAutoscalingAttachments) List(opts v1.ListOptions) (result *v1alpha1.AutoscalingAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(autoscalingattachmentsResource, autoscalingattachmentsKind, c.ns, opts), &v1alpha1.AutoscalingAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutoscalingAttachmentList{ListMeta: obj.(*v1alpha1.AutoscalingAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutoscalingAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autoscalingAttachments.
func (c *FakeAutoscalingAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(autoscalingattachmentsResource, c.ns, opts))

}

// Create takes the representation of a autoscalingAttachment and creates it.  Returns the server's representation of the autoscalingAttachment, and an error, if there is any.
func (c *FakeAutoscalingAttachments) Create(autoscalingAttachment *v1alpha1.AutoscalingAttachment) (result *v1alpha1.AutoscalingAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(autoscalingattachmentsResource, c.ns, autoscalingAttachment), &v1alpha1.AutoscalingAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingAttachment), err
}

// Update takes the representation of a autoscalingAttachment and updates it. Returns the server's representation of the autoscalingAttachment, and an error, if there is any.
func (c *FakeAutoscalingAttachments) Update(autoscalingAttachment *v1alpha1.AutoscalingAttachment) (result *v1alpha1.AutoscalingAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(autoscalingattachmentsResource, c.ns, autoscalingAttachment), &v1alpha1.AutoscalingAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutoscalingAttachments) UpdateStatus(autoscalingAttachment *v1alpha1.AutoscalingAttachment) (*v1alpha1.AutoscalingAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(autoscalingattachmentsResource, "status", c.ns, autoscalingAttachment), &v1alpha1.AutoscalingAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingAttachment), err
}

// Delete takes name of the autoscalingAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAutoscalingAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(autoscalingattachmentsResource, c.ns, name), &v1alpha1.AutoscalingAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutoscalingAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(autoscalingattachmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutoscalingAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched autoscalingAttachment.
func (c *FakeAutoscalingAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(autoscalingattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutoscalingAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoscalingAttachment), err
}
