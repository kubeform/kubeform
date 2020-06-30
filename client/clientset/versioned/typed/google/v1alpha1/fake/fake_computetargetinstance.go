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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeTargetInstances implements ComputeTargetInstanceInterface
type FakeComputeTargetInstances struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computetargetinstancesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computetargetinstances"}

var computetargetinstancesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeTargetInstance"}

// Get takes name of the computeTargetInstance, and returns the corresponding computeTargetInstance object, and an error if there is any.
func (c *FakeComputeTargetInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeTargetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computetargetinstancesResource, c.ns, name), &v1alpha1.ComputeTargetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetInstance), err
}

// List takes label and field selectors, and returns the list of ComputeTargetInstances that match those selectors.
func (c *FakeComputeTargetInstances) List(opts v1.ListOptions) (result *v1alpha1.ComputeTargetInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computetargetinstancesResource, computetargetinstancesKind, c.ns, opts), &v1alpha1.ComputeTargetInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeTargetInstanceList{ListMeta: obj.(*v1alpha1.ComputeTargetInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeTargetInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeTargetInstances.
func (c *FakeComputeTargetInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computetargetinstancesResource, c.ns, opts))

}

// Create takes the representation of a computeTargetInstance and creates it.  Returns the server's representation of the computeTargetInstance, and an error, if there is any.
func (c *FakeComputeTargetInstances) Create(computeTargetInstance *v1alpha1.ComputeTargetInstance) (result *v1alpha1.ComputeTargetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computetargetinstancesResource, c.ns, computeTargetInstance), &v1alpha1.ComputeTargetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetInstance), err
}

// Update takes the representation of a computeTargetInstance and updates it. Returns the server's representation of the computeTargetInstance, and an error, if there is any.
func (c *FakeComputeTargetInstances) Update(computeTargetInstance *v1alpha1.ComputeTargetInstance) (result *v1alpha1.ComputeTargetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computetargetinstancesResource, c.ns, computeTargetInstance), &v1alpha1.ComputeTargetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeTargetInstances) UpdateStatus(computeTargetInstance *v1alpha1.ComputeTargetInstance) (*v1alpha1.ComputeTargetInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computetargetinstancesResource, "status", c.ns, computeTargetInstance), &v1alpha1.ComputeTargetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetInstance), err
}

// Delete takes name of the computeTargetInstance and deletes it. Returns an error if one occurs.
func (c *FakeComputeTargetInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computetargetinstancesResource, c.ns, name), &v1alpha1.ComputeTargetInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeTargetInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computetargetinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeTargetInstanceList{})
	return err
}

// Patch applies the patch and returns the patched computeTargetInstance.
func (c *FakeComputeTargetInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeTargetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computetargetinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeTargetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeTargetInstance), err
}
