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

type NetworkPerformanceMetricSubscriptionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The data aggregation time for the subscription.
	Period *string `json:"period,omitempty" tf:"period,omitempty"`
}

type NetworkPerformanceMetricSubscriptionParameters struct {

	// The target Region or Availability Zone that the metric subscription is enabled for. For example, eu-west-1.
	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// The metric used for the enabled subscription. Valid values: aggregate-latency. Default: aggregate-latency.
	// +kubebuilder:validation:Optional
	Metric *string `json:"metric,omitempty" tf:"metric,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The source Region or Availability Zone that the metric subscription is enabled for. For example, us-east-1.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`

	// The statistic used for the enabled subscription. Valid values: p50. Default: p50.
	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`
}

// NetworkPerformanceMetricSubscriptionSpec defines the desired state of NetworkPerformanceMetricSubscription
type NetworkPerformanceMetricSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkPerformanceMetricSubscriptionParameters `json:"forProvider"`
}

// NetworkPerformanceMetricSubscriptionStatus defines the observed state of NetworkPerformanceMetricSubscription.
type NetworkPerformanceMetricSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkPerformanceMetricSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPerformanceMetricSubscription is the Schema for the NetworkPerformanceMetricSubscriptions API. Provides a resource to manage an Infrastructure Performance subscription.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkPerformanceMetricSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPerformanceMetricSubscriptionSpec   `json:"spec"`
	Status            NetworkPerformanceMetricSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPerformanceMetricSubscriptionList contains a list of NetworkPerformanceMetricSubscriptions
type NetworkPerformanceMetricSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkPerformanceMetricSubscription `json:"items"`
}

// Repository type metadata.
var (
	NetworkPerformanceMetricSubscription_Kind             = "NetworkPerformanceMetricSubscription"
	NetworkPerformanceMetricSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkPerformanceMetricSubscription_Kind}.String()
	NetworkPerformanceMetricSubscription_KindAPIVersion   = NetworkPerformanceMetricSubscription_Kind + "." + CRDGroupVersion.String()
	NetworkPerformanceMetricSubscription_GroupVersionKind = CRDGroupVersion.WithKind(NetworkPerformanceMetricSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkPerformanceMetricSubscription{}, &NetworkPerformanceMetricSubscriptionList{})
}
