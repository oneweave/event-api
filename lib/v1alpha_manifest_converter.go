package lib

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// ParseAndConvertManifest parses YAML/JSON manifest data into the internal PluginManifest representation,
// applying defaults and normalizing fields (e.g. converting omitted slices/arrays to initialized empty arrays).
func ParseAndConvertManifest(data []byte) (PluginManifest, error) {
	// 1. Unmarshal YAML into a generic map to handle yaml-to-json mapping of camelCase fields
	var generic map[string]any
	if err := yaml.Unmarshal(data, &generic); err != nil {
		return PluginManifest{}, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	// 2. Marshal generic map to JSON
	jsonData, err := json.Marshal(generic)
	if err != nil {
		return PluginManifest{}, fmt.Errorf("failed to marshal intermediate json: %w", err)
	}

	// 3. Unmarshal JSON into the pre-initialized external manifest
	ext := NewExtPluginManifest()
	if err := json.Unmarshal(jsonData, &ext); err != nil {
		return PluginManifest{}, fmt.Errorf("failed to decode json into external manifest: %w", err)
	}

	// 4. Validate user input (external representation)
	if err := ValidateStruct(&ext); err != nil {
		return PluginManifest{}, fmt.Errorf("failed to validate external manifest: %w", err)
	}

	// 5. Convert to internal manifest (merges with internal defaults)
	internal, err := ext.toInternal()
	if err != nil {
		return PluginManifest{}, fmt.Errorf("failed to convert to internal manifest: %w", err)
	}

	return internal, nil
}
