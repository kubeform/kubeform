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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// SpannerInstanceIamPoliciesGetter has a method to return a SpannerInstanceIamPolicyInterface.
// A group's client should implement this interface.
type SpannerInstanceIamPoliciesGetter interface {
	SpannerInstanceIamPolicies(namespace string) SpannerInstanceIamPolicyInterface
}

// SpannerInstanceIamPolicyInterface has methods to work with SpannerInstanceIamPolicy resources.
type SpannerInstanceIamPolicyInterface interface {
	Create(*v1alpha1.SpannerInstanceIamPolicy) (*v1alpha1.SpannerInstanceIamPolicy, error)
	Update(*v1alpha1.SpannerInstanceIamPolicy) (*v1alpha1.SpannerInstanceIamPolicy, error)
	UpdateStatus(*v1alpha1.SpannerInstanceIamPolicy) (*v1alpha1.SpannerInstanceIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SpannerInstanceIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.SpannerInstanceIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerInstanceIamPolicy, err error)
	SpannerInstanceIamPolicyExpansion
}

// spannerInstanceIamPolicies implements SpannerInstanceIamPolicyInterface
type spannerInstanceIamPolicies struct {
	client rest.Interface
	ns     string
}

// newSpannerInstanceIamPolicies returns a SpannerInstanceIamPolicies
func newSpannerInstanceIamPolicies(c *GoogleV1alpha1Client, namespace string) *spannerInstanceIamPolicies {
	return &spannerInstanceIamPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the spannerInstanceIamPolicy, and returns the corresponding spannerInstanceIamPolicy object, and an error if there is any.
func (c *spannerInstanceIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.SpannerInstanceIamPolicy, err error) {
	result = &v1alpha1.SpannerInstanceIamPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpannerInstanceIamPolicies that match those selectors.
func (c *spannerInstanceIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.SpannerInstanceIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpannerInstanceIamPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spannerInstanceIamPolicies.
func (c *spannerInstanceIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a spannerInstanceIamPolicy and creates it.  Returns the server's representation of the spannerInstanceIamPolicy, and an error, if there is any.
func (c *spannerInstanceIamPolicies) Create(spannerInstanceIamPolicy *v1alpha1.SpannerInstanceIamPolicy) (result *v1alpha1.SpannerInstanceIamPolicy, err error) {
	result = &v1alpha1.SpannerInstanceIamPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		Body(spannerInstanceIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a spannerInstanceIamPolicy and updates it. Returns the server's representation of the spannerInstanceIamPolicy, and an error, if there is any.
func (c *spannerInstanceIamPolicies) Update(spannerInstanceIamPolicy *v1alpha1.SpannerInstanceIamPolicy) (result *v1alpha1.SpannerInstanceIamPolicy, err error) {
	result = &v1alpha1.SpannerInstanceIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		Name(spannerInstanceIamPolicy.Name).
		Body(spannerInstanceIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *spannerInstanceIamPolicies) UpdateStatus(spannerInstanceIamPolicy *v1alpha1.SpannerInstanceIamPolicy) (result *v1alpha1.SpannerInstanceIamPolicy, err error) {
	result = &v1alpha1.SpannerInstanceIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		Name(spannerInstanceIamPolicy.Name).
		SubResource("status").
		Body(spannerInstanceIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the spannerInstanceIamPolicy and deletes it. Returns an error if one occurs.
func (c *spannerInstanceIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spannerInstanceIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched spannerInstanceIamPolicy.
func (c *spannerInstanceIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerInstanceIamPolicy, err error) {
	result = &v1alpha1.SpannerInstanceIamPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("spannerinstanceiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}