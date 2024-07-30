// Package v1beta1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1beta1

import (
	externalRef1 "github.com/meshery/schemas/models/core"
	"github.com/meshery/schemas/models/v1alpha2"
	"github.com/meshery/schemas/models/v1alpha3"
)

// Defines values for MesheryPatternCatalogDataClass.
const (
	Community MesheryPatternCatalogDataClass = "community"
	Official  MesheryPatternCatalogDataClass = "official"
	Verified  MesheryPatternCatalogDataClass = "verified"
)

// Defines values for MesheryPatternCatalogDataCompatibility.
const (
	Kubernetes MesheryPatternCatalogDataCompatibility = "kubernetes"
)

// Defines values for MesheryPatternCatalogDataType.
const (
	Deployment        MesheryPatternCatalogDataType = "Deployment"
	Observability     MesheryPatternCatalogDataType = "Observability"
	Resiliency        MesheryPatternCatalogDataType = "Resiliency"
	Scaling           MesheryPatternCatalogDataType = "Scaling"
	Security          MesheryPatternCatalogDataType = "Security"
	TrafficManagement MesheryPatternCatalogDataType = "Traffic-management"
	Troubleshooting   MesheryPatternCatalogDataType = "Troubleshooting"
	Workloads         MesheryPatternCatalogDataType = "Workloads"
)

// DeletePatternModel defines model for deletePatternModel.
type DeletePatternModel struct {
	ID   externalRef1.Id   `db:"id" json:"id"`
	Name externalRef1.Text `json:"name,omitempty"`
}

type PatternFile struct {
	Id externalRef1.Id  `json:"id,omitempty" yaml:"id,omitempty"`

	// Components List of component declarations
	Components []*ComponentDefinition `json:"components" yaml:"components"`

	// Name Name of the design; a descriptive, but concise title for the design document.
	Name string `json:"name" yaml:"name"`

	// Preferences Design-level preferences
	Preferences *struct {
		// Layers List of available layers
		Layers []string `json:"layers"`
	} `json:"preferences,omitempty"`

	// Relationships List of relationships between components
	Relationships []*v1alpha3.RelationshipDefinition `json:"relationships" yaml:"relationships"`

	// SchemaVersion Specifies the version of the schema to which the design conforms.
	SchemaVersion string `json:"schemaVersion" yaml:"schemaVersion"`

	// Version Revision of the design as expressed by an auto-incremented, SemVer-compliant version number. May be manually set by a user or third-party system, but will always be required to be of version number higher than the previously defined version number.
	Version string `json:"version" yaml:"version"`
}

// MesheryPattern defines model for mesheryPattern.
type MesheryPattern struct {
	CatalogData *v1alpha2.CatalogData `json:"catalog_data,omitempty" yaml:"catalog_data"`
	CreatedAt externalRef1.Time      `json:"created_at,omitempty"`
	UserID    externalRef1.Id        `db:"user_id" json:"user_id"`
	Location  externalRef1.MapObject `json:"location,omitempty"`
	Name      externalRef1.Text      `json:"name,omitempty"`

	// PatternFile Designs are your primary tool for collaborative authorship of your infrastructure, workflow, and processes.
	PatternFile *PatternFile      `json:"pattern_file,omitempty" yaml:"pattern_file"`
	UpdatedAt   externalRef1.Time `json:"updated_at,omitempty"`
	ID          externalRef1.Id   `db:"id" json:"id"`
	Visibility  externalRef1.Text `json:"visibility,omitempty"`
}

// MesheryPatternCatalogDataClass Published content is classifed by its support level. Content classes help you understand the origin and expected support level for each piece of content. It is important to note that the level of support may vary within each class, and you should exercise discretion when using community-contributed content. Content produced and fully supported by Meshery maintainers. This represents the highest level of support and is considered the most reliable. Content produced by partners and verified by Meshery maintainers. While not directly maintained by Meshery, it has undergone a verification process to ensure quality and compatibility. Content produced and supported by the respective project or organization responsible for the specific technology. This class offers a level of support from the project maintainers themselves. Content produced and shared by Meshery users. This includes a wide range of content, such as performance profiles, test results, filters, patterns, and applications. Community content may have varying levels of support and reliability.
type MesheryPatternCatalogDataClass string

// MesheryPatternCatalogDataCompatibility defines model for MesheryPattern.CatalogData.Compatibility.
type MesheryPatternCatalogDataCompatibility string

// MesheryPatternCatalogDataType Categorization of the type of design or operational flow depicted in this design.
type MesheryPatternCatalogDataType string

// MesheryPatternPage defines model for mesheryPatternPage.
type MesheryPatternPage struct {
	Page       int               `json:"page,omitempty"`
	PageSize   int               `json:"page_size,omitempty"`
	Patterns   *[]MesheryPattern `json:"patterns,omitempty"`
	ResultType string            `json:"resultType,omitempty"`
	TotalCount int               `json:"total_count,omitempty"`
}

// MesheryPatternResource defines model for mesheryPatternResource.
type MesheryPatternResource struct {
	CreatedAt externalRef1.Time `json:"created_at,omitempty"`
	Deleted   *bool             `json:"deleted,omitempty"`
	ID        externalRef1.Id   `db:"id" json:"id"`
	Name      externalRef1.Text `json:"name,omitempty"`
	Namepace  externalRef1.Text `json:"namepace,omitempty"`
	OamType   externalRef1.Text `json:"oam_type,omitempty"`
	Type      externalRef1.Text `json:"type,omitempty"`
	UpdatedAt externalRef1.Time `json:"updated_at,omitempty"`
	UserID    externalRef1.Id   `db:"user_id" json:"user_id"`
}

// MesheryPatternResourcePage defines model for mesheryPatternResourcePage.
type MesheryPatternResourcePage struct {
	Page       int                       `json:"page,omitempty"`
	PageSize   int                       `json:"page_size,omitempty"`
	Resources  *[]MesheryPatternResource `json:"resources,omitempty"`
	ResultType string                    `json:"resultType,omitempty"`
	TotalCount int                       `json:"total_count,omitempty"`
}

// DesignShare defines model for designShare.
type DesignShare struct {
	ContentType string              `json:"content_type"`
	Emails      externalRef1.Emails `json:"emails"`
	ID          externalRef1.Id     `db:"id" json:"id"`
	Share       bool                `json:"share"`
}

// MesheryPatternDeleteRequestBody defines model for mesheryPatternDeleteRequestBody.
type MesheryPatternDeleteRequestBody struct {
	Patterns *[]DeletePatternModel `json:"patterns,omitempty"`
}

// MesheryPatternRequestBody defines model for mesheryPatternRequestBody.
type MesheryPatternRequestBody struct {
	Path        externalRef1.Text     `json:"path,omitempty"`
	PatternData *MesheryPattern       `json:"pattern_data,omitempty"`
	Save        *bool                 `json:"save,omitempty"`
	Url         externalRef1.Endpoint `json:"url,omitempty"`
}
