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

// AwsEc2TransitGatewayRouteTablesGetter has a method to return a AwsEc2TransitGatewayRouteTableInterface.
// A group's client should implement this interface.
type AwsEc2TransitGatewayRouteTablesGetter interface {
	AwsEc2TransitGatewayRouteTables() AwsEc2TransitGatewayRouteTableInterface
}

// AwsEc2TransitGatewayRouteTableInterface has methods to work with AwsEc2TransitGatewayRouteTable resources.
type AwsEc2TransitGatewayRouteTableInterface interface {
	Create(*v1alpha1.AwsEc2TransitGatewayRouteTable) (*v1alpha1.AwsEc2TransitGatewayRouteTable, error)
	Update(*v1alpha1.AwsEc2TransitGatewayRouteTable) (*v1alpha1.AwsEc2TransitGatewayRouteTable, error)
	UpdateStatus(*v1alpha1.AwsEc2TransitGatewayRouteTable) (*v1alpha1.AwsEc2TransitGatewayRouteTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsEc2TransitGatewayRouteTable, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsEc2TransitGatewayRouteTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEc2TransitGatewayRouteTable, err error)
	AwsEc2TransitGatewayRouteTableExpansion
}

// awsEc2TransitGatewayRouteTables implements AwsEc2TransitGatewayRouteTableInterface
type awsEc2TransitGatewayRouteTables struct {
	client rest.Interface
}

// newAwsEc2TransitGatewayRouteTables returns a AwsEc2TransitGatewayRouteTables
func newAwsEc2TransitGatewayRouteTables(c *AwsV1alpha1Client) *awsEc2TransitGatewayRouteTables {
	return &awsEc2TransitGatewayRouteTables{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsEc2TransitGatewayRouteTable, and returns the corresponding awsEc2TransitGatewayRouteTable object, and an error if there is any.
func (c *awsEc2TransitGatewayRouteTables) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsEc2TransitGatewayRouteTable, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayRouteTable{}
	err = c.client.Get().
		Resource("awsec2transitgatewayroutetables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsEc2TransitGatewayRouteTables that match those selectors.
func (c *awsEc2TransitGatewayRouteTables) List(opts v1.ListOptions) (result *v1alpha1.AwsEc2TransitGatewayRouteTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsEc2TransitGatewayRouteTableList{}
	err = c.client.Get().
		Resource("awsec2transitgatewayroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsEc2TransitGatewayRouteTables.
func (c *awsEc2TransitGatewayRouteTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsec2transitgatewayroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsEc2TransitGatewayRouteTable and creates it.  Returns the server's representation of the awsEc2TransitGatewayRouteTable, and an error, if there is any.
func (c *awsEc2TransitGatewayRouteTables) Create(awsEc2TransitGatewayRouteTable *v1alpha1.AwsEc2TransitGatewayRouteTable) (result *v1alpha1.AwsEc2TransitGatewayRouteTable, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayRouteTable{}
	err = c.client.Post().
		Resource("awsec2transitgatewayroutetables").
		Body(awsEc2TransitGatewayRouteTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsEc2TransitGatewayRouteTable and updates it. Returns the server's representation of the awsEc2TransitGatewayRouteTable, and an error, if there is any.
func (c *awsEc2TransitGatewayRouteTables) Update(awsEc2TransitGatewayRouteTable *v1alpha1.AwsEc2TransitGatewayRouteTable) (result *v1alpha1.AwsEc2TransitGatewayRouteTable, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayRouteTable{}
	err = c.client.Put().
		Resource("awsec2transitgatewayroutetables").
		Name(awsEc2TransitGatewayRouteTable.Name).
		Body(awsEc2TransitGatewayRouteTable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsEc2TransitGatewayRouteTables) UpdateStatus(awsEc2TransitGatewayRouteTable *v1alpha1.AwsEc2TransitGatewayRouteTable) (result *v1alpha1.AwsEc2TransitGatewayRouteTable, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayRouteTable{}
	err = c.client.Put().
		Resource("awsec2transitgatewayroutetables").
		Name(awsEc2TransitGatewayRouteTable.Name).
		SubResource("status").
		Body(awsEc2TransitGatewayRouteTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsEc2TransitGatewayRouteTable and deletes it. Returns an error if one occurs.
func (c *awsEc2TransitGatewayRouteTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsec2transitgatewayroutetables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsEc2TransitGatewayRouteTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsec2transitgatewayroutetables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsEc2TransitGatewayRouteTable.
func (c *awsEc2TransitGatewayRouteTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsEc2TransitGatewayRouteTable, err error) {
	result = &v1alpha1.AwsEc2TransitGatewayRouteTable{}
	err = c.client.Patch(pt).
		Resource("awsec2transitgatewayroutetables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
