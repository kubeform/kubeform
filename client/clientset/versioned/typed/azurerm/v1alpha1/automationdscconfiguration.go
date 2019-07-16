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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AutomationDscConfigurationsGetter has a method to return a AutomationDscConfigurationInterface.
// A group's client should implement this interface.
type AutomationDscConfigurationsGetter interface {
	AutomationDscConfigurations() AutomationDscConfigurationInterface
}

// AutomationDscConfigurationInterface has methods to work with AutomationDscConfiguration resources.
type AutomationDscConfigurationInterface interface {
	Create(*v1alpha1.AutomationDscConfiguration) (*v1alpha1.AutomationDscConfiguration, error)
	Update(*v1alpha1.AutomationDscConfiguration) (*v1alpha1.AutomationDscConfiguration, error)
	UpdateStatus(*v1alpha1.AutomationDscConfiguration) (*v1alpha1.AutomationDscConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutomationDscConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.AutomationDscConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationDscConfiguration, err error)
	AutomationDscConfigurationExpansion
}

// automationDscConfigurations implements AutomationDscConfigurationInterface
type automationDscConfigurations struct {
	client rest.Interface
}

// newAutomationDscConfigurations returns a AutomationDscConfigurations
func newAutomationDscConfigurations(c *AzurermV1alpha1Client) *automationDscConfigurations {
	return &automationDscConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the automationDscConfiguration, and returns the corresponding automationDscConfiguration object, and an error if there is any.
func (c *automationDscConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AutomationDscConfiguration, err error) {
	result = &v1alpha1.AutomationDscConfiguration{}
	err = c.client.Get().
		Resource("automationdscconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutomationDscConfigurations that match those selectors.
func (c *automationDscConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AutomationDscConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutomationDscConfigurationList{}
	err = c.client.Get().
		Resource("automationdscconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested automationDscConfigurations.
func (c *automationDscConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("automationdscconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a automationDscConfiguration and creates it.  Returns the server's representation of the automationDscConfiguration, and an error, if there is any.
func (c *automationDscConfigurations) Create(automationDscConfiguration *v1alpha1.AutomationDscConfiguration) (result *v1alpha1.AutomationDscConfiguration, err error) {
	result = &v1alpha1.AutomationDscConfiguration{}
	err = c.client.Post().
		Resource("automationdscconfigurations").
		Body(automationDscConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a automationDscConfiguration and updates it. Returns the server's representation of the automationDscConfiguration, and an error, if there is any.
func (c *automationDscConfigurations) Update(automationDscConfiguration *v1alpha1.AutomationDscConfiguration) (result *v1alpha1.AutomationDscConfiguration, err error) {
	result = &v1alpha1.AutomationDscConfiguration{}
	err = c.client.Put().
		Resource("automationdscconfigurations").
		Name(automationDscConfiguration.Name).
		Body(automationDscConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *automationDscConfigurations) UpdateStatus(automationDscConfiguration *v1alpha1.AutomationDscConfiguration) (result *v1alpha1.AutomationDscConfiguration, err error) {
	result = &v1alpha1.AutomationDscConfiguration{}
	err = c.client.Put().
		Resource("automationdscconfigurations").
		Name(automationDscConfiguration.Name).
		SubResource("status").
		Body(automationDscConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the automationDscConfiguration and deletes it. Returns an error if one occurs.
func (c *automationDscConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("automationdscconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *automationDscConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("automationdscconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched automationDscConfiguration.
func (c *automationDscConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationDscConfiguration, err error) {
	result = &v1alpha1.AutomationDscConfiguration{}
	err = c.client.Patch(pt).
		Resource("automationdscconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
