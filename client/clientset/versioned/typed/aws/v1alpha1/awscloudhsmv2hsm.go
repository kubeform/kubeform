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

// AwsCloudhsmV2HsmsGetter has a method to return a AwsCloudhsmV2HsmInterface.
// A group's client should implement this interface.
type AwsCloudhsmV2HsmsGetter interface {
	AwsCloudhsmV2Hsms() AwsCloudhsmV2HsmInterface
}

// AwsCloudhsmV2HsmInterface has methods to work with AwsCloudhsmV2Hsm resources.
type AwsCloudhsmV2HsmInterface interface {
	Create(*v1alpha1.AwsCloudhsmV2Hsm) (*v1alpha1.AwsCloudhsmV2Hsm, error)
	Update(*v1alpha1.AwsCloudhsmV2Hsm) (*v1alpha1.AwsCloudhsmV2Hsm, error)
	UpdateStatus(*v1alpha1.AwsCloudhsmV2Hsm) (*v1alpha1.AwsCloudhsmV2Hsm, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCloudhsmV2Hsm, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCloudhsmV2HsmList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudhsmV2Hsm, err error)
	AwsCloudhsmV2HsmExpansion
}

// awsCloudhsmV2Hsms implements AwsCloudhsmV2HsmInterface
type awsCloudhsmV2Hsms struct {
	client rest.Interface
}

// newAwsCloudhsmV2Hsms returns a AwsCloudhsmV2Hsms
func newAwsCloudhsmV2Hsms(c *AwsV1alpha1Client) *awsCloudhsmV2Hsms {
	return &awsCloudhsmV2Hsms{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCloudhsmV2Hsm, and returns the corresponding awsCloudhsmV2Hsm object, and an error if there is any.
func (c *awsCloudhsmV2Hsms) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	result = &v1alpha1.AwsCloudhsmV2Hsm{}
	err = c.client.Get().
		Resource("awscloudhsmv2hsms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudhsmV2Hsms that match those selectors.
func (c *awsCloudhsmV2Hsms) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudhsmV2HsmList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCloudhsmV2HsmList{}
	err = c.client.Get().
		Resource("awscloudhsmv2hsms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudhsmV2Hsms.
func (c *awsCloudhsmV2Hsms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscloudhsmv2hsms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCloudhsmV2Hsm and creates it.  Returns the server's representation of the awsCloudhsmV2Hsm, and an error, if there is any.
func (c *awsCloudhsmV2Hsms) Create(awsCloudhsmV2Hsm *v1alpha1.AwsCloudhsmV2Hsm) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	result = &v1alpha1.AwsCloudhsmV2Hsm{}
	err = c.client.Post().
		Resource("awscloudhsmv2hsms").
		Body(awsCloudhsmV2Hsm).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudhsmV2Hsm and updates it. Returns the server's representation of the awsCloudhsmV2Hsm, and an error, if there is any.
func (c *awsCloudhsmV2Hsms) Update(awsCloudhsmV2Hsm *v1alpha1.AwsCloudhsmV2Hsm) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	result = &v1alpha1.AwsCloudhsmV2Hsm{}
	err = c.client.Put().
		Resource("awscloudhsmv2hsms").
		Name(awsCloudhsmV2Hsm.Name).
		Body(awsCloudhsmV2Hsm).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCloudhsmV2Hsms) UpdateStatus(awsCloudhsmV2Hsm *v1alpha1.AwsCloudhsmV2Hsm) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	result = &v1alpha1.AwsCloudhsmV2Hsm{}
	err = c.client.Put().
		Resource("awscloudhsmv2hsms").
		Name(awsCloudhsmV2Hsm.Name).
		SubResource("status").
		Body(awsCloudhsmV2Hsm).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudhsmV2Hsm and deletes it. Returns an error if one occurs.
func (c *awsCloudhsmV2Hsms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscloudhsmv2hsms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudhsmV2Hsms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscloudhsmv2hsms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudhsmV2Hsm.
func (c *awsCloudhsmV2Hsms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudhsmV2Hsm, err error) {
	result = &v1alpha1.AwsCloudhsmV2Hsm{}
	err = c.client.Patch(pt).
		Resource("awscloudhsmv2hsms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
