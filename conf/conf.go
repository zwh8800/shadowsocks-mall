package conf

import (
	"flag"
	"path/filepath"

	"github.com/golang/glog"
	"gopkg.in/gcfg.v1"
)

type config struct {
	DbConf struct {
		Driver        string
		Dsn           string
		MaxConnection int
	}
	Env struct {
		Prod       bool
		ServerPort string
	}
	Site struct {
		Name        string
		FuckIcp     string
		Description string
		BaseUrl     string
		NoteUrl     string
		TagUrl      string
		PageUrl     string
		RssUrl      string
		ArchiveUrl  string
		SearchUrl   string
		StaticUrl   string
		AuthorName  string
		AuthorEmail string
		NotePerPage int64
		Language    string
		LicenseName string
		LicenseUrl  string
		ICP         string
	}
	Social struct {
		Github   string
		Weibo    string
		Twitter  string
		Linkedin string
		About    string
	}
	Haha struct {
		Data []string
	}
}

var Conf config

func ReadConf(filename string) error {
	absFile := filename
	if !filepath.IsAbs(filename) {
		var err error
		absFile, err = filepath.Abs(filename)
		if err != nil {
			return err
		}
	}
	return gcfg.ReadFileInto(&Conf, absFile)
}

func init() {
	configFilename := flag.String("config", "shadowsocks-mall.gcfg", "specify a config file")
	flag.Parse()
	glog.Infoln("configuring...")

	if err := ReadConf(*configFilename); err != nil {
		glog.Fatalf("error occored: %s", err)
		panic(err)
	}
}
