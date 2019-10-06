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

// SnsSmsPreferencesesGetter has a method to return a SnsSmsPreferencesInterface.
// A group's client should implement this interface.
type SnsSmsPreferencesesGetter interface {
	SnsSmsPreferenceses(namespace string) SnsSmsPreferencesInterface
}

// SnsSmsPreferencesInterface has methods to work with SnsSmsPreferences resources.
type SnsSmsPreferencesInterface interface {
	Create(*v1alpha1.SnsSmsPreferences) (*v1alpha1.SnsSmsPreferences, error)
	Update(*v1alpha1.SnsSmsPreferences) (*v1alpha1.SnsSmsPreferences, error)
	UpdateStatus(*v1alpha1.SnsSmsPreferences) (*v1alpha1.SnsSmsPreferences, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SnsSmsPreferences, error)
	List(opts v1.ListOptions) (*v1alpha1.SnsSmsPreferencesList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SnsSmsPreferences, err error)
	SnsSmsPreferencesExpansion
}

// snsSmsPreferenceses implements SnsSmsPreferencesInterface
type snsSmsPreferenceses struct {
	client rest.Interface
	ns     string
}

// newSnsSmsPreferenceses returns a SnsSmsPreferenceses
func newSnsSmsPreferenceses(c *AwsV1alpha1Client, namespace string) *snsSmsPreferenceses {
	return &snsSmsPreferenceses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the snsSmsPreferences, and returns the corresponding snsSmsPreferences object, and an error if there is any.
func (c *snsSmsPreferenceses) Get(name string, options v1.GetOptions) (result *v1alpha1.SnsSmsPreferences, err error) {
	result = &v1alpha1.SnsSmsPreferences{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SnsSmsPreferenceses that match those selectors.
func (c *snsSmsPreferenceses) List(opts v1.ListOptions) (result *v1alpha1.SnsSmsPreferencesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SnsSmsPreferencesList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested snsSmsPreferenceses.
func (c *snsSmsPreferenceses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a snsSmsPreferences and creates it.  Returns the server's representation of the snsSmsPreferences, and an error, if there is any.
func (c *snsSmsPreferenceses) Create(snsSmsPreferences *v1alpha1.SnsSmsPreferences) (result *v1alpha1.SnsSmsPreferences, err error) {
	result = &v1alpha1.SnsSmsPreferences{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		Body(snsSmsPreferences).
		Do().
		Into(result)
	return
}

// Update takes the representation of a snsSmsPreferences and updates it. Returns the server's representation of the snsSmsPreferences, and an error, if there is any.
func (c *snsSmsPreferenceses) Update(snsSmsPreferences *v1alpha1.SnsSmsPreferences) (result *v1alpha1.SnsSmsPreferences, err error) {
	result = &v1alpha1.SnsSmsPreferences{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		Name(snsSmsPreferences.Name).
		Body(snsSmsPreferences).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *snsSmsPreferenceses) UpdateStatus(snsSmsPreferences *v1alpha1.SnsSmsPreferences) (result *v1alpha1.SnsSmsPreferences, err error) {
	result = &v1alpha1.SnsSmsPreferences{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		Name(snsSmsPreferences.Name).
		SubResource("status").
		Body(snsSmsPreferences).
		Do().
		Into(result)
	return
}

// Delete takes name of the snsSmsPreferences and deletes it. Returns an error if one occurs.
func (c *snsSmsPreferenceses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *snsSmsPreferenceses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched snsSmsPreferences.
func (c *snsSmsPreferenceses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SnsSmsPreferences, err error) {
	result = &v1alpha1.SnsSmsPreferences{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("snssmspreferenceses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
