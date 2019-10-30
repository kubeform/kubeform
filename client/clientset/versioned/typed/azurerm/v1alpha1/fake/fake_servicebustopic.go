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

// FakeServicebusTopics implements ServicebusTopicInterface
type FakeServicebusTopics struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var servicebustopicsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicebustopics"}

var servicebustopicsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServicebusTopic"}

// Get takes name of the servicebusTopic, and returns the corresponding servicebusTopic object, and an error if there is any.
func (c *FakeServicebusTopics) Get(name string, options v1.GetOptions) (result *v1alpha1.ServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicebustopicsResource, c.ns, name), &v1alpha1.ServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusTopic), err
}

// List takes label and field selectors, and returns the list of ServicebusTopics that match those selectors.
func (c *FakeServicebusTopics) List(opts v1.ListOptions) (result *v1alpha1.ServicebusTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicebustopicsResource, servicebustopicsKind, c.ns, opts), &v1alpha1.ServicebusTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicebusTopicList{ListMeta: obj.(*v1alpha1.ServicebusTopicList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicebusTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicebusTopics.
func (c *FakeServicebusTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicebustopicsResource, c.ns, opts))

}

// Create takes the representation of a servicebusTopic and creates it.  Returns the server's representation of the servicebusTopic, and an error, if there is any.
func (c *FakeServicebusTopics) Create(servicebusTopic *v1alpha1.ServicebusTopic) (result *v1alpha1.ServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicebustopicsResource, c.ns, servicebusTopic), &v1alpha1.ServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusTopic), err
}

// Update takes the representation of a servicebusTopic and updates it. Returns the server's representation of the servicebusTopic, and an error, if there is any.
func (c *FakeServicebusTopics) Update(servicebusTopic *v1alpha1.ServicebusTopic) (result *v1alpha1.ServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicebustopicsResource, c.ns, servicebusTopic), &v1alpha1.ServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicebusTopics) UpdateStatus(servicebusTopic *v1alpha1.ServicebusTopic) (*v1alpha1.ServicebusTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicebustopicsResource, "status", c.ns, servicebusTopic), &v1alpha1.ServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusTopic), err
}

// Delete takes name of the servicebusTopic and deletes it. Returns an error if one occurs.
func (c *FakeServicebusTopics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicebustopicsResource, c.ns, name), &v1alpha1.ServicebusTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicebusTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicebustopicsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicebusTopicList{})
	return err
}

// Patch applies the patch and returns the patched servicebusTopic.
func (c *FakeServicebusTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicebustopicsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusTopic), err
}
