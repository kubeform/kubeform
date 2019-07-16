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

// VpnConnectionsGetter has a method to return a VpnConnectionInterface.
// A group's client should implement this interface.
type VpnConnectionsGetter interface {
	VpnConnections() VpnConnectionInterface
}

// VpnConnectionInterface has methods to work with VpnConnection resources.
type VpnConnectionInterface interface {
	Create(*v1alpha1.VpnConnection) (*v1alpha1.VpnConnection, error)
	Update(*v1alpha1.VpnConnection) (*v1alpha1.VpnConnection, error)
	UpdateStatus(*v1alpha1.VpnConnection) (*v1alpha1.VpnConnection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VpnConnection, error)
	List(opts v1.ListOptions) (*v1alpha1.VpnConnectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnConnection, err error)
	VpnConnectionExpansion
}

// vpnConnections implements VpnConnectionInterface
type vpnConnections struct {
	client rest.Interface
}

// newVpnConnections returns a VpnConnections
func newVpnConnections(c *AwsV1alpha1Client) *vpnConnections {
	return &vpnConnections{
		client: c.RESTClient(),
	}
}

// Get takes name of the vpnConnection, and returns the corresponding vpnConnection object, and an error if there is any.
func (c *vpnConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.VpnConnection, err error) {
	result = &v1alpha1.VpnConnection{}
	err = c.client.Get().
		Resource("vpnconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpnConnections that match those selectors.
func (c *vpnConnections) List(opts v1.ListOptions) (result *v1alpha1.VpnConnectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpnConnectionList{}
	err = c.client.Get().
		Resource("vpnconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpnConnections.
func (c *vpnConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("vpnconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vpnConnection and creates it.  Returns the server's representation of the vpnConnection, and an error, if there is any.
func (c *vpnConnections) Create(vpnConnection *v1alpha1.VpnConnection) (result *v1alpha1.VpnConnection, err error) {
	result = &v1alpha1.VpnConnection{}
	err = c.client.Post().
		Resource("vpnconnections").
		Body(vpnConnection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vpnConnection and updates it. Returns the server's representation of the vpnConnection, and an error, if there is any.
func (c *vpnConnections) Update(vpnConnection *v1alpha1.VpnConnection) (result *v1alpha1.VpnConnection, err error) {
	result = &v1alpha1.VpnConnection{}
	err = c.client.Put().
		Resource("vpnconnections").
		Name(vpnConnection.Name).
		Body(vpnConnection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vpnConnections) UpdateStatus(vpnConnection *v1alpha1.VpnConnection) (result *v1alpha1.VpnConnection, err error) {
	result = &v1alpha1.VpnConnection{}
	err = c.client.Put().
		Resource("vpnconnections").
		Name(vpnConnection.Name).
		SubResource("status").
		Body(vpnConnection).
		Do().
		Into(result)
	return
}

// Delete takes name of the vpnConnection and deletes it. Returns an error if one occurs.
func (c *vpnConnections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("vpnconnections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpnConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("vpnconnections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vpnConnection.
func (c *vpnConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnConnection, err error) {
	result = &v1alpha1.VpnConnection{}
	err = c.client.Patch(pt).
		Resource("vpnconnections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
