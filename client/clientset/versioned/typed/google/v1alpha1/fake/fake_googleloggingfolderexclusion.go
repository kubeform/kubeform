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

// FakeGoogleLoggingFolderExclusions implements GoogleLoggingFolderExclusionInterface
type FakeGoogleLoggingFolderExclusions struct {
	Fake *FakeGoogleV1alpha1
}

var googleloggingfolderexclusionsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googleloggingfolderexclusions"}

var googleloggingfolderexclusionsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleLoggingFolderExclusion"}

// Get takes name of the googleLoggingFolderExclusion, and returns the corresponding googleLoggingFolderExclusion object, and an error if there is any.
func (c *FakeGoogleLoggingFolderExclusions) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleLoggingFolderExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googleloggingfolderexclusionsResource, name), &v1alpha1.GoogleLoggingFolderExclusion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingFolderExclusion), err
}

// List takes label and field selectors, and returns the list of GoogleLoggingFolderExclusions that match those selectors.
func (c *FakeGoogleLoggingFolderExclusions) List(opts v1.ListOptions) (result *v1alpha1.GoogleLoggingFolderExclusionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googleloggingfolderexclusionsResource, googleloggingfolderexclusionsKind, opts), &v1alpha1.GoogleLoggingFolderExclusionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleLoggingFolderExclusionList{ListMeta: obj.(*v1alpha1.GoogleLoggingFolderExclusionList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleLoggingFolderExclusionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleLoggingFolderExclusions.
func (c *FakeGoogleLoggingFolderExclusions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googleloggingfolderexclusionsResource, opts))
}

// Create takes the representation of a googleLoggingFolderExclusion and creates it.  Returns the server's representation of the googleLoggingFolderExclusion, and an error, if there is any.
func (c *FakeGoogleLoggingFolderExclusions) Create(googleLoggingFolderExclusion *v1alpha1.GoogleLoggingFolderExclusion) (result *v1alpha1.GoogleLoggingFolderExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googleloggingfolderexclusionsResource, googleLoggingFolderExclusion), &v1alpha1.GoogleLoggingFolderExclusion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingFolderExclusion), err
}

// Update takes the representation of a googleLoggingFolderExclusion and updates it. Returns the server's representation of the googleLoggingFolderExclusion, and an error, if there is any.
func (c *FakeGoogleLoggingFolderExclusions) Update(googleLoggingFolderExclusion *v1alpha1.GoogleLoggingFolderExclusion) (result *v1alpha1.GoogleLoggingFolderExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googleloggingfolderexclusionsResource, googleLoggingFolderExclusion), &v1alpha1.GoogleLoggingFolderExclusion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingFolderExclusion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleLoggingFolderExclusions) UpdateStatus(googleLoggingFolderExclusion *v1alpha1.GoogleLoggingFolderExclusion) (*v1alpha1.GoogleLoggingFolderExclusion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googleloggingfolderexclusionsResource, "status", googleLoggingFolderExclusion), &v1alpha1.GoogleLoggingFolderExclusion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingFolderExclusion), err
}

// Delete takes name of the googleLoggingFolderExclusion and deletes it. Returns an error if one occurs.
func (c *FakeGoogleLoggingFolderExclusions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googleloggingfolderexclusionsResource, name), &v1alpha1.GoogleLoggingFolderExclusion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleLoggingFolderExclusions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googleloggingfolderexclusionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleLoggingFolderExclusionList{})
	return err
}

// Patch applies the patch and returns the patched googleLoggingFolderExclusion.
func (c *FakeGoogleLoggingFolderExclusions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleLoggingFolderExclusion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googleloggingfolderexclusionsResource, name, pt, data, subresources...), &v1alpha1.GoogleLoggingFolderExclusion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleLoggingFolderExclusion), err
}
