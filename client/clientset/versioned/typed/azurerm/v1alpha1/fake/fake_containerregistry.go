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

// FakeContainerRegistries implements ContainerRegistryInterface
type FakeContainerRegistries struct {
	Fake *FakeAzurermV1alpha1
}

var containerregistriesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "containerregistries"}

var containerregistriesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ContainerRegistry"}

// Get takes name of the containerRegistry, and returns the corresponding containerRegistry object, and an error if there is any.
func (c *FakeContainerRegistries) Get(name string, options v1.GetOptions) (result *v1alpha1.ContainerRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(containerregistriesResource, name), &v1alpha1.ContainerRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRegistry), err
}

// List takes label and field selectors, and returns the list of ContainerRegistries that match those selectors.
func (c *FakeContainerRegistries) List(opts v1.ListOptions) (result *v1alpha1.ContainerRegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(containerregistriesResource, containerregistriesKind, opts), &v1alpha1.ContainerRegistryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ContainerRegistryList{ListMeta: obj.(*v1alpha1.ContainerRegistryList).ListMeta}
	for _, item := range obj.(*v1alpha1.ContainerRegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerRegistries.
func (c *FakeContainerRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(containerregistriesResource, opts))
}

// Create takes the representation of a containerRegistry and creates it.  Returns the server's representation of the containerRegistry, and an error, if there is any.
func (c *FakeContainerRegistries) Create(containerRegistry *v1alpha1.ContainerRegistry) (result *v1alpha1.ContainerRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(containerregistriesResource, containerRegistry), &v1alpha1.ContainerRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRegistry), err
}

// Update takes the representation of a containerRegistry and updates it. Returns the server's representation of the containerRegistry, and an error, if there is any.
func (c *FakeContainerRegistries) Update(containerRegistry *v1alpha1.ContainerRegistry) (result *v1alpha1.ContainerRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(containerregistriesResource, containerRegistry), &v1alpha1.ContainerRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRegistry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerRegistries) UpdateStatus(containerRegistry *v1alpha1.ContainerRegistry) (*v1alpha1.ContainerRegistry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(containerregistriesResource, "status", containerRegistry), &v1alpha1.ContainerRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRegistry), err
}

// Delete takes name of the containerRegistry and deletes it. Returns an error if one occurs.
func (c *FakeContainerRegistries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(containerregistriesResource, name), &v1alpha1.ContainerRegistry{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerRegistries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(containerregistriesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ContainerRegistryList{})
	return err
}

// Patch applies the patch and returns the patched containerRegistry.
func (c *FakeContainerRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ContainerRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(containerregistriesResource, name, pt, data, subresources...), &v1alpha1.ContainerRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRegistry), err
}
