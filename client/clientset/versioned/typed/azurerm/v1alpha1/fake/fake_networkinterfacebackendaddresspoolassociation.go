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

// FakeNetworkInterfaceBackendAddressPoolAssociations implements NetworkInterfaceBackendAddressPoolAssociationInterface
type FakeNetworkInterfaceBackendAddressPoolAssociations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networkinterfacebackendaddresspoolassociationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkinterfacebackendaddresspoolassociations"}

var networkinterfacebackendaddresspoolassociationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkInterfaceBackendAddressPoolAssociation"}

// Get takes name of the networkInterfaceBackendAddressPoolAssociation, and returns the corresponding networkInterfaceBackendAddressPoolAssociation object, and an error if there is any.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkinterfacebackendaddresspoolassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBackendAddressPoolAssociation), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaceBackendAddressPoolAssociations that match those selectors.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceBackendAddressPoolAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkinterfacebackendaddresspoolassociationsResource, networkinterfacebackendaddresspoolassociationsKind, c.ns, opts), &v1alpha1.NetworkInterfaceBackendAddressPoolAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceBackendAddressPoolAssociationList{ListMeta: obj.(*v1alpha1.NetworkInterfaceBackendAddressPoolAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceBackendAddressPoolAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaceBackendAddressPoolAssociations.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkinterfacebackendaddresspoolassociationsResource, c.ns, opts))

}

// Create takes the representation of a networkInterfaceBackendAddressPoolAssociation and creates it.  Returns the server's representation of the networkInterfaceBackendAddressPoolAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) Create(ctx context.Context, networkInterfaceBackendAddressPoolAssociation *v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, opts v1.CreateOptions) (result *v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkinterfacebackendaddresspoolassociationsResource, c.ns, networkInterfaceBackendAddressPoolAssociation), &v1alpha1.NetworkInterfaceBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBackendAddressPoolAssociation), err
}

// Update takes the representation of a networkInterfaceBackendAddressPoolAssociation and updates it. Returns the server's representation of the networkInterfaceBackendAddressPoolAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) Update(ctx context.Context, networkInterfaceBackendAddressPoolAssociation *v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkinterfacebackendaddresspoolassociationsResource, c.ns, networkInterfaceBackendAddressPoolAssociation), &v1alpha1.NetworkInterfaceBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBackendAddressPoolAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) UpdateStatus(ctx context.Context, networkInterfaceBackendAddressPoolAssociation *v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, opts v1.UpdateOptions) (*v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkinterfacebackendaddresspoolassociationsResource, "status", c.ns, networkInterfaceBackendAddressPoolAssociation), &v1alpha1.NetworkInterfaceBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBackendAddressPoolAssociation), err
}

// Delete takes name of the networkInterfaceBackendAddressPoolAssociation and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkinterfacebackendaddresspoolassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceBackendAddressPoolAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkinterfacebackendaddresspoolassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceBackendAddressPoolAssociationList{})
	return err
}

// Patch applies the patch and returns the patched networkInterfaceBackendAddressPoolAssociation.
func (c *FakeNetworkInterfaceBackendAddressPoolAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkInterfaceBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkinterfacebackendaddresspoolassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkInterfaceBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBackendAddressPoolAssociation), err
}
