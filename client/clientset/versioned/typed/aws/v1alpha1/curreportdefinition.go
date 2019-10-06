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

// CurReportDefinitionsGetter has a method to return a CurReportDefinitionInterface.
// A group's client should implement this interface.
type CurReportDefinitionsGetter interface {
	CurReportDefinitions(namespace string) CurReportDefinitionInterface
}

// CurReportDefinitionInterface has methods to work with CurReportDefinition resources.
type CurReportDefinitionInterface interface {
	Create(*v1alpha1.CurReportDefinition) (*v1alpha1.CurReportDefinition, error)
	Update(*v1alpha1.CurReportDefinition) (*v1alpha1.CurReportDefinition, error)
	UpdateStatus(*v1alpha1.CurReportDefinition) (*v1alpha1.CurReportDefinition, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CurReportDefinition, error)
	List(opts v1.ListOptions) (*v1alpha1.CurReportDefinitionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CurReportDefinition, err error)
	CurReportDefinitionExpansion
}

// curReportDefinitions implements CurReportDefinitionInterface
type curReportDefinitions struct {
	client rest.Interface
	ns     string
}

// newCurReportDefinitions returns a CurReportDefinitions
func newCurReportDefinitions(c *AwsV1alpha1Client, namespace string) *curReportDefinitions {
	return &curReportDefinitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the curReportDefinition, and returns the corresponding curReportDefinition object, and an error if there is any.
func (c *curReportDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.CurReportDefinition, err error) {
	result = &v1alpha1.CurReportDefinition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CurReportDefinitions that match those selectors.
func (c *curReportDefinitions) List(opts v1.ListOptions) (result *v1alpha1.CurReportDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CurReportDefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested curReportDefinitions.
func (c *curReportDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a curReportDefinition and creates it.  Returns the server's representation of the curReportDefinition, and an error, if there is any.
func (c *curReportDefinitions) Create(curReportDefinition *v1alpha1.CurReportDefinition) (result *v1alpha1.CurReportDefinition, err error) {
	result = &v1alpha1.CurReportDefinition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		Body(curReportDefinition).
		Do().
		Into(result)
	return
}

// Update takes the representation of a curReportDefinition and updates it. Returns the server's representation of the curReportDefinition, and an error, if there is any.
func (c *curReportDefinitions) Update(curReportDefinition *v1alpha1.CurReportDefinition) (result *v1alpha1.CurReportDefinition, err error) {
	result = &v1alpha1.CurReportDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		Name(curReportDefinition.Name).
		Body(curReportDefinition).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *curReportDefinitions) UpdateStatus(curReportDefinition *v1alpha1.CurReportDefinition) (result *v1alpha1.CurReportDefinition, err error) {
	result = &v1alpha1.CurReportDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		Name(curReportDefinition.Name).
		SubResource("status").
		Body(curReportDefinition).
		Do().
		Into(result)
	return
}

// Delete takes name of the curReportDefinition and deletes it. Returns an error if one occurs.
func (c *curReportDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *curReportDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("curreportdefinitions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched curReportDefinition.
func (c *curReportDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CurReportDefinition, err error) {
	result = &v1alpha1.CurReportDefinition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("curreportdefinitions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
