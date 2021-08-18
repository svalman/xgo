package meta

type (
	SysTables struct { // tag SYS_TABLES
		SysAttributes

		SqlAlias         string `xml:"SQLALIAS,attr"`
		IsTree           int    `xml:"ISTREE,attr"`
		IsDir            int    `xml:"ISDIR,attr"`
		Order            int    `xml:"ORDER,attr"`
		FieldId          string `xml:"FIELDID,attr"`
		FieldFk          string `xml:"FIELDFK,attr"`
		InfoField        string `xml:"INFOFIELD,attr"`
		FieldsList       string `xml:"FIELDSLIST,attr"`
		MultiSelection   int    `xml:"MULTISELECTION,attr"`
		MultiLookupField string `xml:"MULTILOOKUPFIELD,attr"`
		Uniq             string `xml:"UNIQ,attr"`
		Indexes          string `xml:"INDEXES,attr"`
		Uniques          string `xml:"UNIQUES,attr"`
		MultiIsUnique    int    `xml:"MULTIISUNIQUE,attr"`
		CanEditInChooser int    `xml:"CANEDITINCHOOSER,attr"`
		MultiFilter      string `xml:"MULTIFILTER,attr"`
		KeyGen           string `xml:"KEYGEN,attr"`
		IsFact           int    `xml:"ISFACT,attr"`
		YAxes            string `xml:"YAXES,attr"`
		FieldMeasures    string `xml:"FIELDSMEASURES,attr"`
		XAxes            string `xml:"XAXES,attr"`
		Facts            string `xml:"FACTS,attr"`
		FilterFields     string `xml:"FILTERFIELDS,attr"`
		IsView           int    `xml:"ISVIEW,attr"`
		ViewSql          string `xml:"VIEWSQL,attr"`
		EditTablet       int    `xml:"EDITABLET,attr"`
		ChkNullOnExp     int    `xml:"CHKNULLONEXP,attr"`
		DetailField      string `xml:"DETAILFIELD,attr"`
		DecoGuid         int    `xml:"DECOGUID,attr"`
		OrderBy          string `xml:"ORDERBY,attr"`
		RecChanges       int    `xml:"RECCHANGES,attr"`
		PresetFk         string `xml:"PRESETFK,attr"`
		StyleValue       string `xml:"STYLEVALUE,attr"`
		FullMatching     int    `xml:"FILLMATCHING,attr"`
	}
)
