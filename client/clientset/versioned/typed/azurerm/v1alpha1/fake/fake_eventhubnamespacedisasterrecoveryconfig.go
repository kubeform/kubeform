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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventhubNamespaceDisasterRecoveryConfigs implements EventhubNamespaceDisasterRecoveryConfigInterface
type FakeEventhubNamespaceDisasterRecoveryConfigs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var eventhubnamespacedisasterrecoveryconfigsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "eventhubnamespacedisasterrecoveryconfigs"}

var eventhubnamespacedisasterrecoveryconfigsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "EventhubNamespaceDisasterRecoveryConfig"}

// Get takes name of the eventhubNamespaceDisasterRecoveryConfig, and returns the corresponding eventhubNamespaceDisasterRecoveryConfig object, and an error if there is any.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.EventhubNamespaceDisasterRecoveryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventhubnamespacedisasterrecoveryconfigsResource, c.ns, name), &v1alpha1.EventhubNamespaceDisasterRecoveryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespaceDisasterRecoveryConfig), err
}

// List takes label and field selectors, and returns the list of EventhubNamespaceDisasterRecoveryConfigs that match those selectors.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) List(opts v1.ListOptions) (result *v1alpha1.EventhubNamespaceDisasterRecoveryConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventhubnamespacedisasterrecoveryconfigsResource, eventhubnamespacedisasterrecoveryconfigsKind, c.ns, opts), &v1alpha1.EventhubNamespaceDisasterRecoveryConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventhubNamespaceDisasterRecoveryConfigList{ListMeta: obj.(*v1alpha1.EventhubNamespaceDisasterRecoveryConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventhubNamespaceDisasterRecoveryConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventhubNamespaceDisasterRecoveryConfigs.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventhubnamespacedisasterrecoveryconfigsResource, c.ns, opts))

}

// Create takes the representation of a eventhubNamespaceDisasterRecoveryConfig and creates it.  Returns the server's representation of the eventhubNamespaceDisasterRecoveryConfig, and an error, if there is any.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) Create(eventhubNamespaceDisasterRecoveryConfig *v1alpha1.EventhubNamespaceDisasterRecoveryConfig) (result *v1alpha1.EventhubNamespaceDisasterRecoveryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventhubnamespacedisasterrecoveryconfigsResource, c.ns, eventhubNamespaceDisasterRecoveryConfig), &v1alpha1.EventhubNamespaceDisasterRecoveryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespaceDisasterRecoveryConfig), err
}

// Update takes the representation of a eventhubNamespaceDisasterRecoveryConfig and updates it. Returns the server's representation of the eventhubNamespaceDisasterRecoveryConfig, and an error, if there is any.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) Update(eventhubNamespaceDisasterRecoveryConfig *v1alpha1.EventhubNamespaceDisasterRecoveryConfig) (result *v1alpha1.EventhubNamespaceDisasterRecoveryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventhubnamespacedisasterrecoveryconfigsResource, c.ns, eventhubNamespaceDisasterRecoveryConfig), &v1alpha1.EventhubNamespaceDisasterRecoveryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespaceDisasterRecoveryConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) UpdateStatus(eventhubNamespaceDisasterRecoveryConfig *v1alpha1.EventhubNamespaceDisasterRecoveryConfig) (*v1alpha1.EventhubNamespaceDisasterRecoveryConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventhubnamespacedisasterrecoveryconfigsResource, "status", c.ns, eventhubNamespaceDisasterRecoveryConfig), &v1alpha1.EventhubNamespaceDisasterRecoveryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespaceDisasterRecoveryConfig), err
}

// Delete takes name of the eventhubNamespaceDisasterRecoveryConfig and deletes it. Returns an error if one occurs.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventhubnamespacedisasterrecoveryconfigsResource, c.ns, name), &v1alpha1.EventhubNamespaceDisasterRecoveryConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventhubnamespacedisasterrecoveryconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventhubNamespaceDisasterRecoveryConfigList{})
	return err
}

// Patch applies the patch and returns the patched eventhubNamespaceDisasterRecoveryConfig.
func (c *FakeEventhubNamespaceDisasterRecoveryConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventhubNamespaceDisasterRecoveryConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventhubnamespacedisasterrecoveryconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventhubNamespaceDisasterRecoveryConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespaceDisasterRecoveryConfig), err
}
