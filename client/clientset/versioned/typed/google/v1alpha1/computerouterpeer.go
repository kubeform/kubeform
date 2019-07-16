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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeRouterPeersGetter has a method to return a ComputeRouterPeerInterface.
// A group's client should implement this interface.
type ComputeRouterPeersGetter interface {
	ComputeRouterPeers() ComputeRouterPeerInterface
}

// ComputeRouterPeerInterface has methods to work with ComputeRouterPeer resources.
type ComputeRouterPeerInterface interface {
	Create(*v1alpha1.ComputeRouterPeer) (*v1alpha1.ComputeRouterPeer, error)
	Update(*v1alpha1.ComputeRouterPeer) (*v1alpha1.ComputeRouterPeer, error)
	UpdateStatus(*v1alpha1.ComputeRouterPeer) (*v1alpha1.ComputeRouterPeer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeRouterPeer, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeRouterPeerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouterPeer, err error)
	ComputeRouterPeerExpansion
}

// computeRouterPeers implements ComputeRouterPeerInterface
type computeRouterPeers struct {
	client rest.Interface
}

// newComputeRouterPeers returns a ComputeRouterPeers
func newComputeRouterPeers(c *GoogleV1alpha1Client) *computeRouterPeers {
	return &computeRouterPeers{
		client: c.RESTClient(),
	}
}

// Get takes name of the computeRouterPeer, and returns the corresponding computeRouterPeer object, and an error if there is any.
func (c *computeRouterPeers) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRouterPeer, err error) {
	result = &v1alpha1.ComputeRouterPeer{}
	err = c.client.Get().
		Resource("computerouterpeers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeRouterPeers that match those selectors.
func (c *computeRouterPeers) List(opts v1.ListOptions) (result *v1alpha1.ComputeRouterPeerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeRouterPeerList{}
	err = c.client.Get().
		Resource("computerouterpeers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeRouterPeers.
func (c *computeRouterPeers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("computerouterpeers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeRouterPeer and creates it.  Returns the server's representation of the computeRouterPeer, and an error, if there is any.
func (c *computeRouterPeers) Create(computeRouterPeer *v1alpha1.ComputeRouterPeer) (result *v1alpha1.ComputeRouterPeer, err error) {
	result = &v1alpha1.ComputeRouterPeer{}
	err = c.client.Post().
		Resource("computerouterpeers").
		Body(computeRouterPeer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeRouterPeer and updates it. Returns the server's representation of the computeRouterPeer, and an error, if there is any.
func (c *computeRouterPeers) Update(computeRouterPeer *v1alpha1.ComputeRouterPeer) (result *v1alpha1.ComputeRouterPeer, err error) {
	result = &v1alpha1.ComputeRouterPeer{}
	err = c.client.Put().
		Resource("computerouterpeers").
		Name(computeRouterPeer.Name).
		Body(computeRouterPeer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeRouterPeers) UpdateStatus(computeRouterPeer *v1alpha1.ComputeRouterPeer) (result *v1alpha1.ComputeRouterPeer, err error) {
	result = &v1alpha1.ComputeRouterPeer{}
	err = c.client.Put().
		Resource("computerouterpeers").
		Name(computeRouterPeer.Name).
		SubResource("status").
		Body(computeRouterPeer).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeRouterPeer and deletes it. Returns an error if one occurs.
func (c *computeRouterPeers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("computerouterpeers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeRouterPeers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("computerouterpeers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeRouterPeer.
func (c *computeRouterPeers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouterPeer, err error) {
	result = &v1alpha1.ComputeRouterPeer{}
	err = c.client.Patch(pt).
		Resource("computerouterpeers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
