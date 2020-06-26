package main

import "github.com/pelletier/go-toml"

type Configuration struct {
	Aws_access_key_id string
	Aws_secret_access_key string
}

func LoadConfigFromFile(conf *Configuration, aws_profile string,configfile string) error {
	config, err := toml.LoadFile(configfile)
	if err != nil {
		return err
	}

	conf.Aws_access_key_id = config.Get(aws_profile+".aws_access_key_id").(string)
	conf.Aws_secret_access_key = config.Get(aws_profile+".aws_secret_access_key").(string)

	return nil
}