package config

type MGServer struct {
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
}
