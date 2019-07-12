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

// FakeAwsLbListeners implements AwsLbListenerInterface
type FakeAwsLbListeners struct {
	Fake *FakeAwsV1alpha1
}

var awslblistenersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awslblisteners"}

var awslblistenersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsLbListener"}

// Get takes name of the awsLbListener, and returns the corresponding awsLbListener object, and an error if there is any.
func (c *FakeAwsLbListeners) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsLbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awslblistenersResource, name), &v1alpha1.AwsLbListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListener), err
}

// List takes label and field selectors, and returns the list of AwsLbListeners that match those selectors.
func (c *FakeAwsLbListeners) List(opts v1.ListOptions) (result *v1alpha1.AwsLbListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awslblistenersResource, awslblistenersKind, opts), &v1alpha1.AwsLbListenerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsLbListenerList{ListMeta: obj.(*v1alpha1.AwsLbListenerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsLbListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLbListeners.
func (c *FakeAwsLbListeners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awslblistenersResource, opts))
}

// Create takes the representation of a awsLbListener and creates it.  Returns the server's representation of the awsLbListener, and an error, if there is any.
func (c *FakeAwsLbListeners) Create(awsLbListener *v1alpha1.AwsLbListener) (result *v1alpha1.AwsLbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awslblistenersResource, awsLbListener), &v1alpha1.AwsLbListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListener), err
}

// Update takes the representation of a awsLbListener and updates it. Returns the server's representation of the awsLbListener, and an error, if there is any.
func (c *FakeAwsLbListeners) Update(awsLbListener *v1alpha1.AwsLbListener) (result *v1alpha1.AwsLbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awslblistenersResource, awsLbListener), &v1alpha1.AwsLbListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListener), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsLbListeners) UpdateStatus(awsLbListener *v1alpha1.AwsLbListener) (*v1alpha1.AwsLbListener, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awslblistenersResource, "status", awsLbListener), &v1alpha1.AwsLbListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListener), err
}

// Delete takes name of the awsLbListener and deletes it. Returns an error if one occurs.
func (c *FakeAwsLbListeners) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awslblistenersResource, name), &v1alpha1.AwsLbListener{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLbListeners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awslblistenersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsLbListenerList{})
	return err
}

// Patch applies the patch and returns the patched awsLbListener.
func (c *FakeAwsLbListeners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsLbListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awslblistenersResource, name, pt, data, subresources...), &v1alpha1.AwsLbListener{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsLbListener), err
}
