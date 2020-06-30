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

// CloudfunctionsFunctionIamBindingsGetter has a method to return a CloudfunctionsFunctionIamBindingInterface.
// A group's client should implement this interface.
type CloudfunctionsFunctionIamBindingsGetter interface {
	CloudfunctionsFunctionIamBindings(namespace string) CloudfunctionsFunctionIamBindingInterface
}

// CloudfunctionsFunctionIamBindingInterface has methods to work with CloudfunctionsFunctionIamBinding resources.
type CloudfunctionsFunctionIamBindingInterface interface {
	Create(*v1alpha1.CloudfunctionsFunctionIamBinding) (*v1alpha1.CloudfunctionsFunctionIamBinding, error)
	Update(*v1alpha1.CloudfunctionsFunctionIamBinding) (*v1alpha1.CloudfunctionsFunctionIamBinding, error)
	UpdateStatus(*v1alpha1.CloudfunctionsFunctionIamBinding) (*v1alpha1.CloudfunctionsFunctionIamBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudfunctionsFunctionIamBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudfunctionsFunctionIamBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfunctionsFunctionIamBinding, err error)
	CloudfunctionsFunctionIamBindingExpansion
}

// cloudfunctionsFunctionIamBindings implements CloudfunctionsFunctionIamBindingInterface
type cloudfunctionsFunctionIamBindings struct {
	client rest.Interface
	ns     string
}

// newCloudfunctionsFunctionIamBindings returns a CloudfunctionsFunctionIamBindings
func newCloudfunctionsFunctionIamBindings(c *GoogleV1alpha1Client, namespace string) *cloudfunctionsFunctionIamBindings {
	return &cloudfunctionsFunctionIamBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudfunctionsFunctionIamBinding, and returns the corresponding cloudfunctionsFunctionIamBinding object, and an error if there is any.
func (c *cloudfunctionsFunctionIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudfunctionsFunctionIamBinding, err error) {
	result = &v1alpha1.CloudfunctionsFunctionIamBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudfunctionsFunctionIamBindings that match those selectors.
func (c *cloudfunctionsFunctionIamBindings) List(opts v1.ListOptions) (result *v1alpha1.CloudfunctionsFunctionIamBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudfunctionsFunctionIamBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudfunctionsFunctionIamBindings.
func (c *cloudfunctionsFunctionIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudfunctionsFunctionIamBinding and creates it.  Returns the server's representation of the cloudfunctionsFunctionIamBinding, and an error, if there is any.
func (c *cloudfunctionsFunctionIamBindings) Create(cloudfunctionsFunctionIamBinding *v1alpha1.CloudfunctionsFunctionIamBinding) (result *v1alpha1.CloudfunctionsFunctionIamBinding, err error) {
	result = &v1alpha1.CloudfunctionsFunctionIamBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		Body(cloudfunctionsFunctionIamBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudfunctionsFunctionIamBinding and updates it. Returns the server's representation of the cloudfunctionsFunctionIamBinding, and an error, if there is any.
func (c *cloudfunctionsFunctionIamBindings) Update(cloudfunctionsFunctionIamBinding *v1alpha1.CloudfunctionsFunctionIamBinding) (result *v1alpha1.CloudfunctionsFunctionIamBinding, err error) {
	result = &v1alpha1.CloudfunctionsFunctionIamBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		Name(cloudfunctionsFunctionIamBinding.Name).
		Body(cloudfunctionsFunctionIamBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudfunctionsFunctionIamBindings) UpdateStatus(cloudfunctionsFunctionIamBinding *v1alpha1.CloudfunctionsFunctionIamBinding) (result *v1alpha1.CloudfunctionsFunctionIamBinding, err error) {
	result = &v1alpha1.CloudfunctionsFunctionIamBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		Name(cloudfunctionsFunctionIamBinding.Name).
		SubResource("status").
		Body(cloudfunctionsFunctionIamBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudfunctionsFunctionIamBinding and deletes it. Returns an error if one occurs.
func (c *cloudfunctionsFunctionIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudfunctionsFunctionIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudfunctionsFunctionIamBinding.
func (c *cloudfunctionsFunctionIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfunctionsFunctionIamBinding, err error) {
	result = &v1alpha1.CloudfunctionsFunctionIamBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudfunctionsfunctioniambindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
