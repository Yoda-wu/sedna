/*
Copyright The KubeEdge Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubeedge/sedna/pkg/apis/sedna/v1alpha1"
	scheme "github.com/kubeedge/sedna/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReidJobsGetter has a method to return a ReidJobInterface.
// A group's client should implement this interface.
type ReidJobsGetter interface {
	ReidJobs(namespace string) ReidJobInterface
}

// ReidJobInterface has methods to work with ReidJob resources.
type ReidJobInterface interface {
	Create(ctx context.Context, reidJob *v1alpha1.ReidJob, opts v1.CreateOptions) (*v1alpha1.ReidJob, error)
	Update(ctx context.Context, reidJob *v1alpha1.ReidJob, opts v1.UpdateOptions) (*v1alpha1.ReidJob, error)
	UpdateStatus(ctx context.Context, reidJob *v1alpha1.ReidJob, opts v1.UpdateOptions) (*v1alpha1.ReidJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ReidJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ReidJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReidJob, err error)
	ReidJobExpansion
}

// reidJobs implements ReidJobInterface
type reidJobs struct {
	client rest.Interface
	ns     string
}

// newReidJobs returns a ReidJobs
func newReidJobs(c *SednaV1alpha1Client, namespace string) *reidJobs {
	return &reidJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the reidJob, and returns the corresponding reidJob object, and an error if there is any.
func (c *reidJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ReidJob, err error) {
	result = &v1alpha1.ReidJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reidjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReidJobs that match those selectors.
func (c *reidJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReidJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ReidJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reidjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested reidJobs.
func (c *reidJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("reidjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a reidJob and creates it.  Returns the server's representation of the reidJob, and an error, if there is any.
func (c *reidJobs) Create(ctx context.Context, reidJob *v1alpha1.ReidJob, opts v1.CreateOptions) (result *v1alpha1.ReidJob, err error) {
	result = &v1alpha1.ReidJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("reidjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reidJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a reidJob and updates it. Returns the server's representation of the reidJob, and an error, if there is any.
func (c *reidJobs) Update(ctx context.Context, reidJob *v1alpha1.ReidJob, opts v1.UpdateOptions) (result *v1alpha1.ReidJob, err error) {
	result = &v1alpha1.ReidJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reidjobs").
		Name(reidJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reidJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *reidJobs) UpdateStatus(ctx context.Context, reidJob *v1alpha1.ReidJob, opts v1.UpdateOptions) (result *v1alpha1.ReidJob, err error) {
	result = &v1alpha1.ReidJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reidjobs").
		Name(reidJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(reidJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the reidJob and deletes it. Returns an error if one occurs.
func (c *reidJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reidjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *reidJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reidjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched reidJob.
func (c *reidJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReidJob, err error) {
	result = &v1alpha1.ReidJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("reidjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
