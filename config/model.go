package config

type Config struct {
	Bot  `yaml:"bot"`
	APPs `yaml:"apps"`
}

type Bot struct {
	Website  string `yaml:"website"`
	Token    string `yaml:"token"`
	AdminIDs []int  `yaml:"admin_id"`
	GroupID  int    `yaml:"group_id"`
}

type APPs struct {
	Database `yaml:"database"`
}

type Database struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}
