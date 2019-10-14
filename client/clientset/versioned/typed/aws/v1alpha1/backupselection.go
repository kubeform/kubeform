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

// BackupSelectionsGetter has a method to return a BackupSelectionInterface.
// A group's client should implement this interface.
type BackupSelectionsGetter interface {
	BackupSelections(namespace string) BackupSelectionInterface
}

// BackupSelectionInterface has methods to work with BackupSelection resources.
type BackupSelectionInterface interface {
	Create(*v1alpha1.BackupSelection) (*v1alpha1.BackupSelection, error)
	Update(*v1alpha1.BackupSelection) (*v1alpha1.BackupSelection, error)
	UpdateStatus(*v1alpha1.BackupSelection) (*v1alpha1.BackupSelection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BackupSelection, error)
	List(opts v1.ListOptions) (*v1alpha1.BackupSelectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupSelection, err error)
	BackupSelectionExpansion
}

// backupSelections implements BackupSelectionInterface
type backupSelections struct {
	client rest.Interface
	ns     string
}

// newBackupSelections returns a BackupSelections
func newBackupSelections(c *AwsV1alpha1Client, namespace string) *backupSelections {
	return &backupSelections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupSelection, and returns the corresponding backupSelection object, and an error if there is any.
func (c *backupSelections) Get(name string, options v1.GetOptions) (result *v1alpha1.BackupSelection, err error) {
	result = &v1alpha1.BackupSelection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupselections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupSelections that match those selectors.
func (c *backupSelections) List(opts v1.ListOptions) (result *v1alpha1.BackupSelectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BackupSelectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupselections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupSelections.
func (c *backupSelections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupselections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a backupSelection and creates it.  Returns the server's representation of the backupSelection, and an error, if there is any.
func (c *backupSelections) Create(backupSelection *v1alpha1.BackupSelection) (result *v1alpha1.BackupSelection, err error) {
	result = &v1alpha1.BackupSelection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupselections").
		Body(backupSelection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupSelection and updates it. Returns the server's representation of the backupSelection, and an error, if there is any.
func (c *backupSelections) Update(backupSelection *v1alpha1.BackupSelection) (result *v1alpha1.BackupSelection, err error) {
	result = &v1alpha1.BackupSelection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupselections").
		Name(backupSelection.Name).
		Body(backupSelection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *backupSelections) UpdateStatus(backupSelection *v1alpha1.BackupSelection) (result *v1alpha1.BackupSelection, err error) {
	result = &v1alpha1.BackupSelection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupselections").
		Name(backupSelection.Name).
		SubResource("status").
		Body(backupSelection).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupSelection and deletes it. Returns an error if one occurs.
func (c *backupSelections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupselections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupSelections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupselections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupSelection.
func (c *backupSelections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupSelection, err error) {
	result = &v1alpha1.BackupSelection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupselections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
