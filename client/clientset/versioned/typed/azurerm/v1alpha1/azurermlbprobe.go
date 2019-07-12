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

// AzurermLbProbesGetter has a method to return a AzurermLbProbeInterface.
// A group's client should implement this interface.
type AzurermLbProbesGetter interface {
	AzurermLbProbes() AzurermLbProbeInterface
}

// AzurermLbProbeInterface has methods to work with AzurermLbProbe resources.
type AzurermLbProbeInterface interface {
	Create(*v1alpha1.AzurermLbProbe) (*v1alpha1.AzurermLbProbe, error)
	Update(*v1alpha1.AzurermLbProbe) (*v1alpha1.AzurermLbProbe, error)
	UpdateStatus(*v1alpha1.AzurermLbProbe) (*v1alpha1.AzurermLbProbe, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermLbProbe, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermLbProbeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLbProbe, err error)
	AzurermLbProbeExpansion
}

// azurermLbProbes implements AzurermLbProbeInterface
type azurermLbProbes struct {
	client rest.Interface
}

// newAzurermLbProbes returns a AzurermLbProbes
func newAzurermLbProbes(c *AzurermV1alpha1Client) *azurermLbProbes {
	return &azurermLbProbes{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermLbProbe, and returns the corresponding azurermLbProbe object, and an error if there is any.
func (c *azurermLbProbes) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLbProbe, err error) {
	result = &v1alpha1.AzurermLbProbe{}
	err = c.client.Get().
		Resource("azurermlbprobes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermLbProbes that match those selectors.
func (c *azurermLbProbes) List(opts v1.ListOptions) (result *v1alpha1.AzurermLbProbeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermLbProbeList{}
	err = c.client.Get().
		Resource("azurermlbprobes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermLbProbes.
func (c *azurermLbProbes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermlbprobes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermLbProbe and creates it.  Returns the server's representation of the azurermLbProbe, and an error, if there is any.
func (c *azurermLbProbes) Create(azurermLbProbe *v1alpha1.AzurermLbProbe) (result *v1alpha1.AzurermLbProbe, err error) {
	result = &v1alpha1.AzurermLbProbe{}
	err = c.client.Post().
		Resource("azurermlbprobes").
		Body(azurermLbProbe).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermLbProbe and updates it. Returns the server's representation of the azurermLbProbe, and an error, if there is any.
func (c *azurermLbProbes) Update(azurermLbProbe *v1alpha1.AzurermLbProbe) (result *v1alpha1.AzurermLbProbe, err error) {
	result = &v1alpha1.AzurermLbProbe{}
	err = c.client.Put().
		Resource("azurermlbprobes").
		Name(azurermLbProbe.Name).
		Body(azurermLbProbe).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermLbProbes) UpdateStatus(azurermLbProbe *v1alpha1.AzurermLbProbe) (result *v1alpha1.AzurermLbProbe, err error) {
	result = &v1alpha1.AzurermLbProbe{}
	err = c.client.Put().
		Resource("azurermlbprobes").
		Name(azurermLbProbe.Name).
		SubResource("status").
		Body(azurermLbProbe).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermLbProbe and deletes it. Returns an error if one occurs.
func (c *azurermLbProbes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermlbprobes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermLbProbes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermlbprobes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermLbProbe.
func (c *azurermLbProbes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLbProbe, err error) {
	result = &v1alpha1.AzurermLbProbe{}
	err = c.client.Patch(pt).
		Resource("azurermlbprobes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
