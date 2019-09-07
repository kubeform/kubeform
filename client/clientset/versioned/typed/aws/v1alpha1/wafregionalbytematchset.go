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

// WafregionalByteMatchSetsGetter has a method to return a WafregionalByteMatchSetInterface.
// A group's client should implement this interface.
type WafregionalByteMatchSetsGetter interface {
	WafregionalByteMatchSets(namespace string) WafregionalByteMatchSetInterface
}

// WafregionalByteMatchSetInterface has methods to work with WafregionalByteMatchSet resources.
type WafregionalByteMatchSetInterface interface {
	Create(*v1alpha1.WafregionalByteMatchSet) (*v1alpha1.WafregionalByteMatchSet, error)
	Update(*v1alpha1.WafregionalByteMatchSet) (*v1alpha1.WafregionalByteMatchSet, error)
	UpdateStatus(*v1alpha1.WafregionalByteMatchSet) (*v1alpha1.WafregionalByteMatchSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafregionalByteMatchSet, error)
	List(opts v1.ListOptions) (*v1alpha1.WafregionalByteMatchSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalByteMatchSet, err error)
	WafregionalByteMatchSetExpansion
}

// wafregionalByteMatchSets implements WafregionalByteMatchSetInterface
type wafregionalByteMatchSets struct {
	client rest.Interface
	ns     string
}

// newWafregionalByteMatchSets returns a WafregionalByteMatchSets
func newWafregionalByteMatchSets(c *AwsV1alpha1Client, namespace string) *wafregionalByteMatchSets {
	return &wafregionalByteMatchSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafregionalByteMatchSet, and returns the corresponding wafregionalByteMatchSet object, and an error if there is any.
func (c *wafregionalByteMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalByteMatchSet, err error) {
	result = &v1alpha1.WafregionalByteMatchSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafregionalByteMatchSets that match those selectors.
func (c *wafregionalByteMatchSets) List(opts v1.ListOptions) (result *v1alpha1.WafregionalByteMatchSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafregionalByteMatchSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafregionalByteMatchSets.
func (c *wafregionalByteMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafregionalByteMatchSet and creates it.  Returns the server's representation of the wafregionalByteMatchSet, and an error, if there is any.
func (c *wafregionalByteMatchSets) Create(wafregionalByteMatchSet *v1alpha1.WafregionalByteMatchSet) (result *v1alpha1.WafregionalByteMatchSet, err error) {
	result = &v1alpha1.WafregionalByteMatchSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		Body(wafregionalByteMatchSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafregionalByteMatchSet and updates it. Returns the server's representation of the wafregionalByteMatchSet, and an error, if there is any.
func (c *wafregionalByteMatchSets) Update(wafregionalByteMatchSet *v1alpha1.WafregionalByteMatchSet) (result *v1alpha1.WafregionalByteMatchSet, err error) {
	result = &v1alpha1.WafregionalByteMatchSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		Name(wafregionalByteMatchSet.Name).
		Body(wafregionalByteMatchSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafregionalByteMatchSets) UpdateStatus(wafregionalByteMatchSet *v1alpha1.WafregionalByteMatchSet) (result *v1alpha1.WafregionalByteMatchSet, err error) {
	result = &v1alpha1.WafregionalByteMatchSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		Name(wafregionalByteMatchSet.Name).
		SubResource("status").
		Body(wafregionalByteMatchSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafregionalByteMatchSet and deletes it. Returns an error if one occurs.
func (c *wafregionalByteMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafregionalByteMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafregionalByteMatchSet.
func (c *wafregionalByteMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalByteMatchSet, err error) {
	result = &v1alpha1.WafregionalByteMatchSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafregionalbytematchsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}