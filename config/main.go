package main

import (
	"errors"
	"fmt"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

type Config struct {
	ServerName    string   `json:"server_name"`
	ServerTimeout string   `json:"server-timeout"`
	Endpoints     []string `json:"endpoints"`
}

func main() {
	conf, err := cueToStruct()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", &conf)
}

func cueToStruct() ([]*Config, error) {
	ctx := cuecontext.New()
	buildInstances := load.Instances([]string{"config.cue"}, nil)
	instance := ctx.BuildInstance(buildInstances[0])
	if instance.Err() != nil {
		return nil, instance.Err()
	}

	liberConfig := instance.Lookup("liberConfig")
	if !liberConfig.Exists() {
		return nil, errors.New("liberConfig does not exist")
	}
	iter, err := liberConfig.List()
	if err != nil {
		return nil, err
	}

	var configs []*Config
	for iter.Next() {
		var config Config
		err := iter.Value().Decode(&config)
		if err != nil {
			return nil, err
		}
		configs = append(configs, &config)
	}

	return configs, nil
}
