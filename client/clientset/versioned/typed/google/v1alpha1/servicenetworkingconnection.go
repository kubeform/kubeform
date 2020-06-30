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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceNetworkingConnectionsGetter has a method to return a ServiceNetworkingConnectionInterface.
// A group's client should implement this interface.
type ServiceNetworkingConnectionsGetter interface {
	ServiceNetworkingConnections(namespace string) ServiceNetworkingConnectionInterface
}

// ServiceNetworkingConnectionInterface has methods to work with ServiceNetworkingConnection resources.
type ServiceNetworkingConnectionInterface interface {
	Create(*v1alpha1.ServiceNetworkingConnection) (*v1alpha1.ServiceNetworkingConnection, error)
	Update(*v1alpha1.ServiceNetworkingConnection) (*v1alpha1.ServiceNetworkingConnection, error)
	UpdateStatus(*v1alpha1.ServiceNetworkingConnection) (*v1alpha1.ServiceNetworkingConnection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceNetworkingConnection, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceNetworkingConnectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceNetworkingConnection, err error)
	ServiceNetworkingConnectionExpansion
}

// serviceNetworkingConnections implements ServiceNetworkingConnectionInterface
type serviceNetworkingConnections struct {
	client rest.Interface
	ns     string
}

// newServiceNetworkingConnections returns a ServiceNetworkingConnections
func newServiceNetworkingConnections(c *GoogleV1alpha1Client, namespace string) *serviceNetworkingConnections {
	return &serviceNetworkingConnections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceNetworkingConnection, and returns the corresponding serviceNetworkingConnection object, and an error if there is any.
func (c *serviceNetworkingConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceNetworkingConnection, err error) {
	result = &v1alpha1.ServiceNetworkingConnection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceNetworkingConnections that match those selectors.
func (c *serviceNetworkingConnections) List(opts v1.ListOptions) (result *v1alpha1.ServiceNetworkingConnectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceNetworkingConnectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceNetworkingConnections.
func (c *serviceNetworkingConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a serviceNetworkingConnection and creates it.  Returns the server's representation of the serviceNetworkingConnection, and an error, if there is any.
func (c *serviceNetworkingConnections) Create(serviceNetworkingConnection *v1alpha1.ServiceNetworkingConnection) (result *v1alpha1.ServiceNetworkingConnection, err error) {
	result = &v1alpha1.ServiceNetworkingConnection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		Body(serviceNetworkingConnection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceNetworkingConnection and updates it. Returns the server's representation of the serviceNetworkingConnection, and an error, if there is any.
func (c *serviceNetworkingConnections) Update(serviceNetworkingConnection *v1alpha1.ServiceNetworkingConnection) (result *v1alpha1.ServiceNetworkingConnection, err error) {
	result = &v1alpha1.ServiceNetworkingConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		Name(serviceNetworkingConnection.Name).
		Body(serviceNetworkingConnection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *serviceNetworkingConnections) UpdateStatus(serviceNetworkingConnection *v1alpha1.ServiceNetworkingConnection) (result *v1alpha1.ServiceNetworkingConnection, err error) {
	result = &v1alpha1.ServiceNetworkingConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		Name(serviceNetworkingConnection.Name).
		SubResource("status").
		Body(serviceNetworkingConnection).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceNetworkingConnection and deletes it. Returns an error if one occurs.
func (c *serviceNetworkingConnections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceNetworkingConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceNetworkingConnection.
func (c *serviceNetworkingConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceNetworkingConnection, err error) {
	result = &v1alpha1.ServiceNetworkingConnection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicenetworkingconnections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
