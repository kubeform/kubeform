/*
Copyright 2019 The KubeDB Authors.

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

// AwsAppmeshVirtualRoutersGetter has a method to return a AwsAppmeshVirtualRouterInterface.
// A group's client should implement this interface.
type AwsAppmeshVirtualRoutersGetter interface {
	AwsAppmeshVirtualRouters() AwsAppmeshVirtualRouterInterface
}

// AwsAppmeshVirtualRouterInterface has methods to work with AwsAppmeshVirtualRouter resources.
type AwsAppmeshVirtualRouterInterface interface {
	Create(*v1alpha1.AwsAppmeshVirtualRouter) (*v1alpha1.AwsAppmeshVirtualRouter, error)
	Update(*v1alpha1.AwsAppmeshVirtualRouter) (*v1alpha1.AwsAppmeshVirtualRouter, error)
	UpdateStatus(*v1alpha1.AwsAppmeshVirtualRouter) (*v1alpha1.AwsAppmeshVirtualRouter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsAppmeshVirtualRouter, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsAppmeshVirtualRouterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppmeshVirtualRouter, err error)
	AwsAppmeshVirtualRouterExpansion
}

// awsAppmeshVirtualRouters implements AwsAppmeshVirtualRouterInterface
type awsAppmeshVirtualRouters struct {
	client rest.Interface
}

// newAwsAppmeshVirtualRouters returns a AwsAppmeshVirtualRouters
func newAwsAppmeshVirtualRouters(c *AwsV1alpha1Client) *awsAppmeshVirtualRouters {
	return &awsAppmeshVirtualRouters{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsAppmeshVirtualRouter, and returns the corresponding awsAppmeshVirtualRouter object, and an error if there is any.
func (c *awsAppmeshVirtualRouters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsAppmeshVirtualRouter, err error) {
	result = &v1alpha1.AwsAppmeshVirtualRouter{}
	err = c.client.Get().
		Resource("awsappmeshvirtualrouters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsAppmeshVirtualRouters that match those selectors.
func (c *awsAppmeshVirtualRouters) List(opts v1.ListOptions) (result *v1alpha1.AwsAppmeshVirtualRouterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsAppmeshVirtualRouterList{}
	err = c.client.Get().
		Resource("awsappmeshvirtualrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsAppmeshVirtualRouters.
func (c *awsAppmeshVirtualRouters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsappmeshvirtualrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsAppmeshVirtualRouter and creates it.  Returns the server's representation of the awsAppmeshVirtualRouter, and an error, if there is any.
func (c *awsAppmeshVirtualRouters) Create(awsAppmeshVirtualRouter *v1alpha1.AwsAppmeshVirtualRouter) (result *v1alpha1.AwsAppmeshVirtualRouter, err error) {
	result = &v1alpha1.AwsAppmeshVirtualRouter{}
	err = c.client.Post().
		Resource("awsappmeshvirtualrouters").
		Body(awsAppmeshVirtualRouter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsAppmeshVirtualRouter and updates it. Returns the server's representation of the awsAppmeshVirtualRouter, and an error, if there is any.
func (c *awsAppmeshVirtualRouters) Update(awsAppmeshVirtualRouter *v1alpha1.AwsAppmeshVirtualRouter) (result *v1alpha1.AwsAppmeshVirtualRouter, err error) {
	result = &v1alpha1.AwsAppmeshVirtualRouter{}
	err = c.client.Put().
		Resource("awsappmeshvirtualrouters").
		Name(awsAppmeshVirtualRouter.Name).
		Body(awsAppmeshVirtualRouter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsAppmeshVirtualRouters) UpdateStatus(awsAppmeshVirtualRouter *v1alpha1.AwsAppmeshVirtualRouter) (result *v1alpha1.AwsAppmeshVirtualRouter, err error) {
	result = &v1alpha1.AwsAppmeshVirtualRouter{}
	err = c.client.Put().
		Resource("awsappmeshvirtualrouters").
		Name(awsAppmeshVirtualRouter.Name).
		SubResource("status").
		Body(awsAppmeshVirtualRouter).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsAppmeshVirtualRouter and deletes it. Returns an error if one occurs.
func (c *awsAppmeshVirtualRouters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsappmeshvirtualrouters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsAppmeshVirtualRouters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsappmeshvirtualrouters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsAppmeshVirtualRouter.
func (c *awsAppmeshVirtualRouters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsAppmeshVirtualRouter, err error) {
	result = &v1alpha1.AwsAppmeshVirtualRouter{}
	err = c.client.Patch(pt).
		Resource("awsappmeshvirtualrouters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
