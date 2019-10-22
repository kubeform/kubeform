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

	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodebalancerConfigsGetter has a method to return a NodebalancerConfigInterface.
// A group's client should implement this interface.
type NodebalancerConfigsGetter interface {
	NodebalancerConfigs(namespace string) NodebalancerConfigInterface
}

// NodebalancerConfigInterface has methods to work with NodebalancerConfig resources.
type NodebalancerConfigInterface interface {
	Create(*v1alpha1.NodebalancerConfig) (*v1alpha1.NodebalancerConfig, error)
	Update(*v1alpha1.NodebalancerConfig) (*v1alpha1.NodebalancerConfig, error)
	UpdateStatus(*v1alpha1.NodebalancerConfig) (*v1alpha1.NodebalancerConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NodebalancerConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.NodebalancerConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodebalancerConfig, err error)
	NodebalancerConfigExpansion
}

// nodebalancerConfigs implements NodebalancerConfigInterface
type nodebalancerConfigs struct {
	client rest.Interface
	ns     string
}

// newNodebalancerConfigs returns a NodebalancerConfigs
func newNodebalancerConfigs(c *LinodeV1alpha1Client, namespace string) *nodebalancerConfigs {
	return &nodebalancerConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodebalancerConfig, and returns the corresponding nodebalancerConfig object, and an error if there is any.
func (c *nodebalancerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.NodebalancerConfig, err error) {
	result = &v1alpha1.NodebalancerConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodebalancerConfigs that match those selectors.
func (c *nodebalancerConfigs) List(opts v1.ListOptions) (result *v1alpha1.NodebalancerConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodebalancerConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodebalancerConfigs.
func (c *nodebalancerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nodebalancerConfig and creates it.  Returns the server's representation of the nodebalancerConfig, and an error, if there is any.
func (c *nodebalancerConfigs) Create(nodebalancerConfig *v1alpha1.NodebalancerConfig) (result *v1alpha1.NodebalancerConfig, err error) {
	result = &v1alpha1.NodebalancerConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		Body(nodebalancerConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodebalancerConfig and updates it. Returns the server's representation of the nodebalancerConfig, and an error, if there is any.
func (c *nodebalancerConfigs) Update(nodebalancerConfig *v1alpha1.NodebalancerConfig) (result *v1alpha1.NodebalancerConfig, err error) {
	result = &v1alpha1.NodebalancerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		Name(nodebalancerConfig.Name).
		Body(nodebalancerConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodebalancerConfigs) UpdateStatus(nodebalancerConfig *v1alpha1.NodebalancerConfig) (result *v1alpha1.NodebalancerConfig, err error) {
	result = &v1alpha1.NodebalancerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		Name(nodebalancerConfig.Name).
		SubResource("status").
		Body(nodebalancerConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodebalancerConfig and deletes it. Returns an error if one occurs.
func (c *nodebalancerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodebalancerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodebalancerConfig.
func (c *nodebalancerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodebalancerConfig, err error) {
	result = &v1alpha1.NodebalancerConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodebalancerconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
