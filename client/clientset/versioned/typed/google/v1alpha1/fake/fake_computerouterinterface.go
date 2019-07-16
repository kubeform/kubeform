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

// FakeComputeRouterInterfaces implements ComputeRouterInterfaceInterface
type FakeComputeRouterInterfaces struct {
	Fake *FakeGoogleV1alpha1
}

var computerouterinterfacesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computerouterinterfaces"}

var computerouterinterfacesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeRouterInterface"}

// Get takes name of the computeRouterInterface, and returns the corresponding computeRouterInterface object, and an error if there is any.
func (c *FakeComputeRouterInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(computerouterinterfacesResource, name), &v1alpha1.ComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterInterface), err
}

// List takes label and field selectors, and returns the list of ComputeRouterInterfaces that match those selectors.
func (c *FakeComputeRouterInterfaces) List(opts v1.ListOptions) (result *v1alpha1.ComputeRouterInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(computerouterinterfacesResource, computerouterinterfacesKind, opts), &v1alpha1.ComputeRouterInterfaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeRouterInterfaceList{ListMeta: obj.(*v1alpha1.ComputeRouterInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeRouterInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeRouterInterfaces.
func (c *FakeComputeRouterInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(computerouterinterfacesResource, opts))
}

// Create takes the representation of a computeRouterInterface and creates it.  Returns the server's representation of the computeRouterInterface, and an error, if there is any.
func (c *FakeComputeRouterInterfaces) Create(computeRouterInterface *v1alpha1.ComputeRouterInterface) (result *v1alpha1.ComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(computerouterinterfacesResource, computeRouterInterface), &v1alpha1.ComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterInterface), err
}

// Update takes the representation of a computeRouterInterface and updates it. Returns the server's representation of the computeRouterInterface, and an error, if there is any.
func (c *FakeComputeRouterInterfaces) Update(computeRouterInterface *v1alpha1.ComputeRouterInterface) (result *v1alpha1.ComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(computerouterinterfacesResource, computeRouterInterface), &v1alpha1.ComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeRouterInterfaces) UpdateStatus(computeRouterInterface *v1alpha1.ComputeRouterInterface) (*v1alpha1.ComputeRouterInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(computerouterinterfacesResource, "status", computeRouterInterface), &v1alpha1.ComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterInterface), err
}

// Delete takes name of the computeRouterInterface and deletes it. Returns an error if one occurs.
func (c *FakeComputeRouterInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(computerouterinterfacesResource, name), &v1alpha1.ComputeRouterInterface{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeRouterInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(computerouterinterfacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeRouterInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched computeRouterInterface.
func (c *FakeComputeRouterInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(computerouterinterfacesResource, name, pt, data, subresources...), &v1alpha1.ComputeRouterInterface{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterInterface), err
}
