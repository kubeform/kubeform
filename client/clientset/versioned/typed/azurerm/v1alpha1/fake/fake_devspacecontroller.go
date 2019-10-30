/*
Copyright The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDevspaceControllers implements DevspaceControllerInterface
type FakeDevspaceControllers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var devspacecontrollersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "devspacecontrollers"}

var devspacecontrollersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DevspaceController"}

// Get takes name of the devspaceController, and returns the corresponding devspaceController object, and an error if there is any.
func (c *FakeDevspaceControllers) Get(name string, options v1.GetOptions) (result *v1alpha1.DevspaceController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devspacecontrollersResource, c.ns, name), &v1alpha1.DevspaceController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevspaceController), err
}

// List takes label and field selectors, and returns the list of DevspaceControllers that match those selectors.
func (c *FakeDevspaceControllers) List(opts v1.ListOptions) (result *v1alpha1.DevspaceControllerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devspacecontrollersResource, devspacecontrollersKind, c.ns, opts), &v1alpha1.DevspaceControllerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevspaceControllerList{ListMeta: obj.(*v1alpha1.DevspaceControllerList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevspaceControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devspaceControllers.
func (c *FakeDevspaceControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devspacecontrollersResource, c.ns, opts))

}

// Create takes the representation of a devspaceController and creates it.  Returns the server's representation of the devspaceController, and an error, if there is any.
func (c *FakeDevspaceControllers) Create(devspaceController *v1alpha1.DevspaceController) (result *v1alpha1.DevspaceController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devspacecontrollersResource, c.ns, devspaceController), &v1alpha1.DevspaceController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevspaceController), err
}

// Update takes the representation of a devspaceController and updates it. Returns the server's representation of the devspaceController, and an error, if there is any.
func (c *FakeDevspaceControllers) Update(devspaceController *v1alpha1.DevspaceController) (result *v1alpha1.DevspaceController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devspacecontrollersResource, c.ns, devspaceController), &v1alpha1.DevspaceController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevspaceController), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevspaceControllers) UpdateStatus(devspaceController *v1alpha1.DevspaceController) (*v1alpha1.DevspaceController, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devspacecontrollersResource, "status", c.ns, devspaceController), &v1alpha1.DevspaceController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevspaceController), err
}

// Delete takes name of the devspaceController and deletes it. Returns an error if one occurs.
func (c *FakeDevspaceControllers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devspacecontrollersResource, c.ns, name), &v1alpha1.DevspaceController{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevspaceControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devspacecontrollersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevspaceControllerList{})
	return err
}

// Patch applies the patch and returns the patched devspaceController.
func (c *FakeDevspaceControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevspaceController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devspacecontrollersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevspaceController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevspaceController), err
}
