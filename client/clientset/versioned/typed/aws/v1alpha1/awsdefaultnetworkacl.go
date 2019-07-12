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

// AwsDefaultNetworkAclsGetter has a method to return a AwsDefaultNetworkAclInterface.
// A group's client should implement this interface.
type AwsDefaultNetworkAclsGetter interface {
	AwsDefaultNetworkAcls() AwsDefaultNetworkAclInterface
}

// AwsDefaultNetworkAclInterface has methods to work with AwsDefaultNetworkAcl resources.
type AwsDefaultNetworkAclInterface interface {
	Create(*v1alpha1.AwsDefaultNetworkAcl) (*v1alpha1.AwsDefaultNetworkAcl, error)
	Update(*v1alpha1.AwsDefaultNetworkAcl) (*v1alpha1.AwsDefaultNetworkAcl, error)
	UpdateStatus(*v1alpha1.AwsDefaultNetworkAcl) (*v1alpha1.AwsDefaultNetworkAcl, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDefaultNetworkAcl, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDefaultNetworkAclList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultNetworkAcl, err error)
	AwsDefaultNetworkAclExpansion
}

// awsDefaultNetworkAcls implements AwsDefaultNetworkAclInterface
type awsDefaultNetworkAcls struct {
	client rest.Interface
}

// newAwsDefaultNetworkAcls returns a AwsDefaultNetworkAcls
func newAwsDefaultNetworkAcls(c *AwsV1alpha1Client) *awsDefaultNetworkAcls {
	return &awsDefaultNetworkAcls{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDefaultNetworkAcl, and returns the corresponding awsDefaultNetworkAcl object, and an error if there is any.
func (c *awsDefaultNetworkAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	result = &v1alpha1.AwsDefaultNetworkAcl{}
	err = c.client.Get().
		Resource("awsdefaultnetworkacls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDefaultNetworkAcls that match those selectors.
func (c *awsDefaultNetworkAcls) List(opts v1.ListOptions) (result *v1alpha1.AwsDefaultNetworkAclList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDefaultNetworkAclList{}
	err = c.client.Get().
		Resource("awsdefaultnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDefaultNetworkAcls.
func (c *awsDefaultNetworkAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdefaultnetworkacls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDefaultNetworkAcl and creates it.  Returns the server's representation of the awsDefaultNetworkAcl, and an error, if there is any.
func (c *awsDefaultNetworkAcls) Create(awsDefaultNetworkAcl *v1alpha1.AwsDefaultNetworkAcl) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	result = &v1alpha1.AwsDefaultNetworkAcl{}
	err = c.client.Post().
		Resource("awsdefaultnetworkacls").
		Body(awsDefaultNetworkAcl).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDefaultNetworkAcl and updates it. Returns the server's representation of the awsDefaultNetworkAcl, and an error, if there is any.
func (c *awsDefaultNetworkAcls) Update(awsDefaultNetworkAcl *v1alpha1.AwsDefaultNetworkAcl) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	result = &v1alpha1.AwsDefaultNetworkAcl{}
	err = c.client.Put().
		Resource("awsdefaultnetworkacls").
		Name(awsDefaultNetworkAcl.Name).
		Body(awsDefaultNetworkAcl).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDefaultNetworkAcls) UpdateStatus(awsDefaultNetworkAcl *v1alpha1.AwsDefaultNetworkAcl) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	result = &v1alpha1.AwsDefaultNetworkAcl{}
	err = c.client.Put().
		Resource("awsdefaultnetworkacls").
		Name(awsDefaultNetworkAcl.Name).
		SubResource("status").
		Body(awsDefaultNetworkAcl).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDefaultNetworkAcl and deletes it. Returns an error if one occurs.
func (c *awsDefaultNetworkAcls) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdefaultnetworkacls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDefaultNetworkAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdefaultnetworkacls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDefaultNetworkAcl.
func (c *awsDefaultNetworkAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDefaultNetworkAcl, err error) {
	result = &v1alpha1.AwsDefaultNetworkAcl{}
	err = c.client.Patch(pt).
		Resource("awsdefaultnetworkacls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
