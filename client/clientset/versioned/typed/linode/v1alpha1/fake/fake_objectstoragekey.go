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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeObjectStorageKeys implements ObjectStorageKeyInterface
type FakeObjectStorageKeys struct {
	Fake *FakeLinodeV1alpha1
	ns   string
}

var objectstoragekeysResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "objectstoragekeys"}

var objectstoragekeysKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "ObjectStorageKey"}

// Get takes name of the objectStorageKey, and returns the corresponding objectStorageKey object, and an error if there is any.
func (c *FakeObjectStorageKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ObjectStorageKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(objectstoragekeysResource, c.ns, name), &v1alpha1.ObjectStorageKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectStorageKey), err
}

// List takes label and field selectors, and returns the list of ObjectStorageKeys that match those selectors.
func (c *FakeObjectStorageKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ObjectStorageKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(objectstoragekeysResource, objectstoragekeysKind, c.ns, opts), &v1alpha1.ObjectStorageKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ObjectStorageKeyList{ListMeta: obj.(*v1alpha1.ObjectStorageKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ObjectStorageKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested objectStorageKeys.
func (c *FakeObjectStorageKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(objectstoragekeysResource, c.ns, opts))

}

// Create takes the representation of a objectStorageKey and creates it.  Returns the server's representation of the objectStorageKey, and an error, if there is any.
func (c *FakeObjectStorageKeys) Create(ctx context.Context, objectStorageKey *v1alpha1.ObjectStorageKey, opts v1.CreateOptions) (result *v1alpha1.ObjectStorageKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(objectstoragekeysResource, c.ns, objectStorageKey), &v1alpha1.ObjectStorageKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectStorageKey), err
}

// Update takes the representation of a objectStorageKey and updates it. Returns the server's representation of the objectStorageKey, and an error, if there is any.
func (c *FakeObjectStorageKeys) Update(ctx context.Context, objectStorageKey *v1alpha1.ObjectStorageKey, opts v1.UpdateOptions) (result *v1alpha1.ObjectStorageKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(objectstoragekeysResource, c.ns, objectStorageKey), &v1alpha1.ObjectStorageKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectStorageKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeObjectStorageKeys) UpdateStatus(ctx context.Context, objectStorageKey *v1alpha1.ObjectStorageKey, opts v1.UpdateOptions) (*v1alpha1.ObjectStorageKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(objectstoragekeysResource, "status", c.ns, objectStorageKey), &v1alpha1.ObjectStorageKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectStorageKey), err
}

// Delete takes name of the objectStorageKey and deletes it. Returns an error if one occurs.
func (c *FakeObjectStorageKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(objectstoragekeysResource, c.ns, name), &v1alpha1.ObjectStorageKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeObjectStorageKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(objectstoragekeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ObjectStorageKeyList{})
	return err
}

// Patch applies the patch and returns the patched objectStorageKey.
func (c *FakeObjectStorageKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ObjectStorageKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(objectstoragekeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ObjectStorageKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectStorageKey), err
}