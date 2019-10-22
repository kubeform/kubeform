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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppmeshVirtualNodesGetter has a method to return a AppmeshVirtualNodeInterface.
// A group's client should implement this interface.
type AppmeshVirtualNodesGetter interface {
	AppmeshVirtualNodes(namespace string) AppmeshVirtualNodeInterface
}

// AppmeshVirtualNodeInterface has methods to work with AppmeshVirtualNode resources.
type AppmeshVirtualNodeInterface interface {
	Create(*v1alpha1.AppmeshVirtualNode) (*v1alpha1.AppmeshVirtualNode, error)
	Update(*v1alpha1.AppmeshVirtualNode) (*v1alpha1.AppmeshVirtualNode, error)
	UpdateStatus(*v1alpha1.AppmeshVirtualNode) (*v1alpha1.AppmeshVirtualNode, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppmeshVirtualNode, error)
	List(opts v1.ListOptions) (*v1alpha1.AppmeshVirtualNodeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppmeshVirtualNode, err error)
	AppmeshVirtualNodeExpansion
}

// appmeshVirtualNodes implements AppmeshVirtualNodeInterface
type appmeshVirtualNodes struct {
	client rest.Interface
	ns     string
}

// newAppmeshVirtualNodes returns a AppmeshVirtualNodes
func newAppmeshVirtualNodes(c *AwsV1alpha1Client, namespace string) *appmeshVirtualNodes {
	return &appmeshVirtualNodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appmeshVirtualNode, and returns the corresponding appmeshVirtualNode object, and an error if there is any.
func (c *appmeshVirtualNodes) Get(name string, options v1.GetOptions) (result *v1alpha1.AppmeshVirtualNode, err error) {
	result = &v1alpha1.AppmeshVirtualNode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppmeshVirtualNodes that match those selectors.
func (c *appmeshVirtualNodes) List(opts v1.ListOptions) (result *v1alpha1.AppmeshVirtualNodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppmeshVirtualNodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appmeshVirtualNodes.
func (c *appmeshVirtualNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appmeshVirtualNode and creates it.  Returns the server's representation of the appmeshVirtualNode, and an error, if there is any.
func (c *appmeshVirtualNodes) Create(appmeshVirtualNode *v1alpha1.AppmeshVirtualNode) (result *v1alpha1.AppmeshVirtualNode, err error) {
	result = &v1alpha1.AppmeshVirtualNode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		Body(appmeshVirtualNode).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appmeshVirtualNode and updates it. Returns the server's representation of the appmeshVirtualNode, and an error, if there is any.
func (c *appmeshVirtualNodes) Update(appmeshVirtualNode *v1alpha1.AppmeshVirtualNode) (result *v1alpha1.AppmeshVirtualNode, err error) {
	result = &v1alpha1.AppmeshVirtualNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		Name(appmeshVirtualNode.Name).
		Body(appmeshVirtualNode).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appmeshVirtualNodes) UpdateStatus(appmeshVirtualNode *v1alpha1.AppmeshVirtualNode) (result *v1alpha1.AppmeshVirtualNode, err error) {
	result = &v1alpha1.AppmeshVirtualNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		Name(appmeshVirtualNode.Name).
		SubResource("status").
		Body(appmeshVirtualNode).
		Do().
		Into(result)
	return
}

// Delete takes name of the appmeshVirtualNode and deletes it. Returns an error if one occurs.
func (c *appmeshVirtualNodes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appmeshVirtualNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appmeshVirtualNode.
func (c *appmeshVirtualNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppmeshVirtualNode, err error) {
	result = &v1alpha1.AppmeshVirtualNode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appmeshvirtualnodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
