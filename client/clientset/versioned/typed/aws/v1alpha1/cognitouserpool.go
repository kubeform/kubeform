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

// CognitoUserPoolsGetter has a method to return a CognitoUserPoolInterface.
// A group's client should implement this interface.
type CognitoUserPoolsGetter interface {
	CognitoUserPools() CognitoUserPoolInterface
}

// CognitoUserPoolInterface has methods to work with CognitoUserPool resources.
type CognitoUserPoolInterface interface {
	Create(*v1alpha1.CognitoUserPool) (*v1alpha1.CognitoUserPool, error)
	Update(*v1alpha1.CognitoUserPool) (*v1alpha1.CognitoUserPool, error)
	UpdateStatus(*v1alpha1.CognitoUserPool) (*v1alpha1.CognitoUserPool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CognitoUserPool, error)
	List(opts v1.ListOptions) (*v1alpha1.CognitoUserPoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoUserPool, err error)
	CognitoUserPoolExpansion
}

// cognitoUserPools implements CognitoUserPoolInterface
type cognitoUserPools struct {
	client rest.Interface
}

// newCognitoUserPools returns a CognitoUserPools
func newCognitoUserPools(c *AwsV1alpha1Client) *cognitoUserPools {
	return &cognitoUserPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the cognitoUserPool, and returns the corresponding cognitoUserPool object, and an error if there is any.
func (c *cognitoUserPools) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitoUserPool, err error) {
	result = &v1alpha1.CognitoUserPool{}
	err = c.client.Get().
		Resource("cognitouserpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CognitoUserPools that match those selectors.
func (c *cognitoUserPools) List(opts v1.ListOptions) (result *v1alpha1.CognitoUserPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CognitoUserPoolList{}
	err = c.client.Get().
		Resource("cognitouserpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cognitoUserPools.
func (c *cognitoUserPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cognitouserpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cognitoUserPool and creates it.  Returns the server's representation of the cognitoUserPool, and an error, if there is any.
func (c *cognitoUserPools) Create(cognitoUserPool *v1alpha1.CognitoUserPool) (result *v1alpha1.CognitoUserPool, err error) {
	result = &v1alpha1.CognitoUserPool{}
	err = c.client.Post().
		Resource("cognitouserpools").
		Body(cognitoUserPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cognitoUserPool and updates it. Returns the server's representation of the cognitoUserPool, and an error, if there is any.
func (c *cognitoUserPools) Update(cognitoUserPool *v1alpha1.CognitoUserPool) (result *v1alpha1.CognitoUserPool, err error) {
	result = &v1alpha1.CognitoUserPool{}
	err = c.client.Put().
		Resource("cognitouserpools").
		Name(cognitoUserPool.Name).
		Body(cognitoUserPool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cognitoUserPools) UpdateStatus(cognitoUserPool *v1alpha1.CognitoUserPool) (result *v1alpha1.CognitoUserPool, err error) {
	result = &v1alpha1.CognitoUserPool{}
	err = c.client.Put().
		Resource("cognitouserpools").
		Name(cognitoUserPool.Name).
		SubResource("status").
		Body(cognitoUserPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the cognitoUserPool and deletes it. Returns an error if one occurs.
func (c *cognitoUserPools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cognitouserpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cognitoUserPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cognitouserpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cognitoUserPool.
func (c *cognitoUserPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoUserPool, err error) {
	result = &v1alpha1.CognitoUserPool{}
	err = c.client.Patch(pt).
		Resource("cognitouserpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
