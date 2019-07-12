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

// AwsS3BucketPublicAccessBlocksGetter has a method to return a AwsS3BucketPublicAccessBlockInterface.
// A group's client should implement this interface.
type AwsS3BucketPublicAccessBlocksGetter interface {
	AwsS3BucketPublicAccessBlocks() AwsS3BucketPublicAccessBlockInterface
}

// AwsS3BucketPublicAccessBlockInterface has methods to work with AwsS3BucketPublicAccessBlock resources.
type AwsS3BucketPublicAccessBlockInterface interface {
	Create(*v1alpha1.AwsS3BucketPublicAccessBlock) (*v1alpha1.AwsS3BucketPublicAccessBlock, error)
	Update(*v1alpha1.AwsS3BucketPublicAccessBlock) (*v1alpha1.AwsS3BucketPublicAccessBlock, error)
	UpdateStatus(*v1alpha1.AwsS3BucketPublicAccessBlock) (*v1alpha1.AwsS3BucketPublicAccessBlock, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsS3BucketPublicAccessBlock, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsS3BucketPublicAccessBlockList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsS3BucketPublicAccessBlock, err error)
	AwsS3BucketPublicAccessBlockExpansion
}

// awsS3BucketPublicAccessBlocks implements AwsS3BucketPublicAccessBlockInterface
type awsS3BucketPublicAccessBlocks struct {
	client rest.Interface
}

// newAwsS3BucketPublicAccessBlocks returns a AwsS3BucketPublicAccessBlocks
func newAwsS3BucketPublicAccessBlocks(c *AwsV1alpha1Client) *awsS3BucketPublicAccessBlocks {
	return &awsS3BucketPublicAccessBlocks{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsS3BucketPublicAccessBlock, and returns the corresponding awsS3BucketPublicAccessBlock object, and an error if there is any.
func (c *awsS3BucketPublicAccessBlocks) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsS3BucketPublicAccessBlock, err error) {
	result = &v1alpha1.AwsS3BucketPublicAccessBlock{}
	err = c.client.Get().
		Resource("awss3bucketpublicaccessblocks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsS3BucketPublicAccessBlocks that match those selectors.
func (c *awsS3BucketPublicAccessBlocks) List(opts v1.ListOptions) (result *v1alpha1.AwsS3BucketPublicAccessBlockList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsS3BucketPublicAccessBlockList{}
	err = c.client.Get().
		Resource("awss3bucketpublicaccessblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsS3BucketPublicAccessBlocks.
func (c *awsS3BucketPublicAccessBlocks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awss3bucketpublicaccessblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsS3BucketPublicAccessBlock and creates it.  Returns the server's representation of the awsS3BucketPublicAccessBlock, and an error, if there is any.
func (c *awsS3BucketPublicAccessBlocks) Create(awsS3BucketPublicAccessBlock *v1alpha1.AwsS3BucketPublicAccessBlock) (result *v1alpha1.AwsS3BucketPublicAccessBlock, err error) {
	result = &v1alpha1.AwsS3BucketPublicAccessBlock{}
	err = c.client.Post().
		Resource("awss3bucketpublicaccessblocks").
		Body(awsS3BucketPublicAccessBlock).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsS3BucketPublicAccessBlock and updates it. Returns the server's representation of the awsS3BucketPublicAccessBlock, and an error, if there is any.
func (c *awsS3BucketPublicAccessBlocks) Update(awsS3BucketPublicAccessBlock *v1alpha1.AwsS3BucketPublicAccessBlock) (result *v1alpha1.AwsS3BucketPublicAccessBlock, err error) {
	result = &v1alpha1.AwsS3BucketPublicAccessBlock{}
	err = c.client.Put().
		Resource("awss3bucketpublicaccessblocks").
		Name(awsS3BucketPublicAccessBlock.Name).
		Body(awsS3BucketPublicAccessBlock).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsS3BucketPublicAccessBlocks) UpdateStatus(awsS3BucketPublicAccessBlock *v1alpha1.AwsS3BucketPublicAccessBlock) (result *v1alpha1.AwsS3BucketPublicAccessBlock, err error) {
	result = &v1alpha1.AwsS3BucketPublicAccessBlock{}
	err = c.client.Put().
		Resource("awss3bucketpublicaccessblocks").
		Name(awsS3BucketPublicAccessBlock.Name).
		SubResource("status").
		Body(awsS3BucketPublicAccessBlock).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsS3BucketPublicAccessBlock and deletes it. Returns an error if one occurs.
func (c *awsS3BucketPublicAccessBlocks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awss3bucketpublicaccessblocks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsS3BucketPublicAccessBlocks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awss3bucketpublicaccessblocks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsS3BucketPublicAccessBlock.
func (c *awsS3BucketPublicAccessBlocks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsS3BucketPublicAccessBlock, err error) {
	result = &v1alpha1.AwsS3BucketPublicAccessBlock{}
	err = c.client.Patch(pt).
		Resource("awss3bucketpublicaccessblocks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
