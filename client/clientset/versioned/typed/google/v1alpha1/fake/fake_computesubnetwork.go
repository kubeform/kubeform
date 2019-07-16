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

// FakeComputeSubnetworks implements ComputeSubnetworkInterface
type FakeComputeSubnetworks struct {
	Fake *FakeGoogleV1alpha1
}

var computesubnetworksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computesubnetworks"}

var computesubnetworksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeSubnetwork"}

// Get takes name of the computeSubnetwork, and returns the corresponding computeSubnetwork object, and an error if there is any.
func (c *FakeComputeSubnetworks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(computesubnetworksResource, name), &v1alpha1.ComputeSubnetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetwork), err
}

// List takes label and field selectors, and returns the list of ComputeSubnetworks that match those selectors.
func (c *FakeComputeSubnetworks) List(opts v1.ListOptions) (result *v1alpha1.ComputeSubnetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(computesubnetworksResource, computesubnetworksKind, opts), &v1alpha1.ComputeSubnetworkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeSubnetworkList{ListMeta: obj.(*v1alpha1.ComputeSubnetworkList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeSubnetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeSubnetworks.
func (c *FakeComputeSubnetworks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(computesubnetworksResource, opts))
}

// Create takes the representation of a computeSubnetwork and creates it.  Returns the server's representation of the computeSubnetwork, and an error, if there is any.
func (c *FakeComputeSubnetworks) Create(computeSubnetwork *v1alpha1.ComputeSubnetwork) (result *v1alpha1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(computesubnetworksResource, computeSubnetwork), &v1alpha1.ComputeSubnetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetwork), err
}

// Update takes the representation of a computeSubnetwork and updates it. Returns the server's representation of the computeSubnetwork, and an error, if there is any.
func (c *FakeComputeSubnetworks) Update(computeSubnetwork *v1alpha1.ComputeSubnetwork) (result *v1alpha1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(computesubnetworksResource, computeSubnetwork), &v1alpha1.ComputeSubnetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetwork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeSubnetworks) UpdateStatus(computeSubnetwork *v1alpha1.ComputeSubnetwork) (*v1alpha1.ComputeSubnetwork, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(computesubnetworksResource, "status", computeSubnetwork), &v1alpha1.ComputeSubnetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetwork), err
}

// Delete takes name of the computeSubnetwork and deletes it. Returns an error if one occurs.
func (c *FakeComputeSubnetworks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(computesubnetworksResource, name), &v1alpha1.ComputeSubnetwork{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeSubnetworks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(computesubnetworksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeSubnetworkList{})
	return err
}

// Patch applies the patch and returns the patched computeSubnetwork.
func (c *FakeComputeSubnetworks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSubnetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(computesubnetworksResource, name, pt, data, subresources...), &v1alpha1.ComputeSubnetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSubnetwork), err
}
