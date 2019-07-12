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

// FakeAzurermEventgridTopics implements AzurermEventgridTopicInterface
type FakeAzurermEventgridTopics struct {
	Fake *FakeAzurermV1alpha1
}

var azurermeventgridtopicsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "azurermeventgridtopics"}

var azurermeventgridtopicsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzurermEventgridTopic"}

// Get takes name of the azurermEventgridTopic, and returns the corresponding azurermEventgridTopic object, and an error if there is any.
func (c *FakeAzurermEventgridTopics) Get(name string, options v1.GetOptions) (result *v1alpha1.AzurermEventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azurermeventgridtopicsResource, name), &v1alpha1.AzurermEventgridTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventgridTopic), err
}

// List takes label and field selectors, and returns the list of AzurermEventgridTopics that match those selectors.
func (c *FakeAzurermEventgridTopics) List(opts v1.ListOptions) (result *v1alpha1.AzurermEventgridTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azurermeventgridtopicsResource, azurermeventgridtopicsKind, opts), &v1alpha1.AzurermEventgridTopicList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzurermEventgridTopicList{ListMeta: obj.(*v1alpha1.AzurermEventgridTopicList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzurermEventgridTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azurermEventgridTopics.
func (c *FakeAzurermEventgridTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azurermeventgridtopicsResource, opts))
}

// Create takes the representation of a azurermEventgridTopic and creates it.  Returns the server's representation of the azurermEventgridTopic, and an error, if there is any.
func (c *FakeAzurermEventgridTopics) Create(azurermEventgridTopic *v1alpha1.AzurermEventgridTopic) (result *v1alpha1.AzurermEventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azurermeventgridtopicsResource, azurermEventgridTopic), &v1alpha1.AzurermEventgridTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventgridTopic), err
}

// Update takes the representation of a azurermEventgridTopic and updates it. Returns the server's representation of the azurermEventgridTopic, and an error, if there is any.
func (c *FakeAzurermEventgridTopics) Update(azurermEventgridTopic *v1alpha1.AzurermEventgridTopic) (result *v1alpha1.AzurermEventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azurermeventgridtopicsResource, azurermEventgridTopic), &v1alpha1.AzurermEventgridTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventgridTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzurermEventgridTopics) UpdateStatus(azurermEventgridTopic *v1alpha1.AzurermEventgridTopic) (*v1alpha1.AzurermEventgridTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(azurermeventgridtopicsResource, "status", azurermEventgridTopic), &v1alpha1.AzurermEventgridTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventgridTopic), err
}

// Delete takes name of the azurermEventgridTopic and deletes it. Returns an error if one occurs.
func (c *FakeAzurermEventgridTopics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azurermeventgridtopicsResource, name), &v1alpha1.AzurermEventgridTopic{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzurermEventgridTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azurermeventgridtopicsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzurermEventgridTopicList{})
	return err
}

// Patch applies the patch and returns the patched azurermEventgridTopic.
func (c *FakeAzurermEventgridTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzurermEventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azurermeventgridtopicsResource, name, pt, data, subresources...), &v1alpha1.AzurermEventgridTopic{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzurermEventgridTopic), err
}
