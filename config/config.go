package config

type (
	Config struct { // 定義 Config 結構
		Username string `yaml:"username"` // 使用 yaml 格式的註解，定義屬性名稱
		Password string `yaml:"password"`
		Network  string `yaml:"network"`
		Server   string `yaml:"server"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
	}
)
