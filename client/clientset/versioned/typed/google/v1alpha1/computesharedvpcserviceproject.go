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

// ComputeSharedVpcServiceProjectsGetter has a method to return a ComputeSharedVpcServiceProjectInterface.
// A group's client should implement this interface.
type ComputeSharedVpcServiceProjectsGetter interface {
	ComputeSharedVpcServiceProjects(namespace string) ComputeSharedVpcServiceProjectInterface
}

// ComputeSharedVpcServiceProjectInterface has methods to work with ComputeSharedVpcServiceProject resources.
type ComputeSharedVpcServiceProjectInterface interface {
	Create(*v1alpha1.ComputeSharedVpcServiceProject) (*v1alpha1.ComputeSharedVpcServiceProject, error)
	Update(*v1alpha1.ComputeSharedVpcServiceProject) (*v1alpha1.ComputeSharedVpcServiceProject, error)
	UpdateStatus(*v1alpha1.ComputeSharedVpcServiceProject) (*v1alpha1.ComputeSharedVpcServiceProject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeSharedVpcServiceProject, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeSharedVpcServiceProjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSharedVpcServiceProject, err error)
	ComputeSharedVpcServiceProjectExpansion
}

// computeSharedVpcServiceProjects implements ComputeSharedVpcServiceProjectInterface
type computeSharedVpcServiceProjects struct {
	client rest.Interface
	ns     string
}

// newComputeSharedVpcServiceProjects returns a ComputeSharedVpcServiceProjects
func newComputeSharedVpcServiceProjects(c *GoogleV1alpha1Client, namespace string) *computeSharedVpcServiceProjects {
	return &computeSharedVpcServiceProjects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeSharedVpcServiceProject, and returns the corresponding computeSharedVpcServiceProject object, and an error if there is any.
func (c *computeSharedVpcServiceProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.ComputeSharedVpcServiceProject{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeSharedVpcServiceProjects that match those selectors.
func (c *computeSharedVpcServiceProjects) List(opts v1.ListOptions) (result *v1alpha1.ComputeSharedVpcServiceProjectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeSharedVpcServiceProjectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeSharedVpcServiceProjects.
func (c *computeSharedVpcServiceProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeSharedVpcServiceProject and creates it.  Returns the server's representation of the computeSharedVpcServiceProject, and an error, if there is any.
func (c *computeSharedVpcServiceProjects) Create(computeSharedVpcServiceProject *v1alpha1.ComputeSharedVpcServiceProject) (result *v1alpha1.ComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.ComputeSharedVpcServiceProject{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		Body(computeSharedVpcServiceProject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeSharedVpcServiceProject and updates it. Returns the server's representation of the computeSharedVpcServiceProject, and an error, if there is any.
func (c *computeSharedVpcServiceProjects) Update(computeSharedVpcServiceProject *v1alpha1.ComputeSharedVpcServiceProject) (result *v1alpha1.ComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.ComputeSharedVpcServiceProject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		Name(computeSharedVpcServiceProject.Name).
		Body(computeSharedVpcServiceProject).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeSharedVpcServiceProjects) UpdateStatus(computeSharedVpcServiceProject *v1alpha1.ComputeSharedVpcServiceProject) (result *v1alpha1.ComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.ComputeSharedVpcServiceProject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		Name(computeSharedVpcServiceProject.Name).
		SubResource("status").
		Body(computeSharedVpcServiceProject).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeSharedVpcServiceProject and deletes it. Returns an error if one occurs.
func (c *computeSharedVpcServiceProjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeSharedVpcServiceProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeSharedVpcServiceProject.
func (c *computeSharedVpcServiceProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSharedVpcServiceProject, err error) {
	result = &v1alpha1.ComputeSharedVpcServiceProject{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computesharedvpcserviceprojects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
