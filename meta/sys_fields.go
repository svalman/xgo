package meta

type (
	CommonFieldAttributes struct {
		Visible       int    `xml:"VISIBLE,attr"`
		Readonly      int    `xml:"READONLY,attr"`
		LookupDisplay string `xml:"LOOKUPDISPLAY,attr"`
		LookupFilter  string `xml:"LOOKUPFILTER,attr"`
		Prefix        string `xml:"PREFIX,attr"`
	}

	SysFields struct { // SYS_FIELDS tag
		SysAttributes
		CommonFieldAttributes

		FldType string `xml:"TYPE,attr"`

		Length    int `xml:"LENGTH,attr"`
		Precision int `xml:"PRECISION,attr"`

		DefValue string `xml:"DEFVALUE,attr"`
		Filter   string `xml:"FILTER,attr"`

		ListSource       string `xml:"LISTSOURCE,attr"`
		ListKeys         string `xml:"LISTKEYS,attr"`
		ListStyle        string `xml:"LISTSTYLE,attr"`
		IsUnique         int    `xml:"ISUNIQUE,attr"`
		IsNotNull        int    `xml:"ISNOTNULL,attr"`
		LookupSource     string `xml:"LOOKUPSOURCE,attr"`
		LookupQuery      string `xml:"LOOKUPQUERY,attr"`
		LookupField      string `xml:"LOOKUPFIELD,attr"`
		LookupFieldsList string `xml:"LOOKUPFIELDSLIST,attr"`
		LookupDelim      string `xml:"LOOKUPDELIM,attr"`
		LookupOrder      int    `xml:"LOOKUPORDER,attr"`
		LookupAsForm     int    `xml:"LOOKUPASFORM,attr"`

		MaxFileSize          int64  `xml:"MAXFILESIZE,attr"`
		MaxFileCnt           int    `xml:"MAXFILECNT,attr"`
		CalcScript           string `xml:"CALCSCRIPT,attr"`
		CharCase             string `xml:"CHARCASE,attr"`
		Password             int    `xml:"PASSWORD,attr"`
		Autofill             string `xml:"AUTOFILL,attr"`
		EmptyValue           string `xml:"EMPTYVALUE,attr"`
		NotExportable        int    `xml:"NOTEXPORTABLE,attr"`
		Template             string `xml:"TEMPLATE,attr"`
		Regexp               string `xml:"REGEXP,attr"`
		InputMask            string `xml:"INPUTMASK,attr"`
		CaptionScript        string `xml:"CAPTIONSCRIPT,attr"`
		DisplayFormat        string `xml:"DISPLAYFORMAT,attr"`
		UseThousandSeparator int    `xml:"USETHOUSANDSEPARATOR,attr"`
		RadioOrigin          string `xml:"RADIOORIGIN,attr"`
		RadioColumns         string `xml:"RADIOCOLUMNS,attr"`
		CanEditInChooser     int    `xml:"CANEDITINCHOOSER,attr"`
		WpTemplate           string `xml:"WP_TEMPLATE,attr"`
		Selectable           int    `xml:"SELECTABLE,attr"`
	}
)
