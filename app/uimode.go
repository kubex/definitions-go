package app

type UIMode string

const (
	UIModeFull        UIMode = "full"        // Standard application
	UIModeIntegration UIMode = "integration" // Integrated into existing pages
	UIModeNone        UIMode = "none"        // No UI - flow only
)
