// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKusciaJobs implements KusciaJobInterface
type FakeKusciaJobs struct {
	Fake *FakeKusciaV1alpha1
	ns   string
}

var kusciajobsResource = schema.GroupVersionResource{Group: "kuscia.secretflow", Version: "v1alpha1", Resource: "kusciajobs"}

var kusciajobsKind = schema.GroupVersionKind{Group: "kuscia.secretflow", Version: "v1alpha1", Kind: "KusciaJob"}

// Get takes name of the kusciaJob, and returns the corresponding kusciaJob object, and an error if there is any.
func (c *FakeKusciaJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KusciaJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kusciajobsResource, c.ns, name), &v1alpha1.KusciaJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KusciaJob), err
}

// List takes label and field selectors, and returns the list of KusciaJobs that match those selectors.
func (c *FakeKusciaJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KusciaJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kusciajobsResource, kusciajobsKind, c.ns, opts), &v1alpha1.KusciaJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KusciaJobList{ListMeta: obj.(*v1alpha1.KusciaJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.KusciaJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kusciaJobs.
func (c *FakeKusciaJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kusciajobsResource, c.ns, opts))

}

// Create takes the representation of a kusciaJob and creates it.  Returns the server's representation of the kusciaJob, and an error, if there is any.
func (c *FakeKusciaJobs) Create(ctx context.Context, kusciaJob *v1alpha1.KusciaJob, opts v1.CreateOptions) (result *v1alpha1.KusciaJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kusciajobsResource, c.ns, kusciaJob), &v1alpha1.KusciaJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KusciaJob), err
}

// Update takes the representation of a kusciaJob and updates it. Returns the server's representation of the kusciaJob, and an error, if there is any.
func (c *FakeKusciaJobs) Update(ctx context.Context, kusciaJob *v1alpha1.KusciaJob, opts v1.UpdateOptions) (result *v1alpha1.KusciaJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kusciajobsResource, c.ns, kusciaJob), &v1alpha1.KusciaJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KusciaJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKusciaJobs) UpdateStatus(ctx context.Context, kusciaJob *v1alpha1.KusciaJob, opts v1.UpdateOptions) (*v1alpha1.KusciaJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kusciajobsResource, "status", c.ns, kusciaJob), &v1alpha1.KusciaJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KusciaJob), err
}

// Delete takes name of the kusciaJob and deletes it. Returns an error if one occurs.
func (c *FakeKusciaJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kusciajobsResource, c.ns, name, opts), &v1alpha1.KusciaJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKusciaJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kusciajobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KusciaJobList{})
	return err
}

// Patch applies the patch and returns the patched kusciaJob.
func (c *FakeKusciaJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KusciaJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kusciajobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KusciaJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KusciaJob), err
}
