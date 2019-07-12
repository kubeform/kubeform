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

// AwsGlacierVaultsGetter has a method to return a AwsGlacierVaultInterface.
// A group's client should implement this interface.
type AwsGlacierVaultsGetter interface {
	AwsGlacierVaults() AwsGlacierVaultInterface
}

// AwsGlacierVaultInterface has methods to work with AwsGlacierVault resources.
type AwsGlacierVaultInterface interface {
	Create(*v1alpha1.AwsGlacierVault) (*v1alpha1.AwsGlacierVault, error)
	Update(*v1alpha1.AwsGlacierVault) (*v1alpha1.AwsGlacierVault, error)
	UpdateStatus(*v1alpha1.AwsGlacierVault) (*v1alpha1.AwsGlacierVault, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsGlacierVault, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsGlacierVaultList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlacierVault, err error)
	AwsGlacierVaultExpansion
}

// awsGlacierVaults implements AwsGlacierVaultInterface
type awsGlacierVaults struct {
	client rest.Interface
}

// newAwsGlacierVaults returns a AwsGlacierVaults
func newAwsGlacierVaults(c *AwsV1alpha1Client) *awsGlacierVaults {
	return &awsGlacierVaults{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsGlacierVault, and returns the corresponding awsGlacierVault object, and an error if there is any.
func (c *awsGlacierVaults) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGlacierVault, err error) {
	result = &v1alpha1.AwsGlacierVault{}
	err = c.client.Get().
		Resource("awsglaciervaults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsGlacierVaults that match those selectors.
func (c *awsGlacierVaults) List(opts v1.ListOptions) (result *v1alpha1.AwsGlacierVaultList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsGlacierVaultList{}
	err = c.client.Get().
		Resource("awsglaciervaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsGlacierVaults.
func (c *awsGlacierVaults) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsglaciervaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsGlacierVault and creates it.  Returns the server's representation of the awsGlacierVault, and an error, if there is any.
func (c *awsGlacierVaults) Create(awsGlacierVault *v1alpha1.AwsGlacierVault) (result *v1alpha1.AwsGlacierVault, err error) {
	result = &v1alpha1.AwsGlacierVault{}
	err = c.client.Post().
		Resource("awsglaciervaults").
		Body(awsGlacierVault).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsGlacierVault and updates it. Returns the server's representation of the awsGlacierVault, and an error, if there is any.
func (c *awsGlacierVaults) Update(awsGlacierVault *v1alpha1.AwsGlacierVault) (result *v1alpha1.AwsGlacierVault, err error) {
	result = &v1alpha1.AwsGlacierVault{}
	err = c.client.Put().
		Resource("awsglaciervaults").
		Name(awsGlacierVault.Name).
		Body(awsGlacierVault).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsGlacierVaults) UpdateStatus(awsGlacierVault *v1alpha1.AwsGlacierVault) (result *v1alpha1.AwsGlacierVault, err error) {
	result = &v1alpha1.AwsGlacierVault{}
	err = c.client.Put().
		Resource("awsglaciervaults").
		Name(awsGlacierVault.Name).
		SubResource("status").
		Body(awsGlacierVault).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsGlacierVault and deletes it. Returns an error if one occurs.
func (c *awsGlacierVaults) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsglaciervaults").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsGlacierVaults) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsglaciervaults").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsGlacierVault.
func (c *awsGlacierVaults) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlacierVault, err error) {
	result = &v1alpha1.AwsGlacierVault{}
	err = c.client.Patch(pt).
		Resource("awsglaciervaults").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
