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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsEc2TransitGatewayVpcAttachmentsGetter has a method to return a AwsEc2TransitGatewayVpcAttachmentInterface.
// A group's client should implement this interface.
type AwsEc2TransitGatewayVpcAttachmentsGetter interface {
	AwsEc2TransitGatewayVpcAttachments() AwsEc2TransitGatewayVpcAttachmentInterface
}

// AwsEc2TransitGatewayVpcAttachmentInterface has methods to work with AwsEc2TransitGatewayVpcAttachment resources.
type AwsEc2TransitGatewayVpcAttachmentInterface interface {
	Create(*v1alpha1.AwsEc2TransitGatewayVpcAttachment) (*v1alpha1.AwsEc2TransitGatewayVpcAttachment, error)
	Update(*v1alpha1.AwsEc2TransitGatewayVpcAttachment) (*v1alpha1.AwsEc2TransitGatewayVpcAttachment, error)
	UpdateStatus(*v1alpha1.AwsEc2TransitGatewayVpcAttachment) (*v1alpha1.AwsEc2TransitGatewayVpcAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEc2TransitGatewayVpcAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEc2TransitGatewayVpcAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachment, err error)
	AwsEc2TransitGatewayVpcAttachmentExpansion
}

// awsEc2TransitGatewayVpcAttachments implements AwsEc2TransitGatewayVpcAttachmentInterface
type awsEc2TransitGatewayVpcAttachments struct {
	client rest.Interface
}

// newAwsEc2TransitGatewayVpcAttachments returns a AwsEc2TransitGatewayVpcAttachments
func newAwsEc2TransitGatewayVpcAttachments(c *AwsV1alpha1Client) *awsEc2TransitGatewayVpcAttachments {
	return &awsEc2TransitGatewayVpcAttachments{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsEc2TransitGatewayVpcAttachment, and returns the corresponding awsEc2TransitGatewayVpcAttachment object, and an error if there is any.
func (c *awsEc2TransitGatewayVpcAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachment, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayVpcAttachment{}
	err = c.client.Get().
		Resource("awsec2transitgatewayvpcattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEc2TransitGatewayVpcAttachments that match those selectors.
func (c *awsEc2TransitGatewayVpcAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsEc2TransitGatewayVpcAttachmentList{}
	err = c.client.Get().
		Resource("awsec2transitgatewayvpcattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEc2TransitGatewayVpcAttachments.
func (c *awsEc2TransitGatewayVpcAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsec2transitgatewayvpcattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsEc2TransitGatewayVpcAttachment and creates it.  Returns the server's representation of the awsEc2TransitGatewayVpcAttachment, and an error, if there is any.
func (c *awsEc2TransitGatewayVpcAttachments) Create(awsEc2TransitGatewayVpcAttachment *v1alpha1.AwsEc2TransitGatewayVpcAttachment) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachment, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayVpcAttachment{}
	err = c.client.Post().
		Resource("awsec2transitgatewayvpcattachments").
		Body(awsEc2TransitGatewayVpcAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEc2TransitGatewayVpcAttachment and updates it. Returns the server's representation of the awsEc2TransitGatewayVpcAttachment, and an error, if there is any.
func (c *awsEc2TransitGatewayVpcAttachments) Update(awsEc2TransitGatewayVpcAttachment *v1alpha1.AwsEc2TransitGatewayVpcAttachment) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachment, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayVpcAttachment{}
	err = c.client.Put().
		Resource("awsec2transitgatewayvpcattachments").
		Name(awsEc2TransitGatewayVpcAttachment.Name).
		Body(awsEc2TransitGatewayVpcAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsEc2TransitGatewayVpcAttachments) UpdateStatus(awsEc2TransitGatewayVpcAttachment *v1alpha1.AwsEc2TransitGatewayVpcAttachment) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachment, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayVpcAttachment{}
	err = c.client.Put().
		Resource("awsec2transitgatewayvpcattachments").
		Name(awsEc2TransitGatewayVpcAttachment.Name).
		SubResource("status").
		Body(awsEc2TransitGatewayVpcAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEc2TransitGatewayVpcAttachment and deletes it. Returns an error if one occurs.
func (c *awsEc2TransitGatewayVpcAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsec2transitgatewayvpcattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEc2TransitGatewayVpcAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsec2transitgatewayvpcattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEc2TransitGatewayVpcAttachment.
func (c *awsEc2TransitGatewayVpcAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEc2TransitGatewayVpcAttachment, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayVpcAttachment{}
	err = c.client.Patch(pt).
		Resource("awsec2transitgatewayvpcattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
