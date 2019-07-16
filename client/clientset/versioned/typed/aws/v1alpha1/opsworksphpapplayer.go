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

// OpsworksPhpAppLayersGetter has a method to return a OpsworksPhpAppLayerInterface.
// A group's client should implement this interface.
type OpsworksPhpAppLayersGetter interface {
	OpsworksPhpAppLayers() OpsworksPhpAppLayerInterface
}

// OpsworksPhpAppLayerInterface has methods to work with OpsworksPhpAppLayer resources.
type OpsworksPhpAppLayerInterface interface {
	Create(*v1alpha1.OpsworksPhpAppLayer) (*v1alpha1.OpsworksPhpAppLayer, error)
	Update(*v1alpha1.OpsworksPhpAppLayer) (*v1alpha1.OpsworksPhpAppLayer, error)
	UpdateStatus(*v1alpha1.OpsworksPhpAppLayer) (*v1alpha1.OpsworksPhpAppLayer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpsworksPhpAppLayer, error)
	List(opts v1.ListOptions) (*v1alpha1.OpsworksPhpAppLayerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksPhpAppLayer, err error)
	OpsworksPhpAppLayerExpansion
}

// opsworksPhpAppLayers implements OpsworksPhpAppLayerInterface
type opsworksPhpAppLayers struct {
	client rest.Interface
}

// newOpsworksPhpAppLayers returns a OpsworksPhpAppLayers
func newOpsworksPhpAppLayers(c *AwsV1alpha1Client) *opsworksPhpAppLayers {
	return &opsworksPhpAppLayers{
		client: c.RESTClient(),
	}
}

// Get takes name of the opsworksPhpAppLayer, and returns the corresponding opsworksPhpAppLayer object, and an error if there is any.
func (c *opsworksPhpAppLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	result = &v1alpha1.OpsworksPhpAppLayer{}
	err = c.client.Get().
		Resource("opsworksphpapplayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpsworksPhpAppLayers that match those selectors.
func (c *opsworksPhpAppLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksPhpAppLayerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpsworksPhpAppLayerList{}
	err = c.client.Get().
		Resource("opsworksphpapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested opsworksPhpAppLayers.
func (c *opsworksPhpAppLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("opsworksphpapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a opsworksPhpAppLayer and creates it.  Returns the server's representation of the opsworksPhpAppLayer, and an error, if there is any.
func (c *opsworksPhpAppLayers) Create(opsworksPhpAppLayer *v1alpha1.OpsworksPhpAppLayer) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	result = &v1alpha1.OpsworksPhpAppLayer{}
	err = c.client.Post().
		Resource("opsworksphpapplayers").
		Body(opsworksPhpAppLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a opsworksPhpAppLayer and updates it. Returns the server's representation of the opsworksPhpAppLayer, and an error, if there is any.
func (c *opsworksPhpAppLayers) Update(opsworksPhpAppLayer *v1alpha1.OpsworksPhpAppLayer) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	result = &v1alpha1.OpsworksPhpAppLayer{}
	err = c.client.Put().
		Resource("opsworksphpapplayers").
		Name(opsworksPhpAppLayer.Name).
		Body(opsworksPhpAppLayer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *opsworksPhpAppLayers) UpdateStatus(opsworksPhpAppLayer *v1alpha1.OpsworksPhpAppLayer) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	result = &v1alpha1.OpsworksPhpAppLayer{}
	err = c.client.Put().
		Resource("opsworksphpapplayers").
		Name(opsworksPhpAppLayer.Name).
		SubResource("status").
		Body(opsworksPhpAppLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the opsworksPhpAppLayer and deletes it. Returns an error if one occurs.
func (c *opsworksPhpAppLayers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("opsworksphpapplayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *opsworksPhpAppLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("opsworksphpapplayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched opsworksPhpAppLayer.
func (c *opsworksPhpAppLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	result = &v1alpha1.OpsworksPhpAppLayer{}
	err = c.client.Patch(pt).
		Resource("opsworksphpapplayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
