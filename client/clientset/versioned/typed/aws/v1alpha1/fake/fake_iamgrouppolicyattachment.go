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

// FakeIamGroupPolicyAttachments implements IamGroupPolicyAttachmentInterface
type FakeIamGroupPolicyAttachments struct {
	Fake *FakeAwsV1alpha1
}

var iamgrouppolicyattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamgrouppolicyattachments"}

var iamgrouppolicyattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamGroupPolicyAttachment"}

// Get takes name of the iamGroupPolicyAttachment, and returns the corresponding iamGroupPolicyAttachment object, and an error if there is any.
func (c *FakeIamGroupPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(iamgrouppolicyattachmentsResource, name), &v1alpha1.IamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicyAttachment), err
}

// List takes label and field selectors, and returns the list of IamGroupPolicyAttachments that match those selectors.
func (c *FakeIamGroupPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.IamGroupPolicyAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(iamgrouppolicyattachmentsResource, iamgrouppolicyattachmentsKind, opts), &v1alpha1.IamGroupPolicyAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamGroupPolicyAttachmentList{ListMeta: obj.(*v1alpha1.IamGroupPolicyAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamGroupPolicyAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamGroupPolicyAttachments.
func (c *FakeIamGroupPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(iamgrouppolicyattachmentsResource, opts))
}

// Create takes the representation of a iamGroupPolicyAttachment and creates it.  Returns the server's representation of the iamGroupPolicyAttachment, and an error, if there is any.
func (c *FakeIamGroupPolicyAttachments) Create(iamGroupPolicyAttachment *v1alpha1.IamGroupPolicyAttachment) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(iamgrouppolicyattachmentsResource, iamGroupPolicyAttachment), &v1alpha1.IamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicyAttachment), err
}

// Update takes the representation of a iamGroupPolicyAttachment and updates it. Returns the server's representation of the iamGroupPolicyAttachment, and an error, if there is any.
func (c *FakeIamGroupPolicyAttachments) Update(iamGroupPolicyAttachment *v1alpha1.IamGroupPolicyAttachment) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(iamgrouppolicyattachmentsResource, iamGroupPolicyAttachment), &v1alpha1.IamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicyAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamGroupPolicyAttachments) UpdateStatus(iamGroupPolicyAttachment *v1alpha1.IamGroupPolicyAttachment) (*v1alpha1.IamGroupPolicyAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(iamgrouppolicyattachmentsResource, "status", iamGroupPolicyAttachment), &v1alpha1.IamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicyAttachment), err
}

// Delete takes name of the iamGroupPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *FakeIamGroupPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(iamgrouppolicyattachmentsResource, name), &v1alpha1.IamGroupPolicyAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamGroupPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(iamgrouppolicyattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamGroupPolicyAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched iamGroupPolicyAttachment.
func (c *FakeIamGroupPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(iamgrouppolicyattachmentsResource, name, pt, data, subresources...), &v1alpha1.IamGroupPolicyAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamGroupPolicyAttachment), err
}
