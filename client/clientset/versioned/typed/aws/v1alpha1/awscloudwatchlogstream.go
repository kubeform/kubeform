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

// AwsCloudwatchLogStreamsGetter has a method to return a AwsCloudwatchLogStreamInterface.
// A group's client should implement this interface.
type AwsCloudwatchLogStreamsGetter interface {
	AwsCloudwatchLogStreams() AwsCloudwatchLogStreamInterface
}

// AwsCloudwatchLogStreamInterface has methods to work with AwsCloudwatchLogStream resources.
type AwsCloudwatchLogStreamInterface interface {
	Create(*v1alpha1.AwsCloudwatchLogStream) (*v1alpha1.AwsCloudwatchLogStream, error)
	Update(*v1alpha1.AwsCloudwatchLogStream) (*v1alpha1.AwsCloudwatchLogStream, error)
	UpdateStatus(*v1alpha1.AwsCloudwatchLogStream) (*v1alpha1.AwsCloudwatchLogStream, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCloudwatchLogStream, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCloudwatchLogStreamList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchLogStream, err error)
	AwsCloudwatchLogStreamExpansion
}

// awsCloudwatchLogStreams implements AwsCloudwatchLogStreamInterface
type awsCloudwatchLogStreams struct {
	client rest.Interface
}

// newAwsCloudwatchLogStreams returns a AwsCloudwatchLogStreams
func newAwsCloudwatchLogStreams(c *AwsV1alpha1Client) *awsCloudwatchLogStreams {
	return &awsCloudwatchLogStreams{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCloudwatchLogStream, and returns the corresponding awsCloudwatchLogStream object, and an error if there is any.
func (c *awsCloudwatchLogStreams) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchLogStream, err error) {
	result = &v1alpha1.AwsCloudwatchLogStream{}
	err = c.client.Get().
		Resource("awscloudwatchlogstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudwatchLogStreams that match those selectors.
func (c *awsCloudwatchLogStreams) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchLogStreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCloudwatchLogStreamList{}
	err = c.client.Get().
		Resource("awscloudwatchlogstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchLogStreams.
func (c *awsCloudwatchLogStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscloudwatchlogstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCloudwatchLogStream and creates it.  Returns the server's representation of the awsCloudwatchLogStream, and an error, if there is any.
func (c *awsCloudwatchLogStreams) Create(awsCloudwatchLogStream *v1alpha1.AwsCloudwatchLogStream) (result *v1alpha1.AwsCloudwatchLogStream, err error) {
	result = &v1alpha1.AwsCloudwatchLogStream{}
	err = c.client.Post().
		Resource("awscloudwatchlogstreams").
		Body(awsCloudwatchLogStream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudwatchLogStream and updates it. Returns the server's representation of the awsCloudwatchLogStream, and an error, if there is any.
func (c *awsCloudwatchLogStreams) Update(awsCloudwatchLogStream *v1alpha1.AwsCloudwatchLogStream) (result *v1alpha1.AwsCloudwatchLogStream, err error) {
	result = &v1alpha1.AwsCloudwatchLogStream{}
	err = c.client.Put().
		Resource("awscloudwatchlogstreams").
		Name(awsCloudwatchLogStream.Name).
		Body(awsCloudwatchLogStream).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCloudwatchLogStreams) UpdateStatus(awsCloudwatchLogStream *v1alpha1.AwsCloudwatchLogStream) (result *v1alpha1.AwsCloudwatchLogStream, err error) {
	result = &v1alpha1.AwsCloudwatchLogStream{}
	err = c.client.Put().
		Resource("awscloudwatchlogstreams").
		Name(awsCloudwatchLogStream.Name).
		SubResource("status").
		Body(awsCloudwatchLogStream).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudwatchLogStream and deletes it. Returns an error if one occurs.
func (c *awsCloudwatchLogStreams) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscloudwatchlogstreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudwatchLogStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscloudwatchlogstreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudwatchLogStream.
func (c *awsCloudwatchLogStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchLogStream, err error) {
	result = &v1alpha1.AwsCloudwatchLogStream{}
	err = c.client.Patch(pt).
		Resource("awscloudwatchlogstreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
