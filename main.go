// Copyright 2023 Ainsley Clark. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/ainsleyclark/stock-informer/config"
	"log"
)

func main() {
	cfg, err := config.Load("/Users/ainsley/Desktop/Web/apis/stock-informer")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
}
