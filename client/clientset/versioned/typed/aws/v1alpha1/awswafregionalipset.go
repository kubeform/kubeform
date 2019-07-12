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

// AwsWafregionalIpsetsGetter has a method to return a AwsWafregionalIpsetInterface.
// A group's client should implement this interface.
type AwsWafregionalIpsetsGetter interface {
	AwsWafregionalIpsets() AwsWafregionalIpsetInterface
}

// AwsWafregionalIpsetInterface has methods to work with AwsWafregionalIpset resources.
type AwsWafregionalIpsetInterface interface {
	Create(*v1alpha1.AwsWafregionalIpset) (*v1alpha1.AwsWafregionalIpset, error)
	Update(*v1alpha1.AwsWafregionalIpset) (*v1alpha1.AwsWafregionalIpset, error)
	UpdateStatus(*v1alpha1.AwsWafregionalIpset) (*v1alpha1.AwsWafregionalIpset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsWafregionalIpset, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsWafregionalIpsetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalIpset, err error)
	AwsWafregionalIpsetExpansion
}

// awsWafregionalIpsets implements AwsWafregionalIpsetInterface
type awsWafregionalIpsets struct {
	client rest.Interface
}

// newAwsWafregionalIpsets returns a AwsWafregionalIpsets
func newAwsWafregionalIpsets(c *AwsV1alpha1Client) *awsWafregionalIpsets {
	return &awsWafregionalIpsets{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsWafregionalIpset, and returns the corresponding awsWafregionalIpset object, and an error if there is any.
func (c *awsWafregionalIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsWafregionalIpset, err error) {
	result = &v1alpha1.AwsWafregionalIpset{}
	err = c.client.Get().
		Resource("awswafregionalipsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafregionalIpsets that match those selectors.
func (c *awsWafregionalIpsets) List(opts v1.ListOptions) (result *v1alpha1.AwsWafregionalIpsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsWafregionalIpsetList{}
	err = c.client.Get().
		Resource("awswafregionalipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafregionalIpsets.
func (c *awsWafregionalIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awswafregionalipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsWafregionalIpset and creates it.  Returns the server's representation of the awsWafregionalIpset, and an error, if there is any.
func (c *awsWafregionalIpsets) Create(awsWafregionalIpset *v1alpha1.AwsWafregionalIpset) (result *v1alpha1.AwsWafregionalIpset, err error) {
	result = &v1alpha1.AwsWafregionalIpset{}
	err = c.client.Post().
		Resource("awswafregionalipsets").
		Body(awsWafregionalIpset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafregionalIpset and updates it. Returns the server's representation of the awsWafregionalIpset, and an error, if there is any.
func (c *awsWafregionalIpsets) Update(awsWafregionalIpset *v1alpha1.AwsWafregionalIpset) (result *v1alpha1.AwsWafregionalIpset, err error) {
	result = &v1alpha1.AwsWafregionalIpset{}
	err = c.client.Put().
		Resource("awswafregionalipsets").
		Name(awsWafregionalIpset.Name).
		Body(awsWafregionalIpset).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsWafregionalIpsets) UpdateStatus(awsWafregionalIpset *v1alpha1.AwsWafregionalIpset) (result *v1alpha1.AwsWafregionalIpset, err error) {
	result = &v1alpha1.AwsWafregionalIpset{}
	err = c.client.Put().
		Resource("awswafregionalipsets").
		Name(awsWafregionalIpset.Name).
		SubResource("status").
		Body(awsWafregionalIpset).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafregionalIpset and deletes it. Returns an error if one occurs.
func (c *awsWafregionalIpsets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awswafregionalipsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafregionalIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awswafregionalipsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafregionalIpset.
func (c *awsWafregionalIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsWafregionalIpset, err error) {
	result = &v1alpha1.AwsWafregionalIpset{}
	err = c.client.Patch(pt).
		Resource("awswafregionalipsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
