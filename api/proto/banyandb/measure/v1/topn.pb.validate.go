// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: banyandb/measure/v1/topn.proto

package v1

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

// Validate checks the field values on TopNList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TopNList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TopNList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TopNListMultiError, or nil
// if none found.
func (m *TopNList) ValidateAll() error {
	return m.validate(true)
}

func (m *TopNList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TopNListValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TopNListValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TopNListValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TopNListValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TopNListValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TopNListValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TopNListMultiError(errors)
	}

	return nil
}

// TopNListMultiError is an error wrapping multiple validation errors returned
// by TopNList.ValidateAll() if the designated constraints aren't met.
type TopNListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TopNListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TopNListMultiError) AllErrors() []error { return m }

// TopNListValidationError is the validation error returned by
// TopNList.Validate if the designated constraints aren't met.
type TopNListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TopNListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TopNListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TopNListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TopNListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TopNListValidationError) ErrorName() string { return "TopNListValidationError" }

// Error satisfies the builtin error interface
func (e TopNListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTopNList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TopNListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TopNListValidationError{}

// Validate checks the field values on TopNResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TopNResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TopNResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TopNResponseMultiError, or
// nil if none found.
func (m *TopNResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TopNResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLists() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TopNResponseValidationError{
						field:  fmt.Sprintf("Lists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TopNResponseValidationError{
						field:  fmt.Sprintf("Lists[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TopNResponseValidationError{
					field:  fmt.Sprintf("Lists[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TopNResponseMultiError(errors)
	}

	return nil
}

// TopNResponseMultiError is an error wrapping multiple validation errors
// returned by TopNResponse.ValidateAll() if the designated constraints aren't met.
type TopNResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TopNResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TopNResponseMultiError) AllErrors() []error { return m }

// TopNResponseValidationError is the validation error returned by
// TopNResponse.Validate if the designated constraints aren't met.
type TopNResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TopNResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TopNResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TopNResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TopNResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TopNResponseValidationError) ErrorName() string { return "TopNResponseValidationError" }

// Error satisfies the builtin error interface
func (e TopNResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTopNResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TopNResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TopNResponseValidationError{}

// Validate checks the field values on TopNRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TopNRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TopNRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TopNRequestMultiError, or
// nil if none found.
func (m *TopNRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TopNRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMetadata() == nil {
		err := TopNRequestValidationError{
			field:  "Metadata",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TopNRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TopNRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TopNRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTimeRange() == nil {
		err := TopNRequestValidationError{
			field:  "TimeRange",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTimeRange()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TopNRequestValidationError{
					field:  "TimeRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TopNRequestValidationError{
					field:  "TimeRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeRange()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TopNRequestValidationError{
				field:  "TimeRange",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTopN() <= 0 {
		err := TopNRequestValidationError{
			field:  "TopN",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Agg

	for idx, item := range m.GetConditions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TopNRequestValidationError{
						field:  fmt.Sprintf("Conditions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TopNRequestValidationError{
						field:  fmt.Sprintf("Conditions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TopNRequestValidationError{
					field:  fmt.Sprintf("Conditions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for FieldValueSort

	if len(errors) > 0 {
		return TopNRequestMultiError(errors)
	}

	return nil
}

// TopNRequestMultiError is an error wrapping multiple validation errors
// returned by TopNRequest.ValidateAll() if the designated constraints aren't met.
type TopNRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TopNRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TopNRequestMultiError) AllErrors() []error { return m }

// TopNRequestValidationError is the validation error returned by
// TopNRequest.Validate if the designated constraints aren't met.
type TopNRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TopNRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TopNRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TopNRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TopNRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TopNRequestValidationError) ErrorName() string { return "TopNRequestValidationError" }

// Error satisfies the builtin error interface
func (e TopNRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTopNRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TopNRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TopNRequestValidationError{}

// Validate checks the field values on TopNList_Item with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TopNList_Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TopNList_Item with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TopNList_ItemMultiError, or
// nil if none found.
func (m *TopNList_Item) ValidateAll() error {
	return m.validate(true)
}

func (m *TopNList_Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TopNList_ItemValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TopNList_ItemValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TopNList_ItemValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TopNList_ItemMultiError(errors)
	}

	return nil
}

// TopNList_ItemMultiError is an error wrapping multiple validation errors
// returned by TopNList_Item.ValidateAll() if the designated constraints
// aren't met.
type TopNList_ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TopNList_ItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TopNList_ItemMultiError) AllErrors() []error { return m }

// TopNList_ItemValidationError is the validation error returned by
// TopNList_Item.Validate if the designated constraints aren't met.
type TopNList_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TopNList_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TopNList_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TopNList_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TopNList_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TopNList_ItemValidationError) ErrorName() string { return "TopNList_ItemValidationError" }

// Error satisfies the builtin error interface
func (e TopNList_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTopNList_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TopNList_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TopNList_ItemValidationError{}
