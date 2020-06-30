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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRecoveryNetworkMappings implements RecoveryNetworkMappingInterface
type FakeRecoveryNetworkMappings struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var recoverynetworkmappingsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "recoverynetworkmappings"}

var recoverynetworkmappingsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "RecoveryNetworkMapping"}

// Get takes name of the recoveryNetworkMapping, and returns the corresponding recoveryNetworkMapping object, and an error if there is any.
func (c *FakeRecoveryNetworkMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.RecoveryNetworkMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(recoverynetworkmappingsResource, c.ns, name), &v1alpha1.RecoveryNetworkMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecoveryNetworkMapping), err
}

// List takes label and field selectors, and returns the list of RecoveryNetworkMappings that match those selectors.
func (c *FakeRecoveryNetworkMappings) List(opts v1.ListOptions) (result *v1alpha1.RecoveryNetworkMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(recoverynetworkmappingsResource, recoverynetworkmappingsKind, c.ns, opts), &v1alpha1.RecoveryNetworkMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RecoveryNetworkMappingList{ListMeta: obj.(*v1alpha1.RecoveryNetworkMappingList).ListMeta}
	for _, item := range obj.(*v1alpha1.RecoveryNetworkMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested recoveryNetworkMappings.
func (c *FakeRecoveryNetworkMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(recoverynetworkmappingsResource, c.ns, opts))

}

// Create takes the representation of a recoveryNetworkMapping and creates it.  Returns the server's representation of the recoveryNetworkMapping, and an error, if there is any.
func (c *FakeRecoveryNetworkMappings) Create(recoveryNetworkMapping *v1alpha1.RecoveryNetworkMapping) (result *v1alpha1.RecoveryNetworkMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(recoverynetworkmappingsResource, c.ns, recoveryNetworkMapping), &v1alpha1.RecoveryNetworkMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecoveryNetworkMapping), err
}

// Update takes the representation of a recoveryNetworkMapping and updates it. Returns the server's representation of the recoveryNetworkMapping, and an error, if there is any.
func (c *FakeRecoveryNetworkMappings) Update(recoveryNetworkMapping *v1alpha1.RecoveryNetworkMapping) (result *v1alpha1.RecoveryNetworkMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(recoverynetworkmappingsResource, c.ns, recoveryNetworkMapping), &v1alpha1.RecoveryNetworkMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecoveryNetworkMapping), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRecoveryNetworkMappings) UpdateStatus(recoveryNetworkMapping *v1alpha1.RecoveryNetworkMapping) (*v1alpha1.RecoveryNetworkMapping, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(recoverynetworkmappingsResource, "status", c.ns, recoveryNetworkMapping), &v1alpha1.RecoveryNetworkMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecoveryNetworkMapping), err
}

// Delete takes name of the recoveryNetworkMapping and deletes it. Returns an error if one occurs.
func (c *FakeRecoveryNetworkMappings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(recoverynetworkmappingsResource, c.ns, name), &v1alpha1.RecoveryNetworkMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRecoveryNetworkMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(recoverynetworkmappingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RecoveryNetworkMappingList{})
	return err
}

// Patch applies the patch and returns the patched recoveryNetworkMapping.
func (c *FakeRecoveryNetworkMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RecoveryNetworkMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(recoverynetworkmappingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RecoveryNetworkMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecoveryNetworkMapping), err
}
