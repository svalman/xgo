package meta

type (
	SysRoles struct { // tag SYS_ROLES
		ShortSysAttributes

		Name          string `xml:"NAME,attr"`
		Description   string `xml:"DESCRIPTION,attr"`
		IsAdmin       int    `xml:"ISADMIN,attr"`
		LockReports   int    `xml:"LOCKREPORTS,attr"`
		LockReportGen int    `xml:"LOCKREPORTGEN,attr"`
		LockProcs     int    `xml:"LOCKPROCS,attr"`
		LockImport    int    `xml:"LOCKIMPORT,attr"`
		LockExport    int    `xml:"LOCKEXPORT,attr"`
		LockConnect   int    `xml:"LOCKCONNECT,attr"`
		DefExModel    int    `xml:"DEFEXMODEL,attr"`
	}
)
