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

// AwsSnsSmsPreferencesesGetter has a method to return a AwsSnsSmsPreferencesInterface.
// A group's client should implement this interface.
type AwsSnsSmsPreferencesesGetter interface {
	AwsSnsSmsPreferenceses() AwsSnsSmsPreferencesInterface
}

// AwsSnsSmsPreferencesInterface has methods to work with AwsSnsSmsPreferences resources.
type AwsSnsSmsPreferencesInterface interface {
	Create(*v1alpha1.AwsSnsSmsPreferences) (*v1alpha1.AwsSnsSmsPreferences, error)
	Update(*v1alpha1.AwsSnsSmsPreferences) (*v1alpha1.AwsSnsSmsPreferences, error)
	UpdateStatus(*v1alpha1.AwsSnsSmsPreferences) (*v1alpha1.AwsSnsSmsPreferences, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsSnsSmsPreferences, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsSnsSmsPreferencesList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSnsSmsPreferences, err error)
	AwsSnsSmsPreferencesExpansion
}

// awsSnsSmsPreferenceses implements AwsSnsSmsPreferencesInterface
type awsSnsSmsPreferenceses struct {
	client rest.Interface
}

// newAwsSnsSmsPreferenceses returns a AwsSnsSmsPreferenceses
func newAwsSnsSmsPreferenceses(c *AwsV1alpha1Client) *awsSnsSmsPreferenceses {
	return &awsSnsSmsPreferenceses{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsSnsSmsPreferences, and returns the corresponding awsSnsSmsPreferences object, and an error if there is any.
func (c *awsSnsSmsPreferenceses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSnsSmsPreferences, err error) {
	result = &v1alpha1.AwsSnsSmsPreferences{}
	err = c.client.Get().
		Resource("awssnssmspreferenceses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSnsSmsPreferenceses that match those selectors.
func (c *awsSnsSmsPreferenceses) List(opts v1.ListOptions) (result *v1alpha1.AwsSnsSmsPreferencesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsSnsSmsPreferencesList{}
	err = c.client.Get().
		Resource("awssnssmspreferenceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSnsSmsPreferenceses.
func (c *awsSnsSmsPreferenceses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awssnssmspreferenceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsSnsSmsPreferences and creates it.  Returns the server's representation of the awsSnsSmsPreferences, and an error, if there is any.
func (c *awsSnsSmsPreferenceses) Create(awsSnsSmsPreferences *v1alpha1.AwsSnsSmsPreferences) (result *v1alpha1.AwsSnsSmsPreferences, err error) {
	result = &v1alpha1.AwsSnsSmsPreferences{}
	err = c.client.Post().
		Resource("awssnssmspreferenceses").
		Body(awsSnsSmsPreferences).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSnsSmsPreferences and updates it. Returns the server's representation of the awsSnsSmsPreferences, and an error, if there is any.
func (c *awsSnsSmsPreferenceses) Update(awsSnsSmsPreferences *v1alpha1.AwsSnsSmsPreferences) (result *v1alpha1.AwsSnsSmsPreferences, err error) {
	result = &v1alpha1.AwsSnsSmsPreferences{}
	err = c.client.Put().
		Resource("awssnssmspreferenceses").
		Name(awsSnsSmsPreferences.Name).
		Body(awsSnsSmsPreferences).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsSnsSmsPreferenceses) UpdateStatus(awsSnsSmsPreferences *v1alpha1.AwsSnsSmsPreferences) (result *v1alpha1.AwsSnsSmsPreferences, err error) {
	result = &v1alpha1.AwsSnsSmsPreferences{}
	err = c.client.Put().
		Resource("awssnssmspreferenceses").
		Name(awsSnsSmsPreferences.Name).
		SubResource("status").
		Body(awsSnsSmsPreferences).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSnsSmsPreferences and deletes it. Returns an error if one occurs.
func (c *awsSnsSmsPreferenceses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awssnssmspreferenceses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSnsSmsPreferenceses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awssnssmspreferenceses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSnsSmsPreferences.
func (c *awsSnsSmsPreferenceses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSnsSmsPreferences, err error) {
	result = &v1alpha1.AwsSnsSmsPreferences{}
	err = c.client.Patch(pt).
		Resource("awssnssmspreferenceses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
