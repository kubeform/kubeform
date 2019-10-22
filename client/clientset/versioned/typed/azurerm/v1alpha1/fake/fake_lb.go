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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLbs implements LbInterface
type FakeLbs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var lbsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "lbs"}

var lbsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "Lb"}

// Get takes name of the lb, and returns the corresponding lb object, and an error if there is any.
func (c *FakeLbs) Get(name string, options v1.GetOptions) (result *v1alpha1.Lb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lbsResource, c.ns, name), &v1alpha1.Lb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lb), err
}

// List takes label and field selectors, and returns the list of Lbs that match those selectors.
func (c *FakeLbs) List(opts v1.ListOptions) (result *v1alpha1.LbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lbsResource, lbsKind, c.ns, opts), &v1alpha1.LbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbList{ListMeta: obj.(*v1alpha1.LbList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbs.
func (c *FakeLbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lbsResource, c.ns, opts))

}

// Create takes the representation of a lb and creates it.  Returns the server's representation of the lb, and an error, if there is any.
func (c *FakeLbs) Create(lb *v1alpha1.Lb) (result *v1alpha1.Lb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lbsResource, c.ns, lb), &v1alpha1.Lb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lb), err
}

// Update takes the representation of a lb and updates it. Returns the server's representation of the lb, and an error, if there is any.
func (c *FakeLbs) Update(lb *v1alpha1.Lb) (result *v1alpha1.Lb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lbsResource, c.ns, lb), &v1alpha1.Lb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbs) UpdateStatus(lb *v1alpha1.Lb) (*v1alpha1.Lb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lbsResource, "status", c.ns, lb), &v1alpha1.Lb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lb), err
}

// Delete takes name of the lb and deletes it. Returns an error if one occurs.
func (c *FakeLbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lbsResource, c.ns, name), &v1alpha1.Lb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lbsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbList{})
	return err
}

// Patch applies the patch and returns the patched lb.
func (c *FakeLbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Lb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lbsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Lb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lb), err
}
