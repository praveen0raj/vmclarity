// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package models

import (
	"encoding/json"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// Defines values for CloudProvider.
const (
	AWS CloudProvider = "AWS"
)

// Defines values for MalwareType.
const (
	ADWARE     MalwareType = "ADWARE"
	RANSOMWARE MalwareType = "RANSOMWARE"
	SPYWARE    MalwareType = "SPYWARE"
	TROJAN     MalwareType = "TROJAN"
	VIRUS      MalwareType = "VIRUS"
	WORM       MalwareType = "WORM"
)

// Defines values for RootkitType.
const (
	APPLICATION RootkitType = "APPLICATION"
	FIRMWARE    RootkitType = "FIRMWARE"
	KERNEL      RootkitType = "KERNEL"
	MEMORY      RootkitType = "MEMORY"
)

// Defines values for ScanType.
const (
	EXPLOIT          ScanType = "EXPLOIT"
	MALWARE          ScanType = "MALWARE"
	MISCONFIGURATION ScanType = "MISCONFIGURATION"
	ROOTKIT          ScanType = "ROOTKIT"
	SBOM             ScanType = "SBOM"
	SECRET           ScanType = "SECRET"
	VULNERABILITY    ScanType = "VULNERABILITY"
)

// Defines values for TargetType.
const (
	DIR TargetType = "DIR"
	POD TargetType = "POD"
	VM  TargetType = "VM"
)

// ApiResponse An object that is returned in all cases of failures.
type ApiResponse struct {
	Message *string `json:"message,omitempty"`
}

// CloudProvider defines model for CloudProvider.
type CloudProvider string

// DirInfo defines model for DirInfo.
type DirInfo struct {
	DirName  *string `json:"dirName,omitempty"`
	Location *string `json:"location,omitempty"`
}

// ExploitInfo defines model for ExploitInfo.
type ExploitInfo struct {
	Description     *string   `json:"description,omitempty"`
	Id              *string   `json:"id,omitempty"`
	Vulnerabilities *[]string `json:"vulnerabilities,omitempty"`
}

// ExploitScan defines model for ExploitScan.
type ExploitScan struct {
	Packages *[]ExploitInfo `json:"packages,omitempty"`
}

// MalwareInfo defines model for MalwareInfo.
type MalwareInfo struct {
	Id          *string      `json:"id,omitempty"`
	MalwareName *string      `json:"malwareName,omitempty"`
	MalwareType *MalwareType `json:"malwareType,omitempty"`

	// Path Path of the file that contains malware
	Path *string `json:"path,omitempty"`
}

// MalwareScan defines model for MalwareScan.
type MalwareScan struct {
	Packages *[]MalwareInfo `json:"packages,omitempty"`
}

// MalwareType defines model for MalwareType.
type MalwareType string

// MisconfigurationInfo defines model for MisconfigurationInfo.
type MisconfigurationInfo struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`

	// Path Path of the file that contains misconfigurations
	Path *string `json:"path,omitempty"`
}

// MisconfigurationScan defines model for MisconfigurationScan.
type MisconfigurationScan struct {
	Packages *[]MisconfigurationInfo `json:"packages,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Id          *string      `json:"id,omitempty"`
	PackageInfo *PackageInfo `json:"packageInfo,omitempty"`
}

// PackageInfo defines model for PackageInfo.
type PackageInfo struct {
	Id             *string `json:"id,omitempty"`
	PackageName    *string `json:"packageName,omitempty"`
	PackageVersion *string `json:"packageVersion,omitempty"`
}

// PodInfo defines model for PodInfo.
type PodInfo struct {
	Location *string `json:"location,omitempty"`
	PodName  *string `json:"podName,omitempty"`
}

// RootkitInfo defines model for RootkitInfo.
type RootkitInfo struct {
	Id *string `json:"id,omitempty"`

	// Path Path of the file that contains rootkit
	Path        *string      `json:"path,omitempty"`
	RootkitName *string      `json:"rootkitName,omitempty"`
	RootkitType *RootkitType `json:"rootkitType,omitempty"`
}

// RootkitScan defines model for RootkitScan.
type RootkitScan struct {
	Packages *[]RootkitInfo `json:"packages,omitempty"`
}

// RootkitType defines model for RootkitType.
type RootkitType string

// SbomScan defines model for SbomScan.
type SbomScan struct {
	Packages *[]Package `json:"packages,omitempty"`
}

// ScanResults defines model for ScanResults.
type ScanResults struct {
	Exploits          *ExploitScan          `json:"exploits,omitempty"`
	Id                *string               `json:"id,omitempty"`
	Malwares          *VulnerabilityScan    `json:"malwares,omitempty"`
	Misconfigurations *MisconfigurationScan `json:"misconfigurations,omitempty"`
	Sboms             *SbomScan             `json:"sboms,omitempty"`
	Secrets           *SecretScan           `json:"secrets,omitempty"`
	Vulnerabilities   *VulnerabilityScan    `json:"vulnerabilities,omitempty"`
}

// ScanResultsSummary defines model for ScanResultsSummary.
type ScanResultsSummary struct {
	ExploitsCount          *int `json:"exploitsCount,omitempty"`
	MalwaresCount          *int `json:"malwaresCount,omitempty"`
	MisconfigurationsCount *int `json:"misconfigurationsCount,omitempty"`
	PackagesCount          *int `json:"packagesCount,omitempty"`
	SecretsCount           *int `json:"secretsCount,omitempty"`
	VulnerabilitiesCount   *int `json:"vulnerabilitiesCount,omitempty"`
}

// ScanType defines model for ScanType.
type ScanType string

// SecretInfo defines model for SecretInfo.
type SecretInfo struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`

	// Path Path of the file that contains secrets
	Path *string `json:"path,omitempty"`
}

// SecretScan defines model for SecretScan.
type SecretScan struct {
	Packages *[]SecretInfo `json:"packages,omitempty"`
}

// SuccessResponse An object that is returned in cases of success that returns nothing.
type SuccessResponse struct {
	Message *string `json:"message,omitempty"`
}

// Target defines model for Target.
type Target struct {
	Id          *string            `json:"id,omitempty"`
	ScanResults *uint32            `json:"scanResults,omitempty"`
	TargetInfo  *Target_TargetInfo `json:"targetInfo,omitempty"`
	TargetType  *TargetType        `json:"targetType,omitempty"`
}

// Target_TargetInfo defines model for Target.TargetInfo.
type Target_TargetInfo struct {
	union json.RawMessage
}

// TargetType defines model for TargetType.
type TargetType string

// VMInfo defines model for VMInfo.
type VMInfo struct {
	InstanceName     *string        `json:"instanceName,omitempty"`
	InstanceProvider *CloudProvider `json:"instanceProvider,omitempty"`
	Location         *string        `json:"location,omitempty"`
}

// Vulnerability defines model for Vulnerability.
type Vulnerability struct {
	Id          *string            `json:"id,omitempty"`
	PackageInfo *VulnerabilityInfo `json:"packageInfo,omitempty"`
}

// VulnerabilityInfo defines model for VulnerabilityInfo.
type VulnerabilityInfo struct {
	Description       *string `json:"description,omitempty"`
	Id                *string `json:"id,omitempty"`
	VulnerabilityName *string `json:"vulnerabilityName,omitempty"`
}

// VulnerabilityScan defines model for VulnerabilityScan.
type VulnerabilityScan struct {
	Vulnerabilities *[]Vulnerability `json:"vulnerabilities,omitempty"`
}

// Page defines model for page.
type Page = int

// PageSize defines model for pageSize.
type PageSize = int

// ScanID defines model for scanID.
type ScanID = string

// ScanTypeParam defines model for scanTypeParam.
type ScanTypeParam = ScanType

// TargetID defines model for targetID.
type TargetID = string

// Success An object that is returned in cases of success that returns nothing.
type Success = SuccessResponse

// UnknownError An object that is returned in all cases of failures.
type UnknownError = ApiResponse

// GetTargetsParams defines parameters for GetTargets.
type GetTargetsParams struct {
	// Page Page number of the query
	Page Page `form:"page" json:"page"`

	// PageSize Maximum items to return
	PageSize PageSize `form:"pageSize" json:"pageSize"`
}

// GetTargetsTargetIDScanresultsParams defines parameters for GetTargetsTargetIDScanresults.
type GetTargetsTargetIDScanresultsParams struct {
	// Page Page number of the query
	Page Page `form:"page" json:"page"`

	// PageSize Maximum items to return
	PageSize PageSize `form:"pageSize" json:"pageSize"`
}

// GetTargetsTargetIDScanresultsScanIDParams defines parameters for GetTargetsTargetIDScanresultsScanID.
type GetTargetsTargetIDScanresultsScanIDParams struct {
	ScanType *ScanTypeParam `form:"scanType,omitempty" json:"scanType,omitempty"`
}

// PutTargetTargetIDJSONRequestBody defines body for PutTargetTargetID for application/json ContentType.
type PutTargetTargetIDJSONRequestBody = Target

// PostTargetsJSONRequestBody defines body for PostTargets for application/json ContentType.
type PostTargetsJSONRequestBody = Target

// PostTargetsTargetIDScanresultsJSONRequestBody defines body for PostTargetsTargetIDScanresults for application/json ContentType.
type PostTargetsTargetIDScanresultsJSONRequestBody = ScanResults

// PutTargetsTargetIDScanresultsScanIDJSONRequestBody defines body for PutTargetsTargetIDScanresultsScanID for application/json ContentType.
type PutTargetsTargetIDScanresultsScanIDJSONRequestBody = ScanResults

// AsVMInfo returns the union data inside the Target_TargetInfo as a VMInfo
func (t Target_TargetInfo) AsVMInfo() (VMInfo, error) {
	var body VMInfo
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVMInfo overwrites any union data inside the Target_TargetInfo as the provided VMInfo
func (t *Target_TargetInfo) FromVMInfo(v VMInfo) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVMInfo performs a merge with any union data inside the Target_TargetInfo, using the provided VMInfo
func (t *Target_TargetInfo) MergeVMInfo(v VMInfo) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsPodInfo returns the union data inside the Target_TargetInfo as a PodInfo
func (t Target_TargetInfo) AsPodInfo() (PodInfo, error) {
	var body PodInfo
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPodInfo overwrites any union data inside the Target_TargetInfo as the provided PodInfo
func (t *Target_TargetInfo) FromPodInfo(v PodInfo) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePodInfo performs a merge with any union data inside the Target_TargetInfo, using the provided PodInfo
func (t *Target_TargetInfo) MergePodInfo(v PodInfo) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsDirInfo returns the union data inside the Target_TargetInfo as a DirInfo
func (t Target_TargetInfo) AsDirInfo() (DirInfo, error) {
	var body DirInfo
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDirInfo overwrites any union data inside the Target_TargetInfo as the provided DirInfo
func (t *Target_TargetInfo) FromDirInfo(v DirInfo) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDirInfo performs a merge with any union data inside the Target_TargetInfo, using the provided DirInfo
func (t *Target_TargetInfo) MergeDirInfo(v DirInfo) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t Target_TargetInfo) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Target_TargetInfo) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}