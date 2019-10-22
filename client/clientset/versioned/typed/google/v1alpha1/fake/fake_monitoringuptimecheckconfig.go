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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMonitoringUptimeCheckConfigs implements MonitoringUptimeCheckConfigInterface
type FakeMonitoringUptimeCheckConfigs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var monitoringuptimecheckconfigsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "monitoringuptimecheckconfigs"}

var monitoringuptimecheckconfigsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "MonitoringUptimeCheckConfig"}

// Get takes name of the monitoringUptimeCheckConfig, and returns the corresponding monitoringUptimeCheckConfig object, and an error if there is any.
func (c *FakeMonitoringUptimeCheckConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.MonitoringUptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoringuptimecheckconfigsResource, c.ns, name), &v1alpha1.MonitoringUptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringUptimeCheckConfig), err
}

// List takes label and field selectors, and returns the list of MonitoringUptimeCheckConfigs that match those selectors.
func (c *FakeMonitoringUptimeCheckConfigs) List(opts v1.ListOptions) (result *v1alpha1.MonitoringUptimeCheckConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoringuptimecheckconfigsResource, monitoringuptimecheckconfigsKind, c.ns, opts), &v1alpha1.MonitoringUptimeCheckConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitoringUptimeCheckConfigList{ListMeta: obj.(*v1alpha1.MonitoringUptimeCheckConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitoringUptimeCheckConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitoringUptimeCheckConfigs.
func (c *FakeMonitoringUptimeCheckConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoringuptimecheckconfigsResource, c.ns, opts))

}

// Create takes the representation of a monitoringUptimeCheckConfig and creates it.  Returns the server's representation of the monitoringUptimeCheckConfig, and an error, if there is any.
func (c *FakeMonitoringUptimeCheckConfigs) Create(monitoringUptimeCheckConfig *v1alpha1.MonitoringUptimeCheckConfig) (result *v1alpha1.MonitoringUptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoringuptimecheckconfigsResource, c.ns, monitoringUptimeCheckConfig), &v1alpha1.MonitoringUptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringUptimeCheckConfig), err
}

// Update takes the representation of a monitoringUptimeCheckConfig and updates it. Returns the server's representation of the monitoringUptimeCheckConfig, and an error, if there is any.
func (c *FakeMonitoringUptimeCheckConfigs) Update(monitoringUptimeCheckConfig *v1alpha1.MonitoringUptimeCheckConfig) (result *v1alpha1.MonitoringUptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoringuptimecheckconfigsResource, c.ns, monitoringUptimeCheckConfig), &v1alpha1.MonitoringUptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringUptimeCheckConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitoringUptimeCheckConfigs) UpdateStatus(monitoringUptimeCheckConfig *v1alpha1.MonitoringUptimeCheckConfig) (*v1alpha1.MonitoringUptimeCheckConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoringuptimecheckconfigsResource, "status", c.ns, monitoringUptimeCheckConfig), &v1alpha1.MonitoringUptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringUptimeCheckConfig), err
}

// Delete takes name of the monitoringUptimeCheckConfig and deletes it. Returns an error if one occurs.
func (c *FakeMonitoringUptimeCheckConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitoringuptimecheckconfigsResource, c.ns, name), &v1alpha1.MonitoringUptimeCheckConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitoringUptimeCheckConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoringuptimecheckconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitoringUptimeCheckConfigList{})
	return err
}

// Patch applies the patch and returns the patched monitoringUptimeCheckConfig.
func (c *FakeMonitoringUptimeCheckConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MonitoringUptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoringuptimecheckconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitoringUptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringUptimeCheckConfig), err
}
