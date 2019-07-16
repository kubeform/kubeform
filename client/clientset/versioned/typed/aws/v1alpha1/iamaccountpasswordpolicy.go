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

// IamAccountPasswordPoliciesGetter has a method to return a IamAccountPasswordPolicyInterface.
// A group's client should implement this interface.
type IamAccountPasswordPoliciesGetter interface {
	IamAccountPasswordPolicies() IamAccountPasswordPolicyInterface
}

// IamAccountPasswordPolicyInterface has methods to work with IamAccountPasswordPolicy resources.
type IamAccountPasswordPolicyInterface interface {
	Create(*v1alpha1.IamAccountPasswordPolicy) (*v1alpha1.IamAccountPasswordPolicy, error)
	Update(*v1alpha1.IamAccountPasswordPolicy) (*v1alpha1.IamAccountPasswordPolicy, error)
	UpdateStatus(*v1alpha1.IamAccountPasswordPolicy) (*v1alpha1.IamAccountPasswordPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamAccountPasswordPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.IamAccountPasswordPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamAccountPasswordPolicy, err error)
	IamAccountPasswordPolicyExpansion
}

// iamAccountPasswordPolicies implements IamAccountPasswordPolicyInterface
type iamAccountPasswordPolicies struct {
	client rest.Interface
}

// newIamAccountPasswordPolicies returns a IamAccountPasswordPolicies
func newIamAccountPasswordPolicies(c *AwsV1alpha1Client) *iamAccountPasswordPolicies {
	return &iamAccountPasswordPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the iamAccountPasswordPolicy, and returns the corresponding iamAccountPasswordPolicy object, and an error if there is any.
func (c *iamAccountPasswordPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IamAccountPasswordPolicy, err error) {
	result = &v1alpha1.IamAccountPasswordPolicy{}
	err = c.client.Get().
		Resource("iamaccountpasswordpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamAccountPasswordPolicies that match those selectors.
func (c *iamAccountPasswordPolicies) List(opts v1.ListOptions) (result *v1alpha1.IamAccountPasswordPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamAccountPasswordPolicyList{}
	err = c.client.Get().
		Resource("iamaccountpasswordpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamAccountPasswordPolicies.
func (c *iamAccountPasswordPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("iamaccountpasswordpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamAccountPasswordPolicy and creates it.  Returns the server's representation of the iamAccountPasswordPolicy, and an error, if there is any.
func (c *iamAccountPasswordPolicies) Create(iamAccountPasswordPolicy *v1alpha1.IamAccountPasswordPolicy) (result *v1alpha1.IamAccountPasswordPolicy, err error) {
	result = &v1alpha1.IamAccountPasswordPolicy{}
	err = c.client.Post().
		Resource("iamaccountpasswordpolicies").
		Body(iamAccountPasswordPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamAccountPasswordPolicy and updates it. Returns the server's representation of the iamAccountPasswordPolicy, and an error, if there is any.
func (c *iamAccountPasswordPolicies) Update(iamAccountPasswordPolicy *v1alpha1.IamAccountPasswordPolicy) (result *v1alpha1.IamAccountPasswordPolicy, err error) {
	result = &v1alpha1.IamAccountPasswordPolicy{}
	err = c.client.Put().
		Resource("iamaccountpasswordpolicies").
		Name(iamAccountPasswordPolicy.Name).
		Body(iamAccountPasswordPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamAccountPasswordPolicies) UpdateStatus(iamAccountPasswordPolicy *v1alpha1.IamAccountPasswordPolicy) (result *v1alpha1.IamAccountPasswordPolicy, err error) {
	result = &v1alpha1.IamAccountPasswordPolicy{}
	err = c.client.Put().
		Resource("iamaccountpasswordpolicies").
		Name(iamAccountPasswordPolicy.Name).
		SubResource("status").
		Body(iamAccountPasswordPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamAccountPasswordPolicy and deletes it. Returns an error if one occurs.
func (c *iamAccountPasswordPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("iamaccountpasswordpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamAccountPasswordPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("iamaccountpasswordpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamAccountPasswordPolicy.
func (c *iamAccountPasswordPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamAccountPasswordPolicy, err error) {
	result = &v1alpha1.IamAccountPasswordPolicy{}
	err = c.client.Patch(pt).
		Resource("iamaccountpasswordpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
