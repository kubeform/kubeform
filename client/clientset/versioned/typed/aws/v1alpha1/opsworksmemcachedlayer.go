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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OpsworksMemcachedLayersGetter has a method to return a OpsworksMemcachedLayerInterface.
// A group's client should implement this interface.
type OpsworksMemcachedLayersGetter interface {
	OpsworksMemcachedLayers(namespace string) OpsworksMemcachedLayerInterface
}

// OpsworksMemcachedLayerInterface has methods to work with OpsworksMemcachedLayer resources.
type OpsworksMemcachedLayerInterface interface {
	Create(*v1alpha1.OpsworksMemcachedLayer) (*v1alpha1.OpsworksMemcachedLayer, error)
	Update(*v1alpha1.OpsworksMemcachedLayer) (*v1alpha1.OpsworksMemcachedLayer, error)
	UpdateStatus(*v1alpha1.OpsworksMemcachedLayer) (*v1alpha1.OpsworksMemcachedLayer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpsworksMemcachedLayer, error)
	List(opts v1.ListOptions) (*v1alpha1.OpsworksMemcachedLayerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksMemcachedLayer, err error)
	OpsworksMemcachedLayerExpansion
}

// opsworksMemcachedLayers implements OpsworksMemcachedLayerInterface
type opsworksMemcachedLayers struct {
	client rest.Interface
	ns     string
}

// newOpsworksMemcachedLayers returns a OpsworksMemcachedLayers
func newOpsworksMemcachedLayers(c *AwsV1alpha1Client, namespace string) *opsworksMemcachedLayers {
	return &opsworksMemcachedLayers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the opsworksMemcachedLayer, and returns the corresponding opsworksMemcachedLayer object, and an error if there is any.
func (c *opsworksMemcachedLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksMemcachedLayer, err error) {
	result = &v1alpha1.OpsworksMemcachedLayer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpsworksMemcachedLayers that match those selectors.
func (c *opsworksMemcachedLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksMemcachedLayerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpsworksMemcachedLayerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested opsworksMemcachedLayers.
func (c *opsworksMemcachedLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a opsworksMemcachedLayer and creates it.  Returns the server's representation of the opsworksMemcachedLayer, and an error, if there is any.
func (c *opsworksMemcachedLayers) Create(opsworksMemcachedLayer *v1alpha1.OpsworksMemcachedLayer) (result *v1alpha1.OpsworksMemcachedLayer, err error) {
	result = &v1alpha1.OpsworksMemcachedLayer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		Body(opsworksMemcachedLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a opsworksMemcachedLayer and updates it. Returns the server's representation of the opsworksMemcachedLayer, and an error, if there is any.
func (c *opsworksMemcachedLayers) Update(opsworksMemcachedLayer *v1alpha1.OpsworksMemcachedLayer) (result *v1alpha1.OpsworksMemcachedLayer, err error) {
	result = &v1alpha1.OpsworksMemcachedLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		Name(opsworksMemcachedLayer.Name).
		Body(opsworksMemcachedLayer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *opsworksMemcachedLayers) UpdateStatus(opsworksMemcachedLayer *v1alpha1.OpsworksMemcachedLayer) (result *v1alpha1.OpsworksMemcachedLayer, err error) {
	result = &v1alpha1.OpsworksMemcachedLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		Name(opsworksMemcachedLayer.Name).
		SubResource("status").
		Body(opsworksMemcachedLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the opsworksMemcachedLayer and deletes it. Returns an error if one occurs.
func (c *opsworksMemcachedLayers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *opsworksMemcachedLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched opsworksMemcachedLayer.
func (c *opsworksMemcachedLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksMemcachedLayer, err error) {
	result = &v1alpha1.OpsworksMemcachedLayer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("opsworksmemcachedlayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
