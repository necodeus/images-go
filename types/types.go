package types

type Image struct {
	ID         int    `json:"id"`
	TypeName   string `json:"type_name"`
	ResourceID string `json:"resource_id"`
	MimeType   string `json:"mime_type"`
	Size       int64  `json:"size"`
}

type ImageType struct {
	Name                 string   `json:"name"`
	AvailableResolutions []string `json:"available_resolutions"`
}
