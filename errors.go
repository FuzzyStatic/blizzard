package blizzard

import "errors"

// Sentinel errors
var (
	ErrUnknownLocale          = errors.New("unknown locale")
	ErrInvalidLocaleForRegion = errors.New("locale is not valid for region")
)
