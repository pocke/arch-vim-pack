package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	err := Main()
	if err != nil {
		panic(err)
	}
}

func Main() error {
	dir, err := ioutil.TempDir("", "arch-vim-pack")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	files := []string{
		"archlinux.vim",
		"gvim.desktop",
		"gvim.install",
		"vimrc",
	}

	for _, f := range files {
		if err := PutFile(dir, f); err != nil {
			return err
		}
	}
	if err := PutPKGBUILD(dir); err != nil {
		return err
	}

	prevDir, err := os.Getwd()
	if err != nil {
		return err
	}
	os.Chdir(dir)
	defer os.Chdir(prevDir)
	if err := Exec("makepkg", "-s", "--noconfirm"); err != nil {
		return err
	}

	return Exec("bash", "-c", "cp *pkg.tar.xz "+prevDir)
}

func Exec(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func PutFile(dir, file string) error {
	b, err := Asset("data/" + file)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(dir, file), b, 0644)
}

func PutPKGBUILD(dir string) error {
	v, err := GetLatestVimVersion()
	if err != nil {
		return err
	}

	b, err := Asset("data/PKGBUILD")
	if err != nil {
		return err
	}
	s := strings.Replace(string(b), "VIMVERSION", v, 1)
	return ioutil.WriteFile(path.Join(dir, "PKGBUILD"), []byte(s), 0644)
}

func GetLatestVimVersion() (string, error) {
	doc, err := goquery.NewDocument("https://github.com/vim/vim/releases")
	if err != nil {
		return "", err
	}
	text := doc.Find(".tag-name").First().Text()
	return text[1:], nil
}
