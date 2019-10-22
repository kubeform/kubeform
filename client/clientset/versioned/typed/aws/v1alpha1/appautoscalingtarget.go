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

// AppautoscalingTargetsGetter has a method to return a AppautoscalingTargetInterface.
// A group's client should implement this interface.
type AppautoscalingTargetsGetter interface {
	AppautoscalingTargets(namespace string) AppautoscalingTargetInterface
}

// AppautoscalingTargetInterface has methods to work with AppautoscalingTarget resources.
type AppautoscalingTargetInterface interface {
	Create(*v1alpha1.AppautoscalingTarget) (*v1alpha1.AppautoscalingTarget, error)
	Update(*v1alpha1.AppautoscalingTarget) (*v1alpha1.AppautoscalingTarget, error)
	UpdateStatus(*v1alpha1.AppautoscalingTarget) (*v1alpha1.AppautoscalingTarget, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppautoscalingTarget, error)
	List(opts v1.ListOptions) (*v1alpha1.AppautoscalingTargetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppautoscalingTarget, err error)
	AppautoscalingTargetExpansion
}

// appautoscalingTargets implements AppautoscalingTargetInterface
type appautoscalingTargets struct {
	client rest.Interface
	ns     string
}

// newAppautoscalingTargets returns a AppautoscalingTargets
func newAppautoscalingTargets(c *AwsV1alpha1Client, namespace string) *appautoscalingTargets {
	return &appautoscalingTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appautoscalingTarget, and returns the corresponding appautoscalingTarget object, and an error if there is any.
func (c *appautoscalingTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.AppautoscalingTarget, err error) {
	result = &v1alpha1.AppautoscalingTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppautoscalingTargets that match those selectors.
func (c *appautoscalingTargets) List(opts v1.ListOptions) (result *v1alpha1.AppautoscalingTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppautoscalingTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appautoscalingTargets.
func (c *appautoscalingTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appautoscalingTarget and creates it.  Returns the server's representation of the appautoscalingTarget, and an error, if there is any.
func (c *appautoscalingTargets) Create(appautoscalingTarget *v1alpha1.AppautoscalingTarget) (result *v1alpha1.AppautoscalingTarget, err error) {
	result = &v1alpha1.AppautoscalingTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		Body(appautoscalingTarget).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appautoscalingTarget and updates it. Returns the server's representation of the appautoscalingTarget, and an error, if there is any.
func (c *appautoscalingTargets) Update(appautoscalingTarget *v1alpha1.AppautoscalingTarget) (result *v1alpha1.AppautoscalingTarget, err error) {
	result = &v1alpha1.AppautoscalingTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		Name(appautoscalingTarget.Name).
		Body(appautoscalingTarget).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appautoscalingTargets) UpdateStatus(appautoscalingTarget *v1alpha1.AppautoscalingTarget) (result *v1alpha1.AppautoscalingTarget, err error) {
	result = &v1alpha1.AppautoscalingTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		Name(appautoscalingTarget.Name).
		SubResource("status").
		Body(appautoscalingTarget).
		Do().
		Into(result)
	return
}

// Delete takes name of the appautoscalingTarget and deletes it. Returns an error if one occurs.
func (c *appautoscalingTargets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appautoscalingTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appautoscalingTarget.
func (c *appautoscalingTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppautoscalingTarget, err error) {
	result = &v1alpha1.AppautoscalingTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appautoscalingtargets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
