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

// IamUserPoliciesGetter has a method to return a IamUserPolicyInterface.
// A group's client should implement this interface.
type IamUserPoliciesGetter interface {
	IamUserPolicies(namespace string) IamUserPolicyInterface
}

// IamUserPolicyInterface has methods to work with IamUserPolicy resources.
type IamUserPolicyInterface interface {
	Create(*v1alpha1.IamUserPolicy) (*v1alpha1.IamUserPolicy, error)
	Update(*v1alpha1.IamUserPolicy) (*v1alpha1.IamUserPolicy, error)
	UpdateStatus(*v1alpha1.IamUserPolicy) (*v1alpha1.IamUserPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamUserPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.IamUserPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserPolicy, err error)
	IamUserPolicyExpansion
}

// iamUserPolicies implements IamUserPolicyInterface
type iamUserPolicies struct {
	client rest.Interface
	ns     string
}

// newIamUserPolicies returns a IamUserPolicies
func newIamUserPolicies(c *AwsV1alpha1Client, namespace string) *iamUserPolicies {
	return &iamUserPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamUserPolicy, and returns the corresponding iamUserPolicy object, and an error if there is any.
func (c *iamUserPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IamUserPolicy, err error) {
	result = &v1alpha1.IamUserPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamUserPolicies that match those selectors.
func (c *iamUserPolicies) List(opts v1.ListOptions) (result *v1alpha1.IamUserPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamUserPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamUserPolicies.
func (c *iamUserPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamUserPolicy and creates it.  Returns the server's representation of the iamUserPolicy, and an error, if there is any.
func (c *iamUserPolicies) Create(iamUserPolicy *v1alpha1.IamUserPolicy) (result *v1alpha1.IamUserPolicy, err error) {
	result = &v1alpha1.IamUserPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		Body(iamUserPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamUserPolicy and updates it. Returns the server's representation of the iamUserPolicy, and an error, if there is any.
func (c *iamUserPolicies) Update(iamUserPolicy *v1alpha1.IamUserPolicy) (result *v1alpha1.IamUserPolicy, err error) {
	result = &v1alpha1.IamUserPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		Name(iamUserPolicy.Name).
		Body(iamUserPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamUserPolicies) UpdateStatus(iamUserPolicy *v1alpha1.IamUserPolicy) (result *v1alpha1.IamUserPolicy, err error) {
	result = &v1alpha1.IamUserPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		Name(iamUserPolicy.Name).
		SubResource("status").
		Body(iamUserPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamUserPolicy and deletes it. Returns an error if one occurs.
func (c *iamUserPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamUserPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamuserpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamUserPolicy.
func (c *iamUserPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserPolicy, err error) {
	result = &v1alpha1.IamUserPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamuserpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
