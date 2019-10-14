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

// FakeComputeNetworks implements ComputeNetworkInterface
type FakeComputeNetworks struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computenetworksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computenetworks"}

var computenetworksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeNetwork"}

// Get takes name of the computeNetwork, and returns the corresponding computeNetwork object, and an error if there is any.
func (c *FakeComputeNetworks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computenetworksResource, c.ns, name), &v1alpha1.ComputeNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetwork), err
}

// List takes label and field selectors, and returns the list of ComputeNetworks that match those selectors.
func (c *FakeComputeNetworks) List(opts v1.ListOptions) (result *v1alpha1.ComputeNetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computenetworksResource, computenetworksKind, c.ns, opts), &v1alpha1.ComputeNetworkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeNetworkList{ListMeta: obj.(*v1alpha1.ComputeNetworkList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeNetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeNetworks.
func (c *FakeComputeNetworks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computenetworksResource, c.ns, opts))

}

// Create takes the representation of a computeNetwork and creates it.  Returns the server's representation of the computeNetwork, and an error, if there is any.
func (c *FakeComputeNetworks) Create(computeNetwork *v1alpha1.ComputeNetwork) (result *v1alpha1.ComputeNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computenetworksResource, c.ns, computeNetwork), &v1alpha1.ComputeNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetwork), err
}

// Update takes the representation of a computeNetwork and updates it. Returns the server's representation of the computeNetwork, and an error, if there is any.
func (c *FakeComputeNetworks) Update(computeNetwork *v1alpha1.ComputeNetwork) (result *v1alpha1.ComputeNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computenetworksResource, c.ns, computeNetwork), &v1alpha1.ComputeNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetwork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeNetworks) UpdateStatus(computeNetwork *v1alpha1.ComputeNetwork) (*v1alpha1.ComputeNetwork, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computenetworksResource, "status", c.ns, computeNetwork), &v1alpha1.ComputeNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetwork), err
}

// Delete takes name of the computeNetwork and deletes it. Returns an error if one occurs.
func (c *FakeComputeNetworks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computenetworksResource, c.ns, name), &v1alpha1.ComputeNetwork{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeNetworks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computenetworksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeNetworkList{})
	return err
}

// Patch applies the patch and returns the patched computeNetwork.
func (c *FakeComputeNetworks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computenetworksResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetwork), err
}
