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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBigtableInstances implements BigtableInstanceInterface
type FakeBigtableInstances struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var bigtableinstancesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "bigtableinstances"}

var bigtableinstancesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BigtableInstance"}

// Get takes name of the bigtableInstance, and returns the corresponding bigtableInstance object, and an error if there is any.
func (c *FakeBigtableInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.BigtableInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bigtableinstancesResource, c.ns, name), &v1alpha1.BigtableInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstance), err
}

// List takes label and field selectors, and returns the list of BigtableInstances that match those selectors.
func (c *FakeBigtableInstances) List(opts v1.ListOptions) (result *v1alpha1.BigtableInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bigtableinstancesResource, bigtableinstancesKind, c.ns, opts), &v1alpha1.BigtableInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BigtableInstanceList{ListMeta: obj.(*v1alpha1.BigtableInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.BigtableInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bigtableInstances.
func (c *FakeBigtableInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bigtableinstancesResource, c.ns, opts))

}

// Create takes the representation of a bigtableInstance and creates it.  Returns the server's representation of the bigtableInstance, and an error, if there is any.
func (c *FakeBigtableInstances) Create(bigtableInstance *v1alpha1.BigtableInstance) (result *v1alpha1.BigtableInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bigtableinstancesResource, c.ns, bigtableInstance), &v1alpha1.BigtableInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstance), err
}

// Update takes the representation of a bigtableInstance and updates it. Returns the server's representation of the bigtableInstance, and an error, if there is any.
func (c *FakeBigtableInstances) Update(bigtableInstance *v1alpha1.BigtableInstance) (result *v1alpha1.BigtableInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bigtableinstancesResource, c.ns, bigtableInstance), &v1alpha1.BigtableInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBigtableInstances) UpdateStatus(bigtableInstance *v1alpha1.BigtableInstance) (*v1alpha1.BigtableInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bigtableinstancesResource, "status", c.ns, bigtableInstance), &v1alpha1.BigtableInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstance), err
}

// Delete takes name of the bigtableInstance and deletes it. Returns an error if one occurs.
func (c *FakeBigtableInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bigtableinstancesResource, c.ns, name), &v1alpha1.BigtableInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBigtableInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bigtableinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BigtableInstanceList{})
	return err
}

// Patch applies the patch and returns the patched bigtableInstance.
func (c *FakeBigtableInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BigtableInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bigtableinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BigtableInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BigtableInstance), err
}
