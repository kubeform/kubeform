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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakePublicIPs implements PublicIPInterface
type FakePublicIPs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var publicipsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "publicips"}

var publicipsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PublicIP"}

// Get takes name of the publicIP, and returns the corresponding publicIP object, and an error if there is any.
func (c *FakePublicIPs) Get(name string, options v1.GetOptions) (result *v1alpha1.PublicIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(publicipsResource, c.ns, name), &v1alpha1.PublicIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIP), err
}

// List takes label and field selectors, and returns the list of PublicIPs that match those selectors.
func (c *FakePublicIPs) List(opts v1.ListOptions) (result *v1alpha1.PublicIPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(publicipsResource, publicipsKind, c.ns, opts), &v1alpha1.PublicIPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PublicIPList{ListMeta: obj.(*v1alpha1.PublicIPList).ListMeta}
	for _, item := range obj.(*v1alpha1.PublicIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested publicIPs.
func (c *FakePublicIPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(publicipsResource, c.ns, opts))

}

// Create takes the representation of a publicIP and creates it.  Returns the server's representation of the publicIP, and an error, if there is any.
func (c *FakePublicIPs) Create(publicIP *v1alpha1.PublicIP) (result *v1alpha1.PublicIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(publicipsResource, c.ns, publicIP), &v1alpha1.PublicIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIP), err
}

// Update takes the representation of a publicIP and updates it. Returns the server's representation of the publicIP, and an error, if there is any.
func (c *FakePublicIPs) Update(publicIP *v1alpha1.PublicIP) (result *v1alpha1.PublicIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(publicipsResource, c.ns, publicIP), &v1alpha1.PublicIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePublicIPs) UpdateStatus(publicIP *v1alpha1.PublicIP) (*v1alpha1.PublicIP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(publicipsResource, "status", c.ns, publicIP), &v1alpha1.PublicIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIP), err
}

// Delete takes name of the publicIP and deletes it. Returns an error if one occurs.
func (c *FakePublicIPs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(publicipsResource, c.ns, name), &v1alpha1.PublicIP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePublicIPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(publicipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PublicIPList{})
	return err
}

// Patch applies the patch and returns the patched publicIP.
func (c *FakePublicIPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PublicIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(publicipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PublicIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PublicIP), err
}
