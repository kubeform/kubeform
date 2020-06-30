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

// IapAppEngineVersionIamPoliciesGetter has a method to return a IapAppEngineVersionIamPolicyInterface.
// A group's client should implement this interface.
type IapAppEngineVersionIamPoliciesGetter interface {
	IapAppEngineVersionIamPolicies(namespace string) IapAppEngineVersionIamPolicyInterface
}

// IapAppEngineVersionIamPolicyInterface has methods to work with IapAppEngineVersionIamPolicy resources.
type IapAppEngineVersionIamPolicyInterface interface {
	Create(*v1alpha1.IapAppEngineVersionIamPolicy) (*v1alpha1.IapAppEngineVersionIamPolicy, error)
	Update(*v1alpha1.IapAppEngineVersionIamPolicy) (*v1alpha1.IapAppEngineVersionIamPolicy, error)
	UpdateStatus(*v1alpha1.IapAppEngineVersionIamPolicy) (*v1alpha1.IapAppEngineVersionIamPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IapAppEngineVersionIamPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.IapAppEngineVersionIamPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IapAppEngineVersionIamPolicy, err error)
	IapAppEngineVersionIamPolicyExpansion
}

// iapAppEngineVersionIamPolicies implements IapAppEngineVersionIamPolicyInterface
type iapAppEngineVersionIamPolicies struct {
	client rest.Interface
	ns     string
}

// newIapAppEngineVersionIamPolicies returns a IapAppEngineVersionIamPolicies
func newIapAppEngineVersionIamPolicies(c *GoogleV1alpha1Client, namespace string) *iapAppEngineVersionIamPolicies {
	return &iapAppEngineVersionIamPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iapAppEngineVersionIamPolicy, and returns the corresponding iapAppEngineVersionIamPolicy object, and an error if there is any.
func (c *iapAppEngineVersionIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IapAppEngineVersionIamPolicy, err error) {
	result = &v1alpha1.IapAppEngineVersionIamPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IapAppEngineVersionIamPolicies that match those selectors.
func (c *iapAppEngineVersionIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.IapAppEngineVersionIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IapAppEngineVersionIamPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iapAppEngineVersionIamPolicies.
func (c *iapAppEngineVersionIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iapAppEngineVersionIamPolicy and creates it.  Returns the server's representation of the iapAppEngineVersionIamPolicy, and an error, if there is any.
func (c *iapAppEngineVersionIamPolicies) Create(iapAppEngineVersionIamPolicy *v1alpha1.IapAppEngineVersionIamPolicy) (result *v1alpha1.IapAppEngineVersionIamPolicy, err error) {
	result = &v1alpha1.IapAppEngineVersionIamPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		Body(iapAppEngineVersionIamPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iapAppEngineVersionIamPolicy and updates it. Returns the server's representation of the iapAppEngineVersionIamPolicy, and an error, if there is any.
func (c *iapAppEngineVersionIamPolicies) Update(iapAppEngineVersionIamPolicy *v1alpha1.IapAppEngineVersionIamPolicy) (result *v1alpha1.IapAppEngineVersionIamPolicy, err error) {
	result = &v1alpha1.IapAppEngineVersionIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		Name(iapAppEngineVersionIamPolicy.Name).
		Body(iapAppEngineVersionIamPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iapAppEngineVersionIamPolicies) UpdateStatus(iapAppEngineVersionIamPolicy *v1alpha1.IapAppEngineVersionIamPolicy) (result *v1alpha1.IapAppEngineVersionIamPolicy, err error) {
	result = &v1alpha1.IapAppEngineVersionIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		Name(iapAppEngineVersionIamPolicy.Name).
		SubResource("status").
		Body(iapAppEngineVersionIamPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the iapAppEngineVersionIamPolicy and deletes it. Returns an error if one occurs.
func (c *iapAppEngineVersionIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iapAppEngineVersionIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iapAppEngineVersionIamPolicy.
func (c *iapAppEngineVersionIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IapAppEngineVersionIamPolicy, err error) {
	result = &v1alpha1.IapAppEngineVersionIamPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iapappengineversioniampolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
