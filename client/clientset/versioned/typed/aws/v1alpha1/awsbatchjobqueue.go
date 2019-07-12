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

// AwsBatchJobQueuesGetter has a method to return a AwsBatchJobQueueInterface.
// A group's client should implement this interface.
type AwsBatchJobQueuesGetter interface {
	AwsBatchJobQueues() AwsBatchJobQueueInterface
}

// AwsBatchJobQueueInterface has methods to work with AwsBatchJobQueue resources.
type AwsBatchJobQueueInterface interface {
	Create(*v1alpha1.AwsBatchJobQueue) (*v1alpha1.AwsBatchJobQueue, error)
	Update(*v1alpha1.AwsBatchJobQueue) (*v1alpha1.AwsBatchJobQueue, error)
	UpdateStatus(*v1alpha1.AwsBatchJobQueue) (*v1alpha1.AwsBatchJobQueue, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsBatchJobQueue, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsBatchJobQueueList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBatchJobQueue, err error)
	AwsBatchJobQueueExpansion
}

// awsBatchJobQueues implements AwsBatchJobQueueInterface
type awsBatchJobQueues struct {
	client rest.Interface
}

// newAwsBatchJobQueues returns a AwsBatchJobQueues
func newAwsBatchJobQueues(c *AwsV1alpha1Client) *awsBatchJobQueues {
	return &awsBatchJobQueues{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsBatchJobQueue, and returns the corresponding awsBatchJobQueue object, and an error if there is any.
func (c *awsBatchJobQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsBatchJobQueue, err error) {
	result = &v1alpha1.AwsBatchJobQueue{}
	err = c.client.Get().
		Resource("awsbatchjobqueues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsBatchJobQueues that match those selectors.
func (c *awsBatchJobQueues) List(opts v1.ListOptions) (result *v1alpha1.AwsBatchJobQueueList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsBatchJobQueueList{}
	err = c.client.Get().
		Resource("awsbatchjobqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsBatchJobQueues.
func (c *awsBatchJobQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsbatchjobqueues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsBatchJobQueue and creates it.  Returns the server's representation of the awsBatchJobQueue, and an error, if there is any.
func (c *awsBatchJobQueues) Create(awsBatchJobQueue *v1alpha1.AwsBatchJobQueue) (result *v1alpha1.AwsBatchJobQueue, err error) {
	result = &v1alpha1.AwsBatchJobQueue{}
	err = c.client.Post().
		Resource("awsbatchjobqueues").
		Body(awsBatchJobQueue).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsBatchJobQueue and updates it. Returns the server's representation of the awsBatchJobQueue, and an error, if there is any.
func (c *awsBatchJobQueues) Update(awsBatchJobQueue *v1alpha1.AwsBatchJobQueue) (result *v1alpha1.AwsBatchJobQueue, err error) {
	result = &v1alpha1.AwsBatchJobQueue{}
	err = c.client.Put().
		Resource("awsbatchjobqueues").
		Name(awsBatchJobQueue.Name).
		Body(awsBatchJobQueue).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsBatchJobQueues) UpdateStatus(awsBatchJobQueue *v1alpha1.AwsBatchJobQueue) (result *v1alpha1.AwsBatchJobQueue, err error) {
	result = &v1alpha1.AwsBatchJobQueue{}
	err = c.client.Put().
		Resource("awsbatchjobqueues").
		Name(awsBatchJobQueue.Name).
		SubResource("status").
		Body(awsBatchJobQueue).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsBatchJobQueue and deletes it. Returns an error if one occurs.
func (c *awsBatchJobQueues) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsbatchjobqueues").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsBatchJobQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsbatchjobqueues").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsBatchJobQueue.
func (c *awsBatchJobQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsBatchJobQueue, err error) {
	result = &v1alpha1.AwsBatchJobQueue{}
	err = c.client.Patch(pt).
		Resource("awsbatchjobqueues").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
