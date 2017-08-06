package crd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme add the Redis cluster resource definition to a specified scheme.
	AddToScheme = schemeBuilder.AddToScheme
)

// SchemeGroupVersion is the group version used to register these objects.
var SchemeGroupVersion = schema.GroupVersion{Group: RedisClusterDefinitionGroup, Version: RedisClusterDefinitionVersion}

// Resource takes an unqualified resource and returns a Group-qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds the set of types defined in this package to the supplied scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&RedisCluster{},
		&RedisClusterList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)

	return nil
}
