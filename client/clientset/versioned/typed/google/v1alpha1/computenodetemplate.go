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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeNodeTemplatesGetter has a method to return a ComputeNodeTemplateInterface.
// A group's client should implement this interface.
type ComputeNodeTemplatesGetter interface {
	ComputeNodeTemplates(namespace string) ComputeNodeTemplateInterface
}

// ComputeNodeTemplateInterface has methods to work with ComputeNodeTemplate resources.
type ComputeNodeTemplateInterface interface {
	Create(*v1alpha1.ComputeNodeTemplate) (*v1alpha1.ComputeNodeTemplate, error)
	Update(*v1alpha1.ComputeNodeTemplate) (*v1alpha1.ComputeNodeTemplate, error)
	UpdateStatus(*v1alpha1.ComputeNodeTemplate) (*v1alpha1.ComputeNodeTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeNodeTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeNodeTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeNodeTemplate, err error)
	ComputeNodeTemplateExpansion
}

// computeNodeTemplates implements ComputeNodeTemplateInterface
type computeNodeTemplates struct {
	client rest.Interface
	ns     string
}

// newComputeNodeTemplates returns a ComputeNodeTemplates
func newComputeNodeTemplates(c *GoogleV1alpha1Client, namespace string) *computeNodeTemplates {
	return &computeNodeTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeNodeTemplate, and returns the corresponding computeNodeTemplate object, and an error if there is any.
func (c *computeNodeTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeNodeTemplate, err error) {
	result = &v1alpha1.ComputeNodeTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computenodetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeNodeTemplates that match those selectors.
func (c *computeNodeTemplates) List(opts v1.ListOptions) (result *v1alpha1.ComputeNodeTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeNodeTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computenodetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeNodeTemplates.
func (c *computeNodeTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computenodetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeNodeTemplate and creates it.  Returns the server's representation of the computeNodeTemplate, and an error, if there is any.
func (c *computeNodeTemplates) Create(computeNodeTemplate *v1alpha1.ComputeNodeTemplate) (result *v1alpha1.ComputeNodeTemplate, err error) {
	result = &v1alpha1.ComputeNodeTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computenodetemplates").
		Body(computeNodeTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeNodeTemplate and updates it. Returns the server's representation of the computeNodeTemplate, and an error, if there is any.
func (c *computeNodeTemplates) Update(computeNodeTemplate *v1alpha1.ComputeNodeTemplate) (result *v1alpha1.ComputeNodeTemplate, err error) {
	result = &v1alpha1.ComputeNodeTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computenodetemplates").
		Name(computeNodeTemplate.Name).
		Body(computeNodeTemplate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeNodeTemplates) UpdateStatus(computeNodeTemplate *v1alpha1.ComputeNodeTemplate) (result *v1alpha1.ComputeNodeTemplate, err error) {
	result = &v1alpha1.ComputeNodeTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computenodetemplates").
		Name(computeNodeTemplate.Name).
		SubResource("status").
		Body(computeNodeTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeNodeTemplate and deletes it. Returns an error if one occurs.
func (c *computeNodeTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computenodetemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeNodeTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computenodetemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeNodeTemplate.
func (c *computeNodeTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeNodeTemplate, err error) {
	result = &v1alpha1.ComputeNodeTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computenodetemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
