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

// LightsailInstancesGetter has a method to return a LightsailInstanceInterface.
// A group's client should implement this interface.
type LightsailInstancesGetter interface {
	LightsailInstances(namespace string) LightsailInstanceInterface
}

// LightsailInstanceInterface has methods to work with LightsailInstance resources.
type LightsailInstanceInterface interface {
	Create(*v1alpha1.LightsailInstance) (*v1alpha1.LightsailInstance, error)
	Update(*v1alpha1.LightsailInstance) (*v1alpha1.LightsailInstance, error)
	UpdateStatus(*v1alpha1.LightsailInstance) (*v1alpha1.LightsailInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LightsailInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.LightsailInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailInstance, err error)
	LightsailInstanceExpansion
}

// lightsailInstances implements LightsailInstanceInterface
type lightsailInstances struct {
	client rest.Interface
	ns     string
}

// newLightsailInstances returns a LightsailInstances
func newLightsailInstances(c *AwsV1alpha1Client, namespace string) *lightsailInstances {
	return &lightsailInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lightsailInstance, and returns the corresponding lightsailInstance object, and an error if there is any.
func (c *lightsailInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.LightsailInstance, err error) {
	result = &v1alpha1.LightsailInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lightsailinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LightsailInstances that match those selectors.
func (c *lightsailInstances) List(opts v1.ListOptions) (result *v1alpha1.LightsailInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LightsailInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lightsailinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lightsailInstances.
func (c *lightsailInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lightsailinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lightsailInstance and creates it.  Returns the server's representation of the lightsailInstance, and an error, if there is any.
func (c *lightsailInstances) Create(lightsailInstance *v1alpha1.LightsailInstance) (result *v1alpha1.LightsailInstance, err error) {
	result = &v1alpha1.LightsailInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lightsailinstances").
		Body(lightsailInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lightsailInstance and updates it. Returns the server's representation of the lightsailInstance, and an error, if there is any.
func (c *lightsailInstances) Update(lightsailInstance *v1alpha1.LightsailInstance) (result *v1alpha1.LightsailInstance, err error) {
	result = &v1alpha1.LightsailInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lightsailinstances").
		Name(lightsailInstance.Name).
		Body(lightsailInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lightsailInstances) UpdateStatus(lightsailInstance *v1alpha1.LightsailInstance) (result *v1alpha1.LightsailInstance, err error) {
	result = &v1alpha1.LightsailInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lightsailinstances").
		Name(lightsailInstance.Name).
		SubResource("status").
		Body(lightsailInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the lightsailInstance and deletes it. Returns an error if one occurs.
func (c *lightsailInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lightsailinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lightsailInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lightsailinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lightsailInstance.
func (c *lightsailInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailInstance, err error) {
	result = &v1alpha1.LightsailInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lightsailinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
