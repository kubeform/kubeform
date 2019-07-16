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

// FakeCodedeployApps implements CodedeployAppInterface
type FakeCodedeployApps struct {
	Fake *FakeAwsV1alpha1
}

var codedeployappsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "codedeployapps"}

var codedeployappsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CodedeployApp"}

// Get takes name of the codedeployApp, and returns the corresponding codedeployApp object, and an error if there is any.
func (c *FakeCodedeployApps) Get(name string, options v1.GetOptions) (result *v1alpha1.CodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(codedeployappsResource, name), &v1alpha1.CodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployApp), err
}

// List takes label and field selectors, and returns the list of CodedeployApps that match those selectors.
func (c *FakeCodedeployApps) List(opts v1.ListOptions) (result *v1alpha1.CodedeployAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(codedeployappsResource, codedeployappsKind, opts), &v1alpha1.CodedeployAppList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CodedeployAppList{ListMeta: obj.(*v1alpha1.CodedeployAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.CodedeployAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested codedeployApps.
func (c *FakeCodedeployApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(codedeployappsResource, opts))
}

// Create takes the representation of a codedeployApp and creates it.  Returns the server's representation of the codedeployApp, and an error, if there is any.
func (c *FakeCodedeployApps) Create(codedeployApp *v1alpha1.CodedeployApp) (result *v1alpha1.CodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(codedeployappsResource, codedeployApp), &v1alpha1.CodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployApp), err
}

// Update takes the representation of a codedeployApp and updates it. Returns the server's representation of the codedeployApp, and an error, if there is any.
func (c *FakeCodedeployApps) Update(codedeployApp *v1alpha1.CodedeployApp) (result *v1alpha1.CodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(codedeployappsResource, codedeployApp), &v1alpha1.CodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCodedeployApps) UpdateStatus(codedeployApp *v1alpha1.CodedeployApp) (*v1alpha1.CodedeployApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(codedeployappsResource, "status", codedeployApp), &v1alpha1.CodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployApp), err
}

// Delete takes name of the codedeployApp and deletes it. Returns an error if one occurs.
func (c *FakeCodedeployApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(codedeployappsResource, name), &v1alpha1.CodedeployApp{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCodedeployApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(codedeployappsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CodedeployAppList{})
	return err
}

// Patch applies the patch and returns the patched codedeployApp.
func (c *FakeCodedeployApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CodedeployApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(codedeployappsResource, name, pt, data, subresources...), &v1alpha1.CodedeployApp{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CodedeployApp), err
}
