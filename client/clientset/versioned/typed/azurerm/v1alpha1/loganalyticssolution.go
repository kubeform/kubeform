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

// LogAnalyticsSolutionsGetter has a method to return a LogAnalyticsSolutionInterface.
// A group's client should implement this interface.
type LogAnalyticsSolutionsGetter interface {
	LogAnalyticsSolutions(namespace string) LogAnalyticsSolutionInterface
}

// LogAnalyticsSolutionInterface has methods to work with LogAnalyticsSolution resources.
type LogAnalyticsSolutionInterface interface {
	Create(*v1alpha1.LogAnalyticsSolution) (*v1alpha1.LogAnalyticsSolution, error)
	Update(*v1alpha1.LogAnalyticsSolution) (*v1alpha1.LogAnalyticsSolution, error)
	UpdateStatus(*v1alpha1.LogAnalyticsSolution) (*v1alpha1.LogAnalyticsSolution, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LogAnalyticsSolution, error)
	List(opts v1.ListOptions) (*v1alpha1.LogAnalyticsSolutionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogAnalyticsSolution, err error)
	LogAnalyticsSolutionExpansion
}

// logAnalyticsSolutions implements LogAnalyticsSolutionInterface
type logAnalyticsSolutions struct {
	client rest.Interface
	ns     string
}

// newLogAnalyticsSolutions returns a LogAnalyticsSolutions
func newLogAnalyticsSolutions(c *AzurermV1alpha1Client, namespace string) *logAnalyticsSolutions {
	return &logAnalyticsSolutions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the logAnalyticsSolution, and returns the corresponding logAnalyticsSolution object, and an error if there is any.
func (c *logAnalyticsSolutions) Get(name string, options v1.GetOptions) (result *v1alpha1.LogAnalyticsSolution, err error) {
	result = &v1alpha1.LogAnalyticsSolution{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LogAnalyticsSolutions that match those selectors.
func (c *logAnalyticsSolutions) List(opts v1.ListOptions) (result *v1alpha1.LogAnalyticsSolutionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LogAnalyticsSolutionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested logAnalyticsSolutions.
func (c *logAnalyticsSolutions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a logAnalyticsSolution and creates it.  Returns the server's representation of the logAnalyticsSolution, and an error, if there is any.
func (c *logAnalyticsSolutions) Create(logAnalyticsSolution *v1alpha1.LogAnalyticsSolution) (result *v1alpha1.LogAnalyticsSolution, err error) {
	result = &v1alpha1.LogAnalyticsSolution{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		Body(logAnalyticsSolution).
		Do().
		Into(result)
	return
}

// Update takes the representation of a logAnalyticsSolution and updates it. Returns the server's representation of the logAnalyticsSolution, and an error, if there is any.
func (c *logAnalyticsSolutions) Update(logAnalyticsSolution *v1alpha1.LogAnalyticsSolution) (result *v1alpha1.LogAnalyticsSolution, err error) {
	result = &v1alpha1.LogAnalyticsSolution{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		Name(logAnalyticsSolution.Name).
		Body(logAnalyticsSolution).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *logAnalyticsSolutions) UpdateStatus(logAnalyticsSolution *v1alpha1.LogAnalyticsSolution) (result *v1alpha1.LogAnalyticsSolution, err error) {
	result = &v1alpha1.LogAnalyticsSolution{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		Name(logAnalyticsSolution.Name).
		SubResource("status").
		Body(logAnalyticsSolution).
		Do().
		Into(result)
	return
}

// Delete takes name of the logAnalyticsSolution and deletes it. Returns an error if one occurs.
func (c *logAnalyticsSolutions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *logAnalyticsSolutions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched logAnalyticsSolution.
func (c *logAnalyticsSolutions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogAnalyticsSolution, err error) {
	result = &v1alpha1.LogAnalyticsSolution{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loganalyticssolutions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
