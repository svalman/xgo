package meta

import (
	"xml-diff/meta/types"
)

type (
	SysAu struct { // SYS_AU tag
		ShortSysAttributes

		SysFldOrder int              `xml:"SYS_FLDORDER,attr"`
		Name        string           `xml:"NAME,attr"`
		VersionInfo string           `xml:"VERSIONINFO,attr"`
		VersionDate types.GermanTime `xml:"VERSIONDATE,attr"`
		VersionId   int              `xml:"VERSIONID,attr"`
		UplinkRole  string           `xml:"UPLINKROLE,attr"`
		Conditions  string           `xml:"COND,attr"`
	}
)
