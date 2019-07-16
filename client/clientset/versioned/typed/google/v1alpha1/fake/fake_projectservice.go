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

// FakeProjectServices implements ProjectServiceInterface
type FakeProjectServices struct {
	Fake *FakeGoogleV1alpha1
}

var projectservicesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "projectservices"}

var projectservicesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ProjectService"}

// Get takes name of the projectService, and returns the corresponding projectService object, and an error if there is any.
func (c *FakeProjectServices) Get(name string, options v1.GetOptions) (result *v1alpha1.ProjectService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(projectservicesResource, name), &v1alpha1.ProjectService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectService), err
}

// List takes label and field selectors, and returns the list of ProjectServices that match those selectors.
func (c *FakeProjectServices) List(opts v1.ListOptions) (result *v1alpha1.ProjectServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(projectservicesResource, projectservicesKind, opts), &v1alpha1.ProjectServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProjectServiceList{ListMeta: obj.(*v1alpha1.ProjectServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProjectServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectServices.
func (c *FakeProjectServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(projectservicesResource, opts))
}

// Create takes the representation of a projectService and creates it.  Returns the server's representation of the projectService, and an error, if there is any.
func (c *FakeProjectServices) Create(projectService *v1alpha1.ProjectService) (result *v1alpha1.ProjectService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(projectservicesResource, projectService), &v1alpha1.ProjectService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectService), err
}

// Update takes the representation of a projectService and updates it. Returns the server's representation of the projectService, and an error, if there is any.
func (c *FakeProjectServices) Update(projectService *v1alpha1.ProjectService) (result *v1alpha1.ProjectService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(projectservicesResource, projectService), &v1alpha1.ProjectService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectServices) UpdateStatus(projectService *v1alpha1.ProjectService) (*v1alpha1.ProjectService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(projectservicesResource, "status", projectService), &v1alpha1.ProjectService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectService), err
}

// Delete takes name of the projectService and deletes it. Returns an error if one occurs.
func (c *FakeProjectServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(projectservicesResource, name), &v1alpha1.ProjectService{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(projectservicesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProjectServiceList{})
	return err
}

// Patch applies the patch and returns the patched projectService.
func (c *FakeProjectServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProjectService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(projectservicesResource, name, pt, data, subresources...), &v1alpha1.ProjectService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectService), err
}
