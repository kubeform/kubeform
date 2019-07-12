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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsInternetGateways implements AwsInternetGatewayInterface
type FakeAwsInternetGateways struct {
	Fake *FakeAwsV1alpha1
}

var awsinternetgatewaysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsinternetgateways"}

var awsinternetgatewaysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsInternetGateway"}

// Get takes name of the awsInternetGateway, and returns the corresponding awsInternetGateway object, and an error if there is any.
func (c *FakeAwsInternetGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsinternetgatewaysResource, name), &v1alpha1.AwsInternetGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInternetGateway), err
}

// List takes label and field selectors, and returns the list of AwsInternetGateways that match those selectors.
func (c *FakeAwsInternetGateways) List(opts v1.ListOptions) (result *v1alpha1.AwsInternetGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsinternetgatewaysResource, awsinternetgatewaysKind, opts), &v1alpha1.AwsInternetGatewayList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsInternetGatewayList{ListMeta: obj.(*v1alpha1.AwsInternetGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsInternetGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsInternetGateways.
func (c *FakeAwsInternetGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsinternetgatewaysResource, opts))
}

// Create takes the representation of a awsInternetGateway and creates it.  Returns the server's representation of the awsInternetGateway, and an error, if there is any.
func (c *FakeAwsInternetGateways) Create(awsInternetGateway *v1alpha1.AwsInternetGateway) (result *v1alpha1.AwsInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsinternetgatewaysResource, awsInternetGateway), &v1alpha1.AwsInternetGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInternetGateway), err
}

// Update takes the representation of a awsInternetGateway and updates it. Returns the server's representation of the awsInternetGateway, and an error, if there is any.
func (c *FakeAwsInternetGateways) Update(awsInternetGateway *v1alpha1.AwsInternetGateway) (result *v1alpha1.AwsInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsinternetgatewaysResource, awsInternetGateway), &v1alpha1.AwsInternetGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInternetGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsInternetGateways) UpdateStatus(awsInternetGateway *v1alpha1.AwsInternetGateway) (*v1alpha1.AwsInternetGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsinternetgatewaysResource, "status", awsInternetGateway), &v1alpha1.AwsInternetGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInternetGateway), err
}

// Delete takes name of the awsInternetGateway and deletes it. Returns an error if one occurs.
func (c *FakeAwsInternetGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsinternetgatewaysResource, name), &v1alpha1.AwsInternetGateway{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsInternetGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsinternetgatewaysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsInternetGatewayList{})
	return err
}

// Patch applies the patch and returns the patched awsInternetGateway.
func (c *FakeAwsInternetGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsinternetgatewaysResource, name, pt, data, subresources...), &v1alpha1.AwsInternetGateway{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsInternetGateway), err
}
