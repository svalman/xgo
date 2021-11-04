package types

type (
	DbConnectionParams struct {
		Adapter  string `yaml:"adapter"`
		Dbname   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)
