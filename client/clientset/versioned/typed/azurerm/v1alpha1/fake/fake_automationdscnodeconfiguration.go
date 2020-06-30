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

// FakeAutomationDscNodeconfigurations implements AutomationDscNodeconfigurationInterface
type FakeAutomationDscNodeconfigurations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var automationdscnodeconfigurationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "automationdscnodeconfigurations"}

var automationdscnodeconfigurationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AutomationDscNodeconfiguration"}

// Get takes name of the automationDscNodeconfiguration, and returns the corresponding automationDscNodeconfiguration object, and an error if there is any.
func (c *FakeAutomationDscNodeconfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(automationdscnodeconfigurationsResource, c.ns, name), &v1alpha1.AutomationDscNodeconfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscNodeconfiguration), err
}

// List takes label and field selectors, and returns the list of AutomationDscNodeconfigurations that match those selectors.
func (c *FakeAutomationDscNodeconfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AutomationDscNodeconfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(automationdscnodeconfigurationsResource, automationdscnodeconfigurationsKind, c.ns, opts), &v1alpha1.AutomationDscNodeconfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutomationDscNodeconfigurationList{ListMeta: obj.(*v1alpha1.AutomationDscNodeconfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutomationDscNodeconfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested automationDscNodeconfigurations.
func (c *FakeAutomationDscNodeconfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(automationdscnodeconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a automationDscNodeconfiguration and creates it.  Returns the server's representation of the automationDscNodeconfiguration, and an error, if there is any.
func (c *FakeAutomationDscNodeconfigurations) Create(ctx context.Context, automationDscNodeconfiguration *v1alpha1.AutomationDscNodeconfiguration, opts v1.CreateOptions) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(automationdscnodeconfigurationsResource, c.ns, automationDscNodeconfiguration), &v1alpha1.AutomationDscNodeconfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscNodeconfiguration), err
}

// Update takes the representation of a automationDscNodeconfiguration and updates it. Returns the server's representation of the automationDscNodeconfiguration, and an error, if there is any.
func (c *FakeAutomationDscNodeconfigurations) Update(ctx context.Context, automationDscNodeconfiguration *v1alpha1.AutomationDscNodeconfiguration, opts v1.UpdateOptions) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(automationdscnodeconfigurationsResource, c.ns, automationDscNodeconfiguration), &v1alpha1.AutomationDscNodeconfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscNodeconfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutomationDscNodeconfigurations) UpdateStatus(ctx context.Context, automationDscNodeconfiguration *v1alpha1.AutomationDscNodeconfiguration, opts v1.UpdateOptions) (*v1alpha1.AutomationDscNodeconfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(automationdscnodeconfigurationsResource, "status", c.ns, automationDscNodeconfiguration), &v1alpha1.AutomationDscNodeconfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscNodeconfiguration), err
}

// Delete takes name of the automationDscNodeconfiguration and deletes it. Returns an error if one occurs.
func (c *FakeAutomationDscNodeconfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(automationdscnodeconfigurationsResource, c.ns, name), &v1alpha1.AutomationDscNodeconfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutomationDscNodeconfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(automationdscnodeconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutomationDscNodeconfigurationList{})
	return err
}

// Patch applies the patch and returns the patched automationDscNodeconfiguration.
func (c *FakeAutomationDscNodeconfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AutomationDscNodeconfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(automationdscnodeconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutomationDscNodeconfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationDscNodeconfiguration), err
}
