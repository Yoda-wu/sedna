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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubeedge/sedna/pkg/apis/sedna/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ObjectTrackingServiceLister helps list ObjectTrackingServices.
// All objects returned here must be treated as read-only.
type ObjectTrackingServiceLister interface {
	// List lists all ObjectTrackingServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ObjectTrackingService, err error)
	// ObjectTrackingServices returns an object that can list and get ObjectTrackingServices.
	ObjectTrackingServices(namespace string) ObjectTrackingServiceNamespaceLister
	ObjectTrackingServiceListerExpansion
}

// objectTrackingServiceLister implements the ObjectTrackingServiceLister interface.
type objectTrackingServiceLister struct {
	indexer cache.Indexer
}

// NewObjectTrackingServiceLister returns a new ObjectTrackingServiceLister.
func NewObjectTrackingServiceLister(indexer cache.Indexer) ObjectTrackingServiceLister {
	return &objectTrackingServiceLister{indexer: indexer}
}

// List lists all ObjectTrackingServices in the indexer.
func (s *objectTrackingServiceLister) List(selector labels.Selector) (ret []*v1alpha1.ObjectTrackingService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ObjectTrackingService))
	})
	return ret, err
}

// ObjectTrackingServices returns an object that can list and get ObjectTrackingServices.
func (s *objectTrackingServiceLister) ObjectTrackingServices(namespace string) ObjectTrackingServiceNamespaceLister {
	return objectTrackingServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ObjectTrackingServiceNamespaceLister helps list and get ObjectTrackingServices.
// All objects returned here must be treated as read-only.
type ObjectTrackingServiceNamespaceLister interface {
	// List lists all ObjectTrackingServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ObjectTrackingService, err error)
	// Get retrieves the ObjectTrackingService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ObjectTrackingService, error)
	ObjectTrackingServiceNamespaceListerExpansion
}

// objectTrackingServiceNamespaceLister implements the ObjectTrackingServiceNamespaceLister
// interface.
type objectTrackingServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ObjectTrackingServices in the indexer for a given namespace.
func (s objectTrackingServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ObjectTrackingService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ObjectTrackingService))
	})
	return ret, err
}

// Get retrieves the ObjectTrackingService from the indexer for a given namespace and name.
func (s objectTrackingServiceNamespaceLister) Get(name string) (*v1alpha1.ObjectTrackingService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("objecttrackingservice"), name)
	}
	return obj.(*v1alpha1.ObjectTrackingService), nil
}
