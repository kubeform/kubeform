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

// AzurermHdinsightHbaseClustersGetter has a method to return a AzurermHdinsightHbaseClusterInterface.
// A group's client should implement this interface.
type AzurermHdinsightHbaseClustersGetter interface {
	AzurermHdinsightHbaseClusters() AzurermHdinsightHbaseClusterInterface
}

// AzurermHdinsightHbaseClusterInterface has methods to work with AzurermHdinsightHbaseCluster resources.
type AzurermHdinsightHbaseClusterInterface interface {
	Create(*v1alpha1.AzurermHdinsightHbaseCluster) (*v1alpha1.AzurermHdinsightHbaseCluster, error)
	Update(*v1alpha1.AzurermHdinsightHbaseCluster) (*v1alpha1.AzurermHdinsightHbaseCluster, error)
	UpdateStatus(*v1alpha1.AzurermHdinsightHbaseCluster) (*v1alpha1.AzurermHdinsightHbaseCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermHdinsightHbaseCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermHdinsightHbaseClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermHdinsightHbaseCluster, err error)
	AzurermHdinsightHbaseClusterExpansion
}

// azurermHdinsightHbaseClusters implements AzurermHdinsightHbaseClusterInterface
type azurermHdinsightHbaseClusters struct {
	client rest.Interface
}

// newAzurermHdinsightHbaseClusters returns a AzurermHdinsightHbaseClusters
func newAzurermHdinsightHbaseClusters(c *AzurermV1alpha1Client) *azurermHdinsightHbaseClusters {
	return &azurermHdinsightHbaseClusters{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermHdinsightHbaseCluster, and returns the corresponding azurermHdinsightHbaseCluster object, and an error if there is any.
func (c *azurermHdinsightHbaseClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermHdinsightHbaseCluster, err error) {
	result = &v1alpha1.AzurermHdinsightHbaseCluster{}
	err = c.client.Get().
		Resource("azurermhdinsighthbaseclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermHdinsightHbaseClusters that match those selectors.
func (c *azurermHdinsightHbaseClusters) List(opts v1.ListOptions) (result *v1alpha1.AzurermHdinsightHbaseClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermHdinsightHbaseClusterList{}
	err = c.client.Get().
		Resource("azurermhdinsighthbaseclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermHdinsightHbaseClusters.
func (c *azurermHdinsightHbaseClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermhdinsighthbaseclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermHdinsightHbaseCluster and creates it.  Returns the server's representation of the azurermHdinsightHbaseCluster, and an error, if there is any.
func (c *azurermHdinsightHbaseClusters) Create(azurermHdinsightHbaseCluster *v1alpha1.AzurermHdinsightHbaseCluster) (result *v1alpha1.AzurermHdinsightHbaseCluster, err error) {
	result = &v1alpha1.AzurermHdinsightHbaseCluster{}
	err = c.client.Post().
		Resource("azurermhdinsighthbaseclusters").
		Body(azurermHdinsightHbaseCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermHdinsightHbaseCluster and updates it. Returns the server's representation of the azurermHdinsightHbaseCluster, and an error, if there is any.
func (c *azurermHdinsightHbaseClusters) Update(azurermHdinsightHbaseCluster *v1alpha1.AzurermHdinsightHbaseCluster) (result *v1alpha1.AzurermHdinsightHbaseCluster, err error) {
	result = &v1alpha1.AzurermHdinsightHbaseCluster{}
	err = c.client.Put().
		Resource("azurermhdinsighthbaseclusters").
		Name(azurermHdinsightHbaseCluster.Name).
		Body(azurermHdinsightHbaseCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermHdinsightHbaseClusters) UpdateStatus(azurermHdinsightHbaseCluster *v1alpha1.AzurermHdinsightHbaseCluster) (result *v1alpha1.AzurermHdinsightHbaseCluster, err error) {
	result = &v1alpha1.AzurermHdinsightHbaseCluster{}
	err = c.client.Put().
		Resource("azurermhdinsighthbaseclusters").
		Name(azurermHdinsightHbaseCluster.Name).
		SubResource("status").
		Body(azurermHdinsightHbaseCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermHdinsightHbaseCluster and deletes it. Returns an error if one occurs.
func (c *azurermHdinsightHbaseClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermhdinsighthbaseclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermHdinsightHbaseClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermhdinsighthbaseclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermHdinsightHbaseCluster.
func (c *azurermHdinsightHbaseClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermHdinsightHbaseCluster, err error) {
	result = &v1alpha1.AzurermHdinsightHbaseCluster{}
	err = c.client.Patch(pt).
		Resource("azurermhdinsighthbaseclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
