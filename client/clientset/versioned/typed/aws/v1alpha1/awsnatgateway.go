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

// AwsNatGatewaysGetter has a method to return a AwsNatGatewayInterface.
// A group's client should implement this interface.
type AwsNatGatewaysGetter interface {
	AwsNatGateways() AwsNatGatewayInterface
}

// AwsNatGatewayInterface has methods to work with AwsNatGateway resources.
type AwsNatGatewayInterface interface {
	Create(*v1alpha1.AwsNatGateway) (*v1alpha1.AwsNatGateway, error)
	Update(*v1alpha1.AwsNatGateway) (*v1alpha1.AwsNatGateway, error)
	UpdateStatus(*v1alpha1.AwsNatGateway) (*v1alpha1.AwsNatGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsNatGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsNatGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNatGateway, err error)
	AwsNatGatewayExpansion
}

// awsNatGateways implements AwsNatGatewayInterface
type awsNatGateways struct {
	client rest.Interface
}

// newAwsNatGateways returns a AwsNatGateways
func newAwsNatGateways(c *AwsV1alpha1Client) *awsNatGateways {
	return &awsNatGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsNatGateway, and returns the corresponding awsNatGateway object, and an error if there is any.
func (c *awsNatGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsNatGateway, err error) {
	result = &v1alpha1.AwsNatGateway{}
	err = c.client.Get().
		Resource("awsnatgateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsNatGateways that match those selectors.
func (c *awsNatGateways) List(opts v1.ListOptions) (result *v1alpha1.AwsNatGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsNatGatewayList{}
	err = c.client.Get().
		Resource("awsnatgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsNatGateways.
func (c *awsNatGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsnatgateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsNatGateway and creates it.  Returns the server's representation of the awsNatGateway, and an error, if there is any.
func (c *awsNatGateways) Create(awsNatGateway *v1alpha1.AwsNatGateway) (result *v1alpha1.AwsNatGateway, err error) {
	result = &v1alpha1.AwsNatGateway{}
	err = c.client.Post().
		Resource("awsnatgateways").
		Body(awsNatGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsNatGateway and updates it. Returns the server's representation of the awsNatGateway, and an error, if there is any.
func (c *awsNatGateways) Update(awsNatGateway *v1alpha1.AwsNatGateway) (result *v1alpha1.AwsNatGateway, err error) {
	result = &v1alpha1.AwsNatGateway{}
	err = c.client.Put().
		Resource("awsnatgateways").
		Name(awsNatGateway.Name).
		Body(awsNatGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsNatGateways) UpdateStatus(awsNatGateway *v1alpha1.AwsNatGateway) (result *v1alpha1.AwsNatGateway, err error) {
	result = &v1alpha1.AwsNatGateway{}
	err = c.client.Put().
		Resource("awsnatgateways").
		Name(awsNatGateway.Name).
		SubResource("status").
		Body(awsNatGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsNatGateway and deletes it. Returns an error if one occurs.
func (c *awsNatGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsnatgateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsNatGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsnatgateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsNatGateway.
func (c *awsNatGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsNatGateway, err error) {
	result = &v1alpha1.AwsNatGateway{}
	err = c.client.Patch(pt).
		Resource("awsnatgateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
