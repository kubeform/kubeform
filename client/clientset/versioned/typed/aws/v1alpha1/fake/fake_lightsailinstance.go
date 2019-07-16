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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeLightsailInstances implements LightsailInstanceInterface
type FakeLightsailInstances struct {
	Fake *FakeAwsV1alpha1
}

var lightsailinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lightsailinstances"}

var lightsailinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LightsailInstance"}

// Get takes name of the lightsailInstance, and returns the corresponding lightsailInstance object, and an error if there is any.
func (c *FakeLightsailInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.LightsailInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(lightsailinstancesResource, name), &v1alpha1.LightsailInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailInstance), err
}

// List takes label and field selectors, and returns the list of LightsailInstances that match those selectors.
func (c *FakeLightsailInstances) List(opts v1.ListOptions) (result *v1alpha1.LightsailInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(lightsailinstancesResource, lightsailinstancesKind, opts), &v1alpha1.LightsailInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LightsailInstanceList{ListMeta: obj.(*v1alpha1.LightsailInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.LightsailInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lightsailInstances.
func (c *FakeLightsailInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(lightsailinstancesResource, opts))
}

// Create takes the representation of a lightsailInstance and creates it.  Returns the server's representation of the lightsailInstance, and an error, if there is any.
func (c *FakeLightsailInstances) Create(lightsailInstance *v1alpha1.LightsailInstance) (result *v1alpha1.LightsailInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(lightsailinstancesResource, lightsailInstance), &v1alpha1.LightsailInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailInstance), err
}

// Update takes the representation of a lightsailInstance and updates it. Returns the server's representation of the lightsailInstance, and an error, if there is any.
func (c *FakeLightsailInstances) Update(lightsailInstance *v1alpha1.LightsailInstance) (result *v1alpha1.LightsailInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(lightsailinstancesResource, lightsailInstance), &v1alpha1.LightsailInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLightsailInstances) UpdateStatus(lightsailInstance *v1alpha1.LightsailInstance) (*v1alpha1.LightsailInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(lightsailinstancesResource, "status", lightsailInstance), &v1alpha1.LightsailInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailInstance), err
}

// Delete takes name of the lightsailInstance and deletes it. Returns an error if one occurs.
func (c *FakeLightsailInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(lightsailinstancesResource, name), &v1alpha1.LightsailInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLightsailInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(lightsailinstancesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LightsailInstanceList{})
	return err
}

// Patch applies the patch and returns the patched lightsailInstance.
func (c *FakeLightsailInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(lightsailinstancesResource, name, pt, data, subresources...), &v1alpha1.LightsailInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailInstance), err
}
