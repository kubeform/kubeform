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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MysqlConfigurationsGetter has a method to return a MysqlConfigurationInterface.
// A group's client should implement this interface.
type MysqlConfigurationsGetter interface {
	MysqlConfigurations(namespace string) MysqlConfigurationInterface
}

// MysqlConfigurationInterface has methods to work with MysqlConfiguration resources.
type MysqlConfigurationInterface interface {
	Create(*v1alpha1.MysqlConfiguration) (*v1alpha1.MysqlConfiguration, error)
	Update(*v1alpha1.MysqlConfiguration) (*v1alpha1.MysqlConfiguration, error)
	UpdateStatus(*v1alpha1.MysqlConfiguration) (*v1alpha1.MysqlConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MysqlConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.MysqlConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlConfiguration, err error)
	MysqlConfigurationExpansion
}

// mysqlConfigurations implements MysqlConfigurationInterface
type mysqlConfigurations struct {
	client rest.Interface
	ns     string
}

// newMysqlConfigurations returns a MysqlConfigurations
func newMysqlConfigurations(c *AzurermV1alpha1Client, namespace string) *mysqlConfigurations {
	return &mysqlConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mysqlConfiguration, and returns the corresponding mysqlConfiguration object, and an error if there is any.
func (c *mysqlConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.MysqlConfiguration, err error) {
	result = &v1alpha1.MysqlConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MysqlConfigurations that match those selectors.
func (c *mysqlConfigurations) List(opts v1.ListOptions) (result *v1alpha1.MysqlConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mysqlConfigurations.
func (c *mysqlConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mysqlConfiguration and creates it.  Returns the server's representation of the mysqlConfiguration, and an error, if there is any.
func (c *mysqlConfigurations) Create(mysqlConfiguration *v1alpha1.MysqlConfiguration) (result *v1alpha1.MysqlConfiguration, err error) {
	result = &v1alpha1.MysqlConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		Body(mysqlConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mysqlConfiguration and updates it. Returns the server's representation of the mysqlConfiguration, and an error, if there is any.
func (c *mysqlConfigurations) Update(mysqlConfiguration *v1alpha1.MysqlConfiguration) (result *v1alpha1.MysqlConfiguration, err error) {
	result = &v1alpha1.MysqlConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		Name(mysqlConfiguration.Name).
		Body(mysqlConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mysqlConfigurations) UpdateStatus(mysqlConfiguration *v1alpha1.MysqlConfiguration) (result *v1alpha1.MysqlConfiguration, err error) {
	result = &v1alpha1.MysqlConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		Name(mysqlConfiguration.Name).
		SubResource("status").
		Body(mysqlConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the mysqlConfiguration and deletes it. Returns an error if one occurs.
func (c *mysqlConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mysqlConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mysqlConfiguration.
func (c *mysqlConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlConfiguration, err error) {
	result = &v1alpha1.MysqlConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
