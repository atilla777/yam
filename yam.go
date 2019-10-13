package main

import (
	"fmt"
	"os"
	"io"
	//"io/ioutil"
	"path/filepath"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/src-d/go-billy.v4/memfs"
)

func main() {
	const url = "https://github.com/Yara-Rules/rules"
	
	dirs := []string{
		"Antidebug_AntiVM",
		"CVE_Rules",
		"Capabilities",
		"Crypto",
		"Exploit-Kits",
		"Malicious_Documents",
		//"Mobile_Malware",
		"Packers",
		"Webshells",
		"email",
		"malware",
		"utils",
	}
	
	fs := memfs.New()
	storer := memory.NewStorage()


	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: url,
	})
	CheckIfError(err)
	
	rules, err := os.OpenFile("rules.yara", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	CheckIfError(err)
	defer rules.Close()
	
	for _, dir := range dirs {
		files, err := fs.ReadDir(dir)
		CheckIfError(err)
		fmt.Println(dir)
		
		for _, f := range files {
			path := dir + "/" + f.Name()
			if filepath.Ext(path) == ".yar" {
				file, err := fs.Open(path)
				CheckIfError(err)
				_, err = io.Copy(rules, file)
				CheckIfError(err)
			}
		}
	}
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}