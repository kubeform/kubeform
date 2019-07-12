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

// FakeAzurermApplicationInsightsApiKeys implements AzurermApplicationInsightsApiKeyInterface
type FakeAzurermApplicationInsightsApiKeys struct {
	Fake *FakeAzurermV1alpha1
}

var azurermapplicationinsightsapikeysResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermapplicationinsightsapikeys"}

var azurermapplicationinsightsapikeysKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermApplicationInsightsApiKey"}

// Get takes name of the azurermApplicationInsightsApiKey, and returns the corresponding azurermApplicationInsightsApiKey object, and an error if there is any.
func (c *FakeAzurermApplicationInsightsApiKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermapplicationinsightsapikeysResource, name), &v1alpha1.AzurermApplicationInsightsApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsApiKey), err
}

// List takes label and field selectors, and returns the list of AzurermApplicationInsightsApiKeys that match those selectors.
func (c *FakeAzurermApplicationInsightsApiKeys) List(opts v1.ListOptions) (result *v1alpha1.AzurermApplicationInsightsApiKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermapplicationinsightsapikeysResource, azurermapplicationinsightsapikeysKind, opts), &v1alpha1.AzurermApplicationInsightsApiKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermApplicationInsightsApiKeyList{ListMeta: obj.(*v1alpha1.AzurermApplicationInsightsApiKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermApplicationInsightsApiKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermApplicationInsightsApiKeys.
func (c *FakeAzurermApplicationInsightsApiKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermapplicationinsightsapikeysResource, opts))
}

// Create takes the representation of a azurermApplicationInsightsApiKey and creates it.  Returns the server's representation of the azurermApplicationInsightsApiKey, and an error, if there is any.
func (c *FakeAzurermApplicationInsightsApiKeys) Create(azurermApplicationInsightsApiKey *v1alpha1.AzurermApplicationInsightsApiKey) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermapplicationinsightsapikeysResource, azurermApplicationInsightsApiKey), &v1alpha1.AzurermApplicationInsightsApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsApiKey), err
}

// Update takes the representation of a azurermApplicationInsightsApiKey and updates it. Returns the server's representation of the azurermApplicationInsightsApiKey, and an error, if there is any.
func (c *FakeAzurermApplicationInsightsApiKeys) Update(azurermApplicationInsightsApiKey *v1alpha1.AzurermApplicationInsightsApiKey) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermapplicationinsightsapikeysResource, azurermApplicationInsightsApiKey), &v1alpha1.AzurermApplicationInsightsApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsApiKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermApplicationInsightsApiKeys) UpdateStatus(azurermApplicationInsightsApiKey *v1alpha1.AzurermApplicationInsightsApiKey) (*v1alpha1.AzurermApplicationInsightsApiKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermapplicationinsightsapikeysResource, "status", azurermApplicationInsightsApiKey), &v1alpha1.AzurermApplicationInsightsApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsApiKey), err
}

// Delete takes name of the azurermApplicationInsightsApiKey and deletes it. Returns an error if one occurs.
func (c *FakeAzurermApplicationInsightsApiKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermapplicationinsightsapikeysResource, name), &v1alpha1.AzurermApplicationInsightsApiKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermApplicationInsightsApiKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermapplicationinsightsapikeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermApplicationInsightsApiKeyList{})
	return err
}

// Patch applies the patch and returns the patched azurermApplicationInsightsApiKey.
func (c *FakeAzurermApplicationInsightsApiKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermApplicationInsightsApiKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermapplicationinsightsapikeysResource, name, pt, data, subresources...), &v1alpha1.AzurermApplicationInsightsApiKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermApplicationInsightsApiKey), err
}
