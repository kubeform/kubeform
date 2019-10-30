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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkProfiles implements NetworkProfileInterface
type FakeNetworkProfiles struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networkprofilesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkprofiles"}

var networkprofilesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkProfile"}

// Get takes name of the networkProfile, and returns the corresponding networkProfile object, and an error if there is any.
func (c *FakeNetworkProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkprofilesResource, c.ns, name), &v1alpha1.NetworkProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkProfile), err
}

// List takes label and field selectors, and returns the list of NetworkProfiles that match those selectors.
func (c *FakeNetworkProfiles) List(opts v1.ListOptions) (result *v1alpha1.NetworkProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkprofilesResource, networkprofilesKind, c.ns, opts), &v1alpha1.NetworkProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkProfileList{ListMeta: obj.(*v1alpha1.NetworkProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkProfiles.
func (c *FakeNetworkProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkprofilesResource, c.ns, opts))

}

// Create takes the representation of a networkProfile and creates it.  Returns the server's representation of the networkProfile, and an error, if there is any.
func (c *FakeNetworkProfiles) Create(networkProfile *v1alpha1.NetworkProfile) (result *v1alpha1.NetworkProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkprofilesResource, c.ns, networkProfile), &v1alpha1.NetworkProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkProfile), err
}

// Update takes the representation of a networkProfile and updates it. Returns the server's representation of the networkProfile, and an error, if there is any.
func (c *FakeNetworkProfiles) Update(networkProfile *v1alpha1.NetworkProfile) (result *v1alpha1.NetworkProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkprofilesResource, c.ns, networkProfile), &v1alpha1.NetworkProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkProfiles) UpdateStatus(networkProfile *v1alpha1.NetworkProfile) (*v1alpha1.NetworkProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkprofilesResource, "status", c.ns, networkProfile), &v1alpha1.NetworkProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkProfile), err
}

// Delete takes name of the networkProfile and deletes it. Returns an error if one occurs.
func (c *FakeNetworkProfiles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkprofilesResource, c.ns, name), &v1alpha1.NetworkProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkprofilesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkProfileList{})
	return err
}

// Patch applies the patch and returns the patched networkProfile.
func (c *FakeNetworkProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkProfile), err
}
