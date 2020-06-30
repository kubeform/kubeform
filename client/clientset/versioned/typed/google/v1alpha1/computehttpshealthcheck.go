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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeHTTPSHealthChecksGetter has a method to return a ComputeHTTPSHealthCheckInterface.
// A group's client should implement this interface.
type ComputeHTTPSHealthChecksGetter interface {
	ComputeHTTPSHealthChecks(namespace string) ComputeHTTPSHealthCheckInterface
}

// ComputeHTTPSHealthCheckInterface has methods to work with ComputeHTTPSHealthCheck resources.
type ComputeHTTPSHealthCheckInterface interface {
	Create(ctx context.Context, computeHTTPSHealthCheck *v1alpha1.ComputeHTTPSHealthCheck, opts v1.CreateOptions) (*v1alpha1.ComputeHTTPSHealthCheck, error)
	Update(ctx context.Context, computeHTTPSHealthCheck *v1alpha1.ComputeHTTPSHealthCheck, opts v1.UpdateOptions) (*v1alpha1.ComputeHTTPSHealthCheck, error)
	UpdateStatus(ctx context.Context, computeHTTPSHealthCheck *v1alpha1.ComputeHTTPSHealthCheck, opts v1.UpdateOptions) (*v1alpha1.ComputeHTTPSHealthCheck, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ComputeHTTPSHealthCheck, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ComputeHTTPSHealthCheckList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeHTTPSHealthCheck, err error)
	ComputeHTTPSHealthCheckExpansion
}

// computeHTTPSHealthChecks implements ComputeHTTPSHealthCheckInterface
type computeHTTPSHealthChecks struct {
	client rest.Interface
	ns     string
}

// newComputeHTTPSHealthChecks returns a ComputeHTTPSHealthChecks
func newComputeHTTPSHealthChecks(c *GoogleV1alpha1Client, namespace string) *computeHTTPSHealthChecks {
	return &computeHTTPSHealthChecks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeHTTPSHealthCheck, and returns the corresponding computeHTTPSHealthCheck object, and an error if there is any.
func (c *computeHTTPSHealthChecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeHTTPSHealthCheck, err error) {
	result = &v1alpha1.ComputeHTTPSHealthCheck{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeHTTPSHealthChecks that match those selectors.
func (c *computeHTTPSHealthChecks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeHTTPSHealthCheckList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeHTTPSHealthCheckList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeHTTPSHealthChecks.
func (c *computeHTTPSHealthChecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeHTTPSHealthCheck and creates it.  Returns the server's representation of the computeHTTPSHealthCheck, and an error, if there is any.
func (c *computeHTTPSHealthChecks) Create(ctx context.Context, computeHTTPSHealthCheck *v1alpha1.ComputeHTTPSHealthCheck, opts v1.CreateOptions) (result *v1alpha1.ComputeHTTPSHealthCheck, err error) {
	result = &v1alpha1.ComputeHTTPSHealthCheck{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeHTTPSHealthCheck).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeHTTPSHealthCheck and updates it. Returns the server's representation of the computeHTTPSHealthCheck, and an error, if there is any.
func (c *computeHTTPSHealthChecks) Update(ctx context.Context, computeHTTPSHealthCheck *v1alpha1.ComputeHTTPSHealthCheck, opts v1.UpdateOptions) (result *v1alpha1.ComputeHTTPSHealthCheck, err error) {
	result = &v1alpha1.ComputeHTTPSHealthCheck{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		Name(computeHTTPSHealthCheck.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeHTTPSHealthCheck).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *computeHTTPSHealthChecks) UpdateStatus(ctx context.Context, computeHTTPSHealthCheck *v1alpha1.ComputeHTTPSHealthCheck, opts v1.UpdateOptions) (result *v1alpha1.ComputeHTTPSHealthCheck, err error) {
	result = &v1alpha1.ComputeHTTPSHealthCheck{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		Name(computeHTTPSHealthCheck.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeHTTPSHealthCheck).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeHTTPSHealthCheck and deletes it. Returns an error if one occurs.
func (c *computeHTTPSHealthChecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeHTTPSHealthChecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeHTTPSHealthCheck.
func (c *computeHTTPSHealthChecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeHTTPSHealthCheck, err error) {
	result = &v1alpha1.ComputeHTTPSHealthCheck{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computehttpshealthchecks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
