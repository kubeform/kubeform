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

// AwsIotThingPrincipalAttachmentsGetter has a method to return a AwsIotThingPrincipalAttachmentInterface.
// A group's client should implement this interface.
type AwsIotThingPrincipalAttachmentsGetter interface {
	AwsIotThingPrincipalAttachments() AwsIotThingPrincipalAttachmentInterface
}

// AwsIotThingPrincipalAttachmentInterface has methods to work with AwsIotThingPrincipalAttachment resources.
type AwsIotThingPrincipalAttachmentInterface interface {
	Create(*v1alpha1.AwsIotThingPrincipalAttachment) (*v1alpha1.AwsIotThingPrincipalAttachment, error)
	Update(*v1alpha1.AwsIotThingPrincipalAttachment) (*v1alpha1.AwsIotThingPrincipalAttachment, error)
	UpdateStatus(*v1alpha1.AwsIotThingPrincipalAttachment) (*v1alpha1.AwsIotThingPrincipalAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsIotThingPrincipalAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsIotThingPrincipalAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIotThingPrincipalAttachment, err error)
	AwsIotThingPrincipalAttachmentExpansion
}

// awsIotThingPrincipalAttachments implements AwsIotThingPrincipalAttachmentInterface
type awsIotThingPrincipalAttachments struct {
	client rest.Interface
}

// newAwsIotThingPrincipalAttachments returns a AwsIotThingPrincipalAttachments
func newAwsIotThingPrincipalAttachments(c *AwsV1alpha1Client) *awsIotThingPrincipalAttachments {
	return &awsIotThingPrincipalAttachments{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsIotThingPrincipalAttachment, and returns the corresponding awsIotThingPrincipalAttachment object, and an error if there is any.
func (c *awsIotThingPrincipalAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsIotThingPrincipalAttachment, err error) {
	result = &v1alpha1.AwsIotThingPrincipalAttachment{}
	err = c.client.Get().
		Resource("awsiotthingprincipalattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIotThingPrincipalAttachments that match those selectors.
func (c *awsIotThingPrincipalAttachments) List(opts v1.ListOptions) (result *v1alpha1.AwsIotThingPrincipalAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsIotThingPrincipalAttachmentList{}
	err = c.client.Get().
		Resource("awsiotthingprincipalattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIotThingPrincipalAttachments.
func (c *awsIotThingPrincipalAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsiotthingprincipalattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsIotThingPrincipalAttachment and creates it.  Returns the server's representation of the awsIotThingPrincipalAttachment, and an error, if there is any.
func (c *awsIotThingPrincipalAttachments) Create(awsIotThingPrincipalAttachment *v1alpha1.AwsIotThingPrincipalAttachment) (result *v1alpha1.AwsIotThingPrincipalAttachment, err error) {
	result = &v1alpha1.AwsIotThingPrincipalAttachment{}
	err = c.client.Post().
		Resource("awsiotthingprincipalattachments").
		Body(awsIotThingPrincipalAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIotThingPrincipalAttachment and updates it. Returns the server's representation of the awsIotThingPrincipalAttachment, and an error, if there is any.
func (c *awsIotThingPrincipalAttachments) Update(awsIotThingPrincipalAttachment *v1alpha1.AwsIotThingPrincipalAttachment) (result *v1alpha1.AwsIotThingPrincipalAttachment, err error) {
	result = &v1alpha1.AwsIotThingPrincipalAttachment{}
	err = c.client.Put().
		Resource("awsiotthingprincipalattachments").
		Name(awsIotThingPrincipalAttachment.Name).
		Body(awsIotThingPrincipalAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsIotThingPrincipalAttachments) UpdateStatus(awsIotThingPrincipalAttachment *v1alpha1.AwsIotThingPrincipalAttachment) (result *v1alpha1.AwsIotThingPrincipalAttachment, err error) {
	result = &v1alpha1.AwsIotThingPrincipalAttachment{}
	err = c.client.Put().
		Resource("awsiotthingprincipalattachments").
		Name(awsIotThingPrincipalAttachment.Name).
		SubResource("status").
		Body(awsIotThingPrincipalAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIotThingPrincipalAttachment and deletes it. Returns an error if one occurs.
func (c *awsIotThingPrincipalAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsiotthingprincipalattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIotThingPrincipalAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsiotthingprincipalattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIotThingPrincipalAttachment.
func (c *awsIotThingPrincipalAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsIotThingPrincipalAttachment, err error) {
	result = &v1alpha1.AwsIotThingPrincipalAttachment{}
	err = c.client.Patch(pt).
		Resource("awsiotthingprincipalattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
