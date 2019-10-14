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

// SecretsmanagerSecretVersionsGetter has a method to return a SecretsmanagerSecretVersionInterface.
// A group's client should implement this interface.
type SecretsmanagerSecretVersionsGetter interface {
	SecretsmanagerSecretVersions(namespace string) SecretsmanagerSecretVersionInterface
}

// SecretsmanagerSecretVersionInterface has methods to work with SecretsmanagerSecretVersion resources.
type SecretsmanagerSecretVersionInterface interface {
	Create(*v1alpha1.SecretsmanagerSecretVersion) (*v1alpha1.SecretsmanagerSecretVersion, error)
	Update(*v1alpha1.SecretsmanagerSecretVersion) (*v1alpha1.SecretsmanagerSecretVersion, error)
	UpdateStatus(*v1alpha1.SecretsmanagerSecretVersion) (*v1alpha1.SecretsmanagerSecretVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SecretsmanagerSecretVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.SecretsmanagerSecretVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretsmanagerSecretVersion, err error)
	SecretsmanagerSecretVersionExpansion
}

// secretsmanagerSecretVersions implements SecretsmanagerSecretVersionInterface
type secretsmanagerSecretVersions struct {
	client rest.Interface
	ns     string
}

// newSecretsmanagerSecretVersions returns a SecretsmanagerSecretVersions
func newSecretsmanagerSecretVersions(c *AwsV1alpha1Client, namespace string) *secretsmanagerSecretVersions {
	return &secretsmanagerSecretVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secretsmanagerSecretVersion, and returns the corresponding secretsmanagerSecretVersion object, and an error if there is any.
func (c *secretsmanagerSecretVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.SecretsmanagerSecretVersion, err error) {
	result = &v1alpha1.SecretsmanagerSecretVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecretsmanagerSecretVersions that match those selectors.
func (c *secretsmanagerSecretVersions) List(opts v1.ListOptions) (result *v1alpha1.SecretsmanagerSecretVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecretsmanagerSecretVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secretsmanagerSecretVersions.
func (c *secretsmanagerSecretVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a secretsmanagerSecretVersion and creates it.  Returns the server's representation of the secretsmanagerSecretVersion, and an error, if there is any.
func (c *secretsmanagerSecretVersions) Create(secretsmanagerSecretVersion *v1alpha1.SecretsmanagerSecretVersion) (result *v1alpha1.SecretsmanagerSecretVersion, err error) {
	result = &v1alpha1.SecretsmanagerSecretVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		Body(secretsmanagerSecretVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a secretsmanagerSecretVersion and updates it. Returns the server's representation of the secretsmanagerSecretVersion, and an error, if there is any.
func (c *secretsmanagerSecretVersions) Update(secretsmanagerSecretVersion *v1alpha1.SecretsmanagerSecretVersion) (result *v1alpha1.SecretsmanagerSecretVersion, err error) {
	result = &v1alpha1.SecretsmanagerSecretVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		Name(secretsmanagerSecretVersion.Name).
		Body(secretsmanagerSecretVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *secretsmanagerSecretVersions) UpdateStatus(secretsmanagerSecretVersion *v1alpha1.SecretsmanagerSecretVersion) (result *v1alpha1.SecretsmanagerSecretVersion, err error) {
	result = &v1alpha1.SecretsmanagerSecretVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		Name(secretsmanagerSecretVersion.Name).
		SubResource("status").
		Body(secretsmanagerSecretVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the secretsmanagerSecretVersion and deletes it. Returns an error if one occurs.
func (c *secretsmanagerSecretVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secretsmanagerSecretVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched secretsmanagerSecretVersion.
func (c *secretsmanagerSecretVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretsmanagerSecretVersion, err error) {
	result = &v1alpha1.SecretsmanagerSecretVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secretsmanagersecretversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
