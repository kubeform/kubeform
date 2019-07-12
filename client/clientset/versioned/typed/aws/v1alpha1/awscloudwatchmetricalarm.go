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

// AwsCloudwatchMetricAlarmsGetter has a method to return a AwsCloudwatchMetricAlarmInterface.
// A group's client should implement this interface.
type AwsCloudwatchMetricAlarmsGetter interface {
	AwsCloudwatchMetricAlarms() AwsCloudwatchMetricAlarmInterface
}

// AwsCloudwatchMetricAlarmInterface has methods to work with AwsCloudwatchMetricAlarm resources.
type AwsCloudwatchMetricAlarmInterface interface {
	Create(*v1alpha1.AwsCloudwatchMetricAlarm) (*v1alpha1.AwsCloudwatchMetricAlarm, error)
	Update(*v1alpha1.AwsCloudwatchMetricAlarm) (*v1alpha1.AwsCloudwatchMetricAlarm, error)
	UpdateStatus(*v1alpha1.AwsCloudwatchMetricAlarm) (*v1alpha1.AwsCloudwatchMetricAlarm, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCloudwatchMetricAlarm, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCloudwatchMetricAlarmList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error)
	AwsCloudwatchMetricAlarmExpansion
}

// awsCloudwatchMetricAlarms implements AwsCloudwatchMetricAlarmInterface
type awsCloudwatchMetricAlarms struct {
	client rest.Interface
}

// newAwsCloudwatchMetricAlarms returns a AwsCloudwatchMetricAlarms
func newAwsCloudwatchMetricAlarms(c *AwsV1alpha1Client) *awsCloudwatchMetricAlarms {
	return &awsCloudwatchMetricAlarms{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCloudwatchMetricAlarm, and returns the corresponding awsCloudwatchMetricAlarm object, and an error if there is any.
func (c *awsCloudwatchMetricAlarms) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	result = &v1alpha1.AwsCloudwatchMetricAlarm{}
	err = c.client.Get().
		Resource("awscloudwatchmetricalarms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudwatchMetricAlarms that match those selectors.
func (c *awsCloudwatchMetricAlarms) List(opts v1.ListOptions) (result *v1alpha1.AwsCloudwatchMetricAlarmList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCloudwatchMetricAlarmList{}
	err = c.client.Get().
		Resource("awscloudwatchmetricalarms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchMetricAlarms.
func (c *awsCloudwatchMetricAlarms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscloudwatchmetricalarms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCloudwatchMetricAlarm and creates it.  Returns the server's representation of the awsCloudwatchMetricAlarm, and an error, if there is any.
func (c *awsCloudwatchMetricAlarms) Create(awsCloudwatchMetricAlarm *v1alpha1.AwsCloudwatchMetricAlarm) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	result = &v1alpha1.AwsCloudwatchMetricAlarm{}
	err = c.client.Post().
		Resource("awscloudwatchmetricalarms").
		Body(awsCloudwatchMetricAlarm).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudwatchMetricAlarm and updates it. Returns the server's representation of the awsCloudwatchMetricAlarm, and an error, if there is any.
func (c *awsCloudwatchMetricAlarms) Update(awsCloudwatchMetricAlarm *v1alpha1.AwsCloudwatchMetricAlarm) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	result = &v1alpha1.AwsCloudwatchMetricAlarm{}
	err = c.client.Put().
		Resource("awscloudwatchmetricalarms").
		Name(awsCloudwatchMetricAlarm.Name).
		Body(awsCloudwatchMetricAlarm).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCloudwatchMetricAlarms) UpdateStatus(awsCloudwatchMetricAlarm *v1alpha1.AwsCloudwatchMetricAlarm) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	result = &v1alpha1.AwsCloudwatchMetricAlarm{}
	err = c.client.Put().
		Resource("awscloudwatchmetricalarms").
		Name(awsCloudwatchMetricAlarm.Name).
		SubResource("status").
		Body(awsCloudwatchMetricAlarm).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudwatchMetricAlarm and deletes it. Returns an error if one occurs.
func (c *awsCloudwatchMetricAlarms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscloudwatchmetricalarms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudwatchMetricAlarms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscloudwatchmetricalarms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudwatchMetricAlarm.
func (c *awsCloudwatchMetricAlarms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCloudwatchMetricAlarm, err error) {
	result = &v1alpha1.AwsCloudwatchMetricAlarm{}
	err = c.client.Patch(pt).
		Resource("awscloudwatchmetricalarms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
