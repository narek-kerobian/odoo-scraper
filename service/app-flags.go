package service

import (
	"flag"
)

// Defines and stores command line flag values
type FlagVars struct {
    Scrape  bool
    Serve   bool
}

// Initializes and parses command line flags
func (fv *FlagVars) ParseFlags() {
    scrape := flag.Bool("scrape", false, "Scrape provided url")
    serve := flag.Bool("serve", false, "Start the web server")
    flag.Parse()

    fv.Scrape = *scrape
    fv.Serve = *serve
}
