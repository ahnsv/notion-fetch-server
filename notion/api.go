package notion

import (
	"encoding/json"
	"github.com/kjk/notionapi"
	"io/ioutil"
	"path/filepath"
)

var (
	cacheDir = "cache"
)

func GetNotionPage(c *notionapi.Client, pageID string) (*notionapi.Page, error) {
	res, err := c.DownloadPage(pageID)
	if err == nil {
		return res, nil
	}
	return nil, err
}

func SerializePage(c *notionapi.Client, pageID string) (*notionapi.Page, error) {
	// TODO: Check if it is already loaded
	page, err := GetNotionPage(c, pageID)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(page, "", "	")
	if err != nil {
		return nil, err
	}
	jsonDir := filepath.Join(cacheDir, "/ahnsv-", pageID, ".json")
	err = ioutil.WriteFile(jsonDir, data, 0644)
	if err != nil {
		panic(err)
	}
	return page, nil
}
