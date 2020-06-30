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

// FakeLogicAppTriggerRecurrences implements LogicAppTriggerRecurrenceInterface
type FakeLogicAppTriggerRecurrences struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var logicapptriggerrecurrencesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "logicapptriggerrecurrences"}

var logicapptriggerrecurrencesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogicAppTriggerRecurrence"}

// Get takes name of the logicAppTriggerRecurrence, and returns the corresponding logicAppTriggerRecurrence object, and an error if there is any.
func (c *FakeLogicAppTriggerRecurrences) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logicapptriggerrecurrencesResource, c.ns, name), &v1alpha1.LogicAppTriggerRecurrence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerRecurrence), err
}

// List takes label and field selectors, and returns the list of LogicAppTriggerRecurrences that match those selectors.
func (c *FakeLogicAppTriggerRecurrences) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogicAppTriggerRecurrenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logicapptriggerrecurrencesResource, logicapptriggerrecurrencesKind, c.ns, opts), &v1alpha1.LogicAppTriggerRecurrenceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogicAppTriggerRecurrenceList{ListMeta: obj.(*v1alpha1.LogicAppTriggerRecurrenceList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogicAppTriggerRecurrenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logicAppTriggerRecurrences.
func (c *FakeLogicAppTriggerRecurrences) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logicapptriggerrecurrencesResource, c.ns, opts))

}

// Create takes the representation of a logicAppTriggerRecurrence and creates it.  Returns the server's representation of the logicAppTriggerRecurrence, and an error, if there is any.
func (c *FakeLogicAppTriggerRecurrences) Create(ctx context.Context, logicAppTriggerRecurrence *v1alpha1.LogicAppTriggerRecurrence, opts v1.CreateOptions) (result *v1alpha1.LogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logicapptriggerrecurrencesResource, c.ns, logicAppTriggerRecurrence), &v1alpha1.LogicAppTriggerRecurrence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerRecurrence), err
}

// Update takes the representation of a logicAppTriggerRecurrence and updates it. Returns the server's representation of the logicAppTriggerRecurrence, and an error, if there is any.
func (c *FakeLogicAppTriggerRecurrences) Update(ctx context.Context, logicAppTriggerRecurrence *v1alpha1.LogicAppTriggerRecurrence, opts v1.UpdateOptions) (result *v1alpha1.LogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logicapptriggerrecurrencesResource, c.ns, logicAppTriggerRecurrence), &v1alpha1.LogicAppTriggerRecurrence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerRecurrence), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogicAppTriggerRecurrences) UpdateStatus(ctx context.Context, logicAppTriggerRecurrence *v1alpha1.LogicAppTriggerRecurrence, opts v1.UpdateOptions) (*v1alpha1.LogicAppTriggerRecurrence, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logicapptriggerrecurrencesResource, "status", c.ns, logicAppTriggerRecurrence), &v1alpha1.LogicAppTriggerRecurrence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerRecurrence), err
}

// Delete takes name of the logicAppTriggerRecurrence and deletes it. Returns an error if one occurs.
func (c *FakeLogicAppTriggerRecurrences) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logicapptriggerrecurrencesResource, c.ns, name), &v1alpha1.LogicAppTriggerRecurrence{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogicAppTriggerRecurrences) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logicapptriggerrecurrencesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogicAppTriggerRecurrenceList{})
	return err
}

// Patch applies the patch and returns the patched logicAppTriggerRecurrence.
func (c *FakeLogicAppTriggerRecurrences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogicAppTriggerRecurrence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logicapptriggerrecurrencesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogicAppTriggerRecurrence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerRecurrence), err
}
