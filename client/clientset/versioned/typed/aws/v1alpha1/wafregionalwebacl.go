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

// WafregionalWebACLsGetter has a method to return a WafregionalWebACLInterface.
// A group's client should implement this interface.
type WafregionalWebACLsGetter interface {
	WafregionalWebACLs(namespace string) WafregionalWebACLInterface
}

// WafregionalWebACLInterface has methods to work with WafregionalWebACL resources.
type WafregionalWebACLInterface interface {
	Create(*v1alpha1.WafregionalWebACL) (*v1alpha1.WafregionalWebACL, error)
	Update(*v1alpha1.WafregionalWebACL) (*v1alpha1.WafregionalWebACL, error)
	UpdateStatus(*v1alpha1.WafregionalWebACL) (*v1alpha1.WafregionalWebACL, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafregionalWebACL, error)
	List(opts v1.ListOptions) (*v1alpha1.WafregionalWebACLList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalWebACL, err error)
	WafregionalWebACLExpansion
}

// wafregionalWebACLs implements WafregionalWebACLInterface
type wafregionalWebACLs struct {
	client rest.Interface
	ns     string
}

// newWafregionalWebACLs returns a WafregionalWebACLs
func newWafregionalWebACLs(c *AwsV1alpha1Client, namespace string) *wafregionalWebACLs {
	return &wafregionalWebACLs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafregionalWebACL, and returns the corresponding wafregionalWebACL object, and an error if there is any.
func (c *wafregionalWebACLs) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalWebACL, err error) {
	result = &v1alpha1.WafregionalWebACL{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafregionalWebACLs that match those selectors.
func (c *wafregionalWebACLs) List(opts v1.ListOptions) (result *v1alpha1.WafregionalWebACLList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafregionalWebACLList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafregionalWebACLs.
func (c *wafregionalWebACLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafregionalWebACL and creates it.  Returns the server's representation of the wafregionalWebACL, and an error, if there is any.
func (c *wafregionalWebACLs) Create(wafregionalWebACL *v1alpha1.WafregionalWebACL) (result *v1alpha1.WafregionalWebACL, err error) {
	result = &v1alpha1.WafregionalWebACL{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		Body(wafregionalWebACL).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafregionalWebACL and updates it. Returns the server's representation of the wafregionalWebACL, and an error, if there is any.
func (c *wafregionalWebACLs) Update(wafregionalWebACL *v1alpha1.WafregionalWebACL) (result *v1alpha1.WafregionalWebACL, err error) {
	result = &v1alpha1.WafregionalWebACL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		Name(wafregionalWebACL.Name).
		Body(wafregionalWebACL).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafregionalWebACLs) UpdateStatus(wafregionalWebACL *v1alpha1.WafregionalWebACL) (result *v1alpha1.WafregionalWebACL, err error) {
	result = &v1alpha1.WafregionalWebACL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		Name(wafregionalWebACL.Name).
		SubResource("status").
		Body(wafregionalWebACL).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafregionalWebACL and deletes it. Returns an error if one occurs.
func (c *wafregionalWebACLs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafregionalWebACLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafregionalWebACL.
func (c *wafregionalWebACLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalWebACL, err error) {
	result = &v1alpha1.WafregionalWebACL{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafregionalwebacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
