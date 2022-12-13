// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/ainsleyclark/stock-informer/config"
	"github.com/ainsleyclark/stock-informer/job"
	"log"
)

func main() {
	cfg, err := config.Load("/Users/ainsley/Desktop/Web/apis/stock-informer/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	job := job.New(cfg)
	job.Boot()

	fmt.Println(cfg)

	//scraper := crawl.New()
	//scrape, err := scraper.Scrape("https://www.aersf.com/tokai-pack-white", ".sqs-block-button-element--medium.sqs-button-element--primary.sqs-block-button-element")
	//if err != nil {
	//	log.Fatalln(err)
	//}

}

//
//func check(cfg config.Config) {
//	for _, v := range cfg {
//		if
//	}
//}
