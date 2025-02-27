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

type APICacheObservation struct {

	// AppSync API ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APICacheParameters struct {

	// Caching behavior. Valid values are FULL_REQUEST_CACHING and PER_RESOLVER_CACHING.
	// +kubebuilder:validation:Required
	APICachingBehavior *string `json:"apiCachingBehavior" tf:"api_caching_behavior,omitempty"`

	// GraphQL API ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.GraphQLAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// At-rest encryption flag for cache. You cannot update this setting after creation.
	// +kubebuilder:validation:Optional
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// TTL in seconds for cache entries.
	// +kubebuilder:validation:Required
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`

	// Transit encryption flag when connecting to cache. You cannot update this setting after creation.
	// +kubebuilder:validation:Optional
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`

	// Cache instance type. Valid values are SMALL, MEDIUM, LARGE, XLARGE, LARGE_2X, LARGE_4X, LARGE_8X, LARGE_12X, T2_SMALL, T2_MEDIUM, R4_LARGE, R4_XLARGE, R4_2XLARGE, R4_4XLARGE, R4_8XLARGE.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// APICacheSpec defines the desired state of APICache
type APICacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APICacheParameters `json:"forProvider"`
}

// APICacheStatus defines the observed state of APICache.
type APICacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APICacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APICache is the Schema for the APICaches API. Provides an AppSync API Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type APICache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APICacheSpec   `json:"spec"`
	Status            APICacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APICacheList contains a list of APICaches
type APICacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APICache `json:"items"`
}

// Repository type metadata.
var (
	APICache_Kind             = "APICache"
	APICache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APICache_Kind}.String()
	APICache_KindAPIVersion   = APICache_Kind + "." + CRDGroupVersion.String()
	APICache_GroupVersionKind = CRDGroupVersion.WithKind(APICache_Kind)
)

func init() {
	SchemeBuilder.Register(&APICache{}, &APICacheList{})
}
