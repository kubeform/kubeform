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

// FakeGoogleComputeHealthChecks implements GoogleComputeHealthCheckInterface
type FakeGoogleComputeHealthChecks struct {
	Fake *FakeGoogleV1alpha1
}

var googlecomputehealthchecksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "googlecomputehealthchecks"}

var googlecomputehealthchecksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "GoogleComputeHealthCheck"}

// Get takes name of the googleComputeHealthCheck, and returns the corresponding googleComputeHealthCheck object, and an error if there is any.
func (c *FakeGoogleComputeHealthChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googlecomputehealthchecksResource, name), &v1alpha1.GoogleComputeHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeHealthCheck), err
}

// List takes label and field selectors, and returns the list of GoogleComputeHealthChecks that match those selectors.
func (c *FakeGoogleComputeHealthChecks) List(opts v1.ListOptions) (result *v1alpha1.GoogleComputeHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googlecomputehealthchecksResource, googlecomputehealthchecksKind, opts), &v1alpha1.GoogleComputeHealthCheckList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleComputeHealthCheckList{ListMeta: obj.(*v1alpha1.GoogleComputeHealthCheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleComputeHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleComputeHealthChecks.
func (c *FakeGoogleComputeHealthChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googlecomputehealthchecksResource, opts))
}

// Create takes the representation of a googleComputeHealthCheck and creates it.  Returns the server's representation of the googleComputeHealthCheck, and an error, if there is any.
func (c *FakeGoogleComputeHealthChecks) Create(googleComputeHealthCheck *v1alpha1.GoogleComputeHealthCheck) (result *v1alpha1.GoogleComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googlecomputehealthchecksResource, googleComputeHealthCheck), &v1alpha1.GoogleComputeHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeHealthCheck), err
}

// Update takes the representation of a googleComputeHealthCheck and updates it. Returns the server's representation of the googleComputeHealthCheck, and an error, if there is any.
func (c *FakeGoogleComputeHealthChecks) Update(googleComputeHealthCheck *v1alpha1.GoogleComputeHealthCheck) (result *v1alpha1.GoogleComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googlecomputehealthchecksResource, googleComputeHealthCheck), &v1alpha1.GoogleComputeHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleComputeHealthChecks) UpdateStatus(googleComputeHealthCheck *v1alpha1.GoogleComputeHealthCheck) (*v1alpha1.GoogleComputeHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(googlecomputehealthchecksResource, "status", googleComputeHealthCheck), &v1alpha1.GoogleComputeHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeHealthCheck), err
}

// Delete takes name of the googleComputeHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeGoogleComputeHealthChecks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(googlecomputehealthchecksResource, name), &v1alpha1.GoogleComputeHealthCheck{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleComputeHealthChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googlecomputehealthchecksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleComputeHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched googleComputeHealthCheck.
func (c *FakeGoogleComputeHealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googlecomputehealthchecksResource, name, pt, data, subresources...), &v1alpha1.GoogleComputeHealthCheck{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleComputeHealthCheck), err
}
