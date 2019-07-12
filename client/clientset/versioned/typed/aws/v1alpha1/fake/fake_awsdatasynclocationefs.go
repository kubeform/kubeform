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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsDatasyncLocationEfses implements AwsDatasyncLocationEfsInterface
type FakeAwsDatasyncLocationEfses struct {
	Fake *FakeAwsV1alpha1
}

var awsdatasynclocationefsesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdatasynclocationefses"}

var awsdatasynclocationefsesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDatasyncLocationEfs"}

// Get takes name of the awsDatasyncLocationEfs, and returns the corresponding awsDatasyncLocationEfs object, and an error if there is any.
func (c *FakeAwsDatasyncLocationEfses) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDatasyncLocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdatasynclocationefsesResource, name), &v1alpha1.AwsDatasyncLocationEfs{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationEfs), err
}

// List takes label and field selectors, and returns the list of AwsDatasyncLocationEfses that match those selectors.
func (c *FakeAwsDatasyncLocationEfses) List(opts v1.ListOptions) (result *v1alpha1.AwsDatasyncLocationEfsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdatasynclocationefsesResource, awsdatasynclocationefsesKind, opts), &v1alpha1.AwsDatasyncLocationEfsList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDatasyncLocationEfsList{ListMeta: obj.(*v1alpha1.AwsDatasyncLocationEfsList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDatasyncLocationEfsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDatasyncLocationEfses.
func (c *FakeAwsDatasyncLocationEfses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdatasynclocationefsesResource, opts))
}

// Create takes the representation of a awsDatasyncLocationEfs and creates it.  Returns the server's representation of the awsDatasyncLocationEfs, and an error, if there is any.
func (c *FakeAwsDatasyncLocationEfses) Create(awsDatasyncLocationEfs *v1alpha1.AwsDatasyncLocationEfs) (result *v1alpha1.AwsDatasyncLocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdatasynclocationefsesResource, awsDatasyncLocationEfs), &v1alpha1.AwsDatasyncLocationEfs{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationEfs), err
}

// Update takes the representation of a awsDatasyncLocationEfs and updates it. Returns the server's representation of the awsDatasyncLocationEfs, and an error, if there is any.
func (c *FakeAwsDatasyncLocationEfses) Update(awsDatasyncLocationEfs *v1alpha1.AwsDatasyncLocationEfs) (result *v1alpha1.AwsDatasyncLocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdatasynclocationefsesResource, awsDatasyncLocationEfs), &v1alpha1.AwsDatasyncLocationEfs{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationEfs), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDatasyncLocationEfses) UpdateStatus(awsDatasyncLocationEfs *v1alpha1.AwsDatasyncLocationEfs) (*v1alpha1.AwsDatasyncLocationEfs, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdatasynclocationefsesResource, "status", awsDatasyncLocationEfs), &v1alpha1.AwsDatasyncLocationEfs{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationEfs), err
}

// Delete takes name of the awsDatasyncLocationEfs and deletes it. Returns an error if one occurs.
func (c *FakeAwsDatasyncLocationEfses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdatasynclocationefsesResource, name), &v1alpha1.AwsDatasyncLocationEfs{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDatasyncLocationEfses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdatasynclocationefsesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDatasyncLocationEfsList{})
	return err
}

// Patch applies the patch and returns the patched awsDatasyncLocationEfs.
func (c *FakeAwsDatasyncLocationEfses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDatasyncLocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdatasynclocationefsesResource, name, pt, data, subresources...), &v1alpha1.AwsDatasyncLocationEfs{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationEfs), err
}
