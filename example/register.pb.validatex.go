// Code generated by protoc-gen-validatex. DO NOT EDIT.

// versions:
//  protoc-gen-validatex v0.4.1

package main

import (
	context "context"
	i18n "github.com/nicksnyder/go-i18n/v2/i18n"
	validatex "github.com/protoc-gen/protoc-gen-validatex/pkg/validatex"
)

func (x *SignUpRequest) Validate(ctx context.Context) error {
	if x == nil {
		return nil
	}
	if validatex.ValidEmail(x.Email) != nil {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "EmailInvalid",
				TemplateData: map[string]string{"FieldName": "email"},
			}, "must be a valid email")).
			WithMetadata(map[string]string{"field": "email"})
	}
	if len(x.Password) < 5 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMinLen",
				TemplateData: map[string]string{"MinLen": "5"},
			}, "must be at least 5 characters long")).
			WithMetadata(map[string]string{"field": "password"})
	}
	if len(x.Password) > 50 {
		return validatex.NewError(
			validatex.MustLocalize(ctx, &i18n.LocalizeConfig{MessageID: "StringMaxLen",
				TemplateData: map[string]string{"MaxLen": "50"},
			}, "must be at most 50 characters long")).
			WithMetadata(map[string]string{"field": "password"})
	}
	return nil
}

func init() {
	validatex.Init18n("./example/i18n")
}
