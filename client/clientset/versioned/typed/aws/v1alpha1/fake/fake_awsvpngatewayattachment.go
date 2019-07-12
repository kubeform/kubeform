/*
Copyright 2019 The KubeDB Authors.

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

// FakeAwsVpnGatewayAttachments implements AwsVpnGatewayAttachmentInterface
type FakeAwsVpnGatewayAttachments struct {
	Fake *FakeAwsV1alpha1
}

var awsvpngatewayattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsvpngatewayattachments"}

var awsvpngatewayattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsVpnGatewayAttachment"}

// Get takes name of the awsVpnGatewayAttachment, and returns the corresponding awsVpnGatewayAttachment object, and an error if there is any.
func (c *FakeAwsVpnGatewayAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsvpngatewayattachmentsResource, name), &v1alpha1.AwsVpnGatewayAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpnGatewayAttachment), err
}

// List takes label and field selectors, and returns the list of AwsVpnGatewayAttachments that match those selectors.
func (c *FakeAwsVpnGatewayAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsVpnGatewayAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsvpngatewayattachmentsResource, awsvpngatewayattachmentsKind, opts), &v1alpha1.AwsVpnGatewayAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsVpnGatewayAttachmentList{ListMeta: obj.(*v1alpha1.AwsVpnGatewayAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsVpnGatewayAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpnGatewayAttachments.
func (c *FakeAwsVpnGatewayAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsvpngatewayattachmentsResource, opts))
}

// Create takes the representation of a awsVpnGatewayAttachment and creates it.  Returns the server's representation of the awsVpnGatewayAttachment, and an error, if there is any.
func (c *FakeAwsVpnGatewayAttachments) Create(awsVpnGatewayAttachment *v1alpha1.AwsVpnGatewayAttachment) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsvpngatewayattachmentsResource, awsVpnGatewayAttachment), &v1alpha1.AwsVpnGatewayAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpnGatewayAttachment), err
}

// Update takes the representation of a awsVpnGatewayAttachment and updates it. Returns the server's representation of the awsVpnGatewayAttachment, and an error, if there is any.
func (c *FakeAwsVpnGatewayAttachments) Update(awsVpnGatewayAttachment *v1alpha1.AwsVpnGatewayAttachment) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsvpngatewayattachmentsResource, awsVpnGatewayAttachment), &v1alpha1.AwsVpnGatewayAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpnGatewayAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsVpnGatewayAttachments) UpdateStatus(awsVpnGatewayAttachment *v1alpha1.AwsVpnGatewayAttachment) (*v1alpha1.AwsVpnGatewayAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsvpngatewayattachmentsResource, "status", awsVpnGatewayAttachment), &v1alpha1.AwsVpnGatewayAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpnGatewayAttachment), err
}

// Delete takes name of the awsVpnGatewayAttachment and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpnGatewayAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsvpngatewayattachmentsResource, name), &v1alpha1.AwsVpnGatewayAttachment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpnGatewayAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsvpngatewayattachmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsVpnGatewayAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched awsVpnGatewayAttachment.
func (c *FakeAwsVpnGatewayAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsVpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsvpngatewayattachmentsResource, name, pt, data, subresources...), &v1alpha1.AwsVpnGatewayAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsVpnGatewayAttachment), err
}
