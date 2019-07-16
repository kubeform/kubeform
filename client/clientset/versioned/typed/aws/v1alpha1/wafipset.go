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

// WafIpsetsGetter has a method to return a WafIpsetInterface.
// A group's client should implement this interface.
type WafIpsetsGetter interface {
	WafIpsets() WafIpsetInterface
}

// WafIpsetInterface has methods to work with WafIpset resources.
type WafIpsetInterface interface {
	Create(*v1alpha1.WafIpset) (*v1alpha1.WafIpset, error)
	Update(*v1alpha1.WafIpset) (*v1alpha1.WafIpset, error)
	UpdateStatus(*v1alpha1.WafIpset) (*v1alpha1.WafIpset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafIpset, error)
	List(opts v1.ListOptions) (*v1alpha1.WafIpsetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafIpset, err error)
	WafIpsetExpansion
}

// wafIpsets implements WafIpsetInterface
type wafIpsets struct {
	client rest.Interface
}

// newWafIpsets returns a WafIpsets
func newWafIpsets(c *AwsV1alpha1Client) *wafIpsets {
	return &wafIpsets{
		client: c.RESTClient(),
	}
}

// Get takes name of the wafIpset, and returns the corresponding wafIpset object, and an error if there is any.
func (c *wafIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Get().
		Resource("wafipsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafIpsets that match those selectors.
func (c *wafIpsets) List(opts v1.ListOptions) (result *v1alpha1.WafIpsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafIpsetList{}
	err = c.client.Get().
		Resource("wafipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafIpsets.
func (c *wafIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("wafipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafIpset and creates it.  Returns the server's representation of the wafIpset, and an error, if there is any.
func (c *wafIpsets) Create(wafIpset *v1alpha1.WafIpset) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Post().
		Resource("wafipsets").
		Body(wafIpset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafIpset and updates it. Returns the server's representation of the wafIpset, and an error, if there is any.
func (c *wafIpsets) Update(wafIpset *v1alpha1.WafIpset) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Put().
		Resource("wafipsets").
		Name(wafIpset.Name).
		Body(wafIpset).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafIpsets) UpdateStatus(wafIpset *v1alpha1.WafIpset) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Put().
		Resource("wafipsets").
		Name(wafIpset.Name).
		SubResource("status").
		Body(wafIpset).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafIpset and deletes it. Returns an error if one occurs.
func (c *wafIpsets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("wafipsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("wafipsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafIpset.
func (c *wafIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafIpset, err error) {
	result = &v1alpha1.WafIpset{}
	err = c.client.Patch(pt).
		Resource("wafipsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
