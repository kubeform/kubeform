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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeComputeInstanceGroupManagers implements ComputeInstanceGroupManagerInterface
type FakeComputeInstanceGroupManagers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeinstancegroupmanagersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeinstancegroupmanagers"}

var computeinstancegroupmanagersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeInstanceGroupManager"}

// Get takes name of the computeInstanceGroupManager, and returns the corresponding computeInstanceGroupManager object, and an error if there is any.
func (c *FakeComputeInstanceGroupManagers) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeinstancegroupmanagersResource, c.ns, name), &v1alpha1.ComputeInstanceGroupManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceGroupManager), err
}

// List takes label and field selectors, and returns the list of ComputeInstanceGroupManagers that match those selectors.
func (c *FakeComputeInstanceGroupManagers) List(opts v1.ListOptions) (result *v1alpha1.ComputeInstanceGroupManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeinstancegroupmanagersResource, computeinstancegroupmanagersKind, c.ns, opts), &v1alpha1.ComputeInstanceGroupManagerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeInstanceGroupManagerList{ListMeta: obj.(*v1alpha1.ComputeInstanceGroupManagerList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeInstanceGroupManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeInstanceGroupManagers.
func (c *FakeComputeInstanceGroupManagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeinstancegroupmanagersResource, c.ns, opts))

}

// Create takes the representation of a computeInstanceGroupManager and creates it.  Returns the server's representation of the computeInstanceGroupManager, and an error, if there is any.
func (c *FakeComputeInstanceGroupManagers) Create(computeInstanceGroupManager *v1alpha1.ComputeInstanceGroupManager) (result *v1alpha1.ComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeinstancegroupmanagersResource, c.ns, computeInstanceGroupManager), &v1alpha1.ComputeInstanceGroupManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceGroupManager), err
}

// Update takes the representation of a computeInstanceGroupManager and updates it. Returns the server's representation of the computeInstanceGroupManager, and an error, if there is any.
func (c *FakeComputeInstanceGroupManagers) Update(computeInstanceGroupManager *v1alpha1.ComputeInstanceGroupManager) (result *v1alpha1.ComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeinstancegroupmanagersResource, c.ns, computeInstanceGroupManager), &v1alpha1.ComputeInstanceGroupManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceGroupManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeInstanceGroupManagers) UpdateStatus(computeInstanceGroupManager *v1alpha1.ComputeInstanceGroupManager) (*v1alpha1.ComputeInstanceGroupManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeinstancegroupmanagersResource, "status", c.ns, computeInstanceGroupManager), &v1alpha1.ComputeInstanceGroupManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceGroupManager), err
}

// Delete takes name of the computeInstanceGroupManager and deletes it. Returns an error if one occurs.
func (c *FakeComputeInstanceGroupManagers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeinstancegroupmanagersResource, c.ns, name), &v1alpha1.ComputeInstanceGroupManager{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeInstanceGroupManagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeinstancegroupmanagersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeInstanceGroupManagerList{})
	return err
}

// Patch applies the patch and returns the patched computeInstanceGroupManager.
func (c *FakeComputeInstanceGroupManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeinstancegroupmanagersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeInstanceGroupManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceGroupManager), err
}
