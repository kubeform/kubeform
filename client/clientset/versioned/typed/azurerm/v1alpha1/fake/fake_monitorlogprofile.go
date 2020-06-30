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

// FakeMonitorLogProfiles implements MonitorLogProfileInterface
type FakeMonitorLogProfiles struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var monitorlogprofilesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "monitorlogprofiles"}

var monitorlogprofilesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MonitorLogProfile"}

// Get takes name of the monitorLogProfile, and returns the corresponding monitorLogProfile object, and an error if there is any.
func (c *FakeMonitorLogProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MonitorLogProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitorlogprofilesResource, c.ns, name), &v1alpha1.MonitorLogProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorLogProfile), err
}

// List takes label and field selectors, and returns the list of MonitorLogProfiles that match those selectors.
func (c *FakeMonitorLogProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MonitorLogProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitorlogprofilesResource, monitorlogprofilesKind, c.ns, opts), &v1alpha1.MonitorLogProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitorLogProfileList{ListMeta: obj.(*v1alpha1.MonitorLogProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitorLogProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitorLogProfiles.
func (c *FakeMonitorLogProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitorlogprofilesResource, c.ns, opts))

}

// Create takes the representation of a monitorLogProfile and creates it.  Returns the server's representation of the monitorLogProfile, and an error, if there is any.
func (c *FakeMonitorLogProfiles) Create(ctx context.Context, monitorLogProfile *v1alpha1.MonitorLogProfile, opts v1.CreateOptions) (result *v1alpha1.MonitorLogProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitorlogprofilesResource, c.ns, monitorLogProfile), &v1alpha1.MonitorLogProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorLogProfile), err
}

// Update takes the representation of a monitorLogProfile and updates it. Returns the server's representation of the monitorLogProfile, and an error, if there is any.
func (c *FakeMonitorLogProfiles) Update(ctx context.Context, monitorLogProfile *v1alpha1.MonitorLogProfile, opts v1.UpdateOptions) (result *v1alpha1.MonitorLogProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitorlogprofilesResource, c.ns, monitorLogProfile), &v1alpha1.MonitorLogProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorLogProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitorLogProfiles) UpdateStatus(ctx context.Context, monitorLogProfile *v1alpha1.MonitorLogProfile, opts v1.UpdateOptions) (*v1alpha1.MonitorLogProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitorlogprofilesResource, "status", c.ns, monitorLogProfile), &v1alpha1.MonitorLogProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorLogProfile), err
}

// Delete takes name of the monitorLogProfile and deletes it. Returns an error if one occurs.
func (c *FakeMonitorLogProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitorlogprofilesResource, c.ns, name), &v1alpha1.MonitorLogProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitorLogProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitorlogprofilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitorLogProfileList{})
	return err
}

// Patch applies the patch and returns the patched monitorLogProfile.
func (c *FakeMonitorLogProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MonitorLogProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitorlogprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitorLogProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorLogProfile), err
}
