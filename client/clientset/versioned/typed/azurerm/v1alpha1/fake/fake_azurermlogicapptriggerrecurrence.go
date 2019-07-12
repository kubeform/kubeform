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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAzurermLogicAppTriggerRecurrences implements AzurermLogicAppTriggerRecurrenceInterface
type FakeAzurermLogicAppTriggerRecurrences struct {
	Fake *FakeAzurermV1alpha1
}

var azurermlogicapptriggerrecurrencesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermlogicapptriggerrecurrences"}

var azurermlogicapptriggerrecurrencesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermLogicAppTriggerRecurrence"}

// Get takes name of the azurermLogicAppTriggerRecurrence, and returns the corresponding azurermLogicAppTriggerRecurrence object, and an error if there is any.
func (c *FakeAzurermLogicAppTriggerRecurrences) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermlogicapptriggerrecurrencesResource, name), &v1alpha1.AzurermLogicAppTriggerRecurrence{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerRecurrence), err
}

// List takes label and field selectors, and returns the list of AzurermLogicAppTriggerRecurrences that match those selectors.
func (c *FakeAzurermLogicAppTriggerRecurrences) List(opts v1.ListOptions) (result *v1alpha1.AzurermLogicAppTriggerRecurrenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermlogicapptriggerrecurrencesResource, azurermlogicapptriggerrecurrencesKind, opts), &v1alpha1.AzurermLogicAppTriggerRecurrenceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermLogicAppTriggerRecurrenceList{ListMeta: obj.(*v1alpha1.AzurermLogicAppTriggerRecurrenceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermLogicAppTriggerRecurrenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermLogicAppTriggerRecurrences.
func (c *FakeAzurermLogicAppTriggerRecurrences) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermlogicapptriggerrecurrencesResource, opts))
}

// Create takes the representation of a azurermLogicAppTriggerRecurrence and creates it.  Returns the server's representation of the azurermLogicAppTriggerRecurrence, and an error, if there is any.
func (c *FakeAzurermLogicAppTriggerRecurrences) Create(azurermLogicAppTriggerRecurrence *v1alpha1.AzurermLogicAppTriggerRecurrence) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermlogicapptriggerrecurrencesResource, azurermLogicAppTriggerRecurrence), &v1alpha1.AzurermLogicAppTriggerRecurrence{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerRecurrence), err
}

// Update takes the representation of a azurermLogicAppTriggerRecurrence and updates it. Returns the server's representation of the azurermLogicAppTriggerRecurrence, and an error, if there is any.
func (c *FakeAzurermLogicAppTriggerRecurrences) Update(azurermLogicAppTriggerRecurrence *v1alpha1.AzurermLogicAppTriggerRecurrence) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermlogicapptriggerrecurrencesResource, azurermLogicAppTriggerRecurrence), &v1alpha1.AzurermLogicAppTriggerRecurrence{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerRecurrence), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermLogicAppTriggerRecurrences) UpdateStatus(azurermLogicAppTriggerRecurrence *v1alpha1.AzurermLogicAppTriggerRecurrence) (*v1alpha1.AzurermLogicAppTriggerRecurrence, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermlogicapptriggerrecurrencesResource, "status", azurermLogicAppTriggerRecurrence), &v1alpha1.AzurermLogicAppTriggerRecurrence{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerRecurrence), err
}

// Delete takes name of the azurermLogicAppTriggerRecurrence and deletes it. Returns an error if one occurs.
func (c *FakeAzurermLogicAppTriggerRecurrences) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermlogicapptriggerrecurrencesResource, name), &v1alpha1.AzurermLogicAppTriggerRecurrence{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermLogicAppTriggerRecurrences) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermlogicapptriggerrecurrencesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermLogicAppTriggerRecurrenceList{})
	return err
}

// Patch applies the patch and returns the patched azurermLogicAppTriggerRecurrence.
func (c *FakeAzurermLogicAppTriggerRecurrences) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermLogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermlogicapptriggerrecurrencesResource, name, pt, data, subresources...), &v1alpha1.AzurermLogicAppTriggerRecurrence{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermLogicAppTriggerRecurrence), err
}
