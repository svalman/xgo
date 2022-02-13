package sys_projects

import "xml-diff/xgo/meta"

type (
	SysProjects struct { // tag SYS_PROJECTS
		meta.ShortSysAttributes
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

func SelectAllQuery() string {
	return `select * from "SYS_PROJECTS" where coalesce("SYS_STATE",0)%2=0`
}
