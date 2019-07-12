/*
Copyright 2019 The KubeDB Authors.

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

// AwsStoragegatewayCachedIscsiVolumesGetter has a method to return a AwsStoragegatewayCachedIscsiVolumeInterface.
// A group's client should implement this interface.
type AwsStoragegatewayCachedIscsiVolumesGetter interface {
	AwsStoragegatewayCachedIscsiVolumes() AwsStoragegatewayCachedIscsiVolumeInterface
}

// AwsStoragegatewayCachedIscsiVolumeInterface has methods to work with AwsStoragegatewayCachedIscsiVolume resources.
type AwsStoragegatewayCachedIscsiVolumeInterface interface {
	Create(*v1alpha1.AwsStoragegatewayCachedIscsiVolume) (*v1alpha1.AwsStoragegatewayCachedIscsiVolume, error)
	Update(*v1alpha1.AwsStoragegatewayCachedIscsiVolume) (*v1alpha1.AwsStoragegatewayCachedIscsiVolume, error)
	UpdateStatus(*v1alpha1.AwsStoragegatewayCachedIscsiVolume) (*v1alpha1.AwsStoragegatewayCachedIscsiVolume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsStoragegatewayCachedIscsiVolume, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsStoragegatewayCachedIscsiVolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error)
	AwsStoragegatewayCachedIscsiVolumeExpansion
}

// awsStoragegatewayCachedIscsiVolumes implements AwsStoragegatewayCachedIscsiVolumeInterface
type awsStoragegatewayCachedIscsiVolumes struct {
	client rest.Interface
}

// newAwsStoragegatewayCachedIscsiVolumes returns a AwsStoragegatewayCachedIscsiVolumes
func newAwsStoragegatewayCachedIscsiVolumes(c *AwsV1alpha1Client) *awsStoragegatewayCachedIscsiVolumes {
	return &awsStoragegatewayCachedIscsiVolumes{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsStoragegatewayCachedIscsiVolume, and returns the corresponding awsStoragegatewayCachedIscsiVolume object, and an error if there is any.
func (c *awsStoragegatewayCachedIscsiVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	result = &v1alpha1.AwsStoragegatewayCachedIscsiVolume{}
	err = c.client.Get().
		Resource("awsstoragegatewaycachediscsivolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsStoragegatewayCachedIscsiVolumes that match those selectors.
func (c *awsStoragegatewayCachedIscsiVolumes) List(opts v1.ListOptions) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsStoragegatewayCachedIscsiVolumeList{}
	err = c.client.Get().
		Resource("awsstoragegatewaycachediscsivolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsStoragegatewayCachedIscsiVolumes.
func (c *awsStoragegatewayCachedIscsiVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsstoragegatewaycachediscsivolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsStoragegatewayCachedIscsiVolume and creates it.  Returns the server's representation of the awsStoragegatewayCachedIscsiVolume, and an error, if there is any.
func (c *awsStoragegatewayCachedIscsiVolumes) Create(awsStoragegatewayCachedIscsiVolume *v1alpha1.AwsStoragegatewayCachedIscsiVolume) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	result = &v1alpha1.AwsStoragegatewayCachedIscsiVolume{}
	err = c.client.Post().
		Resource("awsstoragegatewaycachediscsivolumes").
		Body(awsStoragegatewayCachedIscsiVolume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsStoragegatewayCachedIscsiVolume and updates it. Returns the server's representation of the awsStoragegatewayCachedIscsiVolume, and an error, if there is any.
func (c *awsStoragegatewayCachedIscsiVolumes) Update(awsStoragegatewayCachedIscsiVolume *v1alpha1.AwsStoragegatewayCachedIscsiVolume) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	result = &v1alpha1.AwsStoragegatewayCachedIscsiVolume{}
	err = c.client.Put().
		Resource("awsstoragegatewaycachediscsivolumes").
		Name(awsStoragegatewayCachedIscsiVolume.Name).
		Body(awsStoragegatewayCachedIscsiVolume).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsStoragegatewayCachedIscsiVolumes) UpdateStatus(awsStoragegatewayCachedIscsiVolume *v1alpha1.AwsStoragegatewayCachedIscsiVolume) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	result = &v1alpha1.AwsStoragegatewayCachedIscsiVolume{}
	err = c.client.Put().
		Resource("awsstoragegatewaycachediscsivolumes").
		Name(awsStoragegatewayCachedIscsiVolume.Name).
		SubResource("status").
		Body(awsStoragegatewayCachedIscsiVolume).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsStoragegatewayCachedIscsiVolume and deletes it. Returns an error if one occurs.
func (c *awsStoragegatewayCachedIscsiVolumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsstoragegatewaycachediscsivolumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsStoragegatewayCachedIscsiVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsstoragegatewaycachediscsivolumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsStoragegatewayCachedIscsiVolume.
func (c *awsStoragegatewayCachedIscsiVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsStoragegatewayCachedIscsiVolume, err error) {
	result = &v1alpha1.AwsStoragegatewayCachedIscsiVolume{}
	err = c.client.Patch(pt).
		Resource("awsstoragegatewaycachediscsivolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
