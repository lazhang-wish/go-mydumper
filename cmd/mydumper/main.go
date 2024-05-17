/*
 * go-mydumper
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/xelabs/go-mydumper/common"
	"github.com/xelabs/go-mydumper/config"

	"github.com/xelabs/go-mysqlstack/xlog"
)

var (
	flagConfig     string
	debugLogConfig bool

	log = xlog.NewStdLog(xlog.Level(xlog.INFO))
)

func initFlags() {
	flag.StringVar(&flagConfig, "c", "", "config file")
	flag.BoolVar(&debugLogConfig, "D", false, "enable debug log")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " -c conf/mydumper.ini.sample")
	flag.PrintDefaults()
}

func main() {
	initFlags()
	flag.Usage = func() { usage() }
	flag.Parse()

	if debugLogConfig {
		log = xlog.NewStdLog(xlog.Level(xlog.DEBUG))
	}

	log.Info("log level " + strconv.FormatBool(debugLogConfig))

	if flagConfig == "" {
		usage()
		os.Exit(0)
	}

	args, err := config.ParseDumperConfig(flagConfig)
	common.AssertNil(err)

	if _, err := os.Stat(args.Outdir); os.IsNotExist(err) {
		x := os.MkdirAll(args.Outdir, 0o777)
		common.AssertNil(x)
	}

	common.Dumper(log, args)
}
