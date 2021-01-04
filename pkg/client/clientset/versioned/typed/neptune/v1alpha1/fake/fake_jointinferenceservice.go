// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/edgeai-neptune/neptune/pkg/apis/neptune/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJointInferenceServices implements JointInferenceServiceInterface
type FakeJointInferenceServices struct {
	Fake *FakeNeptuneV1alpha1
	ns   string
}

var jointinferenceservicesResource = schema.GroupVersionResource{Group: "neptune.io", Version: "v1alpha1", Resource: "jointinferenceservices"}

var jointinferenceservicesKind = schema.GroupVersionKind{Group: "neptune.io", Version: "v1alpha1", Kind: "JointInferenceService"}

// Get takes name of the jointInferenceService, and returns the corresponding jointInferenceService object, and an error if there is any.
func (c *FakeJointInferenceServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JointInferenceService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jointinferenceservicesResource, c.ns, name), &v1alpha1.JointInferenceService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointInferenceService), err
}

// List takes label and field selectors, and returns the list of JointInferenceServices that match those selectors.
func (c *FakeJointInferenceServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JointInferenceServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jointinferenceservicesResource, jointinferenceservicesKind, c.ns, opts), &v1alpha1.JointInferenceServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JointInferenceServiceList{ListMeta: obj.(*v1alpha1.JointInferenceServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.JointInferenceServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jointInferenceServices.
func (c *FakeJointInferenceServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jointinferenceservicesResource, c.ns, opts))

}

// Create takes the representation of a jointInferenceService and creates it.  Returns the server's representation of the jointInferenceService, and an error, if there is any.
func (c *FakeJointInferenceServices) Create(ctx context.Context, jointInferenceService *v1alpha1.JointInferenceService, opts v1.CreateOptions) (result *v1alpha1.JointInferenceService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jointinferenceservicesResource, c.ns, jointInferenceService), &v1alpha1.JointInferenceService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointInferenceService), err
}

// Update takes the representation of a jointInferenceService and updates it. Returns the server's representation of the jointInferenceService, and an error, if there is any.
func (c *FakeJointInferenceServices) Update(ctx context.Context, jointInferenceService *v1alpha1.JointInferenceService, opts v1.UpdateOptions) (result *v1alpha1.JointInferenceService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jointinferenceservicesResource, c.ns, jointInferenceService), &v1alpha1.JointInferenceService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointInferenceService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJointInferenceServices) UpdateStatus(ctx context.Context, jointInferenceService *v1alpha1.JointInferenceService, opts v1.UpdateOptions) (*v1alpha1.JointInferenceService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jointinferenceservicesResource, "status", c.ns, jointInferenceService), &v1alpha1.JointInferenceService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointInferenceService), err
}

// Delete takes name of the jointInferenceService and deletes it. Returns an error if one occurs.
func (c *FakeJointInferenceServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(jointinferenceservicesResource, c.ns, name), &v1alpha1.JointInferenceService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJointInferenceServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jointinferenceservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.JointInferenceServiceList{})
	return err
}

// Patch applies the patch and returns the patched jointInferenceService.
func (c *FakeJointInferenceServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JointInferenceService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jointinferenceservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.JointInferenceService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JointInferenceService), err
}
