package base

import "github.com/astaxie/beego/toolbox"

func siteMap() error {
	return nil
}

func rssAtom() error {
	return nil
}

func init() {
	toolbox.AddTask(
		"generate sitemap.xml.gz",
		toolbox.NewTask("sitemap.xml.gz", "0 0 3 * * *", siteMap),
	)
	toolbox.AddTask(
		"generate rss.atom",
		toolbox.NewTask("rss.atom", "0 0 */8 * * *", siteMap),
	)
}
