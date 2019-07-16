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

// OpsworksApplicationsGetter has a method to return a OpsworksApplicationInterface.
// A group's client should implement this interface.
type OpsworksApplicationsGetter interface {
	OpsworksApplications() OpsworksApplicationInterface
}

// OpsworksApplicationInterface has methods to work with OpsworksApplication resources.
type OpsworksApplicationInterface interface {
	Create(*v1alpha1.OpsworksApplication) (*v1alpha1.OpsworksApplication, error)
	Update(*v1alpha1.OpsworksApplication) (*v1alpha1.OpsworksApplication, error)
	UpdateStatus(*v1alpha1.OpsworksApplication) (*v1alpha1.OpsworksApplication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpsworksApplication, error)
	List(opts v1.ListOptions) (*v1alpha1.OpsworksApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksApplication, err error)
	OpsworksApplicationExpansion
}

// opsworksApplications implements OpsworksApplicationInterface
type opsworksApplications struct {
	client rest.Interface
}

// newOpsworksApplications returns a OpsworksApplications
func newOpsworksApplications(c *AwsV1alpha1Client) *opsworksApplications {
	return &opsworksApplications{
		client: c.RESTClient(),
	}
}

// Get takes name of the opsworksApplication, and returns the corresponding opsworksApplication object, and an error if there is any.
func (c *opsworksApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksApplication, err error) {
	result = &v1alpha1.OpsworksApplication{}
	err = c.client.Get().
		Resource("opsworksapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpsworksApplications that match those selectors.
func (c *opsworksApplications) List(opts v1.ListOptions) (result *v1alpha1.OpsworksApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpsworksApplicationList{}
	err = c.client.Get().
		Resource("opsworksapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested opsworksApplications.
func (c *opsworksApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("opsworksapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a opsworksApplication and creates it.  Returns the server's representation of the opsworksApplication, and an error, if there is any.
func (c *opsworksApplications) Create(opsworksApplication *v1alpha1.OpsworksApplication) (result *v1alpha1.OpsworksApplication, err error) {
	result = &v1alpha1.OpsworksApplication{}
	err = c.client.Post().
		Resource("opsworksapplications").
		Body(opsworksApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a opsworksApplication and updates it. Returns the server's representation of the opsworksApplication, and an error, if there is any.
func (c *opsworksApplications) Update(opsworksApplication *v1alpha1.OpsworksApplication) (result *v1alpha1.OpsworksApplication, err error) {
	result = &v1alpha1.OpsworksApplication{}
	err = c.client.Put().
		Resource("opsworksapplications").
		Name(opsworksApplication.Name).
		Body(opsworksApplication).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *opsworksApplications) UpdateStatus(opsworksApplication *v1alpha1.OpsworksApplication) (result *v1alpha1.OpsworksApplication, err error) {
	result = &v1alpha1.OpsworksApplication{}
	err = c.client.Put().
		Resource("opsworksapplications").
		Name(opsworksApplication.Name).
		SubResource("status").
		Body(opsworksApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the opsworksApplication and deletes it. Returns an error if one occurs.
func (c *opsworksApplications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("opsworksapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *opsworksApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("opsworksapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched opsworksApplication.
func (c *opsworksApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksApplication, err error) {
	result = &v1alpha1.OpsworksApplication{}
	err = c.client.Patch(pt).
		Resource("opsworksapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
