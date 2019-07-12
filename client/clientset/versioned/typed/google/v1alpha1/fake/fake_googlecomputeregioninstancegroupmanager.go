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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleComputeRegionInstanceGroupManagers implements GoogleComputeRegionInstanceGroupManagerInterface
type FakeGoogleComputeRegionInstanceGroupManagers struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputeregioninstancegroupmanagersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputeregioninstancegroupmanagers"}

var googlecomputeregioninstancegroupmanagersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeRegionInstanceGroupManager"}

// Get takes name of the googleComputeRegionInstanceGroupManager, and returns the corresponding googleComputeRegionInstanceGroupManager object, and an error if there is any.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeRegionInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputeregioninstancegroupmanagersResource, name), &v1alpha1.GoogleComputeRegionInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRegionInstanceGroupManager), err
}

// List takes label and field selectors, and returns the list of GoogleComputeRegionInstanceGroupManagers that match those selectors.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeRegionInstanceGroupManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputeregioninstancegroupmanagersResource, googlecomputeregioninstancegroupmanagersKind, opts), &v1alpha1.GoogleComputeRegionInstanceGroupManagerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeRegionInstanceGroupManagerList{ListMeta: obj.(*v1alpha1.GoogleComputeRegionInstanceGroupManagerList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeRegionInstanceGroupManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeRegionInstanceGroupManagers.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputeregioninstancegroupmanagersResource, opts))
}

// Create takes the representation of a googleComputeRegionInstanceGroupManager and creates it.  Returns the server's representation of the googleComputeRegionInstanceGroupManager, and an error, if there is any.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) Create(googleComputeRegionInstanceGroupManager *v1alpha1.GoogleComputeRegionInstanceGroupManager) (result *v1alpha1.GoogleComputeRegionInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputeregioninstancegroupmanagersResource, googleComputeRegionInstanceGroupManager), &v1alpha1.GoogleComputeRegionInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRegionInstanceGroupManager), err
}

// Update takes the representation of a googleComputeRegionInstanceGroupManager and updates it. Returns the server's representation of the googleComputeRegionInstanceGroupManager, and an error, if there is any.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) Update(googleComputeRegionInstanceGroupManager *v1alpha1.GoogleComputeRegionInstanceGroupManager) (result *v1alpha1.GoogleComputeRegionInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputeregioninstancegroupmanagersResource, googleComputeRegionInstanceGroupManager), &v1alpha1.GoogleComputeRegionInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRegionInstanceGroupManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeRegionInstanceGroupManagers) UpdateStatus(googleComputeRegionInstanceGroupManager *v1alpha1.GoogleComputeRegionInstanceGroupManager) (*v1alpha1.GoogleComputeRegionInstanceGroupManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputeregioninstancegroupmanagersResource, "status", googleComputeRegionInstanceGroupManager), &v1alpha1.GoogleComputeRegionInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRegionInstanceGroupManager), err
}

// Delete takes name of the googleComputeRegionInstanceGroupManager and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputeregioninstancegroupmanagersResource, name), &v1alpha1.GoogleComputeRegionInstanceGroupManager{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputeregioninstancegroupmanagersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeRegionInstanceGroupManagerList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeRegionInstanceGroupManager.
func (c *FakeGoogleComputeRegionInstanceGroupManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeRegionInstanceGroupManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputeregioninstancegroupmanagersResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeRegionInstanceGroupManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeRegionInstanceGroupManager), err
}
