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

// FakeGoogleStorageBucketObjects implements GoogleStorageBucketObjectInterface
type FakeGoogleStorageBucketObjects struct {
	Fake *FakeGoogleV1alpha1
}

var googlestoragebucketobjectsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlestoragebucketobjects"}

var googlestoragebucketobjectsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleStorageBucketObject"}

// Get takes name of the googleStorageBucketObject, and returns the corresponding googleStorageBucketObject object, and an error if there is any.
func (c *FakeGoogleStorageBucketObjects) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleStorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlestoragebucketobjectsResource, name), &v1alpha1.GoogleStorageBucketObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketObject), err
}

// List takes label and field selectors, and returns the list of GoogleStorageBucketObjects that match those selectors.
func (c *FakeGoogleStorageBucketObjects) List(opts v1.ListOptions) (result *v1alpha1.GoogleStorageBucketObjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlestoragebucketobjectsResource, googlestoragebucketobjectsKind, opts), &v1alpha1.GoogleStorageBucketObjectList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleStorageBucketObjectList{ListMeta: obj.(*v1alpha1.GoogleStorageBucketObjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleStorageBucketObjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleStorageBucketObjects.
func (c *FakeGoogleStorageBucketObjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlestoragebucketobjectsResource, opts))
}

// Create takes the representation of a googleStorageBucketObject and creates it.  Returns the server's representation of the googleStorageBucketObject, and an error, if there is any.
func (c *FakeGoogleStorageBucketObjects) Create(googleStorageBucketObject *v1alpha1.GoogleStorageBucketObject) (result *v1alpha1.GoogleStorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlestoragebucketobjectsResource, googleStorageBucketObject), &v1alpha1.GoogleStorageBucketObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketObject), err
}

// Update takes the representation of a googleStorageBucketObject and updates it. Returns the server's representation of the googleStorageBucketObject, and an error, if there is any.
func (c *FakeGoogleStorageBucketObjects) Update(googleStorageBucketObject *v1alpha1.GoogleStorageBucketObject) (result *v1alpha1.GoogleStorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlestoragebucketobjectsResource, googleStorageBucketObject), &v1alpha1.GoogleStorageBucketObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketObject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleStorageBucketObjects) UpdateStatus(googleStorageBucketObject *v1alpha1.GoogleStorageBucketObject) (*v1alpha1.GoogleStorageBucketObject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlestoragebucketobjectsResource, "status", googleStorageBucketObject), &v1alpha1.GoogleStorageBucketObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketObject), err
}

// Delete takes name of the googleStorageBucketObject and deletes it. Returns an error if one occurs.
func (c *FakeGoogleStorageBucketObjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlestoragebucketobjectsResource, name), &v1alpha1.GoogleStorageBucketObject{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleStorageBucketObjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlestoragebucketobjectsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleStorageBucketObjectList{})
	return err
}

// Patch applies the patch and returns the patched googleStorageBucketObject.
func (c *FakeGoogleStorageBucketObjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleStorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlestoragebucketobjectsResource, name, pt, data, subresources...), &v1alpha1.GoogleStorageBucketObject{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleStorageBucketObject), err
}
