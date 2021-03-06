package client

// Code generated by gen.go. DO NOT EDIT.

import (
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// TargetType are the types of targets available in Chrome.
type TargetType string

// TargetType values.
const (
	App            TargetType = "app"
	BackgroundPage TargetType = "background_page"
	Browser        TargetType = "browser"
	External       TargetType = "external"
	Iframe         TargetType = "iframe"
	Other          TargetType = "other"
	Page           TargetType = "page"
	ServiceWorker  TargetType = "service_worker"
	SharedWorker   TargetType = "shared_worker"
	Webview        TargetType = "webview"
	Worker         TargetType = "worker"
)

// String satisfies stringer.
func (tt TargetType) String() string {
	return string(tt)
}

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (tt TargetType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(tt))
}

// MarshalJSON satisfies json.Marshaler.
func (tt TargetType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(tt)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (tt *TargetType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	z := TargetType(in.String())
	switch z {
	case App:
		*tt = App
	case BackgroundPage:
		*tt = BackgroundPage
	case Browser:
		*tt = Browser
	case External:
		*tt = External
	case Iframe:
		*tt = Iframe
	case Other:
		*tt = Other
	case Page:
		*tt = Page
	case ServiceWorker:
		*tt = ServiceWorker
	case SharedWorker:
		*tt = SharedWorker
	case Webview:
		*tt = Webview
	case Worker:
		*tt = Worker

	default:
		*tt = z
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (tt *TargetType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, tt)
}
