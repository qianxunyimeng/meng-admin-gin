package config

type Config struct {
	JWT      JWT          `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System   System       `mapstructure:"system" json:"system" yaml:"system"`
	AutoCode Autocode     `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	Zap      Zap          `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql    Mysql        `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Captcha  Captcha      `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Redis    RedisOptions `mapstructure:"redis" json:"redis" yaml:"redis"`
}
