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

// WafregionalIpsetsGetter has a method to return a WafregionalIpsetInterface.
// A group's client should implement this interface.
type WafregionalIpsetsGetter interface {
	WafregionalIpsets(namespace string) WafregionalIpsetInterface
}

// WafregionalIpsetInterface has methods to work with WafregionalIpset resources.
type WafregionalIpsetInterface interface {
	Create(*v1alpha1.WafregionalIpset) (*v1alpha1.WafregionalIpset, error)
	Update(*v1alpha1.WafregionalIpset) (*v1alpha1.WafregionalIpset, error)
	UpdateStatus(*v1alpha1.WafregionalIpset) (*v1alpha1.WafregionalIpset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafregionalIpset, error)
	List(opts v1.ListOptions) (*v1alpha1.WafregionalIpsetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalIpset, err error)
	WafregionalIpsetExpansion
}

// wafregionalIpsets implements WafregionalIpsetInterface
type wafregionalIpsets struct {
	client rest.Interface
	ns     string
}

// newWafregionalIpsets returns a WafregionalIpsets
func newWafregionalIpsets(c *AwsV1alpha1Client, namespace string) *wafregionalIpsets {
	return &wafregionalIpsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafregionalIpset, and returns the corresponding wafregionalIpset object, and an error if there is any.
func (c *wafregionalIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalIpset, err error) {
	result = &v1alpha1.WafregionalIpset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafregionalIpsets that match those selectors.
func (c *wafregionalIpsets) List(opts v1.ListOptions) (result *v1alpha1.WafregionalIpsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafregionalIpsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafregionalIpsets.
func (c *wafregionalIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafregionalIpset and creates it.  Returns the server's representation of the wafregionalIpset, and an error, if there is any.
func (c *wafregionalIpsets) Create(wafregionalIpset *v1alpha1.WafregionalIpset) (result *v1alpha1.WafregionalIpset, err error) {
	result = &v1alpha1.WafregionalIpset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		Body(wafregionalIpset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafregionalIpset and updates it. Returns the server's representation of the wafregionalIpset, and an error, if there is any.
func (c *wafregionalIpsets) Update(wafregionalIpset *v1alpha1.WafregionalIpset) (result *v1alpha1.WafregionalIpset, err error) {
	result = &v1alpha1.WafregionalIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		Name(wafregionalIpset.Name).
		Body(wafregionalIpset).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafregionalIpsets) UpdateStatus(wafregionalIpset *v1alpha1.WafregionalIpset) (result *v1alpha1.WafregionalIpset, err error) {
	result = &v1alpha1.WafregionalIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		Name(wafregionalIpset.Name).
		SubResource("status").
		Body(wafregionalIpset).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafregionalIpset and deletes it. Returns an error if one occurs.
func (c *wafregionalIpsets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafregionalIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalipsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafregionalIpset.
func (c *wafregionalIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalIpset, err error) {
	result = &v1alpha1.WafregionalIpset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafregionalipsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
