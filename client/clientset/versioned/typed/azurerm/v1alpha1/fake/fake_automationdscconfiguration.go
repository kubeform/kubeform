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

// FakeAutomationDscConfigurations implements AutomationDscConfigurationInterface
type FakeAutomationDscConfigurations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var automationdscconfigurationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "automationdscconfigurations"}

var automationdscconfigurationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AutomationDscConfiguration"}

// Get takes name of the automationDscConfiguration, and returns the corresponding automationDscConfiguration object, and an error if there is any.
func (c *FakeAutomationDscConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.AutomationDscConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(automationdscconfigurationsResource, c.ns, name), &v1alpha1.AutomationDscConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscConfiguration), err
}

// List takes label and field selectors, and returns the list of AutomationDscConfigurations that match those selectors.
func (c *FakeAutomationDscConfigurations) List(opts v1.ListOptions) (result *v1alpha1.AutomationDscConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(automationdscconfigurationsResource, automationdscconfigurationsKind, c.ns, opts), &v1alpha1.AutomationDscConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutomationDscConfigurationList{ListMeta: obj.(*v1alpha1.AutomationDscConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutomationDscConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested automationDscConfigurations.
func (c *FakeAutomationDscConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(automationdscconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a automationDscConfiguration and creates it.  Returns the server's representation of the automationDscConfiguration, and an error, if there is any.
func (c *FakeAutomationDscConfigurations) Create(automationDscConfiguration *v1alpha1.AutomationDscConfiguration) (result *v1alpha1.AutomationDscConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(automationdscconfigurationsResource, c.ns, automationDscConfiguration), &v1alpha1.AutomationDscConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscConfiguration), err
}

// Update takes the representation of a automationDscConfiguration and updates it. Returns the server's representation of the automationDscConfiguration, and an error, if there is any.
func (c *FakeAutomationDscConfigurations) Update(automationDscConfiguration *v1alpha1.AutomationDscConfiguration) (result *v1alpha1.AutomationDscConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(automationdscconfigurationsResource, c.ns, automationDscConfiguration), &v1alpha1.AutomationDscConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutomationDscConfigurations) UpdateStatus(automationDscConfiguration *v1alpha1.AutomationDscConfiguration) (*v1alpha1.AutomationDscConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(automationdscconfigurationsResource, "status", c.ns, automationDscConfiguration), &v1alpha1.AutomationDscConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscConfiguration), err
}

// Delete takes name of the automationDscConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeAutomationDscConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(automationdscconfigurationsResource, c.ns, name), &v1alpha1.AutomationDscConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutomationDscConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(automationdscconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutomationDscConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched automationDscConfiguration.
func (c *FakeAutomationDscConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationDscConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(automationdscconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutomationDscConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscConfiguration), err
}