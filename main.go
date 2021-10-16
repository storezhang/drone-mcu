package main

import (
	`errors`
	`strings`

	`github.com/mcuadros/go-defaults`
	`github.com/storezhang/glog`
	`github.com/storezhang/gox/field`
)

var notSupportLang = errors.New("不支持的语言")

func main() {
	var err error
	// 有错误，输出错误日志
	var logger glog.Logger
	if logger, err = glog.New(); nil != err {
		panic(err)
	}

	// 取各种参数
	conf := new(config)
	conf.lang = lang(env("LANG"))
	conf.filepath = env("FILEPATH")
	conf.version = env("VERSION")
	conf.dependencies = parseMoules(strings.Split(env("DEPENDENCIES"), ",")...)
	conf.replaces = parseReplaces(strings.Split(env("REPLACES"), ",")...)
	defaults.SetDefaults(conf)

	// 记录配置日志信息
	logger.Info(
		"加载配置完成",
		field.String("lang", string(conf.lang)),
		field.String("filepath", conf.filepath),
		field.String("version", conf.version),
		field.Strings("dependencies", conf.dependencyStrings()...),
	)

	switch conf.lang {
	case langGo:
		fallthrough
	case langGolang:
		err = golang(conf, logger)
	case langJavascript:
		fallthrough
	case langJs:
		err = js(conf, logger)
	case langDart:
		err = dart(conf, logger)
	case langMaven:
		err = maven(conf, logger)
	default:
		err = notSupportLang
	}

	if nil != err {
		panic(err)
	}
	logger.Info("修改模块描述文件成功", field.String("filepath", conf.filepath))
}
