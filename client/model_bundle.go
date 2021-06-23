/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing, and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Bundle struct for Bundle
type Bundle struct {
	BundleHeader  BundleHeader  `json:"bundleHeader,omitempty"`
	Checks        []Checks      `json:"checks,omitempty"`
	Returns       []Returns     `json:"returns,omitempty"`
	BundleControl BundleControl `json:"bundleControl,omitempty"`
}
