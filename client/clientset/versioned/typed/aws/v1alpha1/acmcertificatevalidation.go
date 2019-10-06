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

// AcmCertificateValidationsGetter has a method to return a AcmCertificateValidationInterface.
// A group's client should implement this interface.
type AcmCertificateValidationsGetter interface {
	AcmCertificateValidations(namespace string) AcmCertificateValidationInterface
}

// AcmCertificateValidationInterface has methods to work with AcmCertificateValidation resources.
type AcmCertificateValidationInterface interface {
	Create(*v1alpha1.AcmCertificateValidation) (*v1alpha1.AcmCertificateValidation, error)
	Update(*v1alpha1.AcmCertificateValidation) (*v1alpha1.AcmCertificateValidation, error)
	UpdateStatus(*v1alpha1.AcmCertificateValidation) (*v1alpha1.AcmCertificateValidation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AcmCertificateValidation, error)
	List(opts v1.ListOptions) (*v1alpha1.AcmCertificateValidationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AcmCertificateValidation, err error)
	AcmCertificateValidationExpansion
}

// acmCertificateValidations implements AcmCertificateValidationInterface
type acmCertificateValidations struct {
	client rest.Interface
	ns     string
}

// newAcmCertificateValidations returns a AcmCertificateValidations
func newAcmCertificateValidations(c *AwsV1alpha1Client, namespace string) *acmCertificateValidations {
	return &acmCertificateValidations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the acmCertificateValidation, and returns the corresponding acmCertificateValidation object, and an error if there is any.
func (c *acmCertificateValidations) Get(name string, options v1.GetOptions) (result *v1alpha1.AcmCertificateValidation, err error) {
	result = &v1alpha1.AcmCertificateValidation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AcmCertificateValidations that match those selectors.
func (c *acmCertificateValidations) List(opts v1.ListOptions) (result *v1alpha1.AcmCertificateValidationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AcmCertificateValidationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested acmCertificateValidations.
func (c *acmCertificateValidations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a acmCertificateValidation and creates it.  Returns the server's representation of the acmCertificateValidation, and an error, if there is any.
func (c *acmCertificateValidations) Create(acmCertificateValidation *v1alpha1.AcmCertificateValidation) (result *v1alpha1.AcmCertificateValidation, err error) {
	result = &v1alpha1.AcmCertificateValidation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		Body(acmCertificateValidation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a acmCertificateValidation and updates it. Returns the server's representation of the acmCertificateValidation, and an error, if there is any.
func (c *acmCertificateValidations) Update(acmCertificateValidation *v1alpha1.AcmCertificateValidation) (result *v1alpha1.AcmCertificateValidation, err error) {
	result = &v1alpha1.AcmCertificateValidation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		Name(acmCertificateValidation.Name).
		Body(acmCertificateValidation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *acmCertificateValidations) UpdateStatus(acmCertificateValidation *v1alpha1.AcmCertificateValidation) (result *v1alpha1.AcmCertificateValidation, err error) {
	result = &v1alpha1.AcmCertificateValidation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		Name(acmCertificateValidation.Name).
		SubResource("status").
		Body(acmCertificateValidation).
		Do().
		Into(result)
	return
}

// Delete takes name of the acmCertificateValidation and deletes it. Returns an error if one occurs.
func (c *acmCertificateValidations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *acmCertificateValidations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched acmCertificateValidation.
func (c *acmCertificateValidations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AcmCertificateValidation, err error) {
	result = &v1alpha1.AcmCertificateValidation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("acmcertificatevalidations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
