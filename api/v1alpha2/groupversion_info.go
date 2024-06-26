/*

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

// Package v1 contains API Schema definitions for the hnc v1 API group
// +kubebuilder:object:generate=true
// +groupName=hnc.x-k8s.io
package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register the main HNC objects.
	GroupVersion = schema.GroupVersion{Group: "hnc.x-k8s.io", Version: "v1alpha2"}

	// ResourcesGroupVersion is the group used to register resources in an HNC tree.
	ResourcesGroupVersion = schema.GroupVersion{Group: "resources.hnc.x-k8s.io", Version: "v1alpha2"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	HNCConfigurationGR       = schema.GroupResource{Group: GroupVersion.Group, Resource: "hncconfigurations"}
	HierarchyConfigurationGR = schema.GroupResource{Group: GroupVersion.Group, Resource: "hierarchyconfigurations"}
	SubnamespaceAnchorGR     = schema.GroupResource{Group: GroupVersion.Group, Resource: "subnamespaceanchors"}
	HNCConfigurationGK       = schema.GroupKind{Group: GroupVersion.Group, Kind: "HNCConfiguration"}
	HierarchyConfigurationGK = schema.GroupKind{Group: GroupVersion.Group, Kind: "HierarchyConfiguration"}
	SubnamespaceAnchorGK     = schema.GroupKind{Group: GroupVersion.Group, Kind: "SubnamespaceAnchor"}
)
