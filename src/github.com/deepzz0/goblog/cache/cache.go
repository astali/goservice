package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/deepzz0/goblog/models"
	"github.com/deepzz0/logd"
)

type cache struct {
	BackgroundLeftBar  map[string]string
	BackgroundLeftBars []*models.Leftbar
	BuildVersion       string
}

var Cache = NewCache()

func NewCache() *cache {
	return &cache{BackgroundLeftBar: make(map[string]string)}
}

func init() {
	doReadBackLeftBarConfig()
	doReadBuildVersionConfig()
}

var path, _ = os.Getwd()

func doReadBackLeftBarConfig() {
	b, err := ioutil.ReadFile(path + "/conf/backleft.conf")
	if err != nil {
		logd.Fatal(err)
	}
	err = json.Unmarshal(b, &Cache.BackgroundLeftBars)
	if err != nil {
		logd.Fatal(err)
	}
	for _, v := range Cache.BackgroundLeftBars {
		if v.ID != "" {
			Cache.BackgroundLeftBar[v.ID] = v.ID
		}
	}
}

func doReadBuildVersionConfig() {
	b, err := ioutil.ReadFile(path + "/version")
	if err != nil {
		logd.Error(err)
	}
	Cache.BuildVersion = string(b)
}
