package meta

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	ShortSysAttributes struct {
		SysGuid   string `xml:"SYS_GUID,attr"`
		SysGuidfk string `xml:"SYS_GUIDFK,attr"`
		SysState  int    `xml:"SYS_STATE,attr"`
		SysRev    int64  `xml:"SYS_REV,attr"`
	}

	SysAttributes struct {
		ShortSysAttributes
		SysFldOrder   int    `xml:"SYS_FLDORDER,attr"`
		SysParentGuid string `xml:"SYS_PARENTGUID,attr"`
		Name          string `xml:"NAME,attr"`
		Alias         string `xml:"ALIAS,attr"`
		Description   string `xml:"DESCRIPTION,attr"`
		Comment       string `xml:"COMMENT,attr"`
	}
)

const (
	MinFieldLength = 1
	MaxGuidLength  = 50
	MaxFieldLength = 27
	MaxAliasLength = 250
)

func (s ShortSysAttributes) ValidateShortSysAttr() error {
	return validation.ValidateStruct(&s,
		validation.Field(
			&s.SysGuid,
			validation.Required,
			validation.Length(MinFieldLength, MaxGuidLength)),

		validation.Field(&s.SysRev, validation.Required),

		validation.Field(&s.SysState, validation.In(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)),

		validation.Field(&s.SysGuidfk, validation.Length(MinFieldLength, MaxGuidLength)),
	)
}

func (s SysAttributes) ValidateSysAttributes() error {
	return validation.ValidateStruct(&s,
		validation.Field(
			&s.SysGuid,
			validation.Required,
			validation.Length(MinFieldLength, MaxGuidLength)),

		validation.Field(&s.SysRev, validation.Required),

		validation.Field(&s.SysState, validation.In(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)),

		validation.Field(&s.SysGuidfk, validation.Length(MinFieldLength, MaxGuidLength)),

		validation.Field(&s.Name,
			validation.Required,
			validation.Length(MinFieldLength, MaxFieldLength)),

		validation.Field(&s.Alias,
			validation.Required,
			validation.Length(MinFieldLength, MaxAliasLength)),
	)
	/**
	если ненулевой, то должен указывать на существующее поле
	SysParentGuid string `xml:"SYS_PARENTGUID,attr"`

	Name должно быть уникальным
	Name          string `xml:"NAME,attr"`

	*/
}
