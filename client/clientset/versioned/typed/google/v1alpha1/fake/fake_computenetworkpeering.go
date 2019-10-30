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

// FakeComputeNetworkPeerings implements ComputeNetworkPeeringInterface
type FakeComputeNetworkPeerings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computenetworkpeeringsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computenetworkpeerings"}

var computenetworkpeeringsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeNetworkPeering"}

// Get takes name of the computeNetworkPeering, and returns the corresponding computeNetworkPeering object, and an error if there is any.
func (c *FakeComputeNetworkPeerings) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computenetworkpeeringsResource, c.ns, name), &v1alpha1.ComputeNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkPeering), err
}

// List takes label and field selectors, and returns the list of ComputeNetworkPeerings that match those selectors.
func (c *FakeComputeNetworkPeerings) List(opts v1.ListOptions) (result *v1alpha1.ComputeNetworkPeeringList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computenetworkpeeringsResource, computenetworkpeeringsKind, c.ns, opts), &v1alpha1.ComputeNetworkPeeringList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeNetworkPeeringList{ListMeta: obj.(*v1alpha1.ComputeNetworkPeeringList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeNetworkPeeringList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeNetworkPeerings.
func (c *FakeComputeNetworkPeerings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computenetworkpeeringsResource, c.ns, opts))

}

// Create takes the representation of a computeNetworkPeering and creates it.  Returns the server's representation of the computeNetworkPeering, and an error, if there is any.
func (c *FakeComputeNetworkPeerings) Create(computeNetworkPeering *v1alpha1.ComputeNetworkPeering) (result *v1alpha1.ComputeNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computenetworkpeeringsResource, c.ns, computeNetworkPeering), &v1alpha1.ComputeNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkPeering), err
}

// Update takes the representation of a computeNetworkPeering and updates it. Returns the server's representation of the computeNetworkPeering, and an error, if there is any.
func (c *FakeComputeNetworkPeerings) Update(computeNetworkPeering *v1alpha1.ComputeNetworkPeering) (result *v1alpha1.ComputeNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computenetworkpeeringsResource, c.ns, computeNetworkPeering), &v1alpha1.ComputeNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkPeering), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeNetworkPeerings) UpdateStatus(computeNetworkPeering *v1alpha1.ComputeNetworkPeering) (*v1alpha1.ComputeNetworkPeering, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computenetworkpeeringsResource, "status", c.ns, computeNetworkPeering), &v1alpha1.ComputeNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkPeering), err
}

// Delete takes name of the computeNetworkPeering and deletes it. Returns an error if one occurs.
func (c *FakeComputeNetworkPeerings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computenetworkpeeringsResource, c.ns, name), &v1alpha1.ComputeNetworkPeering{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeNetworkPeerings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computenetworkpeeringsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeNetworkPeeringList{})
	return err
}

// Patch applies the patch and returns the patched computeNetworkPeering.
func (c *FakeComputeNetworkPeerings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computenetworkpeeringsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkPeering), err
}
