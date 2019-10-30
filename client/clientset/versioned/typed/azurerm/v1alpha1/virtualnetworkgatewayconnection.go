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

// VirtualNetworkGatewayConnectionsGetter has a method to return a VirtualNetworkGatewayConnectionInterface.
// A group's client should implement this interface.
type VirtualNetworkGatewayConnectionsGetter interface {
	VirtualNetworkGatewayConnections(namespace string) VirtualNetworkGatewayConnectionInterface
}

// VirtualNetworkGatewayConnectionInterface has methods to work with VirtualNetworkGatewayConnection resources.
type VirtualNetworkGatewayConnectionInterface interface {
	Create(*v1alpha1.VirtualNetworkGatewayConnection) (*v1alpha1.VirtualNetworkGatewayConnection, error)
	Update(*v1alpha1.VirtualNetworkGatewayConnection) (*v1alpha1.VirtualNetworkGatewayConnection, error)
	UpdateStatus(*v1alpha1.VirtualNetworkGatewayConnection) (*v1alpha1.VirtualNetworkGatewayConnection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualNetworkGatewayConnection, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualNetworkGatewayConnectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualNetworkGatewayConnection, err error)
	VirtualNetworkGatewayConnectionExpansion
}

// virtualNetworkGatewayConnections implements VirtualNetworkGatewayConnectionInterface
type virtualNetworkGatewayConnections struct {
	client rest.Interface
	ns     string
}

// newVirtualNetworkGatewayConnections returns a VirtualNetworkGatewayConnections
func newVirtualNetworkGatewayConnections(c *AzurermV1alpha1Client, namespace string) *virtualNetworkGatewayConnections {
	return &virtualNetworkGatewayConnections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualNetworkGatewayConnection, and returns the corresponding virtualNetworkGatewayConnection object, and an error if there is any.
func (c *virtualNetworkGatewayConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.VirtualNetworkGatewayConnection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualNetworkGatewayConnections that match those selectors.
func (c *virtualNetworkGatewayConnections) List(opts v1.ListOptions) (result *v1alpha1.VirtualNetworkGatewayConnectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualNetworkGatewayConnectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualNetworkGatewayConnections.
func (c *virtualNetworkGatewayConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualNetworkGatewayConnection and creates it.  Returns the server's representation of the virtualNetworkGatewayConnection, and an error, if there is any.
func (c *virtualNetworkGatewayConnections) Create(virtualNetworkGatewayConnection *v1alpha1.VirtualNetworkGatewayConnection) (result *v1alpha1.VirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.VirtualNetworkGatewayConnection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		Body(virtualNetworkGatewayConnection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualNetworkGatewayConnection and updates it. Returns the server's representation of the virtualNetworkGatewayConnection, and an error, if there is any.
func (c *virtualNetworkGatewayConnections) Update(virtualNetworkGatewayConnection *v1alpha1.VirtualNetworkGatewayConnection) (result *v1alpha1.VirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.VirtualNetworkGatewayConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		Name(virtualNetworkGatewayConnection.Name).
		Body(virtualNetworkGatewayConnection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualNetworkGatewayConnections) UpdateStatus(virtualNetworkGatewayConnection *v1alpha1.VirtualNetworkGatewayConnection) (result *v1alpha1.VirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.VirtualNetworkGatewayConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		Name(virtualNetworkGatewayConnection.Name).
		SubResource("status").
		Body(virtualNetworkGatewayConnection).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualNetworkGatewayConnection and deletes it. Returns an error if one occurs.
func (c *virtualNetworkGatewayConnections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualNetworkGatewayConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualNetworkGatewayConnection.
func (c *virtualNetworkGatewayConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualNetworkGatewayConnection, err error) {
	result = &v1alpha1.VirtualNetworkGatewayConnection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualnetworkgatewayconnections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
