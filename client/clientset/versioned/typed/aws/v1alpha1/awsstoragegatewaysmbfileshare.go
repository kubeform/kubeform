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

// AwsStoragegatewaySmbFileSharesGetter has a method to return a AwsStoragegatewaySmbFileShareInterface.
// A group's client should implement this interface.
type AwsStoragegatewaySmbFileSharesGetter interface {
	AwsStoragegatewaySmbFileShares() AwsStoragegatewaySmbFileShareInterface
}

// AwsStoragegatewaySmbFileShareInterface has methods to work with AwsStoragegatewaySmbFileShare resources.
type AwsStoragegatewaySmbFileShareInterface interface {
	Create(*v1alpha1.AwsStoragegatewaySmbFileShare) (*v1alpha1.AwsStoragegatewaySmbFileShare, error)
	Update(*v1alpha1.AwsStoragegatewaySmbFileShare) (*v1alpha1.AwsStoragegatewaySmbFileShare, error)
	UpdateStatus(*v1alpha1.AwsStoragegatewaySmbFileShare) (*v1alpha1.AwsStoragegatewaySmbFileShare, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsStoragegatewaySmbFileShare, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsStoragegatewaySmbFileShareList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error)
	AwsStoragegatewaySmbFileShareExpansion
}

// awsStoragegatewaySmbFileShares implements AwsStoragegatewaySmbFileShareInterface
type awsStoragegatewaySmbFileShares struct {
	client rest.Interface
}

// newAwsStoragegatewaySmbFileShares returns a AwsStoragegatewaySmbFileShares
func newAwsStoragegatewaySmbFileShares(c *AwsV1alpha1Client) *awsStoragegatewaySmbFileShares {
	return &awsStoragegatewaySmbFileShares{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsStoragegatewaySmbFileShare, and returns the corresponding awsStoragegatewaySmbFileShare object, and an error if there is any.
func (c *awsStoragegatewaySmbFileShares) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	result = &v1alpha1.AwsStoragegatewaySmbFileShare{}
	err = c.client.Get().
		Resource("awsstoragegatewaysmbfileshares").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsStoragegatewaySmbFileShares that match those selectors.
func (c *awsStoragegatewaySmbFileShares) List(opts v1.ListOptions) (result *v1alpha1.AwsStoragegatewaySmbFileShareList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsStoragegatewaySmbFileShareList{}
	err = c.client.Get().
		Resource("awsstoragegatewaysmbfileshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsStoragegatewaySmbFileShares.
func (c *awsStoragegatewaySmbFileShares) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsstoragegatewaysmbfileshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsStoragegatewaySmbFileShare and creates it.  Returns the server's representation of the awsStoragegatewaySmbFileShare, and an error, if there is any.
func (c *awsStoragegatewaySmbFileShares) Create(awsStoragegatewaySmbFileShare *v1alpha1.AwsStoragegatewaySmbFileShare) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	result = &v1alpha1.AwsStoragegatewaySmbFileShare{}
	err = c.client.Post().
		Resource("awsstoragegatewaysmbfileshares").
		Body(awsStoragegatewaySmbFileShare).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsStoragegatewaySmbFileShare and updates it. Returns the server's representation of the awsStoragegatewaySmbFileShare, and an error, if there is any.
func (c *awsStoragegatewaySmbFileShares) Update(awsStoragegatewaySmbFileShare *v1alpha1.AwsStoragegatewaySmbFileShare) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	result = &v1alpha1.AwsStoragegatewaySmbFileShare{}
	err = c.client.Put().
		Resource("awsstoragegatewaysmbfileshares").
		Name(awsStoragegatewaySmbFileShare.Name).
		Body(awsStoragegatewaySmbFileShare).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsStoragegatewaySmbFileShares) UpdateStatus(awsStoragegatewaySmbFileShare *v1alpha1.AwsStoragegatewaySmbFileShare) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	result = &v1alpha1.AwsStoragegatewaySmbFileShare{}
	err = c.client.Put().
		Resource("awsstoragegatewaysmbfileshares").
		Name(awsStoragegatewaySmbFileShare.Name).
		SubResource("status").
		Body(awsStoragegatewaySmbFileShare).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsStoragegatewaySmbFileShare and deletes it. Returns an error if one occurs.
func (c *awsStoragegatewaySmbFileShares) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsstoragegatewaysmbfileshares").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsStoragegatewaySmbFileShares) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsstoragegatewaysmbfileshares").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsStoragegatewaySmbFileShare.
func (c *awsStoragegatewaySmbFileShares) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewaySmbFileShare, err error) {
	result = &v1alpha1.AwsStoragegatewaySmbFileShare{}
	err = c.client.Patch(pt).
		Resource("awsstoragegatewaysmbfileshares").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
