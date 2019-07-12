/*
Copyright 2019 The KubeDB Authors.

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

// AwsMqConfigurationsGetter has a method to return a AwsMqConfigurationInterface.
// A group's client should implement this interface.
type AwsMqConfigurationsGetter interface {
	AwsMqConfigurations() AwsMqConfigurationInterface
}

// AwsMqConfigurationInterface has methods to work with AwsMqConfiguration resources.
type AwsMqConfigurationInterface interface {
	Create(*v1alpha1.AwsMqConfiguration) (*v1alpha1.AwsMqConfiguration, error)
	Update(*v1alpha1.AwsMqConfiguration) (*v1alpha1.AwsMqConfiguration, error)
	UpdateStatus(*v1alpha1.AwsMqConfiguration) (*v1alpha1.AwsMqConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsMqConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsMqConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsMqConfiguration, err error)
	AwsMqConfigurationExpansion
}

// awsMqConfigurations implements AwsMqConfigurationInterface
type awsMqConfigurations struct {
	client rest.Interface
}

// newAwsMqConfigurations returns a AwsMqConfigurations
func newAwsMqConfigurations(c *AwsV1alpha1Client) *awsMqConfigurations {
	return &awsMqConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsMqConfiguration, and returns the corresponding awsMqConfiguration object, and an error if there is any.
func (c *awsMqConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsMqConfiguration, err error) {
	result = &v1alpha1.AwsMqConfiguration{}
	err = c.client.Get().
		Resource("awsmqconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsMqConfigurations that match those selectors.
func (c *awsMqConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AwsMqConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsMqConfigurationList{}
	err = c.client.Get().
		Resource("awsmqconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsMqConfigurations.
func (c *awsMqConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsmqconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsMqConfiguration and creates it.  Returns the server's representation of the awsMqConfiguration, and an error, if there is any.
func (c *awsMqConfigurations) Create(awsMqConfiguration *v1alpha1.AwsMqConfiguration) (result *v1alpha1.AwsMqConfiguration, err error) {
	result = &v1alpha1.AwsMqConfiguration{}
	err = c.client.Post().
		Resource("awsmqconfigurations").
		Body(awsMqConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsMqConfiguration and updates it. Returns the server's representation of the awsMqConfiguration, and an error, if there is any.
func (c *awsMqConfigurations) Update(awsMqConfiguration *v1alpha1.AwsMqConfiguration) (result *v1alpha1.AwsMqConfiguration, err error) {
	result = &v1alpha1.AwsMqConfiguration{}
	err = c.client.Put().
		Resource("awsmqconfigurations").
		Name(awsMqConfiguration.Name).
		Body(awsMqConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsMqConfigurations) UpdateStatus(awsMqConfiguration *v1alpha1.AwsMqConfiguration) (result *v1alpha1.AwsMqConfiguration, err error) {
	result = &v1alpha1.AwsMqConfiguration{}
	err = c.client.Put().
		Resource("awsmqconfigurations").
		Name(awsMqConfiguration.Name).
		SubResource("status").
		Body(awsMqConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsMqConfiguration and deletes it. Returns an error if one occurs.
func (c *awsMqConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsmqconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsMqConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsmqconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsMqConfiguration.
func (c *awsMqConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsMqConfiguration, err error) {
	result = &v1alpha1.AwsMqConfiguration{}
	err = c.client.Patch(pt).
		Resource("awsmqconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
