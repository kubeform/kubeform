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

// FakeLbTargetGroupAttachments implements LbTargetGroupAttachmentInterface
type FakeLbTargetGroupAttachments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var lbtargetgroupattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lbtargetgroupattachments"}

var lbtargetgroupattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LbTargetGroupAttachment"}

// Get takes name of the lbTargetGroupAttachment, and returns the corresponding lbTargetGroupAttachment object, and an error if there is any.
func (c *FakeLbTargetGroupAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.LbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lbtargetgroupattachmentsResource, c.ns, name), &v1alpha1.LbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbTargetGroupAttachment), err
}

// List takes label and field selectors, and returns the list of LbTargetGroupAttachments that match those selectors.
func (c *FakeLbTargetGroupAttachments) List(opts v1.ListOptions) (result *v1alpha1.LbTargetGroupAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lbtargetgroupattachmentsResource, lbtargetgroupattachmentsKind, c.ns, opts), &v1alpha1.LbTargetGroupAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbTargetGroupAttachmentList{ListMeta: obj.(*v1alpha1.LbTargetGroupAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbTargetGroupAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbTargetGroupAttachments.
func (c *FakeLbTargetGroupAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lbtargetgroupattachmentsResource, c.ns, opts))

}

// Create takes the representation of a lbTargetGroupAttachment and creates it.  Returns the server's representation of the lbTargetGroupAttachment, and an error, if there is any.
func (c *FakeLbTargetGroupAttachments) Create(lbTargetGroupAttachment *v1alpha1.LbTargetGroupAttachment) (result *v1alpha1.LbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lbtargetgroupattachmentsResource, c.ns, lbTargetGroupAttachment), &v1alpha1.LbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbTargetGroupAttachment), err
}

// Update takes the representation of a lbTargetGroupAttachment and updates it. Returns the server's representation of the lbTargetGroupAttachment, and an error, if there is any.
func (c *FakeLbTargetGroupAttachments) Update(lbTargetGroupAttachment *v1alpha1.LbTargetGroupAttachment) (result *v1alpha1.LbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lbtargetgroupattachmentsResource, c.ns, lbTargetGroupAttachment), &v1alpha1.LbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbTargetGroupAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbTargetGroupAttachments) UpdateStatus(lbTargetGroupAttachment *v1alpha1.LbTargetGroupAttachment) (*v1alpha1.LbTargetGroupAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lbtargetgroupattachmentsResource, "status", c.ns, lbTargetGroupAttachment), &v1alpha1.LbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbTargetGroupAttachment), err
}

// Delete takes name of the lbTargetGroupAttachment and deletes it. Returns an error if one occurs.
func (c *FakeLbTargetGroupAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lbtargetgroupattachmentsResource, c.ns, name), &v1alpha1.LbTargetGroupAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbTargetGroupAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lbtargetgroupattachmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbTargetGroupAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched lbTargetGroupAttachment.
func (c *FakeLbTargetGroupAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbTargetGroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lbtargetgroupattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LbTargetGroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbTargetGroupAttachment), err
}