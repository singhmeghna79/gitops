package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GitCredentialType string

const (
	GitCredentialTypeSSH           GitCredentialType = "ssh"
	GitCredentialTypeHTTPBasicAuth GitCredentialType = "httpbasicauth"
)

// GitTrackSpec defines the desired state of GitTrack
// +k8s:openapi-gen=true
type GitTrackSpec struct {
	Repository string `json:"repository"`
	// +kubebuilder:validation:Pattern=^[a-zA-Z0-9/\-.]*$
	SubPath   string            `json:"subPath,omitempty"`
	Branch    string            `json:"branch,omitempty"`
	DeployKey GitTrackDeployKey `json:"deployKey,omitempty"`
}

// GitTrackStatus defines the observed state of GitTrack
// +k8s:openapi-gen=true
type GitTrackStatus struct {
	Status            v1.ConditionStatus `json:"status"`
	LastCommitApplied string             `json:"lastCommitApplied"`
	LastCommitFetched string             `json:"lastCommitFetched"`
}

type GitTrackDeployKey struct {
	SecretName      string `json:"secretName"`
	SecretNamespace string `json:"secretNamespace,omitempty"`
	Key             string `json:"key"`
	// +kubebuilder:validation:Enum=SSH,HTTPBasicAuth
	Type GitCredentialType `json:"type,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitTrack is the Schema for the gittracks API
// +k8s:openapi-gen=true
type GitTrack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitTrackSpec   `json:"spec,omitempty"`
	Status GitTrackStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitTrackList contains a list of GitTrack
type GitTrackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitTrack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitTrack{}, &GitTrackList{})
}
