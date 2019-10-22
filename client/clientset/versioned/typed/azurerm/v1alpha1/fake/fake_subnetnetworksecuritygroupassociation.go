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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubnetNetworkSecurityGroupAssociations implements SubnetNetworkSecurityGroupAssociationInterface
type FakeSubnetNetworkSecurityGroupAssociations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var subnetnetworksecuritygroupassociationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "subnetnetworksecuritygroupassociations"}

var subnetnetworksecuritygroupassociationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SubnetNetworkSecurityGroupAssociation"}

// Get takes name of the subnetNetworkSecurityGroupAssociation, and returns the corresponding subnetNetworkSecurityGroupAssociation object, and an error if there is any.
func (c *FakeSubnetNetworkSecurityGroupAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.SubnetNetworkSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subnetnetworksecuritygroupassociationsResource, c.ns, name), &v1alpha1.SubnetNetworkSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociation), err
}

// List takes label and field selectors, and returns the list of SubnetNetworkSecurityGroupAssociations that match those selectors.
func (c *FakeSubnetNetworkSecurityGroupAssociations) List(opts v1.ListOptions) (result *v1alpha1.SubnetNetworkSecurityGroupAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subnetnetworksecuritygroupassociationsResource, subnetnetworksecuritygroupassociationsKind, c.ns, opts), &v1alpha1.SubnetNetworkSecurityGroupAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubnetNetworkSecurityGroupAssociationList{ListMeta: obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subnetNetworkSecurityGroupAssociations.
func (c *FakeSubnetNetworkSecurityGroupAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subnetnetworksecuritygroupassociationsResource, c.ns, opts))

}

// Create takes the representation of a subnetNetworkSecurityGroupAssociation and creates it.  Returns the server's representation of the subnetNetworkSecurityGroupAssociation, and an error, if there is any.
func (c *FakeSubnetNetworkSecurityGroupAssociations) Create(subnetNetworkSecurityGroupAssociation *v1alpha1.SubnetNetworkSecurityGroupAssociation) (result *v1alpha1.SubnetNetworkSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subnetnetworksecuritygroupassociationsResource, c.ns, subnetNetworkSecurityGroupAssociation), &v1alpha1.SubnetNetworkSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociation), err
}

// Update takes the representation of a subnetNetworkSecurityGroupAssociation and updates it. Returns the server's representation of the subnetNetworkSecurityGroupAssociation, and an error, if there is any.
func (c *FakeSubnetNetworkSecurityGroupAssociations) Update(subnetNetworkSecurityGroupAssociation *v1alpha1.SubnetNetworkSecurityGroupAssociation) (result *v1alpha1.SubnetNetworkSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subnetnetworksecuritygroupassociationsResource, c.ns, subnetNetworkSecurityGroupAssociation), &v1alpha1.SubnetNetworkSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubnetNetworkSecurityGroupAssociations) UpdateStatus(subnetNetworkSecurityGroupAssociation *v1alpha1.SubnetNetworkSecurityGroupAssociation) (*v1alpha1.SubnetNetworkSecurityGroupAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subnetnetworksecuritygroupassociationsResource, "status", c.ns, subnetNetworkSecurityGroupAssociation), &v1alpha1.SubnetNetworkSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociation), err
}

// Delete takes name of the subnetNetworkSecurityGroupAssociation and deletes it. Returns an error if one occurs.
func (c *FakeSubnetNetworkSecurityGroupAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(subnetnetworksecuritygroupassociationsResource, c.ns, name), &v1alpha1.SubnetNetworkSecurityGroupAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubnetNetworkSecurityGroupAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subnetnetworksecuritygroupassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubnetNetworkSecurityGroupAssociationList{})
	return err
}

// Patch applies the patch and returns the patched subnetNetworkSecurityGroupAssociation.
func (c *FakeSubnetNetworkSecurityGroupAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SubnetNetworkSecurityGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subnetnetworksecuritygroupassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SubnetNetworkSecurityGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkSecurityGroupAssociation), err
}
