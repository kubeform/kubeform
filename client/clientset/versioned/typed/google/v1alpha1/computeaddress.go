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

// ComputeAddressesGetter has a method to return a ComputeAddressInterface.
// A group's client should implement this interface.
type ComputeAddressesGetter interface {
	ComputeAddresses(namespace string) ComputeAddressInterface
}

// ComputeAddressInterface has methods to work with ComputeAddress resources.
type ComputeAddressInterface interface {
	Create(*v1alpha1.ComputeAddress) (*v1alpha1.ComputeAddress, error)
	Update(*v1alpha1.ComputeAddress) (*v1alpha1.ComputeAddress, error)
	UpdateStatus(*v1alpha1.ComputeAddress) (*v1alpha1.ComputeAddress, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeAddress, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeAddressList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeAddress, err error)
	ComputeAddressExpansion
}

// computeAddresses implements ComputeAddressInterface
type computeAddresses struct {
	client rest.Interface
	ns     string
}

// newComputeAddresses returns a ComputeAddresses
func newComputeAddresses(c *GoogleV1alpha1Client, namespace string) *computeAddresses {
	return &computeAddresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeAddress, and returns the corresponding computeAddress object, and an error if there is any.
func (c *computeAddresses) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeAddress, err error) {
	result = &v1alpha1.ComputeAddress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeaddresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeAddresses that match those selectors.
func (c *computeAddresses) List(opts v1.ListOptions) (result *v1alpha1.ComputeAddressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeAddressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeaddresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeAddresses.
func (c *computeAddresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeaddresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeAddress and creates it.  Returns the server's representation of the computeAddress, and an error, if there is any.
func (c *computeAddresses) Create(computeAddress *v1alpha1.ComputeAddress) (result *v1alpha1.ComputeAddress, err error) {
	result = &v1alpha1.ComputeAddress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeaddresses").
		Body(computeAddress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeAddress and updates it. Returns the server's representation of the computeAddress, and an error, if there is any.
func (c *computeAddresses) Update(computeAddress *v1alpha1.ComputeAddress) (result *v1alpha1.ComputeAddress, err error) {
	result = &v1alpha1.ComputeAddress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeaddresses").
		Name(computeAddress.Name).
		Body(computeAddress).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeAddresses) UpdateStatus(computeAddress *v1alpha1.ComputeAddress) (result *v1alpha1.ComputeAddress, err error) {
	result = &v1alpha1.ComputeAddress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeaddresses").
		Name(computeAddress.Name).
		SubResource("status").
		Body(computeAddress).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeAddress and deletes it. Returns an error if one occurs.
func (c *computeAddresses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeaddresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeAddresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeaddresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeAddress.
func (c *computeAddresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeAddress, err error) {
	result = &v1alpha1.ComputeAddress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeaddresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
