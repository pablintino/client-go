package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSImageStream describes a set of streams and associated images available
// for the MachineConfigPools to be used as base OS images.
//
// The resource is a singleton named "cluster".
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=osimagestreams,scope=Cluster
// +kubebuilder:subresource:status
// +openshift:api-approved.openshift.io=https://github.com/openshift/api/pull/2555
// +openshift:file-pattern=cvoRunLevel=0000_80,operatorName=machine-config,operatorOrdering=01
// +openshift:enable:FeatureGate=OSStreams
// +kubebuilder:metadata:labels=openshift.io/operator-managed=
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'cluster'",message="osimagestream is a singleton, .metadata.name must be 'cluster'"
type OSImageStream struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec contains the desired OSImageStream config configuration.
	// +required
	Spec *OSImageStreamSpec `json:"spec,omitempty"`

	// status describes the last observed state of this OSImageStream.
	// Populated by the MachineConfigOperator after reading release metadata.
	// When not present, the controller has not yet reconciled this resource.
	// +optional
	Status OSImageStreamStatus `json:"status,omitempty,omitzero"`
}

// OSImageStreamStatus describes the current state of a OSImageStream.
// +kubebuilder:validation:XValidation:rule="!has(self.defaultStream) || self.defaultStream in self.availableStreams.map(s, s.name)",message="defaultStream must reference a stream name from availableStreams"
type OSImageStreamStatus struct {
	// availableStreams is a list of the available OS Image Streams
	// available and their associated URLs for both OS and Extensions
	// images.
	//
	// Must have at least one item and may not exceed 100 items.
	// +required
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=100
	// +listType=map
	// +listMapKey=name
	AvailableStreams []OSImageStreamSet `json:"availableStreams,omitempty"`

	// defaultStream is the name of the stream that should be used as the default
	// when no specific stream is requested by a MachineConfigPool.
	// Must reference the name of one of the streams in availableStreams.
	//
	// Must have at least one item and may not exceed 70 items.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=70
	// +kubebuilder:validation:XValidation:rule=`self.matches('^[\\w\\.\\-]+$')`,message="The name must consist only of alphanumeric characters, hyphens ('-') and dots ('.')."
	DefaultStream string `json:"defaultStream,omitempty"`
}

// OSImageStreamSpec defines the desired state of a OSImageStream.
type OSImageStreamSpec struct {
}

type OSImageStreamSet struct {
	// name is the identifier of the stream. This name must be suitable for use
	// as a container image tag or label, which restricts it to simple alphanumeric characters.
	//
	// Must not be empty and must not exceed 70 characters in length.
	// Must only contain alphanumeric characters, hyphens ('-'), or dots ('.').
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=70
	// +kubebuilder:validation:XValidation:rule=`self.matches('^[\\w\\.\\-]+$')`,message="The name must consist only of alphanumeric characters, hyphens ('-') and dots ('.')."
	Name string `json:"name,omitempty"`

	// osImage is an OS Image referenced by digest.
	//
	// The format of the image pull spec is: host[:port][/namespace]/name@sha256:<digest>,
	// where the digest must be 64 characters long, and consist only of lowercase hexadecimal characters, a-f and 0-9.
	// The length of the whole spec must be between 1 to 447 characters.
	// +required
	OSImage ImageDigestFormat `json:"osImage,omitempty"`

	// osExtensionsImage is an OS Extensions Image referenced by digest.
	//
	// The format of the image pull spec is: host[:port][/namespace]/name@sha256:<digest>,
	// where the digest must be 64 characters long, and consist only of lowercase hexadecimal characters, a-f and 0-9.
	// The length of the whole spec must be between 1 to 447 characters.
	// +required
	OSExtensionsImage ImageDigestFormat `json:"osExtensionsImage,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSImageStreamList is a list of OSImageStream resources
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
type OSImageStreamList struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard list's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	Items []OSImageStream `json:"items"`
}
