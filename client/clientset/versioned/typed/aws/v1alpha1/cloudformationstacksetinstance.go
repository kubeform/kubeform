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

// CloudformationStackSetInstancesGetter has a method to return a CloudformationStackSetInstanceInterface.
// A group's client should implement this interface.
type CloudformationStackSetInstancesGetter interface {
	CloudformationStackSetInstances(namespace string) CloudformationStackSetInstanceInterface
}

// CloudformationStackSetInstanceInterface has methods to work with CloudformationStackSetInstance resources.
type CloudformationStackSetInstanceInterface interface {
	Create(*v1alpha1.CloudformationStackSetInstance) (*v1alpha1.CloudformationStackSetInstance, error)
	Update(*v1alpha1.CloudformationStackSetInstance) (*v1alpha1.CloudformationStackSetInstance, error)
	UpdateStatus(*v1alpha1.CloudformationStackSetInstance) (*v1alpha1.CloudformationStackSetInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudformationStackSetInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudformationStackSetInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudformationStackSetInstance, err error)
	CloudformationStackSetInstanceExpansion
}

// cloudformationStackSetInstances implements CloudformationStackSetInstanceInterface
type cloudformationStackSetInstances struct {
	client rest.Interface
	ns     string
}

// newCloudformationStackSetInstances returns a CloudformationStackSetInstances
func newCloudformationStackSetInstances(c *AwsV1alpha1Client, namespace string) *cloudformationStackSetInstances {
	return &cloudformationStackSetInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudformationStackSetInstance, and returns the corresponding cloudformationStackSetInstance object, and an error if there is any.
func (c *cloudformationStackSetInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	result = &v1alpha1.CloudformationStackSetInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudformationStackSetInstances that match those selectors.
func (c *cloudformationStackSetInstances) List(opts v1.ListOptions) (result *v1alpha1.CloudformationStackSetInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudformationStackSetInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudformationStackSetInstances.
func (c *cloudformationStackSetInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudformationStackSetInstance and creates it.  Returns the server's representation of the cloudformationStackSetInstance, and an error, if there is any.
func (c *cloudformationStackSetInstances) Create(cloudformationStackSetInstance *v1alpha1.CloudformationStackSetInstance) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	result = &v1alpha1.CloudformationStackSetInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		Body(cloudformationStackSetInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudformationStackSetInstance and updates it. Returns the server's representation of the cloudformationStackSetInstance, and an error, if there is any.
func (c *cloudformationStackSetInstances) Update(cloudformationStackSetInstance *v1alpha1.CloudformationStackSetInstance) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	result = &v1alpha1.CloudformationStackSetInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		Name(cloudformationStackSetInstance.Name).
		Body(cloudformationStackSetInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudformationStackSetInstances) UpdateStatus(cloudformationStackSetInstance *v1alpha1.CloudformationStackSetInstance) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	result = &v1alpha1.CloudformationStackSetInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		Name(cloudformationStackSetInstance.Name).
		SubResource("status").
		Body(cloudformationStackSetInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudformationStackSetInstance and deletes it. Returns an error if one occurs.
func (c *cloudformationStackSetInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudformationStackSetInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudformationStackSetInstance.
func (c *cloudformationStackSetInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudformationStackSetInstance, err error) {
	result = &v1alpha1.CloudformationStackSetInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudformationstacksetinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}