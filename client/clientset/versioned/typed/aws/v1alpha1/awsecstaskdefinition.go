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

// AwsEcsTaskDefinitionsGetter has a method to return a AwsEcsTaskDefinitionInterface.
// A group's client should implement this interface.
type AwsEcsTaskDefinitionsGetter interface {
	AwsEcsTaskDefinitions() AwsEcsTaskDefinitionInterface
}

// AwsEcsTaskDefinitionInterface has methods to work with AwsEcsTaskDefinition resources.
type AwsEcsTaskDefinitionInterface interface {
	Create(*v1alpha1.AwsEcsTaskDefinition) (*v1alpha1.AwsEcsTaskDefinition, error)
	Update(*v1alpha1.AwsEcsTaskDefinition) (*v1alpha1.AwsEcsTaskDefinition, error)
	UpdateStatus(*v1alpha1.AwsEcsTaskDefinition) (*v1alpha1.AwsEcsTaskDefinition, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEcsTaskDefinition, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEcsTaskDefinitionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcsTaskDefinition, err error)
	AwsEcsTaskDefinitionExpansion
}

// awsEcsTaskDefinitions implements AwsEcsTaskDefinitionInterface
type awsEcsTaskDefinitions struct {
	client rest.Interface
}

// newAwsEcsTaskDefinitions returns a AwsEcsTaskDefinitions
func newAwsEcsTaskDefinitions(c *AwsV1alpha1Client) *awsEcsTaskDefinitions {
	return &awsEcsTaskDefinitions{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsEcsTaskDefinition, and returns the corresponding awsEcsTaskDefinition object, and an error if there is any.
func (c *awsEcsTaskDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEcsTaskDefinition, err error) {
	result = &v1alpha1.AwsEcsTaskDefinition{}
	err = c.client.Get().
		Resource("awsecstaskdefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEcsTaskDefinitions that match those selectors.
func (c *awsEcsTaskDefinitions) List(opts v1.ListOptions) (result *v1alpha1.AwsEcsTaskDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsEcsTaskDefinitionList{}
	err = c.client.Get().
		Resource("awsecstaskdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEcsTaskDefinitions.
func (c *awsEcsTaskDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsecstaskdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsEcsTaskDefinition and creates it.  Returns the server's representation of the awsEcsTaskDefinition, and an error, if there is any.
func (c *awsEcsTaskDefinitions) Create(awsEcsTaskDefinition *v1alpha1.AwsEcsTaskDefinition) (result *v1alpha1.AwsEcsTaskDefinition, err error) {
	result = &v1alpha1.AwsEcsTaskDefinition{}
	err = c.client.Post().
		Resource("awsecstaskdefinitions").
		Body(awsEcsTaskDefinition).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEcsTaskDefinition and updates it. Returns the server's representation of the awsEcsTaskDefinition, and an error, if there is any.
func (c *awsEcsTaskDefinitions) Update(awsEcsTaskDefinition *v1alpha1.AwsEcsTaskDefinition) (result *v1alpha1.AwsEcsTaskDefinition, err error) {
	result = &v1alpha1.AwsEcsTaskDefinition{}
	err = c.client.Put().
		Resource("awsecstaskdefinitions").
		Name(awsEcsTaskDefinition.Name).
		Body(awsEcsTaskDefinition).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsEcsTaskDefinitions) UpdateStatus(awsEcsTaskDefinition *v1alpha1.AwsEcsTaskDefinition) (result *v1alpha1.AwsEcsTaskDefinition, err error) {
	result = &v1alpha1.AwsEcsTaskDefinition{}
	err = c.client.Put().
		Resource("awsecstaskdefinitions").
		Name(awsEcsTaskDefinition.Name).
		SubResource("status").
		Body(awsEcsTaskDefinition).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEcsTaskDefinition and deletes it. Returns an error if one occurs.
func (c *awsEcsTaskDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsecstaskdefinitions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEcsTaskDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsecstaskdefinitions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEcsTaskDefinition.
func (c *awsEcsTaskDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEcsTaskDefinition, err error) {
	result = &v1alpha1.AwsEcsTaskDefinition{}
	err = c.client.Patch(pt).
		Resource("awsecstaskdefinitions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
