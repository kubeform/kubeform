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

// FakeSesEventDestinations implements SesEventDestinationInterface
type FakeSesEventDestinations struct {
	Fake *FakeAwsV1alpha1
}

var seseventdestinationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "seseventdestinations"}

var seseventdestinationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SesEventDestination"}

// Get takes name of the sesEventDestination, and returns the corresponding sesEventDestination object, and an error if there is any.
func (c *FakeSesEventDestinations) Get(name string, options v1.GetOptions) (result *v1alpha1.SesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(seseventdestinationsResource, name), &v1alpha1.SesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesEventDestination), err
}

// List takes label and field selectors, and returns the list of SesEventDestinations that match those selectors.
func (c *FakeSesEventDestinations) List(opts v1.ListOptions) (result *v1alpha1.SesEventDestinationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(seseventdestinationsResource, seseventdestinationsKind, opts), &v1alpha1.SesEventDestinationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SesEventDestinationList{ListMeta: obj.(*v1alpha1.SesEventDestinationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SesEventDestinationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sesEventDestinations.
func (c *FakeSesEventDestinations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(seseventdestinationsResource, opts))
}

// Create takes the representation of a sesEventDestination and creates it.  Returns the server's representation of the sesEventDestination, and an error, if there is any.
func (c *FakeSesEventDestinations) Create(sesEventDestination *v1alpha1.SesEventDestination) (result *v1alpha1.SesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(seseventdestinationsResource, sesEventDestination), &v1alpha1.SesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesEventDestination), err
}

// Update takes the representation of a sesEventDestination and updates it. Returns the server's representation of the sesEventDestination, and an error, if there is any.
func (c *FakeSesEventDestinations) Update(sesEventDestination *v1alpha1.SesEventDestination) (result *v1alpha1.SesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(seseventdestinationsResource, sesEventDestination), &v1alpha1.SesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesEventDestination), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSesEventDestinations) UpdateStatus(sesEventDestination *v1alpha1.SesEventDestination) (*v1alpha1.SesEventDestination, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(seseventdestinationsResource, "status", sesEventDestination), &v1alpha1.SesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesEventDestination), err
}

// Delete takes name of the sesEventDestination and deletes it. Returns an error if one occurs.
func (c *FakeSesEventDestinations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(seseventdestinationsResource, name), &v1alpha1.SesEventDestination{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSesEventDestinations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(seseventdestinationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SesEventDestinationList{})
	return err
}

// Patch applies the patch and returns the patched sesEventDestination.
func (c *FakeSesEventDestinations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesEventDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(seseventdestinationsResource, name, pt, data, subresources...), &v1alpha1.SesEventDestination{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesEventDestination), err
}