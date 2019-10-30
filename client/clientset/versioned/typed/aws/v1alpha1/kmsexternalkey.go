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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KmsExternalKeysGetter has a method to return a KmsExternalKeyInterface.
// A group's client should implement this interface.
type KmsExternalKeysGetter interface {
	KmsExternalKeys(namespace string) KmsExternalKeyInterface
}

// KmsExternalKeyInterface has methods to work with KmsExternalKey resources.
type KmsExternalKeyInterface interface {
	Create(*v1alpha1.KmsExternalKey) (*v1alpha1.KmsExternalKey, error)
	Update(*v1alpha1.KmsExternalKey) (*v1alpha1.KmsExternalKey, error)
	UpdateStatus(*v1alpha1.KmsExternalKey) (*v1alpha1.KmsExternalKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KmsExternalKey, error)
	List(opts v1.ListOptions) (*v1alpha1.KmsExternalKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsExternalKey, err error)
	KmsExternalKeyExpansion
}

// kmsExternalKeys implements KmsExternalKeyInterface
type kmsExternalKeys struct {
	client rest.Interface
	ns     string
}

// newKmsExternalKeys returns a KmsExternalKeys
func newKmsExternalKeys(c *AwsV1alpha1Client, namespace string) *kmsExternalKeys {
	return &kmsExternalKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kmsExternalKey, and returns the corresponding kmsExternalKey object, and an error if there is any.
func (c *kmsExternalKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsExternalKey, err error) {
	result = &v1alpha1.KmsExternalKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KmsExternalKeys that match those selectors.
func (c *kmsExternalKeys) List(opts v1.ListOptions) (result *v1alpha1.KmsExternalKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KmsExternalKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kmsExternalKeys.
func (c *kmsExternalKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kmsExternalKey and creates it.  Returns the server's representation of the kmsExternalKey, and an error, if there is any.
func (c *kmsExternalKeys) Create(kmsExternalKey *v1alpha1.KmsExternalKey) (result *v1alpha1.KmsExternalKey, err error) {
	result = &v1alpha1.KmsExternalKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		Body(kmsExternalKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kmsExternalKey and updates it. Returns the server's representation of the kmsExternalKey, and an error, if there is any.
func (c *kmsExternalKeys) Update(kmsExternalKey *v1alpha1.KmsExternalKey) (result *v1alpha1.KmsExternalKey, err error) {
	result = &v1alpha1.KmsExternalKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		Name(kmsExternalKey.Name).
		Body(kmsExternalKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kmsExternalKeys) UpdateStatus(kmsExternalKey *v1alpha1.KmsExternalKey) (result *v1alpha1.KmsExternalKey, err error) {
	result = &v1alpha1.KmsExternalKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		Name(kmsExternalKey.Name).
		SubResource("status").
		Body(kmsExternalKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the kmsExternalKey and deletes it. Returns an error if one occurs.
func (c *kmsExternalKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kmsExternalKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kmsExternalKey.
func (c *kmsExternalKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsExternalKey, err error) {
	result = &v1alpha1.KmsExternalKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kmsexternalkeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
