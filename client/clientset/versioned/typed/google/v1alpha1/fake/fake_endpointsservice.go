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

// FakeEndpointsServices implements EndpointsServiceInterface
type FakeEndpointsServices struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var endpointsservicesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "endpointsservices"}

var endpointsservicesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "EndpointsService"}

// Get takes name of the endpointsService, and returns the corresponding endpointsService object, and an error if there is any.
func (c *FakeEndpointsServices) Get(name string, options v1.GetOptions) (result *v1alpha1.EndpointsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsservicesResource, c.ns, name), &v1alpha1.EndpointsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointsService), err
}

// List takes label and field selectors, and returns the list of EndpointsServices that match those selectors.
func (c *FakeEndpointsServices) List(opts v1.ListOptions) (result *v1alpha1.EndpointsServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsservicesResource, endpointsservicesKind, c.ns, opts), &v1alpha1.EndpointsServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EndpointsServiceList{ListMeta: obj.(*v1alpha1.EndpointsServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.EndpointsServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpointsServices.
func (c *FakeEndpointsServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsservicesResource, c.ns, opts))

}

// Create takes the representation of a endpointsService and creates it.  Returns the server's representation of the endpointsService, and an error, if there is any.
func (c *FakeEndpointsServices) Create(endpointsService *v1alpha1.EndpointsService) (result *v1alpha1.EndpointsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsservicesResource, c.ns, endpointsService), &v1alpha1.EndpointsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointsService), err
}

// Update takes the representation of a endpointsService and updates it. Returns the server's representation of the endpointsService, and an error, if there is any.
func (c *FakeEndpointsServices) Update(endpointsService *v1alpha1.EndpointsService) (result *v1alpha1.EndpointsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsservicesResource, c.ns, endpointsService), &v1alpha1.EndpointsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointsService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEndpointsServices) UpdateStatus(endpointsService *v1alpha1.EndpointsService) (*v1alpha1.EndpointsService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(endpointsservicesResource, "status", c.ns, endpointsService), &v1alpha1.EndpointsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointsService), err
}

// Delete takes name of the endpointsService and deletes it. Returns an error if one occurs.
func (c *FakeEndpointsServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(endpointsservicesResource, c.ns, name), &v1alpha1.EndpointsService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpointsServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EndpointsServiceList{})
	return err
}

// Patch applies the patch and returns the patched endpointsService.
func (c *FakeEndpointsServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EndpointsService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EndpointsService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointsService), err
}
