package meta

type (
	SysProjects struct { // tag SYS_PROJECTS
		ShortSysAttributes
		Name          string `xml:"NAME,attr"`
		DbName        string `xml:"DBNAME,attr"`
		NodeUrl       string `xml:"NODEURL,attr"`
		MinWinPVer    int    `xml:"MINWINPVER,attr"`
		MinWebPVer    int    `xml:"MINWEBPVER,attr"`
		MinNodeVer    int    `xml:"MINNODEVER,attr"`
		BaseInit      string `xml:"BASEINIT,attr"`
		Filename      string `xml:"FILENAME,attr"`
		BlobStoreFile string `xml:"BLOBSTOREFILE,attr"`
		SchemaVersion string `xml:"SCHEMAVERSION,attr"`
	}
)
