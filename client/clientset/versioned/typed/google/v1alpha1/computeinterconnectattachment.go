/*
Copyright The Kubeform Authors.

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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeInterconnectAttachmentsGetter has a method to return a ComputeInterconnectAttachmentInterface.
// A group's client should implement this interface.
type ComputeInterconnectAttachmentsGetter interface {
	ComputeInterconnectAttachments(namespace string) ComputeInterconnectAttachmentInterface
}

// ComputeInterconnectAttachmentInterface has methods to work with ComputeInterconnectAttachment resources.
type ComputeInterconnectAttachmentInterface interface {
	Create(*v1alpha1.ComputeInterconnectAttachment) (*v1alpha1.ComputeInterconnectAttachment, error)
	Update(*v1alpha1.ComputeInterconnectAttachment) (*v1alpha1.ComputeInterconnectAttachment, error)
	UpdateStatus(*v1alpha1.ComputeInterconnectAttachment) (*v1alpha1.ComputeInterconnectAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeInterconnectAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeInterconnectAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeInterconnectAttachment, err error)
	ComputeInterconnectAttachmentExpansion
}

// computeInterconnectAttachments implements ComputeInterconnectAttachmentInterface
type computeInterconnectAttachments struct {
	client rest.Interface
	ns     string
}

// newComputeInterconnectAttachments returns a ComputeInterconnectAttachments
func newComputeInterconnectAttachments(c *GoogleV1alpha1Client, namespace string) *computeInterconnectAttachments {
	return &computeInterconnectAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeInterconnectAttachment, and returns the corresponding computeInterconnectAttachment object, and an error if there is any.
func (c *computeInterconnectAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeInterconnectAttachment, err error) {
	result = &v1alpha1.ComputeInterconnectAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeInterconnectAttachments that match those selectors.
func (c *computeInterconnectAttachments) List(opts v1.ListOptions) (result *v1alpha1.ComputeInterconnectAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeInterconnectAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeInterconnectAttachments.
func (c *computeInterconnectAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeInterconnectAttachment and creates it.  Returns the server's representation of the computeInterconnectAttachment, and an error, if there is any.
func (c *computeInterconnectAttachments) Create(computeInterconnectAttachment *v1alpha1.ComputeInterconnectAttachment) (result *v1alpha1.ComputeInterconnectAttachment, err error) {
	result = &v1alpha1.ComputeInterconnectAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		Body(computeInterconnectAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeInterconnectAttachment and updates it. Returns the server's representation of the computeInterconnectAttachment, and an error, if there is any.
func (c *computeInterconnectAttachments) Update(computeInterconnectAttachment *v1alpha1.ComputeInterconnectAttachment) (result *v1alpha1.ComputeInterconnectAttachment, err error) {
	result = &v1alpha1.ComputeInterconnectAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		Name(computeInterconnectAttachment.Name).
		Body(computeInterconnectAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeInterconnectAttachments) UpdateStatus(computeInterconnectAttachment *v1alpha1.ComputeInterconnectAttachment) (result *v1alpha1.ComputeInterconnectAttachment, err error) {
	result = &v1alpha1.ComputeInterconnectAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		Name(computeInterconnectAttachment.Name).
		SubResource("status").
		Body(computeInterconnectAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeInterconnectAttachment and deletes it. Returns an error if one occurs.
func (c *computeInterconnectAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeInterconnectAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeInterconnectAttachment.
func (c *computeInterconnectAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeInterconnectAttachment, err error) {
	result = &v1alpha1.ComputeInterconnectAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeinterconnectattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
