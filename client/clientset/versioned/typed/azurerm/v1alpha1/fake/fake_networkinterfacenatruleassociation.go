/*
Copyright The Kubeform Authors.

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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkInterfaceNATRuleAssociations implements NetworkInterfaceNATRuleAssociationInterface
type FakeNetworkInterfaceNATRuleAssociations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networkinterfacenatruleassociationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkinterfacenatruleassociations"}

var networkinterfacenatruleassociationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkInterfaceNATRuleAssociation"}

// Get takes name of the networkInterfaceNATRuleAssociation, and returns the corresponding networkInterfaceNATRuleAssociation object, and an error if there is any.
func (c *FakeNetworkInterfaceNATRuleAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkinterfacenatruleassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceNATRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceNATRuleAssociation), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaceNATRuleAssociations that match those selectors.
func (c *FakeNetworkInterfaceNATRuleAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkinterfacenatruleassociationsResource, networkinterfacenatruleassociationsKind, c.ns, opts), &v1alpha1.NetworkInterfaceNATRuleAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceNATRuleAssociationList{ListMeta: obj.(*v1alpha1.NetworkInterfaceNATRuleAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceNATRuleAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaceNATRuleAssociations.
func (c *FakeNetworkInterfaceNATRuleAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkinterfacenatruleassociationsResource, c.ns, opts))

}

// Create takes the representation of a networkInterfaceNATRuleAssociation and creates it.  Returns the server's representation of the networkInterfaceNATRuleAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceNATRuleAssociations) Create(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.CreateOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkinterfacenatruleassociationsResource, c.ns, networkInterfaceNATRuleAssociation), &v1alpha1.NetworkInterfaceNATRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceNATRuleAssociation), err
}

// Update takes the representation of a networkInterfaceNATRuleAssociation and updates it. Returns the server's representation of the networkInterfaceNATRuleAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceNATRuleAssociations) Update(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkinterfacenatruleassociationsResource, c.ns, networkInterfaceNATRuleAssociation), &v1alpha1.NetworkInterfaceNATRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceNATRuleAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkInterfaceNATRuleAssociations) UpdateStatus(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.UpdateOptions) (*v1alpha1.NetworkInterfaceNATRuleAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkinterfacenatruleassociationsResource, "status", c.ns, networkInterfaceNATRuleAssociation), &v1alpha1.NetworkInterfaceNATRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceNATRuleAssociation), err
}

// Delete takes name of the networkInterfaceNATRuleAssociation and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaceNATRuleAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkinterfacenatruleassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceNATRuleAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaceNATRuleAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkinterfacenatruleassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceNATRuleAssociationList{})
	return err
}

// Patch applies the patch and returns the patched networkInterfaceNATRuleAssociation.
func (c *FakeNetworkInterfaceNATRuleAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkinterfacenatruleassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkInterfaceNATRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceNATRuleAssociation), err
}
