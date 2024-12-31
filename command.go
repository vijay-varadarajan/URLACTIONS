package main

import (
	"flag"
	"fmt"
)

const version = "1.0.0"

type cmdFlags struct {
	shorten  string
	qrcode   string
	save     bool
	filename string
	validate string
}

func newCmdFlag() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.shorten, "shorten", "", "Shorten the provided url string")

	flag.StringVar(&cf.qrcode, "qrcode", "", "Generate a QR code for the provided URL (Use -save and -filename to save image to a PNG file)")
	flag.BoolVar(&cf.save, "save", false, "Save the qrcode with a default name, PNG format (Works only while generating QR code for URL)")
	flag.StringVar(&cf.filename, "filename", "", "Specify the filename to save the qrcode (Works only while generating QR code for URL)")

	flag.StringVar(&cf.validate, "validate", "", "Check if the provided URL is valid and the server is responding")

	help := flag.Bool("help", false, "Display help")
	versionFlag := flag.Bool("version", false, "Display version")

	flag.Parse()

	if *help {
		flag.Usage()
		return nil
	}

	if *versionFlag {
		fmt.Println("Version:", version)
		return nil
	}

	return &cf
}

func (cf *cmdFlags) execute() {
	
	if cf == nil {
		return
	}

	switch {
	case cf.shorten != "":
		goshorten(cf.shorten)
	case cf.qrcode != "":
		if cf.save {
			if cf.filename != "" {
				goqrgen(cf.qrcode, true, cf.filename)
			} else {
				goqrgen(cf.qrcode, true, "")
			}
		} else {
			goqrgen(cf.qrcode, false, "")
		}
	case cf.validate != "":
		govalidateURL(cf.validate)
	default:
		fmt.Println("Invalid arguments. Use -help for list of options")
	}
}