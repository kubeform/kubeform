/*
Copyright The Kubeform Authors.

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

// Route53ResolverRulesGetter has a method to return a Route53ResolverRuleInterface.
// A group's client should implement this interface.
type Route53ResolverRulesGetter interface {
	Route53ResolverRules(namespace string) Route53ResolverRuleInterface
}

// Route53ResolverRuleInterface has methods to work with Route53ResolverRule resources.
type Route53ResolverRuleInterface interface {
	Create(*v1alpha1.Route53ResolverRule) (*v1alpha1.Route53ResolverRule, error)
	Update(*v1alpha1.Route53ResolverRule) (*v1alpha1.Route53ResolverRule, error)
	UpdateStatus(*v1alpha1.Route53ResolverRule) (*v1alpha1.Route53ResolverRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Route53ResolverRule, error)
	List(opts v1.ListOptions) (*v1alpha1.Route53ResolverRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ResolverRule, err error)
	Route53ResolverRuleExpansion
}

// route53ResolverRules implements Route53ResolverRuleInterface
type route53ResolverRules struct {
	client rest.Interface
	ns     string
}

// newRoute53ResolverRules returns a Route53ResolverRules
func newRoute53ResolverRules(c *AwsV1alpha1Client, namespace string) *route53ResolverRules {
	return &route53ResolverRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the route53ResolverRule, and returns the corresponding route53ResolverRule object, and an error if there is any.
func (c *route53ResolverRules) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53ResolverRule, err error) {
	result = &v1alpha1.Route53ResolverRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("route53resolverrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Route53ResolverRules that match those selectors.
func (c *route53ResolverRules) List(opts v1.ListOptions) (result *v1alpha1.Route53ResolverRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Route53ResolverRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("route53resolverrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested route53ResolverRules.
func (c *route53ResolverRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("route53resolverrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a route53ResolverRule and creates it.  Returns the server's representation of the route53ResolverRule, and an error, if there is any.
func (c *route53ResolverRules) Create(route53ResolverRule *v1alpha1.Route53ResolverRule) (result *v1alpha1.Route53ResolverRule, err error) {
	result = &v1alpha1.Route53ResolverRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("route53resolverrules").
		Body(route53ResolverRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a route53ResolverRule and updates it. Returns the server's representation of the route53ResolverRule, and an error, if there is any.
func (c *route53ResolverRules) Update(route53ResolverRule *v1alpha1.Route53ResolverRule) (result *v1alpha1.Route53ResolverRule, err error) {
	result = &v1alpha1.Route53ResolverRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("route53resolverrules").
		Name(route53ResolverRule.Name).
		Body(route53ResolverRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *route53ResolverRules) UpdateStatus(route53ResolverRule *v1alpha1.Route53ResolverRule) (result *v1alpha1.Route53ResolverRule, err error) {
	result = &v1alpha1.Route53ResolverRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("route53resolverrules").
		Name(route53ResolverRule.Name).
		SubResource("status").
		Body(route53ResolverRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the route53ResolverRule and deletes it. Returns an error if one occurs.
func (c *route53ResolverRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("route53resolverrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *route53ResolverRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("route53resolverrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched route53ResolverRule.
func (c *route53ResolverRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ResolverRule, err error) {
	result = &v1alpha1.Route53ResolverRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("route53resolverrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
