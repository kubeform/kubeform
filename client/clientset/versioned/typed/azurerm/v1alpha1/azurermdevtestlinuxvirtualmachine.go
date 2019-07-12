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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AzurermDevTestLinuxVirtualMachinesGetter has a method to return a AzurermDevTestLinuxVirtualMachineInterface.
// A group's client should implement this interface.
type AzurermDevTestLinuxVirtualMachinesGetter interface {
	AzurermDevTestLinuxVirtualMachines() AzurermDevTestLinuxVirtualMachineInterface
}

// AzurermDevTestLinuxVirtualMachineInterface has methods to work with AzurermDevTestLinuxVirtualMachine resources.
type AzurermDevTestLinuxVirtualMachineInterface interface {
	Create(*v1alpha1.AzurermDevTestLinuxVirtualMachine) (*v1alpha1.AzurermDevTestLinuxVirtualMachine, error)
	Update(*v1alpha1.AzurermDevTestLinuxVirtualMachine) (*v1alpha1.AzurermDevTestLinuxVirtualMachine, error)
	UpdateStatus(*v1alpha1.AzurermDevTestLinuxVirtualMachine) (*v1alpha1.AzurermDevTestLinuxVirtualMachine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzurermDevTestLinuxVirtualMachine, error)
	List(opts v1.ListOptions) (*v1alpha1.AzurermDevTestLinuxVirtualMachineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDevTestLinuxVirtualMachine, err error)
	AzurermDevTestLinuxVirtualMachineExpansion
}

// azurermDevTestLinuxVirtualMachines implements AzurermDevTestLinuxVirtualMachineInterface
type azurermDevTestLinuxVirtualMachines struct {
	client rest.Interface
}

// newAzurermDevTestLinuxVirtualMachines returns a AzurermDevTestLinuxVirtualMachines
func newAzurermDevTestLinuxVirtualMachines(c *AzurermV1alpha1Client) *azurermDevTestLinuxVirtualMachines {
	return &azurermDevTestLinuxVirtualMachines{
		client: c.RESTClient(),
	}
}

// Get takes name of the azurermDevTestLinuxVirtualMachine, and returns the corresponding azurermDevTestLinuxVirtualMachine object, and an error if there is any.
func (c *azurermDevTestLinuxVirtualMachines) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermDevTestLinuxVirtualMachine, err error) {
	result = &v1alpha1.AzurermDevTestLinuxVirtualMachine{}
	err = c.client.Get().
		Resource("azurermdevtestlinuxvirtualmachines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzurermDevTestLinuxVirtualMachines that match those selectors.
func (c *azurermDevTestLinuxVirtualMachines) List(opts v1.ListOptions) (result *v1alpha1.AzurermDevTestLinuxVirtualMachineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzurermDevTestLinuxVirtualMachineList{}
	err = c.client.Get().
		Resource("azurermdevtestlinuxvirtualmachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azurermDevTestLinuxVirtualMachines.
func (c *azurermDevTestLinuxVirtualMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("azurermdevtestlinuxvirtualmachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a azurermDevTestLinuxVirtualMachine and creates it.  Returns the server's representation of the azurermDevTestLinuxVirtualMachine, and an error, if there is any.
func (c *azurermDevTestLinuxVirtualMachines) Create(azurermDevTestLinuxVirtualMachine *v1alpha1.AzurermDevTestLinuxVirtualMachine) (result *v1alpha1.AzurermDevTestLinuxVirtualMachine, err error) {
	result = &v1alpha1.AzurermDevTestLinuxVirtualMachine{}
	err = c.client.Post().
		Resource("azurermdevtestlinuxvirtualmachines").
		Body(azurermDevTestLinuxVirtualMachine).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azurermDevTestLinuxVirtualMachine and updates it. Returns the server's representation of the azurermDevTestLinuxVirtualMachine, and an error, if there is any.
func (c *azurermDevTestLinuxVirtualMachines) Update(azurermDevTestLinuxVirtualMachine *v1alpha1.AzurermDevTestLinuxVirtualMachine) (result *v1alpha1.AzurermDevTestLinuxVirtualMachine, err error) {
	result = &v1alpha1.AzurermDevTestLinuxVirtualMachine{}
	err = c.client.Put().
		Resource("azurermdevtestlinuxvirtualmachines").
		Name(azurermDevTestLinuxVirtualMachine.Name).
		Body(azurermDevTestLinuxVirtualMachine).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *azurermDevTestLinuxVirtualMachines) UpdateStatus(azurermDevTestLinuxVirtualMachine *v1alpha1.AzurermDevTestLinuxVirtualMachine) (result *v1alpha1.AzurermDevTestLinuxVirtualMachine, err error) {
	result = &v1alpha1.AzurermDevTestLinuxVirtualMachine{}
	err = c.client.Put().
		Resource("azurermdevtestlinuxvirtualmachines").
		Name(azurermDevTestLinuxVirtualMachine.Name).
		SubResource("status").
		Body(azurermDevTestLinuxVirtualMachine).
		Do().
		Into(result)
	return
}

// Delete takes name of the azurermDevTestLinuxVirtualMachine and deletes it. Returns an error if one occurs.
func (c *azurermDevTestLinuxVirtualMachines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("azurermdevtestlinuxvirtualmachines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azurermDevTestLinuxVirtualMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("azurermdevtestlinuxvirtualmachines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azurermDevTestLinuxVirtualMachine.
func (c *azurermDevTestLinuxVirtualMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermDevTestLinuxVirtualMachine, err error) {
	result = &v1alpha1.AzurermDevTestLinuxVirtualMachine{}
	err = c.client.Patch(pt).
		Resource("azurermdevtestlinuxvirtualmachines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
