package meta

type (
	SysAuTables struct { // tag SYS_AU_TABLES
		ShortSysAttributes
		CommonFieldAttributes

		SysFldOrder   int    `xml:"SYS_FLDORDER,attr"`
		SysParentGuid string `xml:"SYS_PARENTGUID,attr"`
		Alias         string `xml:"ALIAS,attr"`

		IsGroup   int    `xml:"ISGROUP,attr"`
		TableGuid string `xml:"TABLE_GUID,attr"`
		ExtraSQL  string `xml:"EXRTRASQL,attr"`
		Filter    string `xml:"FILTER,attr"`

		ExportFilter     string `xml:"EXPORTFILTER,attr"`
		DetailField      string `xml:"DETAILFIELD,attr"`
		BaseDetailSet    string `xml:"BASEDETAILSET,attr"`
		DetailScript     string `xml:"DETAILSCRIPT,attr"`
		DetailAlias      string `xml:"DETAILALIAS,attr"`
		MinCount         int    `xml:"MINCOUNT,attr"`
		MaxCount         int    `xml:"MAXCOUNT,attr"`
		GenFltQuery      string `xml:"GENFLTQUERY,attr"`
		UserIdFld        string `xml:"USERIDFLD,attr"`
		NameFld          string `xml:"NAMEFLD,attr"`
		RoleFk           string `xml:"ROLEFK,attr"`
		UserGenFilter    string `xml:"USERGENFILTER,attr"`
		CanAppend        string `xml:"CANAPPEND,attr"`
		CanUpdate        string `xml:"CANUPDATE,attr"`
		CanDelete        string `xml:"CANDELETE,attr"`
		SubTabs          string `xml:"SUBTABS,attr"`
		MultiSelection   string `xml:"MULTISELECTION,attr"`
		MultiLookupField string `xml:"MULTILOOKUPFIELD,attr"`
		MultiIsUnique    int    `xml:"MULTIISUNIQUE,attr"`
		CanEditInChooser string `xml:"CANEDITINCHOOSER,attr"`
		MultiFilter      string `xml:"MULTIFILTER,attr"`
		FieldsList       string `xml:"FIELDSLIST,attr"`
		OrderBy          string `xml:"ORDERBY,attr"`
		GroupBy          string `xml:"GROUPBY,attr"`
		HideCard         int    `xml:"HIDECARD,attr"`
		HideFromTree     int    `xml:"HIDEFROMTREE,attr"`
		MultiSelFields   string `xml:"MULTISELFIELDS,attr"`
		WpTemplate       string `xml:"WP_TEMPLATE,attr"`
		Expanded         string `xml:"EXPANDED,attr"`
		FillMatching     string `xml:"FILLMATCHING,attr"`
		FilteredOnly     string `xml:"FILTEREDONLY,attr"`
	}
)
