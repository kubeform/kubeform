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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// S3BucketPoliciesGetter has a method to return a S3BucketPolicyInterface.
// A group's client should implement this interface.
type S3BucketPoliciesGetter interface {
	S3BucketPolicies(namespace string) S3BucketPolicyInterface
}

// S3BucketPolicyInterface has methods to work with S3BucketPolicy resources.
type S3BucketPolicyInterface interface {
	Create(*v1alpha1.S3BucketPolicy) (*v1alpha1.S3BucketPolicy, error)
	Update(*v1alpha1.S3BucketPolicy) (*v1alpha1.S3BucketPolicy, error)
	UpdateStatus(*v1alpha1.S3BucketPolicy) (*v1alpha1.S3BucketPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.S3BucketPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.S3BucketPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketPolicy, err error)
	S3BucketPolicyExpansion
}

// s3BucketPolicies implements S3BucketPolicyInterface
type s3BucketPolicies struct {
	client rest.Interface
	ns     string
}

// newS3BucketPolicies returns a S3BucketPolicies
func newS3BucketPolicies(c *AwsV1alpha1Client, namespace string) *s3BucketPolicies {
	return &s3BucketPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s3BucketPolicy, and returns the corresponding s3BucketPolicy object, and an error if there is any.
func (c *s3BucketPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.S3BucketPolicy, err error) {
	result = &v1alpha1.S3BucketPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S3BucketPolicies that match those selectors.
func (c *s3BucketPolicies) List(opts v1.ListOptions) (result *v1alpha1.S3BucketPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.S3BucketPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s3BucketPolicies.
func (c *s3BucketPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a s3BucketPolicy and creates it.  Returns the server's representation of the s3BucketPolicy, and an error, if there is any.
func (c *s3BucketPolicies) Create(s3BucketPolicy *v1alpha1.S3BucketPolicy) (result *v1alpha1.S3BucketPolicy, err error) {
	result = &v1alpha1.S3BucketPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		Body(s3BucketPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s3BucketPolicy and updates it. Returns the server's representation of the s3BucketPolicy, and an error, if there is any.
func (c *s3BucketPolicies) Update(s3BucketPolicy *v1alpha1.S3BucketPolicy) (result *v1alpha1.S3BucketPolicy, err error) {
	result = &v1alpha1.S3BucketPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		Name(s3BucketPolicy.Name).
		Body(s3BucketPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *s3BucketPolicies) UpdateStatus(s3BucketPolicy *v1alpha1.S3BucketPolicy) (result *v1alpha1.S3BucketPolicy, err error) {
	result = &v1alpha1.S3BucketPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		Name(s3BucketPolicy.Name).
		SubResource("status").
		Body(s3BucketPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the s3BucketPolicy and deletes it. Returns an error if one occurs.
func (c *s3BucketPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s3BucketPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s3BucketPolicy.
func (c *s3BucketPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketPolicy, err error) {
	result = &v1alpha1.S3BucketPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s3bucketpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
