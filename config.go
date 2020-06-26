package main

import (
	    "gopkg.in/ini.v1"
	)

type Configuration struct {
	Aws_access_key_id string
	Aws_secret_access_key string
}

func LoadConfigFromFile(conf *Configuration, aws_profile string,configfile string) error {
	config, err := ini.Load(configfile)
	if err != nil {
		return err
	}

	conf.Aws_access_key_id = config.Section(aws_profile).Key("aws_access_key_id").String()
	conf.Aws_secret_access_key = config.Section(aws_profile).Key("aws_secret_access_key").String()

	return nil
}