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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeDigitaloceanVolumes implements DigitaloceanVolumeInterface
type FakeDigitaloceanVolumes struct {
	Fake *FakeDigitaloceanV1alpha1
}

var digitaloceanvolumesResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "digitaloceanvolumes"}

var digitaloceanvolumesKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DigitaloceanVolume"}

// Get takes name of the digitaloceanVolume, and returns the corresponding digitaloceanVolume object, and an error if there is any.
func (c *FakeDigitaloceanVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.DigitaloceanVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(digitaloceanvolumesResource, name), &v1alpha1.DigitaloceanVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolume), err
}

// List takes label and field selectors, and returns the list of DigitaloceanVolumes that match those selectors.
func (c *FakeDigitaloceanVolumes) List(opts v1.ListOptions) (result *v1alpha1.DigitaloceanVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(digitaloceanvolumesResource, digitaloceanvolumesKind, opts), &v1alpha1.DigitaloceanVolumeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DigitaloceanVolumeList{ListMeta: obj.(*v1alpha1.DigitaloceanVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.DigitaloceanVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested digitaloceanVolumes.
func (c *FakeDigitaloceanVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(digitaloceanvolumesResource, opts))
}

// Create takes the representation of a digitaloceanVolume and creates it.  Returns the server's representation of the digitaloceanVolume, and an error, if there is any.
func (c *FakeDigitaloceanVolumes) Create(digitaloceanVolume *v1alpha1.DigitaloceanVolume) (result *v1alpha1.DigitaloceanVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(digitaloceanvolumesResource, digitaloceanVolume), &v1alpha1.DigitaloceanVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolume), err
}

// Update takes the representation of a digitaloceanVolume and updates it. Returns the server's representation of the digitaloceanVolume, and an error, if there is any.
func (c *FakeDigitaloceanVolumes) Update(digitaloceanVolume *v1alpha1.DigitaloceanVolume) (result *v1alpha1.DigitaloceanVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(digitaloceanvolumesResource, digitaloceanVolume), &v1alpha1.DigitaloceanVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDigitaloceanVolumes) UpdateStatus(digitaloceanVolume *v1alpha1.DigitaloceanVolume) (*v1alpha1.DigitaloceanVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(digitaloceanvolumesResource, "status", digitaloceanVolume), &v1alpha1.DigitaloceanVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolume), err
}

// Delete takes name of the digitaloceanVolume and deletes it. Returns an error if one occurs.
func (c *FakeDigitaloceanVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(digitaloceanvolumesResource, name), &v1alpha1.DigitaloceanVolume{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDigitaloceanVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(digitaloceanvolumesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DigitaloceanVolumeList{})
	return err
}

// Patch applies the patch and returns the patched digitaloceanVolume.
func (c *FakeDigitaloceanVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DigitaloceanVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(digitaloceanvolumesResource, name, pt, data, subresources...), &v1alpha1.DigitaloceanVolume{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DigitaloceanVolume), err
}
