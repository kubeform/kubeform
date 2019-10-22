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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIamPolicyAttachments implements IamPolicyAttachmentInterface
type FakeIamPolicyAttachments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iampolicyattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iampolicyattachments"}

var iampolicyattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamPolicyAttachment"}

// Get takes name of the iamPolicyAttachment, and returns the corresponding iamPolicyAttachment object, and an error if there is any.
func (c *FakeIamPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.IamPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iampolicyattachmentsResource, c.ns, name), &v1alpha1.IamPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicyAttachment), err
}

// List takes label and field selectors, and returns the list of IamPolicyAttachments that match those selectors.
func (c *FakeIamPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.IamPolicyAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iampolicyattachmentsResource, iampolicyattachmentsKind, c.ns, opts), &v1alpha1.IamPolicyAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamPolicyAttachmentList{ListMeta: obj.(*v1alpha1.IamPolicyAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamPolicyAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamPolicyAttachments.
func (c *FakeIamPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iampolicyattachmentsResource, c.ns, opts))

}

// Create takes the representation of a iamPolicyAttachment and creates it.  Returns the server's representation of the iamPolicyAttachment, and an error, if there is any.
func (c *FakeIamPolicyAttachments) Create(iamPolicyAttachment *v1alpha1.IamPolicyAttachment) (result *v1alpha1.IamPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iampolicyattachmentsResource, c.ns, iamPolicyAttachment), &v1alpha1.IamPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicyAttachment), err
}

// Update takes the representation of a iamPolicyAttachment and updates it. Returns the server's representation of the iamPolicyAttachment, and an error, if there is any.
func (c *FakeIamPolicyAttachments) Update(iamPolicyAttachment *v1alpha1.IamPolicyAttachment) (result *v1alpha1.IamPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iampolicyattachmentsResource, c.ns, iamPolicyAttachment), &v1alpha1.IamPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicyAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamPolicyAttachments) UpdateStatus(iamPolicyAttachment *v1alpha1.IamPolicyAttachment) (*v1alpha1.IamPolicyAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iampolicyattachmentsResource, "status", c.ns, iamPolicyAttachment), &v1alpha1.IamPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicyAttachment), err
}

// Delete takes name of the iamPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *FakeIamPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iampolicyattachmentsResource, c.ns, name), &v1alpha1.IamPolicyAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iampolicyattachmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamPolicyAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched iamPolicyAttachment.
func (c *FakeIamPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iampolicyattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicyAttachment), err
}
