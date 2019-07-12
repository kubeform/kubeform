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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeLinodeStackscripts implements LinodeStackscriptInterface
type FakeLinodeStackscripts struct {
	Fake *FakeLinodeV1alpha1
}

var linodestackscriptsResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "linodestackscripts"}

var linodestackscriptsKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "LinodeStackscript"}

// Get takes name of the linodeStackscript, and returns the corresponding linodeStackscript object, and an error if there is any.
func (c *FakeLinodeStackscripts) Get(name string, options v1.GetOptions) (result *v1alpha1.LinodeStackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(linodestackscriptsResource, name), &v1alpha1.LinodeStackscript{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeStackscript), err
}

// List takes label and field selectors, and returns the list of LinodeStackscripts that match those selectors.
func (c *FakeLinodeStackscripts) List(opts v1.ListOptions) (result *v1alpha1.LinodeStackscriptList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(linodestackscriptsResource, linodestackscriptsKind, opts), &v1alpha1.LinodeStackscriptList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LinodeStackscriptList{ListMeta: obj.(*v1alpha1.LinodeStackscriptList).ListMeta}
	for _, item := range obj.(*v1alpha1.LinodeStackscriptList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested linodeStackscripts.
func (c *FakeLinodeStackscripts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(linodestackscriptsResource, opts))
}

// Create takes the representation of a linodeStackscript and creates it.  Returns the server's representation of the linodeStackscript, and an error, if there is any.
func (c *FakeLinodeStackscripts) Create(linodeStackscript *v1alpha1.LinodeStackscript) (result *v1alpha1.LinodeStackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(linodestackscriptsResource, linodeStackscript), &v1alpha1.LinodeStackscript{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeStackscript), err
}

// Update takes the representation of a linodeStackscript and updates it. Returns the server's representation of the linodeStackscript, and an error, if there is any.
func (c *FakeLinodeStackscripts) Update(linodeStackscript *v1alpha1.LinodeStackscript) (result *v1alpha1.LinodeStackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(linodestackscriptsResource, linodeStackscript), &v1alpha1.LinodeStackscript{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeStackscript), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLinodeStackscripts) UpdateStatus(linodeStackscript *v1alpha1.LinodeStackscript) (*v1alpha1.LinodeStackscript, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(linodestackscriptsResource, "status", linodeStackscript), &v1alpha1.LinodeStackscript{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeStackscript), err
}

// Delete takes name of the linodeStackscript and deletes it. Returns an error if one occurs.
func (c *FakeLinodeStackscripts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(linodestackscriptsResource, name), &v1alpha1.LinodeStackscript{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLinodeStackscripts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(linodestackscriptsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LinodeStackscriptList{})
	return err
}

// Patch applies the patch and returns the patched linodeStackscript.
func (c *FakeLinodeStackscripts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LinodeStackscript, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(linodestackscriptsResource, name, pt, data, subresources...), &v1alpha1.LinodeStackscript{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinodeStackscript), err
}
