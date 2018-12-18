package notion

import (
	"github.com/kjk/notionapi"
	"log"
)

func Init() *notionapi.Page {
	client := &notionapi.Client{}
	pageID := "4ceee3ab2dc74897b9b2c6850655cc3e"
	page, err := SerializePage(client, pageID)
	if err != nil {
		log.Fatalf("DownloadPage() failed with %s\n", err)
	}
	return page
}
