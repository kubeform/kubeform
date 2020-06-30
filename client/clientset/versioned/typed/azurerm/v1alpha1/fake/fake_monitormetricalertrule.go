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

// FakeMonitorMetricAlertrules implements MonitorMetricAlertruleInterface
type FakeMonitorMetricAlertrules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var monitormetricalertrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "monitormetricalertrules"}

var monitormetricalertrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MonitorMetricAlertrule"}

// Get takes name of the monitorMetricAlertrule, and returns the corresponding monitorMetricAlertrule object, and an error if there is any.
func (c *FakeMonitorMetricAlertrules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitormetricalertrulesResource, c.ns, name), &v1alpha1.MonitorMetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorMetricAlertrule), err
}

// List takes label and field selectors, and returns the list of MonitorMetricAlertrules that match those selectors.
func (c *FakeMonitorMetricAlertrules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MonitorMetricAlertruleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitormetricalertrulesResource, monitormetricalertrulesKind, c.ns, opts), &v1alpha1.MonitorMetricAlertruleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitorMetricAlertruleList{ListMeta: obj.(*v1alpha1.MonitorMetricAlertruleList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitorMetricAlertruleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitorMetricAlertrules.
func (c *FakeMonitorMetricAlertrules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitormetricalertrulesResource, c.ns, opts))

}

// Create takes the representation of a monitorMetricAlertrule and creates it.  Returns the server's representation of the monitorMetricAlertrule, and an error, if there is any.
func (c *FakeMonitorMetricAlertrules) Create(ctx context.Context, monitorMetricAlertrule *v1alpha1.MonitorMetricAlertrule, opts v1.CreateOptions) (result *v1alpha1.MonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitormetricalertrulesResource, c.ns, monitorMetricAlertrule), &v1alpha1.MonitorMetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorMetricAlertrule), err
}

// Update takes the representation of a monitorMetricAlertrule and updates it. Returns the server's representation of the monitorMetricAlertrule, and an error, if there is any.
func (c *FakeMonitorMetricAlertrules) Update(ctx context.Context, monitorMetricAlertrule *v1alpha1.MonitorMetricAlertrule, opts v1.UpdateOptions) (result *v1alpha1.MonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitormetricalertrulesResource, c.ns, monitorMetricAlertrule), &v1alpha1.MonitorMetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorMetricAlertrule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitorMetricAlertrules) UpdateStatus(ctx context.Context, monitorMetricAlertrule *v1alpha1.MonitorMetricAlertrule, opts v1.UpdateOptions) (*v1alpha1.MonitorMetricAlertrule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitormetricalertrulesResource, "status", c.ns, monitorMetricAlertrule), &v1alpha1.MonitorMetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorMetricAlertrule), err
}

// Delete takes name of the monitorMetricAlertrule and deletes it. Returns an error if one occurs.
func (c *FakeMonitorMetricAlertrules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitormetricalertrulesResource, c.ns, name), &v1alpha1.MonitorMetricAlertrule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitorMetricAlertrules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitormetricalertrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitorMetricAlertruleList{})
	return err
}

// Patch applies the patch and returns the patched monitorMetricAlertrule.
func (c *FakeMonitorMetricAlertrules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MonitorMetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitormetricalertrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitorMetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorMetricAlertrule), err
}
