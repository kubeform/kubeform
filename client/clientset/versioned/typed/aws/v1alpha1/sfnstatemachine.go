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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SfnStateMachinesGetter has a method to return a SfnStateMachineInterface.
// A group's client should implement this interface.
type SfnStateMachinesGetter interface {
	SfnStateMachines(namespace string) SfnStateMachineInterface
}

// SfnStateMachineInterface has methods to work with SfnStateMachine resources.
type SfnStateMachineInterface interface {
	Create(*v1alpha1.SfnStateMachine) (*v1alpha1.SfnStateMachine, error)
	Update(*v1alpha1.SfnStateMachine) (*v1alpha1.SfnStateMachine, error)
	UpdateStatus(*v1alpha1.SfnStateMachine) (*v1alpha1.SfnStateMachine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SfnStateMachine, error)
	List(opts v1.ListOptions) (*v1alpha1.SfnStateMachineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SfnStateMachine, err error)
	SfnStateMachineExpansion
}

// sfnStateMachines implements SfnStateMachineInterface
type sfnStateMachines struct {
	client rest.Interface
	ns     string
}

// newSfnStateMachines returns a SfnStateMachines
func newSfnStateMachines(c *AwsV1alpha1Client, namespace string) *sfnStateMachines {
	return &sfnStateMachines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sfnStateMachine, and returns the corresponding sfnStateMachine object, and an error if there is any.
func (c *sfnStateMachines) Get(name string, options v1.GetOptions) (result *v1alpha1.SfnStateMachine, err error) {
	result = &v1alpha1.SfnStateMachine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SfnStateMachines that match those selectors.
func (c *sfnStateMachines) List(opts v1.ListOptions) (result *v1alpha1.SfnStateMachineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SfnStateMachineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sfnStateMachines.
func (c *sfnStateMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sfnStateMachine and creates it.  Returns the server's representation of the sfnStateMachine, and an error, if there is any.
func (c *sfnStateMachines) Create(sfnStateMachine *v1alpha1.SfnStateMachine) (result *v1alpha1.SfnStateMachine, err error) {
	result = &v1alpha1.SfnStateMachine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		Body(sfnStateMachine).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sfnStateMachine and updates it. Returns the server's representation of the sfnStateMachine, and an error, if there is any.
func (c *sfnStateMachines) Update(sfnStateMachine *v1alpha1.SfnStateMachine) (result *v1alpha1.SfnStateMachine, err error) {
	result = &v1alpha1.SfnStateMachine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		Name(sfnStateMachine.Name).
		Body(sfnStateMachine).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sfnStateMachines) UpdateStatus(sfnStateMachine *v1alpha1.SfnStateMachine) (result *v1alpha1.SfnStateMachine, err error) {
	result = &v1alpha1.SfnStateMachine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		Name(sfnStateMachine.Name).
		SubResource("status").
		Body(sfnStateMachine).
		Do().
		Into(result)
	return
}

// Delete takes name of the sfnStateMachine and deletes it. Returns an error if one occurs.
func (c *sfnStateMachines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sfnStateMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sfnstatemachines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sfnStateMachine.
func (c *sfnStateMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SfnStateMachine, err error) {
	result = &v1alpha1.SfnStateMachine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sfnstatemachines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
