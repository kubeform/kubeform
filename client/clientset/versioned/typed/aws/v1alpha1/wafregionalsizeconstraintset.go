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

// WafregionalSizeConstraintSetsGetter has a method to return a WafregionalSizeConstraintSetInterface.
// A group's client should implement this interface.
type WafregionalSizeConstraintSetsGetter interface {
	WafregionalSizeConstraintSets(namespace string) WafregionalSizeConstraintSetInterface
}

// WafregionalSizeConstraintSetInterface has methods to work with WafregionalSizeConstraintSet resources.
type WafregionalSizeConstraintSetInterface interface {
	Create(*v1alpha1.WafregionalSizeConstraintSet) (*v1alpha1.WafregionalSizeConstraintSet, error)
	Update(*v1alpha1.WafregionalSizeConstraintSet) (*v1alpha1.WafregionalSizeConstraintSet, error)
	UpdateStatus(*v1alpha1.WafregionalSizeConstraintSet) (*v1alpha1.WafregionalSizeConstraintSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafregionalSizeConstraintSet, error)
	List(opts v1.ListOptions) (*v1alpha1.WafregionalSizeConstraintSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalSizeConstraintSet, err error)
	WafregionalSizeConstraintSetExpansion
}

// wafregionalSizeConstraintSets implements WafregionalSizeConstraintSetInterface
type wafregionalSizeConstraintSets struct {
	client rest.Interface
	ns     string
}

// newWafregionalSizeConstraintSets returns a WafregionalSizeConstraintSets
func newWafregionalSizeConstraintSets(c *AwsV1alpha1Client, namespace string) *wafregionalSizeConstraintSets {
	return &wafregionalSizeConstraintSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafregionalSizeConstraintSet, and returns the corresponding wafregionalSizeConstraintSet object, and an error if there is any.
func (c *wafregionalSizeConstraintSets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalSizeConstraintSet, err error) {
	result = &v1alpha1.WafregionalSizeConstraintSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafregionalSizeConstraintSets that match those selectors.
func (c *wafregionalSizeConstraintSets) List(opts v1.ListOptions) (result *v1alpha1.WafregionalSizeConstraintSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafregionalSizeConstraintSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafregionalSizeConstraintSets.
func (c *wafregionalSizeConstraintSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafregionalSizeConstraintSet and creates it.  Returns the server's representation of the wafregionalSizeConstraintSet, and an error, if there is any.
func (c *wafregionalSizeConstraintSets) Create(wafregionalSizeConstraintSet *v1alpha1.WafregionalSizeConstraintSet) (result *v1alpha1.WafregionalSizeConstraintSet, err error) {
	result = &v1alpha1.WafregionalSizeConstraintSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		Body(wafregionalSizeConstraintSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafregionalSizeConstraintSet and updates it. Returns the server's representation of the wafregionalSizeConstraintSet, and an error, if there is any.
func (c *wafregionalSizeConstraintSets) Update(wafregionalSizeConstraintSet *v1alpha1.WafregionalSizeConstraintSet) (result *v1alpha1.WafregionalSizeConstraintSet, err error) {
	result = &v1alpha1.WafregionalSizeConstraintSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		Name(wafregionalSizeConstraintSet.Name).
		Body(wafregionalSizeConstraintSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafregionalSizeConstraintSets) UpdateStatus(wafregionalSizeConstraintSet *v1alpha1.WafregionalSizeConstraintSet) (result *v1alpha1.WafregionalSizeConstraintSet, err error) {
	result = &v1alpha1.WafregionalSizeConstraintSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		Name(wafregionalSizeConstraintSet.Name).
		SubResource("status").
		Body(wafregionalSizeConstraintSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafregionalSizeConstraintSet and deletes it. Returns an error if one occurs.
func (c *wafregionalSizeConstraintSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafregionalSizeConstraintSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafregionalSizeConstraintSet.
func (c *wafregionalSizeConstraintSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalSizeConstraintSet, err error) {
	result = &v1alpha1.WafregionalSizeConstraintSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafregionalsizeconstraintsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
