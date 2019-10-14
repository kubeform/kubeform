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

// LoggingOrganizationExclusionsGetter has a method to return a LoggingOrganizationExclusionInterface.
// A group's client should implement this interface.
type LoggingOrganizationExclusionsGetter interface {
	LoggingOrganizationExclusions(namespace string) LoggingOrganizationExclusionInterface
}

// LoggingOrganizationExclusionInterface has methods to work with LoggingOrganizationExclusion resources.
type LoggingOrganizationExclusionInterface interface {
	Create(*v1alpha1.LoggingOrganizationExclusion) (*v1alpha1.LoggingOrganizationExclusion, error)
	Update(*v1alpha1.LoggingOrganizationExclusion) (*v1alpha1.LoggingOrganizationExclusion, error)
	UpdateStatus(*v1alpha1.LoggingOrganizationExclusion) (*v1alpha1.LoggingOrganizationExclusion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LoggingOrganizationExclusion, error)
	List(opts v1.ListOptions) (*v1alpha1.LoggingOrganizationExclusionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingOrganizationExclusion, err error)
	LoggingOrganizationExclusionExpansion
}

// loggingOrganizationExclusions implements LoggingOrganizationExclusionInterface
type loggingOrganizationExclusions struct {
	client rest.Interface
	ns     string
}

// newLoggingOrganizationExclusions returns a LoggingOrganizationExclusions
func newLoggingOrganizationExclusions(c *GoogleV1alpha1Client, namespace string) *loggingOrganizationExclusions {
	return &loggingOrganizationExclusions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loggingOrganizationExclusion, and returns the corresponding loggingOrganizationExclusion object, and an error if there is any.
func (c *loggingOrganizationExclusions) Get(name string, options v1.GetOptions) (result *v1alpha1.LoggingOrganizationExclusion, err error) {
	result = &v1alpha1.LoggingOrganizationExclusion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoggingOrganizationExclusions that match those selectors.
func (c *loggingOrganizationExclusions) List(opts v1.ListOptions) (result *v1alpha1.LoggingOrganizationExclusionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LoggingOrganizationExclusionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loggingOrganizationExclusions.
func (c *loggingOrganizationExclusions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a loggingOrganizationExclusion and creates it.  Returns the server's representation of the loggingOrganizationExclusion, and an error, if there is any.
func (c *loggingOrganizationExclusions) Create(loggingOrganizationExclusion *v1alpha1.LoggingOrganizationExclusion) (result *v1alpha1.LoggingOrganizationExclusion, err error) {
	result = &v1alpha1.LoggingOrganizationExclusion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		Body(loggingOrganizationExclusion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a loggingOrganizationExclusion and updates it. Returns the server's representation of the loggingOrganizationExclusion, and an error, if there is any.
func (c *loggingOrganizationExclusions) Update(loggingOrganizationExclusion *v1alpha1.LoggingOrganizationExclusion) (result *v1alpha1.LoggingOrganizationExclusion, err error) {
	result = &v1alpha1.LoggingOrganizationExclusion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		Name(loggingOrganizationExclusion.Name).
		Body(loggingOrganizationExclusion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *loggingOrganizationExclusions) UpdateStatus(loggingOrganizationExclusion *v1alpha1.LoggingOrganizationExclusion) (result *v1alpha1.LoggingOrganizationExclusion, err error) {
	result = &v1alpha1.LoggingOrganizationExclusion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		Name(loggingOrganizationExclusion.Name).
		SubResource("status").
		Body(loggingOrganizationExclusion).
		Do().
		Into(result)
	return
}

// Delete takes name of the loggingOrganizationExclusion and deletes it. Returns an error if one occurs.
func (c *loggingOrganizationExclusions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loggingOrganizationExclusions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched loggingOrganizationExclusion.
func (c *loggingOrganizationExclusions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingOrganizationExclusion, err error) {
	result = &v1alpha1.LoggingOrganizationExclusion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loggingorganizationexclusions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
