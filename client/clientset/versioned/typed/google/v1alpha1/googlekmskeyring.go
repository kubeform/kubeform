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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleKmsKeyRingsGetter has a method to return a GoogleKmsKeyRingInterface.
// A group's client should implement this interface.
type GoogleKmsKeyRingsGetter interface {
	GoogleKmsKeyRings() GoogleKmsKeyRingInterface
}

// GoogleKmsKeyRingInterface has methods to work with GoogleKmsKeyRing resources.
type GoogleKmsKeyRingInterface interface {
	Create(*v1alpha1.GoogleKmsKeyRing) (*v1alpha1.GoogleKmsKeyRing, error)
	Update(*v1alpha1.GoogleKmsKeyRing) (*v1alpha1.GoogleKmsKeyRing, error)
	UpdateStatus(*v1alpha1.GoogleKmsKeyRing) (*v1alpha1.GoogleKmsKeyRing, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleKmsKeyRing, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleKmsKeyRingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleKmsKeyRing, err error)
	GoogleKmsKeyRingExpansion
}

// googleKmsKeyRings implements GoogleKmsKeyRingInterface
type googleKmsKeyRings struct {
	client rest.Interface
}

// newGoogleKmsKeyRings returns a GoogleKmsKeyRings
func newGoogleKmsKeyRings(c *GoogleV1alpha1Client) *googleKmsKeyRings {
	return &googleKmsKeyRings{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleKmsKeyRing, and returns the corresponding googleKmsKeyRing object, and an error if there is any.
func (c *googleKmsKeyRings) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleKmsKeyRing, err error) {
	result = &v1alpha1.GoogleKmsKeyRing{}
	err = c.client.Get().
		Resource("googlekmskeyrings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleKmsKeyRings that match those selectors.
func (c *googleKmsKeyRings) List(opts v1.ListOptions) (result *v1alpha1.GoogleKmsKeyRingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleKmsKeyRingList{}
	err = c.client.Get().
		Resource("googlekmskeyrings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleKmsKeyRings.
func (c *googleKmsKeyRings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googlekmskeyrings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleKmsKeyRing and creates it.  Returns the server's representation of the googleKmsKeyRing, and an error, if there is any.
func (c *googleKmsKeyRings) Create(googleKmsKeyRing *v1alpha1.GoogleKmsKeyRing) (result *v1alpha1.GoogleKmsKeyRing, err error) {
	result = &v1alpha1.GoogleKmsKeyRing{}
	err = c.client.Post().
		Resource("googlekmskeyrings").
		Body(googleKmsKeyRing).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleKmsKeyRing and updates it. Returns the server's representation of the googleKmsKeyRing, and an error, if there is any.
func (c *googleKmsKeyRings) Update(googleKmsKeyRing *v1alpha1.GoogleKmsKeyRing) (result *v1alpha1.GoogleKmsKeyRing, err error) {
	result = &v1alpha1.GoogleKmsKeyRing{}
	err = c.client.Put().
		Resource("googlekmskeyrings").
		Name(googleKmsKeyRing.Name).
		Body(googleKmsKeyRing).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleKmsKeyRings) UpdateStatus(googleKmsKeyRing *v1alpha1.GoogleKmsKeyRing) (result *v1alpha1.GoogleKmsKeyRing, err error) {
	result = &v1alpha1.GoogleKmsKeyRing{}
	err = c.client.Put().
		Resource("googlekmskeyrings").
		Name(googleKmsKeyRing.Name).
		SubResource("status").
		Body(googleKmsKeyRing).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleKmsKeyRing and deletes it. Returns an error if one occurs.
func (c *googleKmsKeyRings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googlekmskeyrings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleKmsKeyRings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googlekmskeyrings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleKmsKeyRing.
func (c *googleKmsKeyRings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleKmsKeyRing, err error) {
	result = &v1alpha1.GoogleKmsKeyRing{}
	err = c.client.Patch(pt).
		Resource("googlekmskeyrings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
