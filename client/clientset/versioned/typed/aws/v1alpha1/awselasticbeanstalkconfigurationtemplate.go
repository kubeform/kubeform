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

// AwsElasticBeanstalkConfigurationTemplatesGetter has a method to return a AwsElasticBeanstalkConfigurationTemplateInterface.
// A group's client should implement this interface.
type AwsElasticBeanstalkConfigurationTemplatesGetter interface {
	AwsElasticBeanstalkConfigurationTemplates() AwsElasticBeanstalkConfigurationTemplateInterface
}

// AwsElasticBeanstalkConfigurationTemplateInterface has methods to work with AwsElasticBeanstalkConfigurationTemplate resources.
type AwsElasticBeanstalkConfigurationTemplateInterface interface {
	Create(*v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (*v1alpha1.AwsElasticBeanstalkConfigurationTemplate, error)
	Update(*v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (*v1alpha1.AwsElasticBeanstalkConfigurationTemplate, error)
	UpdateStatus(*v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (*v1alpha1.AwsElasticBeanstalkConfigurationTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsElasticBeanstalkConfigurationTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsElasticBeanstalkConfigurationTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error)
	AwsElasticBeanstalkConfigurationTemplateExpansion
}

// awsElasticBeanstalkConfigurationTemplates implements AwsElasticBeanstalkConfigurationTemplateInterface
type awsElasticBeanstalkConfigurationTemplates struct {
	client rest.Interface
}

// newAwsElasticBeanstalkConfigurationTemplates returns a AwsElasticBeanstalkConfigurationTemplates
func newAwsElasticBeanstalkConfigurationTemplates(c *AwsV1alpha1Client) *awsElasticBeanstalkConfigurationTemplates {
	return &awsElasticBeanstalkConfigurationTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsElasticBeanstalkConfigurationTemplate, and returns the corresponding awsElasticBeanstalkConfigurationTemplate object, and an error if there is any.
func (c *awsElasticBeanstalkConfigurationTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{}
	err = c.client.Get().
		Resource("awselasticbeanstalkconfigurationtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsElasticBeanstalkConfigurationTemplates that match those selectors.
func (c *awsElasticBeanstalkConfigurationTemplates) List(opts v1.ListOptions) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsElasticBeanstalkConfigurationTemplateList{}
	err = c.client.Get().
		Resource("awselasticbeanstalkconfigurationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsElasticBeanstalkConfigurationTemplates.
func (c *awsElasticBeanstalkConfigurationTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awselasticbeanstalkconfigurationtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsElasticBeanstalkConfigurationTemplate and creates it.  Returns the server's representation of the awsElasticBeanstalkConfigurationTemplate, and an error, if there is any.
func (c *awsElasticBeanstalkConfigurationTemplates) Create(awsElasticBeanstalkConfigurationTemplate *v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{}
	err = c.client.Post().
		Resource("awselasticbeanstalkconfigurationtemplates").
		Body(awsElasticBeanstalkConfigurationTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsElasticBeanstalkConfigurationTemplate and updates it. Returns the server's representation of the awsElasticBeanstalkConfigurationTemplate, and an error, if there is any.
func (c *awsElasticBeanstalkConfigurationTemplates) Update(awsElasticBeanstalkConfigurationTemplate *v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{}
	err = c.client.Put().
		Resource("awselasticbeanstalkconfigurationtemplates").
		Name(awsElasticBeanstalkConfigurationTemplate.Name).
		Body(awsElasticBeanstalkConfigurationTemplate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsElasticBeanstalkConfigurationTemplates) UpdateStatus(awsElasticBeanstalkConfigurationTemplate *v1alpha1.AwsElasticBeanstalkConfigurationTemplate) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{}
	err = c.client.Put().
		Resource("awselasticbeanstalkconfigurationtemplates").
		Name(awsElasticBeanstalkConfigurationTemplate.Name).
		SubResource("status").
		Body(awsElasticBeanstalkConfigurationTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsElasticBeanstalkConfigurationTemplate and deletes it. Returns an error if one occurs.
func (c *awsElasticBeanstalkConfigurationTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awselasticbeanstalkconfigurationtemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsElasticBeanstalkConfigurationTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awselasticbeanstalkconfigurationtemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsElasticBeanstalkConfigurationTemplate.
func (c *awsElasticBeanstalkConfigurationTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsElasticBeanstalkConfigurationTemplate, err error) {
	result = &v1alpha1.AwsElasticBeanstalkConfigurationTemplate{}
	err = c.client.Patch(pt).
		Resource("awselasticbeanstalkconfigurationtemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
