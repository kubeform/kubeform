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

// AwsCodedeployDeploymentGroupsGetter has a method to return a AwsCodedeployDeploymentGroupInterface.
// A group's client should implement this interface.
type AwsCodedeployDeploymentGroupsGetter interface {
	AwsCodedeployDeploymentGroups() AwsCodedeployDeploymentGroupInterface
}

// AwsCodedeployDeploymentGroupInterface has methods to work with AwsCodedeployDeploymentGroup resources.
type AwsCodedeployDeploymentGroupInterface interface {
	Create(*v1alpha1.AwsCodedeployDeploymentGroup) (*v1alpha1.AwsCodedeployDeploymentGroup, error)
	Update(*v1alpha1.AwsCodedeployDeploymentGroup) (*v1alpha1.AwsCodedeployDeploymentGroup, error)
	UpdateStatus(*v1alpha1.AwsCodedeployDeploymentGroup) (*v1alpha1.AwsCodedeployDeploymentGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsCodedeployDeploymentGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsCodedeployDeploymentGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error)
	AwsCodedeployDeploymentGroupExpansion
}

// awsCodedeployDeploymentGroups implements AwsCodedeployDeploymentGroupInterface
type awsCodedeployDeploymentGroups struct {
	client rest.Interface
}

// newAwsCodedeployDeploymentGroups returns a AwsCodedeployDeploymentGroups
func newAwsCodedeployDeploymentGroups(c *AwsV1alpha1Client) *awsCodedeployDeploymentGroups {
	return &awsCodedeployDeploymentGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsCodedeployDeploymentGroup, and returns the corresponding awsCodedeployDeploymentGroup object, and an error if there is any.
func (c *awsCodedeployDeploymentGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	result = &v1alpha1.AwsCodedeployDeploymentGroup{}
	err = c.client.Get().
		Resource("awscodedeploydeploymentgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCodedeployDeploymentGroups that match those selectors.
func (c *awsCodedeployDeploymentGroups) List(opts v1.ListOptions) (result *v1alpha1.AwsCodedeployDeploymentGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsCodedeployDeploymentGroupList{}
	err = c.client.Get().
		Resource("awscodedeploydeploymentgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCodedeployDeploymentGroups.
func (c *awsCodedeployDeploymentGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awscodedeploydeploymentgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsCodedeployDeploymentGroup and creates it.  Returns the server's representation of the awsCodedeployDeploymentGroup, and an error, if there is any.
func (c *awsCodedeployDeploymentGroups) Create(awsCodedeployDeploymentGroup *v1alpha1.AwsCodedeployDeploymentGroup) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	result = &v1alpha1.AwsCodedeployDeploymentGroup{}
	err = c.client.Post().
		Resource("awscodedeploydeploymentgroups").
		Body(awsCodedeployDeploymentGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCodedeployDeploymentGroup and updates it. Returns the server's representation of the awsCodedeployDeploymentGroup, and an error, if there is any.
func (c *awsCodedeployDeploymentGroups) Update(awsCodedeployDeploymentGroup *v1alpha1.AwsCodedeployDeploymentGroup) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	result = &v1alpha1.AwsCodedeployDeploymentGroup{}
	err = c.client.Put().
		Resource("awscodedeploydeploymentgroups").
		Name(awsCodedeployDeploymentGroup.Name).
		Body(awsCodedeployDeploymentGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsCodedeployDeploymentGroups) UpdateStatus(awsCodedeployDeploymentGroup *v1alpha1.AwsCodedeployDeploymentGroup) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	result = &v1alpha1.AwsCodedeployDeploymentGroup{}
	err = c.client.Put().
		Resource("awscodedeploydeploymentgroups").
		Name(awsCodedeployDeploymentGroup.Name).
		SubResource("status").
		Body(awsCodedeployDeploymentGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCodedeployDeploymentGroup and deletes it. Returns an error if one occurs.
func (c *awsCodedeployDeploymentGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awscodedeploydeploymentgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCodedeployDeploymentGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awscodedeploydeploymentgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCodedeployDeploymentGroup.
func (c *awsCodedeployDeploymentGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsCodedeployDeploymentGroup, err error) {
	result = &v1alpha1.AwsCodedeployDeploymentGroup{}
	err = c.client.Patch(pt).
		Resource("awscodedeploydeploymentgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
