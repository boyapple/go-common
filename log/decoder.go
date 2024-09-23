package log

import "gopkg.in/yaml.v3"

type Decoder interface {
	Decode(cfg interface{}) error
}

type YamlDecoderNode struct {
	Node *yaml.Node
}

func (y *YamlDecoderNode) Decode(cfg interface{}) error {
	return y.Node.Decode(cfg)
}

func Setup(name string, dec Decoder) error {
	var cfg []OutputConfig
	if err := dec.Decode(&cfg); err != nil {
		return err
	}
	Register(name, NewZapLog(cfg))
	return nil
}
