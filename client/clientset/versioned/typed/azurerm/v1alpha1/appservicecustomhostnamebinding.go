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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AppServiceCustomHostnameBindingsGetter has a method to return a AppServiceCustomHostnameBindingInterface.
// A group's client should implement this interface.
type AppServiceCustomHostnameBindingsGetter interface {
	AppServiceCustomHostnameBindings(namespace string) AppServiceCustomHostnameBindingInterface
}

// AppServiceCustomHostnameBindingInterface has methods to work with AppServiceCustomHostnameBinding resources.
type AppServiceCustomHostnameBindingInterface interface {
	Create(*v1alpha1.AppServiceCustomHostnameBinding) (*v1alpha1.AppServiceCustomHostnameBinding, error)
	Update(*v1alpha1.AppServiceCustomHostnameBinding) (*v1alpha1.AppServiceCustomHostnameBinding, error)
	UpdateStatus(*v1alpha1.AppServiceCustomHostnameBinding) (*v1alpha1.AppServiceCustomHostnameBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppServiceCustomHostnameBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.AppServiceCustomHostnameBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceCustomHostnameBinding, err error)
	AppServiceCustomHostnameBindingExpansion
}

// appServiceCustomHostnameBindings implements AppServiceCustomHostnameBindingInterface
type appServiceCustomHostnameBindings struct {
	client rest.Interface
	ns     string
}

// newAppServiceCustomHostnameBindings returns a AppServiceCustomHostnameBindings
func newAppServiceCustomHostnameBindings(c *AzurermV1alpha1Client, namespace string) *appServiceCustomHostnameBindings {
	return &appServiceCustomHostnameBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appServiceCustomHostnameBinding, and returns the corresponding appServiceCustomHostnameBinding object, and an error if there is any.
func (c *appServiceCustomHostnameBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.AppServiceCustomHostnameBinding, err error) {
	result = &v1alpha1.AppServiceCustomHostnameBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppServiceCustomHostnameBindings that match those selectors.
func (c *appServiceCustomHostnameBindings) List(opts v1.ListOptions) (result *v1alpha1.AppServiceCustomHostnameBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppServiceCustomHostnameBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appServiceCustomHostnameBindings.
func (c *appServiceCustomHostnameBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appServiceCustomHostnameBinding and creates it.  Returns the server's representation of the appServiceCustomHostnameBinding, and an error, if there is any.
func (c *appServiceCustomHostnameBindings) Create(appServiceCustomHostnameBinding *v1alpha1.AppServiceCustomHostnameBinding) (result *v1alpha1.AppServiceCustomHostnameBinding, err error) {
	result = &v1alpha1.AppServiceCustomHostnameBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		Body(appServiceCustomHostnameBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appServiceCustomHostnameBinding and updates it. Returns the server's representation of the appServiceCustomHostnameBinding, and an error, if there is any.
func (c *appServiceCustomHostnameBindings) Update(appServiceCustomHostnameBinding *v1alpha1.AppServiceCustomHostnameBinding) (result *v1alpha1.AppServiceCustomHostnameBinding, err error) {
	result = &v1alpha1.AppServiceCustomHostnameBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		Name(appServiceCustomHostnameBinding.Name).
		Body(appServiceCustomHostnameBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appServiceCustomHostnameBindings) UpdateStatus(appServiceCustomHostnameBinding *v1alpha1.AppServiceCustomHostnameBinding) (result *v1alpha1.AppServiceCustomHostnameBinding, err error) {
	result = &v1alpha1.AppServiceCustomHostnameBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		Name(appServiceCustomHostnameBinding.Name).
		SubResource("status").
		Body(appServiceCustomHostnameBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the appServiceCustomHostnameBinding and deletes it. Returns an error if one occurs.
func (c *appServiceCustomHostnameBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appServiceCustomHostnameBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appServiceCustomHostnameBinding.
func (c *appServiceCustomHostnameBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceCustomHostnameBinding, err error) {
	result = &v1alpha1.AppServiceCustomHostnameBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appservicecustomhostnamebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}