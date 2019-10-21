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

// IamGroupPolicyAttachmentsGetter has a method to return a IamGroupPolicyAttachmentInterface.
// A group's client should implement this interface.
type IamGroupPolicyAttachmentsGetter interface {
	IamGroupPolicyAttachments(namespace string) IamGroupPolicyAttachmentInterface
}

// IamGroupPolicyAttachmentInterface has methods to work with IamGroupPolicyAttachment resources.
type IamGroupPolicyAttachmentInterface interface {
	Create(*v1alpha1.IamGroupPolicyAttachment) (*v1alpha1.IamGroupPolicyAttachment, error)
	Update(*v1alpha1.IamGroupPolicyAttachment) (*v1alpha1.IamGroupPolicyAttachment, error)
	UpdateStatus(*v1alpha1.IamGroupPolicyAttachment) (*v1alpha1.IamGroupPolicyAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamGroupPolicyAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.IamGroupPolicyAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamGroupPolicyAttachment, err error)
	IamGroupPolicyAttachmentExpansion
}

// iamGroupPolicyAttachments implements IamGroupPolicyAttachmentInterface
type iamGroupPolicyAttachments struct {
	client rest.Interface
	ns     string
}

// newIamGroupPolicyAttachments returns a IamGroupPolicyAttachments
func newIamGroupPolicyAttachments(c *AwsV1alpha1Client, namespace string) *iamGroupPolicyAttachments {
	return &iamGroupPolicyAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamGroupPolicyAttachment, and returns the corresponding iamGroupPolicyAttachment object, and an error if there is any.
func (c *iamGroupPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	result = &v1alpha1.IamGroupPolicyAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamGroupPolicyAttachments that match those selectors.
func (c *iamGroupPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.IamGroupPolicyAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamGroupPolicyAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamGroupPolicyAttachments.
func (c *iamGroupPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamGroupPolicyAttachment and creates it.  Returns the server's representation of the iamGroupPolicyAttachment, and an error, if there is any.
func (c *iamGroupPolicyAttachments) Create(iamGroupPolicyAttachment *v1alpha1.IamGroupPolicyAttachment) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	result = &v1alpha1.IamGroupPolicyAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		Body(iamGroupPolicyAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamGroupPolicyAttachment and updates it. Returns the server's representation of the iamGroupPolicyAttachment, and an error, if there is any.
func (c *iamGroupPolicyAttachments) Update(iamGroupPolicyAttachment *v1alpha1.IamGroupPolicyAttachment) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	result = &v1alpha1.IamGroupPolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		Name(iamGroupPolicyAttachment.Name).
		Body(iamGroupPolicyAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamGroupPolicyAttachments) UpdateStatus(iamGroupPolicyAttachment *v1alpha1.IamGroupPolicyAttachment) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	result = &v1alpha1.IamGroupPolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		Name(iamGroupPolicyAttachment.Name).
		SubResource("status").
		Body(iamGroupPolicyAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamGroupPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *iamGroupPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamGroupPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamGroupPolicyAttachment.
func (c *iamGroupPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamGroupPolicyAttachment, err error) {
	result = &v1alpha1.IamGroupPolicyAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamgrouppolicyattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}