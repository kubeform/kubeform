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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeHealthChecks implements ComputeHealthCheckInterface
type FakeComputeHealthChecks struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computehealthchecksResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computehealthchecks"}

var computehealthchecksKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeHealthCheck"}

// Get takes name of the computeHealthCheck, and returns the corresponding computeHealthCheck object, and an error if there is any.
func (c *FakeComputeHealthChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computehealthchecksResource, c.ns, name), &v1alpha1.ComputeHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHealthCheck), err
}

// List takes label and field selectors, and returns the list of ComputeHealthChecks that match those selectors.
func (c *FakeComputeHealthChecks) List(opts v1.ListOptions) (result *v1alpha1.ComputeHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computehealthchecksResource, computehealthchecksKind, c.ns, opts), &v1alpha1.ComputeHealthCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeHealthCheckList{ListMeta: obj.(*v1alpha1.ComputeHealthCheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeHealthChecks.
func (c *FakeComputeHealthChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computehealthchecksResource, c.ns, opts))

}

// Create takes the representation of a computeHealthCheck and creates it.  Returns the server's representation of the computeHealthCheck, and an error, if there is any.
func (c *FakeComputeHealthChecks) Create(computeHealthCheck *v1alpha1.ComputeHealthCheck) (result *v1alpha1.ComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computehealthchecksResource, c.ns, computeHealthCheck), &v1alpha1.ComputeHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHealthCheck), err
}

// Update takes the representation of a computeHealthCheck and updates it. Returns the server's representation of the computeHealthCheck, and an error, if there is any.
func (c *FakeComputeHealthChecks) Update(computeHealthCheck *v1alpha1.ComputeHealthCheck) (result *v1alpha1.ComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computehealthchecksResource, c.ns, computeHealthCheck), &v1alpha1.ComputeHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeHealthChecks) UpdateStatus(computeHealthCheck *v1alpha1.ComputeHealthCheck) (*v1alpha1.ComputeHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computehealthchecksResource, "status", c.ns, computeHealthCheck), &v1alpha1.ComputeHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHealthCheck), err
}

// Delete takes name of the computeHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeComputeHealthChecks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computehealthchecksResource, c.ns, name), &v1alpha1.ComputeHealthCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeHealthChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computehealthchecksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched computeHealthCheck.
func (c *FakeComputeHealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computehealthchecksResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeHealthCheck), err
}
