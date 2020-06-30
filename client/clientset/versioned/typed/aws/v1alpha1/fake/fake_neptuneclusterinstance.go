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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNeptuneClusterInstances implements NeptuneClusterInstanceInterface
type FakeNeptuneClusterInstances struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var neptuneclusterinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "neptuneclusterinstances"}

var neptuneclusterinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "NeptuneClusterInstance"}

// Get takes name of the neptuneClusterInstance, and returns the corresponding neptuneClusterInstance object, and an error if there is any.
func (c *FakeNeptuneClusterInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NeptuneClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(neptuneclusterinstancesResource, c.ns, name), &v1alpha1.NeptuneClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneClusterInstance), err
}

// List takes label and field selectors, and returns the list of NeptuneClusterInstances that match those selectors.
func (c *FakeNeptuneClusterInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NeptuneClusterInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(neptuneclusterinstancesResource, neptuneclusterinstancesKind, c.ns, opts), &v1alpha1.NeptuneClusterInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NeptuneClusterInstanceList{ListMeta: obj.(*v1alpha1.NeptuneClusterInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.NeptuneClusterInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested neptuneClusterInstances.
func (c *FakeNeptuneClusterInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(neptuneclusterinstancesResource, c.ns, opts))

}

// Create takes the representation of a neptuneClusterInstance and creates it.  Returns the server's representation of the neptuneClusterInstance, and an error, if there is any.
func (c *FakeNeptuneClusterInstances) Create(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.CreateOptions) (result *v1alpha1.NeptuneClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(neptuneclusterinstancesResource, c.ns, neptuneClusterInstance), &v1alpha1.NeptuneClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneClusterInstance), err
}

// Update takes the representation of a neptuneClusterInstance and updates it. Returns the server's representation of the neptuneClusterInstance, and an error, if there is any.
func (c *FakeNeptuneClusterInstances) Update(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.UpdateOptions) (result *v1alpha1.NeptuneClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(neptuneclusterinstancesResource, c.ns, neptuneClusterInstance), &v1alpha1.NeptuneClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneClusterInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNeptuneClusterInstances) UpdateStatus(ctx context.Context, neptuneClusterInstance *v1alpha1.NeptuneClusterInstance, opts v1.UpdateOptions) (*v1alpha1.NeptuneClusterInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(neptuneclusterinstancesResource, "status", c.ns, neptuneClusterInstance), &v1alpha1.NeptuneClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneClusterInstance), err
}

// Delete takes name of the neptuneClusterInstance and deletes it. Returns an error if one occurs.
func (c *FakeNeptuneClusterInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(neptuneclusterinstancesResource, c.ns, name), &v1alpha1.NeptuneClusterInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNeptuneClusterInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(neptuneclusterinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NeptuneClusterInstanceList{})
	return err
}

// Patch applies the patch and returns the patched neptuneClusterInstance.
func (c *FakeNeptuneClusterInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NeptuneClusterInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(neptuneclusterinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NeptuneClusterInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NeptuneClusterInstance), err
}
