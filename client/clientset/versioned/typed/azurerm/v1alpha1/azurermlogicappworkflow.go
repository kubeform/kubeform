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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermLogicAppWorkflowsGetter has a method to return a AzurermLogicAppWorkflowInterface.
// A group's client should implement this interface.
type AzurermLogicAppWorkflowsGetter interface {
	AzurermLogicAppWorkflows() AzurermLogicAppWorkflowInterface
}

// AzurermLogicAppWorkflowInterface has methods to work with AzurermLogicAppWorkflow resources.
type AzurermLogicAppWorkflowInterface interface {
	Create(*v1alpha1.AzurermLogicAppWorkflow) (*v1alpha1.AzurermLogicAppWorkflow, error)
	Update(*v1alpha1.AzurermLogicAppWorkflow) (*v1alpha1.AzurermLogicAppWorkflow, error)
	UpdateStatus(*v1alpha1.AzurermLogicAppWorkflow) (*v1alpha1.AzurermLogicAppWorkflow, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermLogicAppWorkflow, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermLogicAppWorkflowList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppWorkflow, err error)
	AzurermLogicAppWorkflowExpansion
}

// azurermLogicAppWorkflows implements AzurermLogicAppWorkflowInterface
type azurermLogicAppWorkflows struct {
	client rest.Interface
}

// newAzurermLogicAppWorkflows returns a AzurermLogicAppWorkflows
func newAzurermLogicAppWorkflows(c *AzurermV1alpha1Client) *azurermLogicAppWorkflows {
	return &azurermLogicAppWorkflows{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermLogicAppWorkflow, and returns the corresponding azurermLogicAppWorkflow object, and an error if there is any.
func (c *azurermLogicAppWorkflows) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLogicAppWorkflow, err error) {
	result = &v1alpha1.AzurermLogicAppWorkflow{}
	err = c.client.Get().
		Resource("azurermlogicappworkflows").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermLogicAppWorkflows that match those selectors.
func (c *azurermLogicAppWorkflows) List(opts v1.ListOptions) (result *v1alpha1.AzurermLogicAppWorkflowList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermLogicAppWorkflowList{}
	err = c.client.Get().
		Resource("azurermlogicappworkflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermLogicAppWorkflows.
func (c *azurermLogicAppWorkflows) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermlogicappworkflows").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermLogicAppWorkflow and creates it.  Returns the server's representation of the azurermLogicAppWorkflow, and an error, if there is any.
func (c *azurermLogicAppWorkflows) Create(azurermLogicAppWorkflow *v1alpha1.AzurermLogicAppWorkflow) (result *v1alpha1.AzurermLogicAppWorkflow, err error) {
	result = &v1alpha1.AzurermLogicAppWorkflow{}
	err = c.client.Post().
		Resource("azurermlogicappworkflows").
		Body(azurermLogicAppWorkflow).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermLogicAppWorkflow and updates it. Returns the server's representation of the azurermLogicAppWorkflow, and an error, if there is any.
func (c *azurermLogicAppWorkflows) Update(azurermLogicAppWorkflow *v1alpha1.AzurermLogicAppWorkflow) (result *v1alpha1.AzurermLogicAppWorkflow, err error) {
	result = &v1alpha1.AzurermLogicAppWorkflow{}
	err = c.client.Put().
		Resource("azurermlogicappworkflows").
		Name(azurermLogicAppWorkflow.Name).
		Body(azurermLogicAppWorkflow).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermLogicAppWorkflows) UpdateStatus(azurermLogicAppWorkflow *v1alpha1.AzurermLogicAppWorkflow) (result *v1alpha1.AzurermLogicAppWorkflow, err error) {
	result = &v1alpha1.AzurermLogicAppWorkflow{}
	err = c.client.Put().
		Resource("azurermlogicappworkflows").
		Name(azurermLogicAppWorkflow.Name).
		SubResource("status").
		Body(azurermLogicAppWorkflow).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermLogicAppWorkflow and deletes it. Returns an error if one occurs.
func (c *azurermLogicAppWorkflows) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermlogicappworkflows").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermLogicAppWorkflows) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermlogicappworkflows").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermLogicAppWorkflow.
func (c *azurermLogicAppWorkflows) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppWorkflow, err error) {
	result = &v1alpha1.AzurermLogicAppWorkflow{}
	err = c.client.Patch(pt).
		Resource("azurermlogicappworkflows").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
