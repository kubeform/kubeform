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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeGoogleFilestoreInstances implements GoogleFilestoreInstanceInterface
type FakeGoogleFilestoreInstances struct {
	Fake *FakeGoogleV1alpha1
}

var googlefilestoreinstancesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlefilestoreinstances"}

var googlefilestoreinstancesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleFilestoreInstance"}

// Get takes name of the googleFilestoreInstance, and returns the corresponding googleFilestoreInstance object, and an error if there is any.
func (c *FakeGoogleFilestoreInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleFilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlefilestoreinstancesResource, name), &v1alpha1.GoogleFilestoreInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFilestoreInstance), err
}

// List takes label and field selectors, and returns the list of GoogleFilestoreInstances that match those selectors.
func (c *FakeGoogleFilestoreInstances) List(opts v1.ListOptions) (result *v1alpha1.GoogleFilestoreInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlefilestoreinstancesResource, googlefilestoreinstancesKind, opts), &v1alpha1.GoogleFilestoreInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleFilestoreInstanceList{ListMeta: obj.(*v1alpha1.GoogleFilestoreInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleFilestoreInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleFilestoreInstances.
func (c *FakeGoogleFilestoreInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlefilestoreinstancesResource, opts))
}

// Create takes the representation of a googleFilestoreInstance and creates it.  Returns the server's representation of the googleFilestoreInstance, and an error, if there is any.
func (c *FakeGoogleFilestoreInstances) Create(googleFilestoreInstance *v1alpha1.GoogleFilestoreInstance) (result *v1alpha1.GoogleFilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlefilestoreinstancesResource, googleFilestoreInstance), &v1alpha1.GoogleFilestoreInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFilestoreInstance), err
}

// Update takes the representation of a googleFilestoreInstance and updates it. Returns the server's representation of the googleFilestoreInstance, and an error, if there is any.
func (c *FakeGoogleFilestoreInstances) Update(googleFilestoreInstance *v1alpha1.GoogleFilestoreInstance) (result *v1alpha1.GoogleFilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlefilestoreinstancesResource, googleFilestoreInstance), &v1alpha1.GoogleFilestoreInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFilestoreInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleFilestoreInstances) UpdateStatus(googleFilestoreInstance *v1alpha1.GoogleFilestoreInstance) (*v1alpha1.GoogleFilestoreInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlefilestoreinstancesResource, "status", googleFilestoreInstance), &v1alpha1.GoogleFilestoreInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFilestoreInstance), err
}

// Delete takes name of the googleFilestoreInstance and deletes it. Returns an error if one occurs.
func (c *FakeGoogleFilestoreInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlefilestoreinstancesResource, name), &v1alpha1.GoogleFilestoreInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleFilestoreInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlefilestoreinstancesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleFilestoreInstanceList{})
	return err
}

// Patch applies the patch and returns the patched googleFilestoreInstance.
func (c *FakeGoogleFilestoreInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleFilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlefilestoreinstancesResource, name, pt, data, subresources...), &v1alpha1.GoogleFilestoreInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleFilestoreInstance), err
}
