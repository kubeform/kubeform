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

// FakeAzurermMonitorMetricAlertrules implements AzurermMonitorMetricAlertruleInterface
type FakeAzurermMonitorMetricAlertrules struct {
	Fake *FakeAzurermV1alpha1
}

var azurermmonitormetricalertrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermmonitormetricalertrules"}

var azurermmonitormetricalertrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermMonitorMetricAlertrule"}

// Get takes name of the azurermMonitorMetricAlertrule, and returns the corresponding azurermMonitorMetricAlertrule object, and an error if there is any.
func (c *FakeAzurermMonitorMetricAlertrules) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermMonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermmonitormetricalertrulesResource, name), &v1alpha1.AzurermMonitorMetricAlertrule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorMetricAlertrule), err
}

// List takes label and field selectors, and returns the list of AzurermMonitorMetricAlertrules that match those selectors.
func (c *FakeAzurermMonitorMetricAlertrules) List(opts v1.ListOptions) (result *v1alpha1.AzurermMonitorMetricAlertruleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermmonitormetricalertrulesResource, azurermmonitormetricalertrulesKind, opts), &v1alpha1.AzurermMonitorMetricAlertruleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermMonitorMetricAlertruleList{ListMeta: obj.(*v1alpha1.AzurermMonitorMetricAlertruleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermMonitorMetricAlertruleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermMonitorMetricAlertrules.
func (c *FakeAzurermMonitorMetricAlertrules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermmonitormetricalertrulesResource, opts))
}

// Create takes the representation of a azurermMonitorMetricAlertrule and creates it.  Returns the server's representation of the azurermMonitorMetricAlertrule, and an error, if there is any.
func (c *FakeAzurermMonitorMetricAlertrules) Create(azurermMonitorMetricAlertrule *v1alpha1.AzurermMonitorMetricAlertrule) (result *v1alpha1.AzurermMonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermmonitormetricalertrulesResource, azurermMonitorMetricAlertrule), &v1alpha1.AzurermMonitorMetricAlertrule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorMetricAlertrule), err
}

// Update takes the representation of a azurermMonitorMetricAlertrule and updates it. Returns the server's representation of the azurermMonitorMetricAlertrule, and an error, if there is any.
func (c *FakeAzurermMonitorMetricAlertrules) Update(azurermMonitorMetricAlertrule *v1alpha1.AzurermMonitorMetricAlertrule) (result *v1alpha1.AzurermMonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermmonitormetricalertrulesResource, azurermMonitorMetricAlertrule), &v1alpha1.AzurermMonitorMetricAlertrule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorMetricAlertrule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermMonitorMetricAlertrules) UpdateStatus(azurermMonitorMetricAlertrule *v1alpha1.AzurermMonitorMetricAlertrule) (*v1alpha1.AzurermMonitorMetricAlertrule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermmonitormetricalertrulesResource, "status", azurermMonitorMetricAlertrule), &v1alpha1.AzurermMonitorMetricAlertrule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorMetricAlertrule), err
}

// Delete takes name of the azurermMonitorMetricAlertrule and deletes it. Returns an error if one occurs.
func (c *FakeAzurermMonitorMetricAlertrules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermmonitormetricalertrulesResource, name), &v1alpha1.AzurermMonitorMetricAlertrule{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermMonitorMetricAlertrules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermmonitormetricalertrulesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermMonitorMetricAlertruleList{})
	return err
}

// Patch applies the patch and returns the patched azurermMonitorMetricAlertrule.
func (c *FakeAzurermMonitorMetricAlertrules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermMonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermmonitormetricalertrulesResource, name, pt, data, subresources...), &v1alpha1.AzurermMonitorMetricAlertrule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermMonitorMetricAlertrule), err
}
