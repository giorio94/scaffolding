// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/websocket.proto

package cilium

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on WebSocketClient with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WebSocketClient) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WebSocketClient with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WebSocketClientMultiError, or nil if none found.
func (m *WebSocketClient) ValidateAll() error {
	return m.validate(true)
}

func (m *WebSocketClient) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessLogPath

	if utf8.RuneCountInString(m.GetHost()) < 2 {
		err := WebSocketClientValidationError{
			field:  "Host",
			reason: "value length must be at least 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Path

	// no validation rules for Key

	// no validation rules for Version

	// no validation rules for Origin

	if all {
		switch v := interface{}(m.GetHandshakeTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WebSocketClientValidationError{
					field:  "HandshakeTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WebSocketClientValidationError{
					field:  "HandshakeTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHandshakeTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WebSocketClientValidationError{
				field:  "HandshakeTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPingInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WebSocketClientValidationError{
					field:  "PingInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WebSocketClientValidationError{
					field:  "PingInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPingInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WebSocketClientValidationError{
				field:  "PingInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PingWhenIdle

	if len(errors) > 0 {
		return WebSocketClientMultiError(errors)
	}

	return nil
}

// WebSocketClientMultiError is an error wrapping multiple validation errors
// returned by WebSocketClient.ValidateAll() if the designated constraints
// aren't met.
type WebSocketClientMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WebSocketClientMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WebSocketClientMultiError) AllErrors() []error { return m }

// WebSocketClientValidationError is the validation error returned by
// WebSocketClient.Validate if the designated constraints aren't met.
type WebSocketClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WebSocketClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WebSocketClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WebSocketClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WebSocketClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WebSocketClientValidationError) ErrorName() string { return "WebSocketClientValidationError" }

// Error satisfies the builtin error interface
func (e WebSocketClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWebSocketClient.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WebSocketClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WebSocketClientValidationError{}

// Validate checks the field values on WebSocketServer with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WebSocketServer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WebSocketServer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WebSocketServerMultiError, or nil if none found.
func (m *WebSocketServer) ValidateAll() error {
	return m.validate(true)
}

func (m *WebSocketServer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessLogPath

	// no validation rules for Host

	// no validation rules for Path

	// no validation rules for Key

	// no validation rules for Version

	// no validation rules for Origin

	if all {
		switch v := interface{}(m.GetHandshakeTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WebSocketServerValidationError{
					field:  "HandshakeTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WebSocketServerValidationError{
					field:  "HandshakeTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHandshakeTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WebSocketServerValidationError{
				field:  "HandshakeTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPingInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WebSocketServerValidationError{
					field:  "PingInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WebSocketServerValidationError{
					field:  "PingInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPingInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WebSocketServerValidationError{
				field:  "PingInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PingWhenIdle

	if len(errors) > 0 {
		return WebSocketServerMultiError(errors)
	}

	return nil
}

// WebSocketServerMultiError is an error wrapping multiple validation errors
// returned by WebSocketServer.ValidateAll() if the designated constraints
// aren't met.
type WebSocketServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WebSocketServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WebSocketServerMultiError) AllErrors() []error { return m }

// WebSocketServerValidationError is the validation error returned by
// WebSocketServer.Validate if the designated constraints aren't met.
type WebSocketServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WebSocketServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WebSocketServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WebSocketServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WebSocketServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WebSocketServerValidationError) ErrorName() string { return "WebSocketServerValidationError" }

// Error satisfies the builtin error interface
func (e WebSocketServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWebSocketServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WebSocketServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WebSocketServerValidationError{}
