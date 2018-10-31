package goenv

import (
	"fmt"
	"testing"
)

// func Test_html(t *testing.T) {
//  res, _ := yunwangke.BaiduSearch("大连一对一课外辅导班", 1)
//  fmt.Println(res)
// }

func Test_Read(t *testing.T) {
	myConfig := new(Config)
	myConfig.Load(".env")
	fmt.Println(myConfig.Env("API_HOME"))
	// fmt.Printf("%v", myConfig.Mymap)
}
