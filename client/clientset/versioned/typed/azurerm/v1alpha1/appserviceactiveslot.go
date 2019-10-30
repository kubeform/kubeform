/*
Copyright The Kubeform Authors.

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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppServiceActiveSlotsGetter has a method to return a AppServiceActiveSlotInterface.
// A group's client should implement this interface.
type AppServiceActiveSlotsGetter interface {
	AppServiceActiveSlots(namespace string) AppServiceActiveSlotInterface
}

// AppServiceActiveSlotInterface has methods to work with AppServiceActiveSlot resources.
type AppServiceActiveSlotInterface interface {
	Create(*v1alpha1.AppServiceActiveSlot) (*v1alpha1.AppServiceActiveSlot, error)
	Update(*v1alpha1.AppServiceActiveSlot) (*v1alpha1.AppServiceActiveSlot, error)
	UpdateStatus(*v1alpha1.AppServiceActiveSlot) (*v1alpha1.AppServiceActiveSlot, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppServiceActiveSlot, error)
	List(opts v1.ListOptions) (*v1alpha1.AppServiceActiveSlotList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceActiveSlot, err error)
	AppServiceActiveSlotExpansion
}

// appServiceActiveSlots implements AppServiceActiveSlotInterface
type appServiceActiveSlots struct {
	client rest.Interface
	ns     string
}

// newAppServiceActiveSlots returns a AppServiceActiveSlots
func newAppServiceActiveSlots(c *AzurermV1alpha1Client, namespace string) *appServiceActiveSlots {
	return &appServiceActiveSlots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appServiceActiveSlot, and returns the corresponding appServiceActiveSlot object, and an error if there is any.
func (c *appServiceActiveSlots) Get(name string, options v1.GetOptions) (result *v1alpha1.AppServiceActiveSlot, err error) {
	result = &v1alpha1.AppServiceActiveSlot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppServiceActiveSlots that match those selectors.
func (c *appServiceActiveSlots) List(opts v1.ListOptions) (result *v1alpha1.AppServiceActiveSlotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppServiceActiveSlotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appServiceActiveSlots.
func (c *appServiceActiveSlots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appServiceActiveSlot and creates it.  Returns the server's representation of the appServiceActiveSlot, and an error, if there is any.
func (c *appServiceActiveSlots) Create(appServiceActiveSlot *v1alpha1.AppServiceActiveSlot) (result *v1alpha1.AppServiceActiveSlot, err error) {
	result = &v1alpha1.AppServiceActiveSlot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		Body(appServiceActiveSlot).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appServiceActiveSlot and updates it. Returns the server's representation of the appServiceActiveSlot, and an error, if there is any.
func (c *appServiceActiveSlots) Update(appServiceActiveSlot *v1alpha1.AppServiceActiveSlot) (result *v1alpha1.AppServiceActiveSlot, err error) {
	result = &v1alpha1.AppServiceActiveSlot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		Name(appServiceActiveSlot.Name).
		Body(appServiceActiveSlot).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appServiceActiveSlots) UpdateStatus(appServiceActiveSlot *v1alpha1.AppServiceActiveSlot) (result *v1alpha1.AppServiceActiveSlot, err error) {
	result = &v1alpha1.AppServiceActiveSlot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		Name(appServiceActiveSlot.Name).
		SubResource("status").
		Body(appServiceActiveSlot).
		Do().
		Into(result)
	return
}

// Delete takes name of the appServiceActiveSlot and deletes it. Returns an error if one occurs.
func (c *appServiceActiveSlots) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appServiceActiveSlots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appServiceActiveSlot.
func (c *appServiceActiveSlots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceActiveSlot, err error) {
	result = &v1alpha1.AppServiceActiveSlot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appserviceactiveslots").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
