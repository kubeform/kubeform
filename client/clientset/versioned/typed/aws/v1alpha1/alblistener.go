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
	"context"
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AlbListenersGetter has a method to return a AlbListenerInterface.
// A group's client should implement this interface.
type AlbListenersGetter interface {
	AlbListeners(namespace string) AlbListenerInterface
}

// AlbListenerInterface has methods to work with AlbListener resources.
type AlbListenerInterface interface {
	Create(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.CreateOptions) (*v1alpha1.AlbListener, error)
	Update(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.UpdateOptions) (*v1alpha1.AlbListener, error)
	UpdateStatus(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.UpdateOptions) (*v1alpha1.AlbListener, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AlbListener, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AlbListenerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlbListener, err error)
	AlbListenerExpansion
}

// albListeners implements AlbListenerInterface
type albListeners struct {
	client rest.Interface
	ns     string
}

// newAlbListeners returns a AlbListeners
func newAlbListeners(c *AwsV1alpha1Client, namespace string) *albListeners {
	return &albListeners{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the albListener, and returns the corresponding albListener object, and an error if there is any.
func (c *albListeners) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlbListener, err error) {
	result = &v1alpha1.AlbListener{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alblisteners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AlbListeners that match those selectors.
func (c *albListeners) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlbListenerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AlbListenerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alblisteners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested albListeners.
func (c *albListeners) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alblisteners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a albListener and creates it.  Returns the server's representation of the albListener, and an error, if there is any.
func (c *albListeners) Create(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.CreateOptions) (result *v1alpha1.AlbListener, err error) {
	result = &v1alpha1.AlbListener{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alblisteners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(albListener).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a albListener and updates it. Returns the server's representation of the albListener, and an error, if there is any.
func (c *albListeners) Update(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.UpdateOptions) (result *v1alpha1.AlbListener, err error) {
	result = &v1alpha1.AlbListener{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alblisteners").
		Name(albListener.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(albListener).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *albListeners) UpdateStatus(ctx context.Context, albListener *v1alpha1.AlbListener, opts v1.UpdateOptions) (result *v1alpha1.AlbListener, err error) {
	result = &v1alpha1.AlbListener{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alblisteners").
		Name(albListener.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(albListener).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the albListener and deletes it. Returns an error if one occurs.
func (c *albListeners) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alblisteners").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *albListeners) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alblisteners").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched albListener.
func (c *albListeners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlbListener, err error) {
	result = &v1alpha1.AlbListener{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alblisteners").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
