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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeDatasyncLocationNfses implements DatasyncLocationNfsInterface
type FakeDatasyncLocationNfses struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var datasynclocationnfsesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "datasynclocationnfses"}

var datasynclocationnfsesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DatasyncLocationNfs"}

// Get takes name of the datasyncLocationNfs, and returns the corresponding datasyncLocationNfs object, and an error if there is any.
func (c *FakeDatasyncLocationNfses) Get(name string, options v1.GetOptions) (result *v1alpha1.DatasyncLocationNfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasynclocationnfsesResource, c.ns, name), &v1alpha1.DatasyncLocationNfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationNfs), err
}

// List takes label and field selectors, and returns the list of DatasyncLocationNfses that match those selectors.
func (c *FakeDatasyncLocationNfses) List(opts v1.ListOptions) (result *v1alpha1.DatasyncLocationNfsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasynclocationnfsesResource, datasynclocationnfsesKind, c.ns, opts), &v1alpha1.DatasyncLocationNfsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatasyncLocationNfsList{ListMeta: obj.(*v1alpha1.DatasyncLocationNfsList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatasyncLocationNfsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasyncLocationNfses.
func (c *FakeDatasyncLocationNfses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasynclocationnfsesResource, c.ns, opts))

}

// Create takes the representation of a datasyncLocationNfs and creates it.  Returns the server's representation of the datasyncLocationNfs, and an error, if there is any.
func (c *FakeDatasyncLocationNfses) Create(datasyncLocationNfs *v1alpha1.DatasyncLocationNfs) (result *v1alpha1.DatasyncLocationNfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasynclocationnfsesResource, c.ns, datasyncLocationNfs), &v1alpha1.DatasyncLocationNfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationNfs), err
}

// Update takes the representation of a datasyncLocationNfs and updates it. Returns the server's representation of the datasyncLocationNfs, and an error, if there is any.
func (c *FakeDatasyncLocationNfses) Update(datasyncLocationNfs *v1alpha1.DatasyncLocationNfs) (result *v1alpha1.DatasyncLocationNfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasynclocationnfsesResource, c.ns, datasyncLocationNfs), &v1alpha1.DatasyncLocationNfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationNfs), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatasyncLocationNfses) UpdateStatus(datasyncLocationNfs *v1alpha1.DatasyncLocationNfs) (*v1alpha1.DatasyncLocationNfs, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datasynclocationnfsesResource, "status", c.ns, datasyncLocationNfs), &v1alpha1.DatasyncLocationNfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationNfs), err
}

// Delete takes name of the datasyncLocationNfs and deletes it. Returns an error if one occurs.
func (c *FakeDatasyncLocationNfses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasynclocationnfsesResource, c.ns, name), &v1alpha1.DatasyncLocationNfs{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasyncLocationNfses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasynclocationnfsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatasyncLocationNfsList{})
	return err
}

// Patch applies the patch and returns the patched datasyncLocationNfs.
func (c *FakeDatasyncLocationNfses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatasyncLocationNfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasynclocationnfsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatasyncLocationNfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasyncLocationNfs), err
}