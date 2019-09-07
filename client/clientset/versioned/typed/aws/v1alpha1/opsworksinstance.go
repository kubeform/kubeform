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

// OpsworksInstancesGetter has a method to return a OpsworksInstanceInterface.
// A group's client should implement this interface.
type OpsworksInstancesGetter interface {
	OpsworksInstances(namespace string) OpsworksInstanceInterface
}

// OpsworksInstanceInterface has methods to work with OpsworksInstance resources.
type OpsworksInstanceInterface interface {
	Create(*v1alpha1.OpsworksInstance) (*v1alpha1.OpsworksInstance, error)
	Update(*v1alpha1.OpsworksInstance) (*v1alpha1.OpsworksInstance, error)
	UpdateStatus(*v1alpha1.OpsworksInstance) (*v1alpha1.OpsworksInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpsworksInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.OpsworksInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksInstance, err error)
	OpsworksInstanceExpansion
}

// opsworksInstances implements OpsworksInstanceInterface
type opsworksInstances struct {
	client rest.Interface
	ns     string
}

// newOpsworksInstances returns a OpsworksInstances
func newOpsworksInstances(c *AwsV1alpha1Client, namespace string) *opsworksInstances {
	return &opsworksInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the opsworksInstance, and returns the corresponding opsworksInstance object, and an error if there is any.
func (c *opsworksInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksInstance, err error) {
	result = &v1alpha1.OpsworksInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opsworksinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpsworksInstances that match those selectors.
func (c *opsworksInstances) List(opts v1.ListOptions) (result *v1alpha1.OpsworksInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpsworksInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opsworksinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested opsworksInstances.
func (c *opsworksInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("opsworksinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a opsworksInstance and creates it.  Returns the server's representation of the opsworksInstance, and an error, if there is any.
func (c *opsworksInstances) Create(opsworksInstance *v1alpha1.OpsworksInstance) (result *v1alpha1.OpsworksInstance, err error) {
	result = &v1alpha1.OpsworksInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("opsworksinstances").
		Body(opsworksInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a opsworksInstance and updates it. Returns the server's representation of the opsworksInstance, and an error, if there is any.
func (c *opsworksInstances) Update(opsworksInstance *v1alpha1.OpsworksInstance) (result *v1alpha1.OpsworksInstance, err error) {
	result = &v1alpha1.OpsworksInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opsworksinstances").
		Name(opsworksInstance.Name).
		Body(opsworksInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *opsworksInstances) UpdateStatus(opsworksInstance *v1alpha1.OpsworksInstance) (result *v1alpha1.OpsworksInstance, err error) {
	result = &v1alpha1.OpsworksInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opsworksinstances").
		Name(opsworksInstance.Name).
		SubResource("status").
		Body(opsworksInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the opsworksInstance and deletes it. Returns an error if one occurs.
func (c *opsworksInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opsworksinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *opsworksInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opsworksinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched opsworksInstance.
func (c *opsworksInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksInstance, err error) {
	result = &v1alpha1.OpsworksInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("opsworksinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}