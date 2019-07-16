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

// FakeAppmeshMeshes implements AppmeshMeshInterface
type FakeAppmeshMeshes struct {
	Fake *FakeAwsV1alpha1
}

var appmeshmeshesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "appmeshmeshes"}

var appmeshmeshesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AppmeshMesh"}

// Get takes name of the appmeshMesh, and returns the corresponding appmeshMesh object, and an error if there is any.
func (c *FakeAppmeshMeshes) Get(name string, options v1.GetOptions) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(appmeshmeshesResource, name), &v1alpha1.AppmeshMesh{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// List takes label and field selectors, and returns the list of AppmeshMeshes that match those selectors.
func (c *FakeAppmeshMeshes) List(opts v1.ListOptions) (result *v1alpha1.AppmeshMeshList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(appmeshmeshesResource, appmeshmeshesKind, opts), &v1alpha1.AppmeshMeshList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppmeshMeshList{ListMeta: obj.(*v1alpha1.AppmeshMeshList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppmeshMeshList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appmeshMeshes.
func (c *FakeAppmeshMeshes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(appmeshmeshesResource, opts))
}

// Create takes the representation of a appmeshMesh and creates it.  Returns the server's representation of the appmeshMesh, and an error, if there is any.
func (c *FakeAppmeshMeshes) Create(appmeshMesh *v1alpha1.AppmeshMesh) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(appmeshmeshesResource, appmeshMesh), &v1alpha1.AppmeshMesh{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// Update takes the representation of a appmeshMesh and updates it. Returns the server's representation of the appmeshMesh, and an error, if there is any.
func (c *FakeAppmeshMeshes) Update(appmeshMesh *v1alpha1.AppmeshMesh) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(appmeshmeshesResource, appmeshMesh), &v1alpha1.AppmeshMesh{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppmeshMeshes) UpdateStatus(appmeshMesh *v1alpha1.AppmeshMesh) (*v1alpha1.AppmeshMesh, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(appmeshmeshesResource, "status", appmeshMesh), &v1alpha1.AppmeshMesh{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}

// Delete takes name of the appmeshMesh and deletes it. Returns an error if one occurs.
func (c *FakeAppmeshMeshes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(appmeshmeshesResource, name), &v1alpha1.AppmeshMesh{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppmeshMeshes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(appmeshmeshesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppmeshMeshList{})
	return err
}

// Patch applies the patch and returns the patched appmeshMesh.
func (c *FakeAppmeshMeshes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppmeshMesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(appmeshmeshesResource, name, pt, data, subresources...), &v1alpha1.AppmeshMesh{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppmeshMesh), err
}