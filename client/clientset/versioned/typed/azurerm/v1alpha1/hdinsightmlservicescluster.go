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

// HdinsightMlServicesClustersGetter has a method to return a HdinsightMlServicesClusterInterface.
// A group's client should implement this interface.
type HdinsightMlServicesClustersGetter interface {
	HdinsightMlServicesClusters(namespace string) HdinsightMlServicesClusterInterface
}

// HdinsightMlServicesClusterInterface has methods to work with HdinsightMlServicesCluster resources.
type HdinsightMlServicesClusterInterface interface {
	Create(*v1alpha1.HdinsightMlServicesCluster) (*v1alpha1.HdinsightMlServicesCluster, error)
	Update(*v1alpha1.HdinsightMlServicesCluster) (*v1alpha1.HdinsightMlServicesCluster, error)
	UpdateStatus(*v1alpha1.HdinsightMlServicesCluster) (*v1alpha1.HdinsightMlServicesCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HdinsightMlServicesCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.HdinsightMlServicesClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HdinsightMlServicesCluster, err error)
	HdinsightMlServicesClusterExpansion
}

// hdinsightMlServicesClusters implements HdinsightMlServicesClusterInterface
type hdinsightMlServicesClusters struct {
	client rest.Interface
	ns     string
}

// newHdinsightMlServicesClusters returns a HdinsightMlServicesClusters
func newHdinsightMlServicesClusters(c *AzurermV1alpha1Client, namespace string) *hdinsightMlServicesClusters {
	return &hdinsightMlServicesClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hdinsightMlServicesCluster, and returns the corresponding hdinsightMlServicesCluster object, and an error if there is any.
func (c *hdinsightMlServicesClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	result = &v1alpha1.HdinsightMlServicesCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HdinsightMlServicesClusters that match those selectors.
func (c *hdinsightMlServicesClusters) List(opts v1.ListOptions) (result *v1alpha1.HdinsightMlServicesClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HdinsightMlServicesClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hdinsightMlServicesClusters.
func (c *hdinsightMlServicesClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a hdinsightMlServicesCluster and creates it.  Returns the server's representation of the hdinsightMlServicesCluster, and an error, if there is any.
func (c *hdinsightMlServicesClusters) Create(hdinsightMlServicesCluster *v1alpha1.HdinsightMlServicesCluster) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	result = &v1alpha1.HdinsightMlServicesCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		Body(hdinsightMlServicesCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hdinsightMlServicesCluster and updates it. Returns the server's representation of the hdinsightMlServicesCluster, and an error, if there is any.
func (c *hdinsightMlServicesClusters) Update(hdinsightMlServicesCluster *v1alpha1.HdinsightMlServicesCluster) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	result = &v1alpha1.HdinsightMlServicesCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		Name(hdinsightMlServicesCluster.Name).
		Body(hdinsightMlServicesCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *hdinsightMlServicesClusters) UpdateStatus(hdinsightMlServicesCluster *v1alpha1.HdinsightMlServicesCluster) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	result = &v1alpha1.HdinsightMlServicesCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		Name(hdinsightMlServicesCluster.Name).
		SubResource("status").
		Body(hdinsightMlServicesCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the hdinsightMlServicesCluster and deletes it. Returns an error if one occurs.
func (c *hdinsightMlServicesClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hdinsightMlServicesClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hdinsightMlServicesCluster.
func (c *hdinsightMlServicesClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HdinsightMlServicesCluster, err error) {
	result = &v1alpha1.HdinsightMlServicesCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hdinsightmlservicesclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
