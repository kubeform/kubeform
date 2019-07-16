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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeSubnetworkIamPoliciesGetter has a method to return a ComputeSubnetworkIamPolicyInterface.
// A group's client should implement this interface.
type ComputeSubnetworkIamPoliciesGetter interface {
	ComputeSubnetworkIamPolicies() ComputeSubnetworkIamPolicyInterface
}

// ComputeSubnetworkIamPolicyInterface has methods to work with ComputeSubnetworkIamPolicy resources.
type ComputeSubnetworkIamPolicyInterface interface {
	Create(*v1alpha1.ComputeSubnetworkIamPolicy) (*v1alpha1.ComputeSubnetworkIamPolicy, error)
	Update(*v1alpha1.ComputeSubnetworkIamPolicy) (*v1alpha1.ComputeSubnetworkIamPolicy, error)
	UpdateStatus(*v1alpha1.ComputeSubnetworkIamPolicy) (*v1alpha1.ComputeSubnetworkIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeSubnetworkIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeSubnetworkIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error)
	ComputeSubnetworkIamPolicyExpansion
}

// computeSubnetworkIamPolicies implements ComputeSubnetworkIamPolicyInterface
type computeSubnetworkIamPolicies struct {
	client rest.Interface
}

// newComputeSubnetworkIamPolicies returns a ComputeSubnetworkIamPolicies
func newComputeSubnetworkIamPolicies(c *GoogleV1alpha1Client) *computeSubnetworkIamPolicies {
	return &computeSubnetworkIamPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the computeSubnetworkIamPolicy, and returns the corresponding computeSubnetworkIamPolicy object, and an error if there is any.
func (c *computeSubnetworkIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.ComputeSubnetworkIamPolicy{}
	err = c.client.Get().
		Resource("computesubnetworkiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeSubnetworkIamPolicies that match those selectors.
func (c *computeSubnetworkIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.ComputeSubnetworkIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeSubnetworkIamPolicyList{}
	err = c.client.Get().
		Resource("computesubnetworkiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeSubnetworkIamPolicies.
func (c *computeSubnetworkIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("computesubnetworkiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeSubnetworkIamPolicy and creates it.  Returns the server's representation of the computeSubnetworkIamPolicy, and an error, if there is any.
func (c *computeSubnetworkIamPolicies) Create(computeSubnetworkIamPolicy *v1alpha1.ComputeSubnetworkIamPolicy) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.ComputeSubnetworkIamPolicy{}
	err = c.client.Post().
		Resource("computesubnetworkiampolicies").
		Body(computeSubnetworkIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeSubnetworkIamPolicy and updates it. Returns the server's representation of the computeSubnetworkIamPolicy, and an error, if there is any.
func (c *computeSubnetworkIamPolicies) Update(computeSubnetworkIamPolicy *v1alpha1.ComputeSubnetworkIamPolicy) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.ComputeSubnetworkIamPolicy{}
	err = c.client.Put().
		Resource("computesubnetworkiampolicies").
		Name(computeSubnetworkIamPolicy.Name).
		Body(computeSubnetworkIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeSubnetworkIamPolicies) UpdateStatus(computeSubnetworkIamPolicy *v1alpha1.ComputeSubnetworkIamPolicy) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.ComputeSubnetworkIamPolicy{}
	err = c.client.Put().
		Resource("computesubnetworkiampolicies").
		Name(computeSubnetworkIamPolicy.Name).
		SubResource("status").
		Body(computeSubnetworkIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeSubnetworkIamPolicy and deletes it. Returns an error if one occurs.
func (c *computeSubnetworkIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("computesubnetworkiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeSubnetworkIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("computesubnetworkiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeSubnetworkIamPolicy.
func (c *computeSubnetworkIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSubnetworkIamPolicy, err error) {
	result = &v1alpha1.ComputeSubnetworkIamPolicy{}
	err = c.client.Patch(pt).
		Resource("computesubnetworkiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
