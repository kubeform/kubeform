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

// AwsOpsworksApplicationsGetter has a method to return a AwsOpsworksApplicationInterface.
// A group's client should implement this interface.
type AwsOpsworksApplicationsGetter interface {
	AwsOpsworksApplications() AwsOpsworksApplicationInterface
}

// AwsOpsworksApplicationInterface has methods to work with AwsOpsworksApplication resources.
type AwsOpsworksApplicationInterface interface {
	Create(*v1alpha1.AwsOpsworksApplication) (*v1alpha1.AwsOpsworksApplication, error)
	Update(*v1alpha1.AwsOpsworksApplication) (*v1alpha1.AwsOpsworksApplication, error)
	UpdateStatus(*v1alpha1.AwsOpsworksApplication) (*v1alpha1.AwsOpsworksApplication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsOpsworksApplication, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsOpsworksApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksApplication, err error)
	AwsOpsworksApplicationExpansion
}

// awsOpsworksApplications implements AwsOpsworksApplicationInterface
type awsOpsworksApplications struct {
	client rest.Interface
}

// newAwsOpsworksApplications returns a AwsOpsworksApplications
func newAwsOpsworksApplications(c *AwsV1alpha1Client) *awsOpsworksApplications {
	return &awsOpsworksApplications{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsOpsworksApplication, and returns the corresponding awsOpsworksApplication object, and an error if there is any.
func (c *awsOpsworksApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsOpsworksApplication, err error) {
	result = &v1alpha1.AwsOpsworksApplication{}
	err = c.client.Get().
		Resource("awsopsworksapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsOpsworksApplications that match those selectors.
func (c *awsOpsworksApplications) List(opts v1.ListOptions) (result *v1alpha1.AwsOpsworksApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsOpsworksApplicationList{}
	err = c.client.Get().
		Resource("awsopsworksapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsOpsworksApplications.
func (c *awsOpsworksApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsopsworksapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsOpsworksApplication and creates it.  Returns the server's representation of the awsOpsworksApplication, and an error, if there is any.
func (c *awsOpsworksApplications) Create(awsOpsworksApplication *v1alpha1.AwsOpsworksApplication) (result *v1alpha1.AwsOpsworksApplication, err error) {
	result = &v1alpha1.AwsOpsworksApplication{}
	err = c.client.Post().
		Resource("awsopsworksapplications").
		Body(awsOpsworksApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsOpsworksApplication and updates it. Returns the server's representation of the awsOpsworksApplication, and an error, if there is any.
func (c *awsOpsworksApplications) Update(awsOpsworksApplication *v1alpha1.AwsOpsworksApplication) (result *v1alpha1.AwsOpsworksApplication, err error) {
	result = &v1alpha1.AwsOpsworksApplication{}
	err = c.client.Put().
		Resource("awsopsworksapplications").
		Name(awsOpsworksApplication.Name).
		Body(awsOpsworksApplication).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsOpsworksApplications) UpdateStatus(awsOpsworksApplication *v1alpha1.AwsOpsworksApplication) (result *v1alpha1.AwsOpsworksApplication, err error) {
	result = &v1alpha1.AwsOpsworksApplication{}
	err = c.client.Put().
		Resource("awsopsworksapplications").
		Name(awsOpsworksApplication.Name).
		SubResource("status").
		Body(awsOpsworksApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsOpsworksApplication and deletes it. Returns an error if one occurs.
func (c *awsOpsworksApplications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsopsworksapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsOpsworksApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsopsworksapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsOpsworksApplication.
func (c *awsOpsworksApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsOpsworksApplication, err error) {
	result = &v1alpha1.AwsOpsworksApplication{}
	err = c.client.Patch(pt).
		Resource("awsopsworksapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
