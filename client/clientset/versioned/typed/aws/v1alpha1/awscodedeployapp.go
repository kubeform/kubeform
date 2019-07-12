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

// AwsCodedeployAppsGetter has a method to return a AwsCodedeployAppInterface.
// A group's client should implement this interface.
type AwsCodedeployAppsGetter interface {
	AwsCodedeployApps() AwsCodedeployAppInterface
}

// AwsCodedeployAppInterface has methods to work with AwsCodedeployApp resources.
type AwsCodedeployAppInterface interface {
	Create(*v1alpha1.AwsCodedeployApp) (*v1alpha1.AwsCodedeployApp, error)
	Update(*v1alpha1.AwsCodedeployApp) (*v1alpha1.AwsCodedeployApp, error)
	UpdateStatus(*v1alpha1.AwsCodedeployApp) (*v1alpha1.AwsCodedeployApp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCodedeployApp, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCodedeployAppList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodedeployApp, err error)
	AwsCodedeployAppExpansion
}

// awsCodedeployApps implements AwsCodedeployAppInterface
type awsCodedeployApps struct {
	client rest.Interface
}

// newAwsCodedeployApps returns a AwsCodedeployApps
func newAwsCodedeployApps(c *AwsV1alpha1Client) *awsCodedeployApps {
	return &awsCodedeployApps{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCodedeployApp, and returns the corresponding awsCodedeployApp object, and an error if there is any.
func (c *awsCodedeployApps) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodedeployApp, err error) {
	result = &v1alpha1.AwsCodedeployApp{}
	err = c.client.Get().
		Resource("awscodedeployapps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCodedeployApps that match those selectors.
func (c *awsCodedeployApps) List(opts v1.ListOptions) (result *v1alpha1.AwsCodedeployAppList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCodedeployAppList{}
	err = c.client.Get().
		Resource("awscodedeployapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCodedeployApps.
func (c *awsCodedeployApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscodedeployapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCodedeployApp and creates it.  Returns the server's representation of the awsCodedeployApp, and an error, if there is any.
func (c *awsCodedeployApps) Create(awsCodedeployApp *v1alpha1.AwsCodedeployApp) (result *v1alpha1.AwsCodedeployApp, err error) {
	result = &v1alpha1.AwsCodedeployApp{}
	err = c.client.Post().
		Resource("awscodedeployapps").
		Body(awsCodedeployApp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCodedeployApp and updates it. Returns the server's representation of the awsCodedeployApp, and an error, if there is any.
func (c *awsCodedeployApps) Update(awsCodedeployApp *v1alpha1.AwsCodedeployApp) (result *v1alpha1.AwsCodedeployApp, err error) {
	result = &v1alpha1.AwsCodedeployApp{}
	err = c.client.Put().
		Resource("awscodedeployapps").
		Name(awsCodedeployApp.Name).
		Body(awsCodedeployApp).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCodedeployApps) UpdateStatus(awsCodedeployApp *v1alpha1.AwsCodedeployApp) (result *v1alpha1.AwsCodedeployApp, err error) {
	result = &v1alpha1.AwsCodedeployApp{}
	err = c.client.Put().
		Resource("awscodedeployapps").
		Name(awsCodedeployApp.Name).
		SubResource("status").
		Body(awsCodedeployApp).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCodedeployApp and deletes it. Returns an error if one occurs.
func (c *awsCodedeployApps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscodedeployapps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCodedeployApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscodedeployapps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCodedeployApp.
func (c *awsCodedeployApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodedeployApp, err error) {
	result = &v1alpha1.AwsCodedeployApp{}
	err = c.client.Patch(pt).
		Resource("awscodedeployapps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
