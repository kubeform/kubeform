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

// FakeLoggingProjectSinks implements LoggingProjectSinkInterface
type FakeLoggingProjectSinks struct {
	Fake *FakeGoogleV1alpha1
}

var loggingprojectsinksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "loggingprojectsinks"}

var loggingprojectsinksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "LoggingProjectSink"}

// Get takes name of the loggingProjectSink, and returns the corresponding loggingProjectSink object, and an error if there is any.
func (c *FakeLoggingProjectSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.LoggingProjectSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(loggingprojectsinksResource, name), &v1alpha1.LoggingProjectSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingProjectSink), err
}

// List takes label and field selectors, and returns the list of LoggingProjectSinks that match those selectors.
func (c *FakeLoggingProjectSinks) List(opts v1.ListOptions) (result *v1alpha1.LoggingProjectSinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(loggingprojectsinksResource, loggingprojectsinksKind, opts), &v1alpha1.LoggingProjectSinkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoggingProjectSinkList{ListMeta: obj.(*v1alpha1.LoggingProjectSinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoggingProjectSinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loggingProjectSinks.
func (c *FakeLoggingProjectSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(loggingprojectsinksResource, opts))
}

// Create takes the representation of a loggingProjectSink and creates it.  Returns the server's representation of the loggingProjectSink, and an error, if there is any.
func (c *FakeLoggingProjectSinks) Create(loggingProjectSink *v1alpha1.LoggingProjectSink) (result *v1alpha1.LoggingProjectSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(loggingprojectsinksResource, loggingProjectSink), &v1alpha1.LoggingProjectSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingProjectSink), err
}

// Update takes the representation of a loggingProjectSink and updates it. Returns the server's representation of the loggingProjectSink, and an error, if there is any.
func (c *FakeLoggingProjectSinks) Update(loggingProjectSink *v1alpha1.LoggingProjectSink) (result *v1alpha1.LoggingProjectSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(loggingprojectsinksResource, loggingProjectSink), &v1alpha1.LoggingProjectSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingProjectSink), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoggingProjectSinks) UpdateStatus(loggingProjectSink *v1alpha1.LoggingProjectSink) (*v1alpha1.LoggingProjectSink, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(loggingprojectsinksResource, "status", loggingProjectSink), &v1alpha1.LoggingProjectSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingProjectSink), err
}

// Delete takes name of the loggingProjectSink and deletes it. Returns an error if one occurs.
func (c *FakeLoggingProjectSinks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(loggingprojectsinksResource, name), &v1alpha1.LoggingProjectSink{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoggingProjectSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(loggingprojectsinksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoggingProjectSinkList{})
	return err
}

// Patch applies the patch and returns the patched loggingProjectSink.
func (c *FakeLoggingProjectSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingProjectSink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(loggingprojectsinksResource, name, pt, data, subresources...), &v1alpha1.LoggingProjectSink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoggingProjectSink), err
}
