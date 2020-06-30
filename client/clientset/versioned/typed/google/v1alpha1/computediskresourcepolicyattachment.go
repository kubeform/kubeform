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

// ComputeDiskResourcePolicyAttachmentsGetter has a method to return a ComputeDiskResourcePolicyAttachmentInterface.
// A group's client should implement this interface.
type ComputeDiskResourcePolicyAttachmentsGetter interface {
	ComputeDiskResourcePolicyAttachments(namespace string) ComputeDiskResourcePolicyAttachmentInterface
}

// ComputeDiskResourcePolicyAttachmentInterface has methods to work with ComputeDiskResourcePolicyAttachment resources.
type ComputeDiskResourcePolicyAttachmentInterface interface {
	Create(*v1alpha1.ComputeDiskResourcePolicyAttachment) (*v1alpha1.ComputeDiskResourcePolicyAttachment, error)
	Update(*v1alpha1.ComputeDiskResourcePolicyAttachment) (*v1alpha1.ComputeDiskResourcePolicyAttachment, error)
	UpdateStatus(*v1alpha1.ComputeDiskResourcePolicyAttachment) (*v1alpha1.ComputeDiskResourcePolicyAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeDiskResourcePolicyAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeDiskResourcePolicyAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error)
	ComputeDiskResourcePolicyAttachmentExpansion
}

// computeDiskResourcePolicyAttachments implements ComputeDiskResourcePolicyAttachmentInterface
type computeDiskResourcePolicyAttachments struct {
	client rest.Interface
	ns     string
}

// newComputeDiskResourcePolicyAttachments returns a ComputeDiskResourcePolicyAttachments
func newComputeDiskResourcePolicyAttachments(c *GoogleV1alpha1Client, namespace string) *computeDiskResourcePolicyAttachments {
	return &computeDiskResourcePolicyAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeDiskResourcePolicyAttachment, and returns the corresponding computeDiskResourcePolicyAttachment object, and an error if there is any.
func (c *computeDiskResourcePolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	result = &v1alpha1.ComputeDiskResourcePolicyAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeDiskResourcePolicyAttachments that match those selectors.
func (c *computeDiskResourcePolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.ComputeDiskResourcePolicyAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeDiskResourcePolicyAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeDiskResourcePolicyAttachments.
func (c *computeDiskResourcePolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeDiskResourcePolicyAttachment and creates it.  Returns the server's representation of the computeDiskResourcePolicyAttachment, and an error, if there is any.
func (c *computeDiskResourcePolicyAttachments) Create(computeDiskResourcePolicyAttachment *v1alpha1.ComputeDiskResourcePolicyAttachment) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	result = &v1alpha1.ComputeDiskResourcePolicyAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		Body(computeDiskResourcePolicyAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeDiskResourcePolicyAttachment and updates it. Returns the server's representation of the computeDiskResourcePolicyAttachment, and an error, if there is any.
func (c *computeDiskResourcePolicyAttachments) Update(computeDiskResourcePolicyAttachment *v1alpha1.ComputeDiskResourcePolicyAttachment) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	result = &v1alpha1.ComputeDiskResourcePolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		Name(computeDiskResourcePolicyAttachment.Name).
		Body(computeDiskResourcePolicyAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeDiskResourcePolicyAttachments) UpdateStatus(computeDiskResourcePolicyAttachment *v1alpha1.ComputeDiskResourcePolicyAttachment) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	result = &v1alpha1.ComputeDiskResourcePolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		Name(computeDiskResourcePolicyAttachment.Name).
		SubResource("status").
		Body(computeDiskResourcePolicyAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeDiskResourcePolicyAttachment and deletes it. Returns an error if one occurs.
func (c *computeDiskResourcePolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeDiskResourcePolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeDiskResourcePolicyAttachment.
func (c *computeDiskResourcePolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	result = &v1alpha1.ComputeDiskResourcePolicyAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computediskresourcepolicyattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
