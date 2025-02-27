/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IPGroupObservation struct {

	// The IP group identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type IPGroupParameters struct {

	// The description of the IP group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the IP group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	// +kubebuilder:validation:Optional
	Rules []RulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// The description of the IP group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP address range, in CIDR notation, e.g., 10.0.0.0/16
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

// IPGroupSpec defines the desired state of IPGroup
type IPGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPGroupParameters `json:"forProvider"`
}

// IPGroupStatus defines the observed state of IPGroup.
type IPGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroup is the Schema for the IPGroups API. Provides an IP access control group in AWS WorkSpaces Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IPGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPGroupSpec   `json:"spec"`
	Status            IPGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroupList contains a list of IPGroups
type IPGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPGroup `json:"items"`
}

// Repository type metadata.
var (
	IPGroup_Kind             = "IPGroup"
	IPGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPGroup_Kind}.String()
	IPGroup_KindAPIVersion   = IPGroup_Kind + "." + CRDGroupVersion.String()
	IPGroup_GroupVersionKind = CRDGroupVersion.WithKind(IPGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IPGroup{}, &IPGroupList{})
}
