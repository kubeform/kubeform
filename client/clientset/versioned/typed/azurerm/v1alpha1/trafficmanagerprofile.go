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

// TrafficManagerProfilesGetter has a method to return a TrafficManagerProfileInterface.
// A group's client should implement this interface.
type TrafficManagerProfilesGetter interface {
	TrafficManagerProfiles(namespace string) TrafficManagerProfileInterface
}

// TrafficManagerProfileInterface has methods to work with TrafficManagerProfile resources.
type TrafficManagerProfileInterface interface {
	Create(*v1alpha1.TrafficManagerProfile) (*v1alpha1.TrafficManagerProfile, error)
	Update(*v1alpha1.TrafficManagerProfile) (*v1alpha1.TrafficManagerProfile, error)
	UpdateStatus(*v1alpha1.TrafficManagerProfile) (*v1alpha1.TrafficManagerProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TrafficManagerProfile, error)
	List(opts v1.ListOptions) (*v1alpha1.TrafficManagerProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TrafficManagerProfile, err error)
	TrafficManagerProfileExpansion
}

// trafficManagerProfiles implements TrafficManagerProfileInterface
type trafficManagerProfiles struct {
	client rest.Interface
	ns     string
}

// newTrafficManagerProfiles returns a TrafficManagerProfiles
func newTrafficManagerProfiles(c *AzurermV1alpha1Client, namespace string) *trafficManagerProfiles {
	return &trafficManagerProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the trafficManagerProfile, and returns the corresponding trafficManagerProfile object, and an error if there is any.
func (c *trafficManagerProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.TrafficManagerProfile, err error) {
	result = &v1alpha1.TrafficManagerProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TrafficManagerProfiles that match those selectors.
func (c *trafficManagerProfiles) List(opts v1.ListOptions) (result *v1alpha1.TrafficManagerProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TrafficManagerProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested trafficManagerProfiles.
func (c *trafficManagerProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a trafficManagerProfile and creates it.  Returns the server's representation of the trafficManagerProfile, and an error, if there is any.
func (c *trafficManagerProfiles) Create(trafficManagerProfile *v1alpha1.TrafficManagerProfile) (result *v1alpha1.TrafficManagerProfile, err error) {
	result = &v1alpha1.TrafficManagerProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		Body(trafficManagerProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a trafficManagerProfile and updates it. Returns the server's representation of the trafficManagerProfile, and an error, if there is any.
func (c *trafficManagerProfiles) Update(trafficManagerProfile *v1alpha1.TrafficManagerProfile) (result *v1alpha1.TrafficManagerProfile, err error) {
	result = &v1alpha1.TrafficManagerProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		Name(trafficManagerProfile.Name).
		Body(trafficManagerProfile).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *trafficManagerProfiles) UpdateStatus(trafficManagerProfile *v1alpha1.TrafficManagerProfile) (result *v1alpha1.TrafficManagerProfile, err error) {
	result = &v1alpha1.TrafficManagerProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		Name(trafficManagerProfile.Name).
		SubResource("status").
		Body(trafficManagerProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the trafficManagerProfile and deletes it. Returns an error if one occurs.
func (c *trafficManagerProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *trafficManagerProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched trafficManagerProfile.
func (c *trafficManagerProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TrafficManagerProfile, err error) {
	result = &v1alpha1.TrafficManagerProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("trafficmanagerprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
