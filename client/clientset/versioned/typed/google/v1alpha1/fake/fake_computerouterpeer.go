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

// FakeComputeRouterPeers implements ComputeRouterPeerInterface
type FakeComputeRouterPeers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computerouterpeersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computerouterpeers"}

var computerouterpeersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeRouterPeer"}

// Get takes name of the computeRouterPeer, and returns the corresponding computeRouterPeer object, and an error if there is any.
func (c *FakeComputeRouterPeers) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computerouterpeersResource, c.ns, name), &v1alpha1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterPeer), err
}

// List takes label and field selectors, and returns the list of ComputeRouterPeers that match those selectors.
func (c *FakeComputeRouterPeers) List(opts v1.ListOptions) (result *v1alpha1.ComputeRouterPeerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computerouterpeersResource, computerouterpeersKind, c.ns, opts), &v1alpha1.ComputeRouterPeerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeRouterPeerList{ListMeta: obj.(*v1alpha1.ComputeRouterPeerList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeRouterPeerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeRouterPeers.
func (c *FakeComputeRouterPeers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computerouterpeersResource, c.ns, opts))

}

// Create takes the representation of a computeRouterPeer and creates it.  Returns the server's representation of the computeRouterPeer, and an error, if there is any.
func (c *FakeComputeRouterPeers) Create(computeRouterPeer *v1alpha1.ComputeRouterPeer) (result *v1alpha1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computerouterpeersResource, c.ns, computeRouterPeer), &v1alpha1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterPeer), err
}

// Update takes the representation of a computeRouterPeer and updates it. Returns the server's representation of the computeRouterPeer, and an error, if there is any.
func (c *FakeComputeRouterPeers) Update(computeRouterPeer *v1alpha1.ComputeRouterPeer) (result *v1alpha1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computerouterpeersResource, c.ns, computeRouterPeer), &v1alpha1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterPeer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeRouterPeers) UpdateStatus(computeRouterPeer *v1alpha1.ComputeRouterPeer) (*v1alpha1.ComputeRouterPeer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computerouterpeersResource, "status", c.ns, computeRouterPeer), &v1alpha1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterPeer), err
}

// Delete takes name of the computeRouterPeer and deletes it. Returns an error if one occurs.
func (c *FakeComputeRouterPeers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computerouterpeersResource, c.ns, name), &v1alpha1.ComputeRouterPeer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeRouterPeers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computerouterpeersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeRouterPeerList{})
	return err
}

// Patch applies the patch and returns the patched computeRouterPeer.
func (c *FakeComputeRouterPeers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouterPeer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computerouterpeersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeRouterPeer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeRouterPeer), err
}
