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

// LaunchConfigurationsGetter has a method to return a LaunchConfigurationInterface.
// A group's client should implement this interface.
type LaunchConfigurationsGetter interface {
	LaunchConfigurations(namespace string) LaunchConfigurationInterface
}

// LaunchConfigurationInterface has methods to work with LaunchConfiguration resources.
type LaunchConfigurationInterface interface {
	Create(*v1alpha1.LaunchConfiguration) (*v1alpha1.LaunchConfiguration, error)
	Update(*v1alpha1.LaunchConfiguration) (*v1alpha1.LaunchConfiguration, error)
	UpdateStatus(*v1alpha1.LaunchConfiguration) (*v1alpha1.LaunchConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LaunchConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.LaunchConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LaunchConfiguration, err error)
	LaunchConfigurationExpansion
}

// launchConfigurations implements LaunchConfigurationInterface
type launchConfigurations struct {
	client rest.Interface
	ns     string
}

// newLaunchConfigurations returns a LaunchConfigurations
func newLaunchConfigurations(c *AwsV1alpha1Client, namespace string) *launchConfigurations {
	return &launchConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the launchConfiguration, and returns the corresponding launchConfiguration object, and an error if there is any.
func (c *launchConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.LaunchConfiguration, err error) {
	result = &v1alpha1.LaunchConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("launchconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LaunchConfigurations that match those selectors.
func (c *launchConfigurations) List(opts v1.ListOptions) (result *v1alpha1.LaunchConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LaunchConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("launchconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested launchConfigurations.
func (c *launchConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("launchconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a launchConfiguration and creates it.  Returns the server's representation of the launchConfiguration, and an error, if there is any.
func (c *launchConfigurations) Create(launchConfiguration *v1alpha1.LaunchConfiguration) (result *v1alpha1.LaunchConfiguration, err error) {
	result = &v1alpha1.LaunchConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("launchconfigurations").
		Body(launchConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a launchConfiguration and updates it. Returns the server's representation of the launchConfiguration, and an error, if there is any.
func (c *launchConfigurations) Update(launchConfiguration *v1alpha1.LaunchConfiguration) (result *v1alpha1.LaunchConfiguration, err error) {
	result = &v1alpha1.LaunchConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("launchconfigurations").
		Name(launchConfiguration.Name).
		Body(launchConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *launchConfigurations) UpdateStatus(launchConfiguration *v1alpha1.LaunchConfiguration) (result *v1alpha1.LaunchConfiguration, err error) {
	result = &v1alpha1.LaunchConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("launchconfigurations").
		Name(launchConfiguration.Name).
		SubResource("status").
		Body(launchConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the launchConfiguration and deletes it. Returns an error if one occurs.
func (c *launchConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("launchconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *launchConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("launchconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched launchConfiguration.
func (c *launchConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LaunchConfiguration, err error) {
	result = &v1alpha1.LaunchConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("launchconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
