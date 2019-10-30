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

// SsmDocumentsGetter has a method to return a SsmDocumentInterface.
// A group's client should implement this interface.
type SsmDocumentsGetter interface {
	SsmDocuments(namespace string) SsmDocumentInterface
}

// SsmDocumentInterface has methods to work with SsmDocument resources.
type SsmDocumentInterface interface {
	Create(*v1alpha1.SsmDocument) (*v1alpha1.SsmDocument, error)
	Update(*v1alpha1.SsmDocument) (*v1alpha1.SsmDocument, error)
	UpdateStatus(*v1alpha1.SsmDocument) (*v1alpha1.SsmDocument, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SsmDocument, error)
	List(opts v1.ListOptions) (*v1alpha1.SsmDocumentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SsmDocument, err error)
	SsmDocumentExpansion
}

// ssmDocuments implements SsmDocumentInterface
type ssmDocuments struct {
	client rest.Interface
	ns     string
}

// newSsmDocuments returns a SsmDocuments
func newSsmDocuments(c *AwsV1alpha1Client, namespace string) *ssmDocuments {
	return &ssmDocuments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ssmDocument, and returns the corresponding ssmDocument object, and an error if there is any.
func (c *ssmDocuments) Get(name string, options v1.GetOptions) (result *v1alpha1.SsmDocument, err error) {
	result = &v1alpha1.SsmDocument{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ssmdocuments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SsmDocuments that match those selectors.
func (c *ssmDocuments) List(opts v1.ListOptions) (result *v1alpha1.SsmDocumentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SsmDocumentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ssmdocuments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ssmDocuments.
func (c *ssmDocuments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ssmdocuments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ssmDocument and creates it.  Returns the server's representation of the ssmDocument, and an error, if there is any.
func (c *ssmDocuments) Create(ssmDocument *v1alpha1.SsmDocument) (result *v1alpha1.SsmDocument, err error) {
	result = &v1alpha1.SsmDocument{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ssmdocuments").
		Body(ssmDocument).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ssmDocument and updates it. Returns the server's representation of the ssmDocument, and an error, if there is any.
func (c *ssmDocuments) Update(ssmDocument *v1alpha1.SsmDocument) (result *v1alpha1.SsmDocument, err error) {
	result = &v1alpha1.SsmDocument{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ssmdocuments").
		Name(ssmDocument.Name).
		Body(ssmDocument).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ssmDocuments) UpdateStatus(ssmDocument *v1alpha1.SsmDocument) (result *v1alpha1.SsmDocument, err error) {
	result = &v1alpha1.SsmDocument{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ssmdocuments").
		Name(ssmDocument.Name).
		SubResource("status").
		Body(ssmDocument).
		Do().
		Into(result)
	return
}

// Delete takes name of the ssmDocument and deletes it. Returns an error if one occurs.
func (c *ssmDocuments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ssmdocuments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ssmDocuments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ssmdocuments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ssmDocument.
func (c *ssmDocuments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SsmDocument, err error) {
	result = &v1alpha1.SsmDocument{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ssmdocuments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
