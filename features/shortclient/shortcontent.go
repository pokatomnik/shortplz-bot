package shortclient

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/samber/mo"
)

const (
	metaTag     = "meta[name=\"description\"]"
	contentAttr = "content"
)

func getShortContent(document *goquery.Document) mo.Result[string] {
	var el = document.Find(metaTag)
	if el == nil {
		return mo.Err[string](errors.New(errorMetaTagMissing))
	}
	content, exists := el.Attr(contentAttr)
	if !exists {
		return mo.Err[string](errors.New(errorMissingAttrContent))
	}
	return mo.Ok(content)
}
