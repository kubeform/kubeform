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

// FakeAzurermAutomationVariableBools implements AzurermAutomationVariableBoolInterface
type FakeAzurermAutomationVariableBools struct {
	Fake *FakeAzurermV1alpha1
}

var azurermautomationvariableboolsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermautomationvariablebools"}

var azurermautomationvariableboolsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermAutomationVariableBool"}

// Get takes name of the azurermAutomationVariableBool, and returns the corresponding azurermAutomationVariableBool object, and an error if there is any.
func (c *FakeAzurermAutomationVariableBools) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermAutomationVariableBool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermautomationvariableboolsResource, name), &v1alpha1.AzurermAutomationVariableBool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationVariableBool), err
}

// List takes label and field selectors, and returns the list of AzurermAutomationVariableBools that match those selectors.
func (c *FakeAzurermAutomationVariableBools) List(opts v1.ListOptions) (result *v1alpha1.AzurermAutomationVariableBoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermautomationvariableboolsResource, azurermautomationvariableboolsKind, opts), &v1alpha1.AzurermAutomationVariableBoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermAutomationVariableBoolList{ListMeta: obj.(*v1alpha1.AzurermAutomationVariableBoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermAutomationVariableBoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermAutomationVariableBools.
func (c *FakeAzurermAutomationVariableBools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermautomationvariableboolsResource, opts))
}

// Create takes the representation of a azurermAutomationVariableBool and creates it.  Returns the server's representation of the azurermAutomationVariableBool, and an error, if there is any.
func (c *FakeAzurermAutomationVariableBools) Create(azurermAutomationVariableBool *v1alpha1.AzurermAutomationVariableBool) (result *v1alpha1.AzurermAutomationVariableBool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermautomationvariableboolsResource, azurermAutomationVariableBool), &v1alpha1.AzurermAutomationVariableBool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationVariableBool), err
}

// Update takes the representation of a azurermAutomationVariableBool and updates it. Returns the server's representation of the azurermAutomationVariableBool, and an error, if there is any.
func (c *FakeAzurermAutomationVariableBools) Update(azurermAutomationVariableBool *v1alpha1.AzurermAutomationVariableBool) (result *v1alpha1.AzurermAutomationVariableBool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermautomationvariableboolsResource, azurermAutomationVariableBool), &v1alpha1.AzurermAutomationVariableBool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationVariableBool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermAutomationVariableBools) UpdateStatus(azurermAutomationVariableBool *v1alpha1.AzurermAutomationVariableBool) (*v1alpha1.AzurermAutomationVariableBool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermautomationvariableboolsResource, "status", azurermAutomationVariableBool), &v1alpha1.AzurermAutomationVariableBool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationVariableBool), err
}

// Delete takes name of the azurermAutomationVariableBool and deletes it. Returns an error if one occurs.
func (c *FakeAzurermAutomationVariableBools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermautomationvariableboolsResource, name), &v1alpha1.AzurermAutomationVariableBool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermAutomationVariableBools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermautomationvariableboolsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermAutomationVariableBoolList{})
	return err
}

// Patch applies the patch and returns the patched azurermAutomationVariableBool.
func (c *FakeAzurermAutomationVariableBools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermAutomationVariableBool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermautomationvariableboolsResource, name, pt, data, subresources...), &v1alpha1.AzurermAutomationVariableBool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermAutomationVariableBool), err
}
