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

// FakeVpcEndpointRouteTableAssociations implements VpcEndpointRouteTableAssociationInterface
type FakeVpcEndpointRouteTableAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpcendpointroutetableassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpcendpointroutetableassociations"}

var vpcendpointroutetableassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpcEndpointRouteTableAssociation"}

// Get takes name of the vpcEndpointRouteTableAssociation, and returns the corresponding vpcEndpointRouteTableAssociation object, and an error if there is any.
func (c *FakeVpcEndpointRouteTableAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.VpcEndpointRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcendpointroutetableassociationsResource, c.ns, name), &v1alpha1.VpcEndpointRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointRouteTableAssociation), err
}

// List takes label and field selectors, and returns the list of VpcEndpointRouteTableAssociations that match those selectors.
func (c *FakeVpcEndpointRouteTableAssociations) List(opts v1.ListOptions) (result *v1alpha1.VpcEndpointRouteTableAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcendpointroutetableassociationsResource, vpcendpointroutetableassociationsKind, c.ns, opts), &v1alpha1.VpcEndpointRouteTableAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcEndpointRouteTableAssociationList{ListMeta: obj.(*v1alpha1.VpcEndpointRouteTableAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcEndpointRouteTableAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcEndpointRouteTableAssociations.
func (c *FakeVpcEndpointRouteTableAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcendpointroutetableassociationsResource, c.ns, opts))

}

// Create takes the representation of a vpcEndpointRouteTableAssociation and creates it.  Returns the server's representation of the vpcEndpointRouteTableAssociation, and an error, if there is any.
func (c *FakeVpcEndpointRouteTableAssociations) Create(vpcEndpointRouteTableAssociation *v1alpha1.VpcEndpointRouteTableAssociation) (result *v1alpha1.VpcEndpointRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcendpointroutetableassociationsResource, c.ns, vpcEndpointRouteTableAssociation), &v1alpha1.VpcEndpointRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointRouteTableAssociation), err
}

// Update takes the representation of a vpcEndpointRouteTableAssociation and updates it. Returns the server's representation of the vpcEndpointRouteTableAssociation, and an error, if there is any.
func (c *FakeVpcEndpointRouteTableAssociations) Update(vpcEndpointRouteTableAssociation *v1alpha1.VpcEndpointRouteTableAssociation) (result *v1alpha1.VpcEndpointRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcendpointroutetableassociationsResource, c.ns, vpcEndpointRouteTableAssociation), &v1alpha1.VpcEndpointRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointRouteTableAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcEndpointRouteTableAssociations) UpdateStatus(vpcEndpointRouteTableAssociation *v1alpha1.VpcEndpointRouteTableAssociation) (*v1alpha1.VpcEndpointRouteTableAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcendpointroutetableassociationsResource, "status", c.ns, vpcEndpointRouteTableAssociation), &v1alpha1.VpcEndpointRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointRouteTableAssociation), err
}

// Delete takes name of the vpcEndpointRouteTableAssociation and deletes it. Returns an error if one occurs.
func (c *FakeVpcEndpointRouteTableAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcendpointroutetableassociationsResource, c.ns, name), &v1alpha1.VpcEndpointRouteTableAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcEndpointRouteTableAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcendpointroutetableassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcEndpointRouteTableAssociationList{})
	return err
}

// Patch applies the patch and returns the patched vpcEndpointRouteTableAssociation.
func (c *FakeVpcEndpointRouteTableAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcEndpointRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcendpointroutetableassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcEndpointRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcEndpointRouteTableAssociation), err
}
