package config

type Config struct {
	Bot struct {
		Website  string  `yaml:"website"`
		Token    string  `yaml:"token"`
		AdminIDs []int64 `yaml:"admin_id"`
		GroupID  int64   `yaml:"group_id"`
	} `yaml:"bot"`
	Apps struct {
		Database struct {
			IP   string `yaml:"ip"`
			Port int    `yaml:"port"`
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
			Name string `yaml:"name"`
		} `yaml:"database"`
	} `yaml:"apps"`
}
