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

// FakeEgressOnlyInternetGateways implements EgressOnlyInternetGatewayInterface
type FakeEgressOnlyInternetGateways struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var egressonlyinternetgatewaysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "egressonlyinternetgateways"}

var egressonlyinternetgatewaysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EgressOnlyInternetGateway"}

// Get takes name of the egressOnlyInternetGateway, and returns the corresponding egressOnlyInternetGateway object, and an error if there is any.
func (c *FakeEgressOnlyInternetGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.EgressOnlyInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(egressonlyinternetgatewaysResource, c.ns, name), &v1alpha1.EgressOnlyInternetGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EgressOnlyInternetGateway), err
}

// List takes label and field selectors, and returns the list of EgressOnlyInternetGateways that match those selectors.
func (c *FakeEgressOnlyInternetGateways) List(opts v1.ListOptions) (result *v1alpha1.EgressOnlyInternetGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(egressonlyinternetgatewaysResource, egressonlyinternetgatewaysKind, c.ns, opts), &v1alpha1.EgressOnlyInternetGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EgressOnlyInternetGatewayList{ListMeta: obj.(*v1alpha1.EgressOnlyInternetGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.EgressOnlyInternetGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested egressOnlyInternetGateways.
func (c *FakeEgressOnlyInternetGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(egressonlyinternetgatewaysResource, c.ns, opts))

}

// Create takes the representation of a egressOnlyInternetGateway and creates it.  Returns the server's representation of the egressOnlyInternetGateway, and an error, if there is any.
func (c *FakeEgressOnlyInternetGateways) Create(egressOnlyInternetGateway *v1alpha1.EgressOnlyInternetGateway) (result *v1alpha1.EgressOnlyInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(egressonlyinternetgatewaysResource, c.ns, egressOnlyInternetGateway), &v1alpha1.EgressOnlyInternetGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EgressOnlyInternetGateway), err
}

// Update takes the representation of a egressOnlyInternetGateway and updates it. Returns the server's representation of the egressOnlyInternetGateway, and an error, if there is any.
func (c *FakeEgressOnlyInternetGateways) Update(egressOnlyInternetGateway *v1alpha1.EgressOnlyInternetGateway) (result *v1alpha1.EgressOnlyInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(egressonlyinternetgatewaysResource, c.ns, egressOnlyInternetGateway), &v1alpha1.EgressOnlyInternetGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EgressOnlyInternetGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEgressOnlyInternetGateways) UpdateStatus(egressOnlyInternetGateway *v1alpha1.EgressOnlyInternetGateway) (*v1alpha1.EgressOnlyInternetGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(egressonlyinternetgatewaysResource, "status", c.ns, egressOnlyInternetGateway), &v1alpha1.EgressOnlyInternetGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EgressOnlyInternetGateway), err
}

// Delete takes name of the egressOnlyInternetGateway and deletes it. Returns an error if one occurs.
func (c *FakeEgressOnlyInternetGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(egressonlyinternetgatewaysResource, c.ns, name), &v1alpha1.EgressOnlyInternetGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEgressOnlyInternetGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(egressonlyinternetgatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EgressOnlyInternetGatewayList{})
	return err
}

// Patch applies the patch and returns the patched egressOnlyInternetGateway.
func (c *FakeEgressOnlyInternetGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EgressOnlyInternetGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(egressonlyinternetgatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.EgressOnlyInternetGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EgressOnlyInternetGateway), err
}