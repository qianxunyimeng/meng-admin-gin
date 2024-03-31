// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/3/29 22:11:00
// @Desc
package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/redis/go-redis/v9"
	"io/ioutil"
)

type RedisOptions struct {
	Network    string `yaml:"network" json:"network"` // 连接网络类型，如: tcp、udp、unix等方式 如果为空 默认tcp
	Addr       string `yaml:"addr" json:"addr"`       // redis服务器地址，ip:port格式，比如：192.168.1.100:6379
	Username   string `yaml:"username" json:"username"`
	Password   string `yaml:"password" json:"password"`
	DB         int    `yaml:"db" json:"db"`
	PoolSize   int    `yaml:"pool_size" json:"pool_size"` // 连接池最大连接数量，注意：这里不包括 pub/sub，pub/sub 将使用独立的网络连接 默认为 10 * runtime.GOMAXPROCS
	Tls        *Tls   `yaml:"tls" json:"tls"`
	MaxRetries int    `yaml:"max_retries" json:"max_retries"` // 命令最大重试次数， 默认为3
}

type Tls struct {
	Cert string `yaml:"cert" json:"cert"`
	Key  string `yaml:"key" json:"key"`
	Ca   string `yaml:"ca" json:"ca"`
}

func (e RedisOptions) GetRedisOptions() (*redis.Options, error) {
	r := &redis.Options{
		Network:    e.Network,
		Addr:       e.Addr,
		Username:   e.Username,
		Password:   e.Password,
		DB:         e.DB,
		MaxRetries: e.MaxRetries,
		PoolSize:   e.PoolSize,
	}
	var err error
	r.TLSConfig, err = getTLS(e.Tls)
	return r, err
}

func getTLS(c *Tls) (*tls.Config, error) {
	if c != nil && c.Cert != "" {
		// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
		cert, err := tls.LoadX509KeyPair(c.Cert, c.Key)
		if err != nil {
			fmt.Printf("tls.LoadX509KeyPair err: %v\n", err)
			return nil, err
		}
		// 创建一个新的、空的 CertPool，并尝试解析 PEM 编码的证书，解析成功会将其加到 CertPool 中
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(c.Ca)
		if err != nil {
			fmt.Printf("ioutil.ReadFile err: %v\n", err)
			return nil, err
		}

		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			fmt.Println("certPool.AppendCertsFromPEM err")
			return nil, err
		}
		return &tls.Config{
			// 设置证书链，允许包含一个或多个
			Certificates: []tls.Certificate{cert},
			// 要求必须校验客户端的证书
			ClientAuth: tls.RequireAndVerifyClientCert,
			// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
			ClientCAs: certPool,
		}, nil
	}
	return nil, nil
}
