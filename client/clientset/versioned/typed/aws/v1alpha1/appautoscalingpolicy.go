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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppautoscalingPoliciesGetter has a method to return a AppautoscalingPolicyInterface.
// A group's client should implement this interface.
type AppautoscalingPoliciesGetter interface {
	AppautoscalingPolicies(namespace string) AppautoscalingPolicyInterface
}

// AppautoscalingPolicyInterface has methods to work with AppautoscalingPolicy resources.
type AppautoscalingPolicyInterface interface {
	Create(*v1alpha1.AppautoscalingPolicy) (*v1alpha1.AppautoscalingPolicy, error)
	Update(*v1alpha1.AppautoscalingPolicy) (*v1alpha1.AppautoscalingPolicy, error)
	UpdateStatus(*v1alpha1.AppautoscalingPolicy) (*v1alpha1.AppautoscalingPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppautoscalingPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.AppautoscalingPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppautoscalingPolicy, err error)
	AppautoscalingPolicyExpansion
}

// appautoscalingPolicies implements AppautoscalingPolicyInterface
type appautoscalingPolicies struct {
	client rest.Interface
	ns     string
}

// newAppautoscalingPolicies returns a AppautoscalingPolicies
func newAppautoscalingPolicies(c *AwsV1alpha1Client, namespace string) *appautoscalingPolicies {
	return &appautoscalingPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appautoscalingPolicy, and returns the corresponding appautoscalingPolicy object, and an error if there is any.
func (c *appautoscalingPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.AppautoscalingPolicy, err error) {
	result = &v1alpha1.AppautoscalingPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppautoscalingPolicies that match those selectors.
func (c *appautoscalingPolicies) List(opts v1.ListOptions) (result *v1alpha1.AppautoscalingPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppautoscalingPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appautoscalingPolicies.
func (c *appautoscalingPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appautoscalingPolicy and creates it.  Returns the server's representation of the appautoscalingPolicy, and an error, if there is any.
func (c *appautoscalingPolicies) Create(appautoscalingPolicy *v1alpha1.AppautoscalingPolicy) (result *v1alpha1.AppautoscalingPolicy, err error) {
	result = &v1alpha1.AppautoscalingPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		Body(appautoscalingPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appautoscalingPolicy and updates it. Returns the server's representation of the appautoscalingPolicy, and an error, if there is any.
func (c *appautoscalingPolicies) Update(appautoscalingPolicy *v1alpha1.AppautoscalingPolicy) (result *v1alpha1.AppautoscalingPolicy, err error) {
	result = &v1alpha1.AppautoscalingPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		Name(appautoscalingPolicy.Name).
		Body(appautoscalingPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appautoscalingPolicies) UpdateStatus(appautoscalingPolicy *v1alpha1.AppautoscalingPolicy) (result *v1alpha1.AppautoscalingPolicy, err error) {
	result = &v1alpha1.AppautoscalingPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		Name(appautoscalingPolicy.Name).
		SubResource("status").
		Body(appautoscalingPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the appautoscalingPolicy and deletes it. Returns an error if one occurs.
func (c *appautoscalingPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appautoscalingPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appautoscalingPolicy.
func (c *appautoscalingPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppautoscalingPolicy, err error) {
	result = &v1alpha1.AppautoscalingPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appautoscalingpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
