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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleServiceAccountIamPoliciesGetter has a method to return a GoogleServiceAccountIamPolicyInterface.
// A group's client should implement this interface.
type GoogleServiceAccountIamPoliciesGetter interface {
	GoogleServiceAccountIamPolicies() GoogleServiceAccountIamPolicyInterface
}

// GoogleServiceAccountIamPolicyInterface has methods to work with GoogleServiceAccountIamPolicy resources.
type GoogleServiceAccountIamPolicyInterface interface {
	Create(*v1alpha1.GoogleServiceAccountIamPolicy) (*v1alpha1.GoogleServiceAccountIamPolicy, error)
	Update(*v1alpha1.GoogleServiceAccountIamPolicy) (*v1alpha1.GoogleServiceAccountIamPolicy, error)
	UpdateStatus(*v1alpha1.GoogleServiceAccountIamPolicy) (*v1alpha1.GoogleServiceAccountIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleServiceAccountIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleServiceAccountIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleServiceAccountIamPolicy, err error)
	GoogleServiceAccountIamPolicyExpansion
}

// googleServiceAccountIamPolicies implements GoogleServiceAccountIamPolicyInterface
type googleServiceAccountIamPolicies struct {
	client rest.Interface
}

// newGoogleServiceAccountIamPolicies returns a GoogleServiceAccountIamPolicies
func newGoogleServiceAccountIamPolicies(c *GoogleV1alpha1Client) *googleServiceAccountIamPolicies {
	return &googleServiceAccountIamPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the googleServiceAccountIamPolicy, and returns the corresponding googleServiceAccountIamPolicy object, and an error if there is any.
func (c *googleServiceAccountIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleServiceAccountIamPolicy, err error) {
	result = &v1alpha1.GoogleServiceAccountIamPolicy{}
	err = c.client.Get().
		Resource("googleserviceaccountiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleServiceAccountIamPolicies that match those selectors.
func (c *googleServiceAccountIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.GoogleServiceAccountIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleServiceAccountIamPolicyList{}
	err = c.client.Get().
		Resource("googleserviceaccountiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleServiceAccountIamPolicies.
func (c *googleServiceAccountIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("googleserviceaccountiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleServiceAccountIamPolicy and creates it.  Returns the server's representation of the googleServiceAccountIamPolicy, and an error, if there is any.
func (c *googleServiceAccountIamPolicies) Create(googleServiceAccountIamPolicy *v1alpha1.GoogleServiceAccountIamPolicy) (result *v1alpha1.GoogleServiceAccountIamPolicy, err error) {
	result = &v1alpha1.GoogleServiceAccountIamPolicy{}
	err = c.client.Post().
		Resource("googleserviceaccountiampolicies").
		Body(googleServiceAccountIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleServiceAccountIamPolicy and updates it. Returns the server's representation of the googleServiceAccountIamPolicy, and an error, if there is any.
func (c *googleServiceAccountIamPolicies) Update(googleServiceAccountIamPolicy *v1alpha1.GoogleServiceAccountIamPolicy) (result *v1alpha1.GoogleServiceAccountIamPolicy, err error) {
	result = &v1alpha1.GoogleServiceAccountIamPolicy{}
	err = c.client.Put().
		Resource("googleserviceaccountiampolicies").
		Name(googleServiceAccountIamPolicy.Name).
		Body(googleServiceAccountIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleServiceAccountIamPolicies) UpdateStatus(googleServiceAccountIamPolicy *v1alpha1.GoogleServiceAccountIamPolicy) (result *v1alpha1.GoogleServiceAccountIamPolicy, err error) {
	result = &v1alpha1.GoogleServiceAccountIamPolicy{}
	err = c.client.Put().
		Resource("googleserviceaccountiampolicies").
		Name(googleServiceAccountIamPolicy.Name).
		SubResource("status").
		Body(googleServiceAccountIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleServiceAccountIamPolicy and deletes it. Returns an error if one occurs.
func (c *googleServiceAccountIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("googleserviceaccountiampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleServiceAccountIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("googleserviceaccountiampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleServiceAccountIamPolicy.
func (c *googleServiceAccountIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleServiceAccountIamPolicy, err error) {
	result = &v1alpha1.GoogleServiceAccountIamPolicy{}
	err = c.client.Patch(pt).
		Resource("googleserviceaccountiampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
