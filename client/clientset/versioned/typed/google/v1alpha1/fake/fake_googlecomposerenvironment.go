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

// FakeGoogleComposerEnvironments implements GoogleComposerEnvironmentInterface
type FakeGoogleComposerEnvironments struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomposerenvironmentsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomposerenvironments"}

var googlecomposerenvironmentsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComposerEnvironment"}

// Get takes name of the googleComposerEnvironment, and returns the corresponding googleComposerEnvironment object, and an error if there is any.
func (c *FakeGoogleComposerEnvironments) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComposerEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomposerenvironmentsResource, name), &v1alpha1.GoogleComposerEnvironment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComposerEnvironment), err
}

// List takes label and field selectors, and returns the list of GoogleComposerEnvironments that match those selectors.
func (c *FakeGoogleComposerEnvironments) List(opts v1.ListOptions) (result *v1alpha1.GoogleComposerEnvironmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomposerenvironmentsResource, googlecomposerenvironmentsKind, opts), &v1alpha1.GoogleComposerEnvironmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComposerEnvironmentList{ListMeta: obj.(*v1alpha1.GoogleComposerEnvironmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComposerEnvironmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComposerEnvironments.
func (c *FakeGoogleComposerEnvironments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomposerenvironmentsResource, opts))
}

// Create takes the representation of a googleComposerEnvironment and creates it.  Returns the server's representation of the googleComposerEnvironment, and an error, if there is any.
func (c *FakeGoogleComposerEnvironments) Create(googleComposerEnvironment *v1alpha1.GoogleComposerEnvironment) (result *v1alpha1.GoogleComposerEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomposerenvironmentsResource, googleComposerEnvironment), &v1alpha1.GoogleComposerEnvironment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComposerEnvironment), err
}

// Update takes the representation of a googleComposerEnvironment and updates it. Returns the server's representation of the googleComposerEnvironment, and an error, if there is any.
func (c *FakeGoogleComposerEnvironments) Update(googleComposerEnvironment *v1alpha1.GoogleComposerEnvironment) (result *v1alpha1.GoogleComposerEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomposerenvironmentsResource, googleComposerEnvironment), &v1alpha1.GoogleComposerEnvironment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComposerEnvironment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComposerEnvironments) UpdateStatus(googleComposerEnvironment *v1alpha1.GoogleComposerEnvironment) (*v1alpha1.GoogleComposerEnvironment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomposerenvironmentsResource, "status", googleComposerEnvironment), &v1alpha1.GoogleComposerEnvironment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComposerEnvironment), err
}

// Delete takes name of the googleComposerEnvironment and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComposerEnvironments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomposerenvironmentsResource, name), &v1alpha1.GoogleComposerEnvironment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComposerEnvironments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomposerenvironmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComposerEnvironmentList{})
	return err
}

// Patch applies the patch and returns the patched googleComposerEnvironment.
func (c *FakeGoogleComposerEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComposerEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomposerenvironmentsResource, name, pt, data, subresources...), &v1alpha1.GoogleComposerEnvironment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComposerEnvironment), err
}
