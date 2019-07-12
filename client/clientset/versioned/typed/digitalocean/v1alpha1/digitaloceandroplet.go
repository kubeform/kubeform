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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// DigitaloceanDropletsGetter has a method to return a DigitaloceanDropletInterface.
// A group's client should implement this interface.
type DigitaloceanDropletsGetter interface {
	DigitaloceanDroplets() DigitaloceanDropletInterface
}

// DigitaloceanDropletInterface has methods to work with DigitaloceanDroplet resources.
type DigitaloceanDropletInterface interface {
	Create(*v1alpha1.DigitaloceanDroplet) (*v1alpha1.DigitaloceanDroplet, error)
	Update(*v1alpha1.DigitaloceanDroplet) (*v1alpha1.DigitaloceanDroplet, error)
	UpdateStatus(*v1alpha1.DigitaloceanDroplet) (*v1alpha1.DigitaloceanDroplet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DigitaloceanDroplet, error)
	List(opts v1.ListOptions) (*v1alpha1.DigitaloceanDropletList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanDroplet, err error)
	DigitaloceanDropletExpansion
}

// digitaloceanDroplets implements DigitaloceanDropletInterface
type digitaloceanDroplets struct {
	client rest.Interface
}

// newDigitaloceanDroplets returns a DigitaloceanDroplets
func newDigitaloceanDroplets(c *DigitaloceanV1alpha1Client) *digitaloceanDroplets {
	return &digitaloceanDroplets{
		client: c.RESTClient(),
	}
}

// Get takes name of the digitaloceanDroplet, and returns the corresponding digitaloceanDroplet object, and an error if there is any.
func (c *digitaloceanDroplets) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanDroplet, err error) {
	result = &v1alpha1.DigitaloceanDroplet{}
	err = c.client.Get().
		Resource("digitaloceandroplets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DigitaloceanDroplets that match those selectors.
func (c *digitaloceanDroplets) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanDropletList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DigitaloceanDropletList{}
	err = c.client.Get().
		Resource("digitaloceandroplets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested digitaloceanDroplets.
func (c *digitaloceanDroplets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("digitaloceandroplets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a digitaloceanDroplet and creates it.  Returns the server's representation of the digitaloceanDroplet, and an error, if there is any.
func (c *digitaloceanDroplets) Create(digitaloceanDroplet *v1alpha1.DigitaloceanDroplet) (result *v1alpha1.DigitaloceanDroplet, err error) {
	result = &v1alpha1.DigitaloceanDroplet{}
	err = c.client.Post().
		Resource("digitaloceandroplets").
		Body(digitaloceanDroplet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a digitaloceanDroplet and updates it. Returns the server's representation of the digitaloceanDroplet, and an error, if there is any.
func (c *digitaloceanDroplets) Update(digitaloceanDroplet *v1alpha1.DigitaloceanDroplet) (result *v1alpha1.DigitaloceanDroplet, err error) {
	result = &v1alpha1.DigitaloceanDroplet{}
	err = c.client.Put().
		Resource("digitaloceandroplets").
		Name(digitaloceanDroplet.Name).
		Body(digitaloceanDroplet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *digitaloceanDroplets) UpdateStatus(digitaloceanDroplet *v1alpha1.DigitaloceanDroplet) (result *v1alpha1.DigitaloceanDroplet, err error) {
	result = &v1alpha1.DigitaloceanDroplet{}
	err = c.client.Put().
		Resource("digitaloceandroplets").
		Name(digitaloceanDroplet.Name).
		SubResource("status").
		Body(digitaloceanDroplet).
		Do().
		Into(result)
	return
}

// Delete takes name of the digitaloceanDroplet and deletes it. Returns an error if one occurs.
func (c *digitaloceanDroplets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("digitaloceandroplets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *digitaloceanDroplets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("digitaloceandroplets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched digitaloceanDroplet.
func (c *digitaloceanDroplets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanDroplet, err error) {
	result = &v1alpha1.DigitaloceanDroplet{}
	err = c.client.Patch(pt).
		Resource("digitaloceandroplets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
