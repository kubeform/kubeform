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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigConfigurationAggregatorsGetter has a method to return a ConfigConfigurationAggregatorInterface.
// A group's client should implement this interface.
type ConfigConfigurationAggregatorsGetter interface {
	ConfigConfigurationAggregators(namespace string) ConfigConfigurationAggregatorInterface
}

// ConfigConfigurationAggregatorInterface has methods to work with ConfigConfigurationAggregator resources.
type ConfigConfigurationAggregatorInterface interface {
	Create(*v1alpha1.ConfigConfigurationAggregator) (*v1alpha1.ConfigConfigurationAggregator, error)
	Update(*v1alpha1.ConfigConfigurationAggregator) (*v1alpha1.ConfigConfigurationAggregator, error)
	UpdateStatus(*v1alpha1.ConfigConfigurationAggregator) (*v1alpha1.ConfigConfigurationAggregator, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ConfigConfigurationAggregator, error)
	List(opts v1.ListOptions) (*v1alpha1.ConfigConfigurationAggregatorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigConfigurationAggregator, err error)
	ConfigConfigurationAggregatorExpansion
}

// configConfigurationAggregators implements ConfigConfigurationAggregatorInterface
type configConfigurationAggregators struct {
	client rest.Interface
	ns     string
}

// newConfigConfigurationAggregators returns a ConfigConfigurationAggregators
func newConfigConfigurationAggregators(c *AwsV1alpha1Client, namespace string) *configConfigurationAggregators {
	return &configConfigurationAggregators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configConfigurationAggregator, and returns the corresponding configConfigurationAggregator object, and an error if there is any.
func (c *configConfigurationAggregators) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigConfigurationAggregator, err error) {
	result = &v1alpha1.ConfigConfigurationAggregator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigConfigurationAggregators that match those selectors.
func (c *configConfigurationAggregators) List(opts v1.ListOptions) (result *v1alpha1.ConfigConfigurationAggregatorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigConfigurationAggregatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configConfigurationAggregators.
func (c *configConfigurationAggregators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a configConfigurationAggregator and creates it.  Returns the server's representation of the configConfigurationAggregator, and an error, if there is any.
func (c *configConfigurationAggregators) Create(configConfigurationAggregator *v1alpha1.ConfigConfigurationAggregator) (result *v1alpha1.ConfigConfigurationAggregator, err error) {
	result = &v1alpha1.ConfigConfigurationAggregator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		Body(configConfigurationAggregator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a configConfigurationAggregator and updates it. Returns the server's representation of the configConfigurationAggregator, and an error, if there is any.
func (c *configConfigurationAggregators) Update(configConfigurationAggregator *v1alpha1.ConfigConfigurationAggregator) (result *v1alpha1.ConfigConfigurationAggregator, err error) {
	result = &v1alpha1.ConfigConfigurationAggregator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		Name(configConfigurationAggregator.Name).
		Body(configConfigurationAggregator).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *configConfigurationAggregators) UpdateStatus(configConfigurationAggregator *v1alpha1.ConfigConfigurationAggregator) (result *v1alpha1.ConfigConfigurationAggregator, err error) {
	result = &v1alpha1.ConfigConfigurationAggregator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		Name(configConfigurationAggregator.Name).
		SubResource("status").
		Body(configConfigurationAggregator).
		Do().
		Into(result)
	return
}

// Delete takes name of the configConfigurationAggregator and deletes it. Returns an error if one occurs.
func (c *configConfigurationAggregators) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configConfigurationAggregators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched configConfigurationAggregator.
func (c *configConfigurationAggregators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigConfigurationAggregator, err error) {
	result = &v1alpha1.ConfigConfigurationAggregator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configconfigurationaggregators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
