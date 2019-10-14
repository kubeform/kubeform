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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzureadServicePrincipalPasswords implements AzureadServicePrincipalPasswordInterface
type FakeAzureadServicePrincipalPasswords struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var azureadserviceprincipalpasswordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azureadserviceprincipalpasswords"}

var azureadserviceprincipalpasswordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzureadServicePrincipalPassword"}

// Get takes name of the azureadServicePrincipalPassword, and returns the corresponding azureadServicePrincipalPassword object, and an error if there is any.
func (c *FakeAzureadServicePrincipalPasswords) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureadServicePrincipalPassword, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureadserviceprincipalpasswordsResource, c.ns, name), &v1alpha1.AzureadServicePrincipalPassword{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadServicePrincipalPassword), err
}

// List takes label and field selectors, and returns the list of AzureadServicePrincipalPasswords that match those selectors.
func (c *FakeAzureadServicePrincipalPasswords) List(opts v1.ListOptions) (result *v1alpha1.AzureadServicePrincipalPasswordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureadserviceprincipalpasswordsResource, azureadserviceprincipalpasswordsKind, c.ns, opts), &v1alpha1.AzureadServicePrincipalPasswordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureadServicePrincipalPasswordList{ListMeta: obj.(*v1alpha1.AzureadServicePrincipalPasswordList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureadServicePrincipalPasswordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureadServicePrincipalPasswords.
func (c *FakeAzureadServicePrincipalPasswords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureadserviceprincipalpasswordsResource, c.ns, opts))

}

// Create takes the representation of a azureadServicePrincipalPassword and creates it.  Returns the server's representation of the azureadServicePrincipalPassword, and an error, if there is any.
func (c *FakeAzureadServicePrincipalPasswords) Create(azureadServicePrincipalPassword *v1alpha1.AzureadServicePrincipalPassword) (result *v1alpha1.AzureadServicePrincipalPassword, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureadserviceprincipalpasswordsResource, c.ns, azureadServicePrincipalPassword), &v1alpha1.AzureadServicePrincipalPassword{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadServicePrincipalPassword), err
}

// Update takes the representation of a azureadServicePrincipalPassword and updates it. Returns the server's representation of the azureadServicePrincipalPassword, and an error, if there is any.
func (c *FakeAzureadServicePrincipalPasswords) Update(azureadServicePrincipalPassword *v1alpha1.AzureadServicePrincipalPassword) (result *v1alpha1.AzureadServicePrincipalPassword, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureadserviceprincipalpasswordsResource, c.ns, azureadServicePrincipalPassword), &v1alpha1.AzureadServicePrincipalPassword{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadServicePrincipalPassword), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureadServicePrincipalPasswords) UpdateStatus(azureadServicePrincipalPassword *v1alpha1.AzureadServicePrincipalPassword) (*v1alpha1.AzureadServicePrincipalPassword, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azureadserviceprincipalpasswordsResource, "status", c.ns, azureadServicePrincipalPassword), &v1alpha1.AzureadServicePrincipalPassword{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadServicePrincipalPassword), err
}

// Delete takes name of the azureadServicePrincipalPassword and deletes it. Returns an error if one occurs.
func (c *FakeAzureadServicePrincipalPasswords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azureadserviceprincipalpasswordsResource, c.ns, name), &v1alpha1.AzureadServicePrincipalPassword{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureadServicePrincipalPasswords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureadserviceprincipalpasswordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureadServicePrincipalPasswordList{})
	return err
}

// Patch applies the patch and returns the patched azureadServicePrincipalPassword.
func (c *FakeAzureadServicePrincipalPasswords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureadServicePrincipalPassword, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureadserviceprincipalpasswordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureadServicePrincipalPassword{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureadServicePrincipalPassword), err
}
