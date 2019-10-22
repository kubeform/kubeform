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

// SesReceiptRuleSetsGetter has a method to return a SesReceiptRuleSetInterface.
// A group's client should implement this interface.
type SesReceiptRuleSetsGetter interface {
	SesReceiptRuleSets(namespace string) SesReceiptRuleSetInterface
}

// SesReceiptRuleSetInterface has methods to work with SesReceiptRuleSet resources.
type SesReceiptRuleSetInterface interface {
	Create(*v1alpha1.SesReceiptRuleSet) (*v1alpha1.SesReceiptRuleSet, error)
	Update(*v1alpha1.SesReceiptRuleSet) (*v1alpha1.SesReceiptRuleSet, error)
	UpdateStatus(*v1alpha1.SesReceiptRuleSet) (*v1alpha1.SesReceiptRuleSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesReceiptRuleSet, error)
	List(opts v1.ListOptions) (*v1alpha1.SesReceiptRuleSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesReceiptRuleSet, err error)
	SesReceiptRuleSetExpansion
}

// sesReceiptRuleSets implements SesReceiptRuleSetInterface
type sesReceiptRuleSets struct {
	client rest.Interface
	ns     string
}

// newSesReceiptRuleSets returns a SesReceiptRuleSets
func newSesReceiptRuleSets(c *AwsV1alpha1Client, namespace string) *sesReceiptRuleSets {
	return &sesReceiptRuleSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sesReceiptRuleSet, and returns the corresponding sesReceiptRuleSet object, and an error if there is any.
func (c *sesReceiptRuleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.SesReceiptRuleSet, err error) {
	result = &v1alpha1.SesReceiptRuleSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesReceiptRuleSets that match those selectors.
func (c *sesReceiptRuleSets) List(opts v1.ListOptions) (result *v1alpha1.SesReceiptRuleSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesReceiptRuleSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesReceiptRuleSets.
func (c *sesReceiptRuleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesReceiptRuleSet and creates it.  Returns the server's representation of the sesReceiptRuleSet, and an error, if there is any.
func (c *sesReceiptRuleSets) Create(sesReceiptRuleSet *v1alpha1.SesReceiptRuleSet) (result *v1alpha1.SesReceiptRuleSet, err error) {
	result = &v1alpha1.SesReceiptRuleSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		Body(sesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesReceiptRuleSet and updates it. Returns the server's representation of the sesReceiptRuleSet, and an error, if there is any.
func (c *sesReceiptRuleSets) Update(sesReceiptRuleSet *v1alpha1.SesReceiptRuleSet) (result *v1alpha1.SesReceiptRuleSet, err error) {
	result = &v1alpha1.SesReceiptRuleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		Name(sesReceiptRuleSet.Name).
		Body(sesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesReceiptRuleSets) UpdateStatus(sesReceiptRuleSet *v1alpha1.SesReceiptRuleSet) (result *v1alpha1.SesReceiptRuleSet, err error) {
	result = &v1alpha1.SesReceiptRuleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		Name(sesReceiptRuleSet.Name).
		SubResource("status").
		Body(sesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesReceiptRuleSet and deletes it. Returns an error if one occurs.
func (c *sesReceiptRuleSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesReceiptRuleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesReceiptRuleSet.
func (c *sesReceiptRuleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesReceiptRuleSet, err error) {
	result = &v1alpha1.SesReceiptRuleSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sesreceiptrulesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
