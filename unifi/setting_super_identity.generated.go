// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"fmt"
)

// just to fix compile issues with the import
var (
	_ fmt.Formatter
	_ context.Context
)

type SettingSuperIdentity struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Hostname string `json:"hostname,omitempty"`
	Name     string `json:"name,omitempty"`
}
