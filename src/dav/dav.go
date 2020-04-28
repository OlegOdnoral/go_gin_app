package dav

import (
	"fmt"
	//"golang.org/x/net/webdav"
	"io/ioutil"
	"os"
)

func GetFsInfo() {
	fmt.Println("\n")
	files, err := ioutil.ReadDir("./files")
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
	os.Chdir("files")
	//os.Mkdir("texts", 666)
	//os.Create("./texts/kek.txt")
	//webdav.Dir =
	//fmt.Printf("%+v\n", webdav.NewMemFS())
	//fmt.Printf("%+v\n", webdav.NewMemFS().Mkdir())
	fmt.Println("\n")
}
