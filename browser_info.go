package adyen

// BrowserInfo represents the "browserInfo" object as defined in the Adyen documentation.
//
// Read more: https://docs.adyen.com/online-payments/3d-secure/api-reference#browserinfo
type BrowserInfo struct {
	// AcceptHeader is the value of the browser's Accept header.
	AcceptHeader string `json:"acceptHeader"`

	// ColorDepth is the browser's color depth in bits per pixel.
	// This is the "screen.colorDepth" value.
	ColorDepth int `json:"colorDepth"`

	// JavaEnabled is if the browser is capable of executing Java.
	JavaEnabled bool `json:"javaEnabled"`

	// JavaScriptEnabled is an optional field indicating if the browser
	// is capable of executing JavaScript. If not specified, the Adyen
	// API assumes this value is true.
	//
	// To use a boolean value, use Bool.
	JavaScriptEnabled *bool `json:"javaScriptEnabled,omitempty"`

	// Language is the name of the language used by the browser.
	// This is the value of "navigator.language".
	Language string `json:"language"`

	// ScreenHeight is the browser's screen height.
	ScreenHeight int `json:"screenHeight"`

	// ScreenWidth is the browser's screen width.
	ScreenWidth int `json:"screenWidth"`

	// TimeZoneOffset is the time difference between UTC time and the browser's
	// local time in minutes.
	TimeZoneOffset int `json:"timeZoneOffset"`

	// UserAgent is the browser's user agent.
	UserAgent string `json:"userAgent"`
}

// Bool returns a pointer to value.
func Bool(value bool) *bool {
	return &value
}
