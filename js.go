package main

import (
	`fmt`
	`path`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
)

func js(conf *config, _ glog.Logger) (err error) {
	if dir, dirErr := gox.IsDir(conf.path); nil != dirErr {
		panic(dirErr)
	} else if dir {
		conf.path = path.Join(conf.path, "package.json")
	}

	dependencies := make([]jsonElement, 0, len(conf.dependencies)+1)
	dependencies = append(dependencies, jsonElement{
		path:  "version",
		value: conf.version,
	})
	for _, _dependency := range conf.dependencies {
		dependencies = append(dependencies, jsonElement{
			path:  fmt.Sprintf("dependencies.%s", _dependency.module),
			value: _dependency.version,
		})
	}
	err = json(conf.path, dependencies...)

	return
}
