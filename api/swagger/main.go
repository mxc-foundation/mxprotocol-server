// simple tool to merge different swagger definition into a single file
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const apiVersion = "1.0.0"
