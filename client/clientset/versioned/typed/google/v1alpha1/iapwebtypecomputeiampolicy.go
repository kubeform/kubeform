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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IapWebTypeComputeIamPoliciesGetter has a method to return a IapWebTypeComputeIamPolicyInterface.
// A group's client should implement this interface.
type IapWebTypeComputeIamPoliciesGetter interface {
	IapWebTypeComputeIamPolicies(namespace string) IapWebTypeComputeIamPolicyInterface
}

// IapWebTypeComputeIamPolicyInterface has methods to work with IapWebTypeComputeIamPolicy resources.
type IapWebTypeComputeIamPolicyInterface interface {
	Create(*v1alpha1.IapWebTypeComputeIamPolicy) (*v1alpha1.IapWebTypeComputeIamPolicy, error)
	Update(*v1alpha1.IapWebTypeComputeIamPolicy) (*v1alpha1.IapWebTypeComputeIamPolicy, error)
	UpdateStatus(*v1alpha1.IapWebTypeComputeIamPolicy) (*v1alpha1.IapWebTypeComputeIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IapWebTypeComputeIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.IapWebTypeComputeIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IapWebTypeComputeIamPolicy, err error)
	IapWebTypeComputeIamPolicyExpansion
}

// iapWebTypeComputeIamPolicies implements IapWebTypeComputeIamPolicyInterface
type iapWebTypeComputeIamPolicies struct {
	client rest.Interface
	ns     string
}

// newIapWebTypeComputeIamPolicies returns a IapWebTypeComputeIamPolicies
func newIapWebTypeComputeIamPolicies(c *GoogleV1alpha1Client, namespace string) *iapWebTypeComputeIamPolicies {
	return &iapWebTypeComputeIamPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iapWebTypeComputeIamPolicy, and returns the corresponding iapWebTypeComputeIamPolicy object, and an error if there is any.
func (c *iapWebTypeComputeIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IapWebTypeComputeIamPolicy, err error) {
	result = &v1alpha1.IapWebTypeComputeIamPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IapWebTypeComputeIamPolicies that match those selectors.
func (c *iapWebTypeComputeIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.IapWebTypeComputeIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IapWebTypeComputeIamPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iapWebTypeComputeIamPolicies.
func (c *iapWebTypeComputeIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iapWebTypeComputeIamPolicy and creates it.  Returns the server's representation of the iapWebTypeComputeIamPolicy, and an error, if there is any.
func (c *iapWebTypeComputeIamPolicies) Create(iapWebTypeComputeIamPolicy *v1alpha1.IapWebTypeComputeIamPolicy) (result *v1alpha1.IapWebTypeComputeIamPolicy, err error) {
	result = &v1alpha1.IapWebTypeComputeIamPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		Body(iapWebTypeComputeIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iapWebTypeComputeIamPolicy and updates it. Returns the server's representation of the iapWebTypeComputeIamPolicy, and an error, if there is any.
func (c *iapWebTypeComputeIamPolicies) Update(iapWebTypeComputeIamPolicy *v1alpha1.IapWebTypeComputeIamPolicy) (result *v1alpha1.IapWebTypeComputeIamPolicy, err error) {
	result = &v1alpha1.IapWebTypeComputeIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		Name(iapWebTypeComputeIamPolicy.Name).
		Body(iapWebTypeComputeIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iapWebTypeComputeIamPolicies) UpdateStatus(iapWebTypeComputeIamPolicy *v1alpha1.IapWebTypeComputeIamPolicy) (result *v1alpha1.IapWebTypeComputeIamPolicy, err error) {
	result = &v1alpha1.IapWebTypeComputeIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		Name(iapWebTypeComputeIamPolicy.Name).
		SubResource("status").
		Body(iapWebTypeComputeIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the iapWebTypeComputeIamPolicy and deletes it. Returns an error if one occurs.
func (c *iapWebTypeComputeIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iapWebTypeComputeIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iapWebTypeComputeIamPolicy.
func (c *iapWebTypeComputeIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IapWebTypeComputeIamPolicy, err error) {
	result = &v1alpha1.IapWebTypeComputeIamPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iapwebtypecomputeiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
