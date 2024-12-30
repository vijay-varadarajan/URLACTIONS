package main

import "flag"

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

type cmdFlags struct {
	shorten string
	qrcode string 
	save string
}

func newCmdFlag() *cmdFlags {
	cf := cmdFlags{} 

	flag.StringVar(&cf.shorten, "shorten", "", "Shorten the provided url string")
	flag.StringVar(&cf.qrcode, "qrcode", "", "Convert the provided url to a qrcode")
	flag.StringVar(&cf.save, "save", "", "Save the qrcode in given filename")

	flag.Parse()

	return &cf
}

