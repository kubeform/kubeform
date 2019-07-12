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

// FakeAwsSesEventDestinations implements AwsSesEventDestinationInterface
type FakeAwsSesEventDestinations struct {
	Fake *FakeAwsV1alpha1
}

var awsseseventdestinationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsseseventdestinations"}

var awsseseventdestinationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSesEventDestination"}

// Get takes name of the awsSesEventDestination, and returns the corresponding awsSesEventDestination object, and an error if there is any.
func (c *FakeAwsSesEventDestinations) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsseseventdestinationsResource, name), &v1alpha1.AwsSesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEventDestination), err
}

// List takes label and field selectors, and returns the list of AwsSesEventDestinations that match those selectors.
func (c *FakeAwsSesEventDestinations) List(opts v1.ListOptions) (result *v1alpha1.AwsSesEventDestinationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsseseventdestinationsResource, awsseseventdestinationsKind, opts), &v1alpha1.AwsSesEventDestinationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSesEventDestinationList{ListMeta: obj.(*v1alpha1.AwsSesEventDestinationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSesEventDestinationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSesEventDestinations.
func (c *FakeAwsSesEventDestinations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsseseventdestinationsResource, opts))
}

// Create takes the representation of a awsSesEventDestination and creates it.  Returns the server's representation of the awsSesEventDestination, and an error, if there is any.
func (c *FakeAwsSesEventDestinations) Create(awsSesEventDestination *v1alpha1.AwsSesEventDestination) (result *v1alpha1.AwsSesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsseseventdestinationsResource, awsSesEventDestination), &v1alpha1.AwsSesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEventDestination), err
}

// Update takes the representation of a awsSesEventDestination and updates it. Returns the server's representation of the awsSesEventDestination, and an error, if there is any.
func (c *FakeAwsSesEventDestinations) Update(awsSesEventDestination *v1alpha1.AwsSesEventDestination) (result *v1alpha1.AwsSesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsseseventdestinationsResource, awsSesEventDestination), &v1alpha1.AwsSesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEventDestination), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSesEventDestinations) UpdateStatus(awsSesEventDestination *v1alpha1.AwsSesEventDestination) (*v1alpha1.AwsSesEventDestination, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsseseventdestinationsResource, "status", awsSesEventDestination), &v1alpha1.AwsSesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEventDestination), err
}

// Delete takes name of the awsSesEventDestination and deletes it. Returns an error if one occurs.
func (c *FakeAwsSesEventDestinations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsseseventdestinationsResource, name), &v1alpha1.AwsSesEventDestination{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSesEventDestinations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsseseventdestinationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSesEventDestinationList{})
	return err
}

// Patch applies the patch and returns the patched awsSesEventDestination.
func (c *FakeAwsSesEventDestinations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsseseventdestinationsResource, name, pt, data, subresources...), &v1alpha1.AwsSesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSesEventDestination), err
}
