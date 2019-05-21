package models

type Functions struct {
	Functions []Function `yaml:"functions"`
}

type Function struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Url         string   `yaml:"url"`
	Parser      string   `yaml:"response_parser"`
	Headers     []Header `yaml:"headers"`
}

type Header struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
	Env   bool   `yaml:"env"`
}
