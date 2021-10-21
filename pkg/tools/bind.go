package tools

import (
	"gopkg.in/yaml.v2"
	"insurance/pkg/json"
)

const (
	// YAML 绑定方式
	YAML = "yaml"
)

// Bind 绑定数据
func Bind(dst, src interface{}, method ...string) error {
	if len(method) == 0 {
		return bindJSON(dst, src)
	}
	switch method[0] {
	case YAML:
		return bindYaml(dst, src)
	default:
		return bindJSON(dst, src)
	}
}

func bindJSON(dst, src interface{}) error {
	str, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(str, dst)
}

func bindYaml(dst, src interface{}) error {
	str, err := yaml.Marshal(src)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(str, dst)
}
