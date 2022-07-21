package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"gitlab.com/bracketnco/odoo-scraper/model"
	"gorm.io/gorm"
)

// Prompts for inputs and runs the scraper
func InitScraper(dbPath string) {
    var parentId uint

    // Ask if page is a parent-level
    first := CliPrompt("First level? (y/n): [y]", "y") 
    fmt.Printf("You have chosen: %s \n", first)

    // If no, show a list to select from
    if first == "n" {
        // Get pages
        db := InitDb(dbPath)
        var pages []model.Page
        db.Select("id", "title").Find(&pages)

        // Create a list of strings
        list := []string{}
        listID := []uint{}

        for _, v := range pages {
            title := fmt.Sprintf("%s", v.Title)
            list = append(list, title)
            listID = append(listID, v.ID)
        }

        // Prompt for seletion
        parentIdx := CliSelect(
            "Please choose a parent page from the list: ", 
            list,
            0,
        ) 

        // Get choice ID 
        parentId = listID[parentIdx-1]

        fmt.Printf("You have chosen: %d, with ID: %d \n", parentIdx, parentId)
    }
    
    // Url
    url := CliPrompt("Please provide the url: ", "") 
    fmt.Printf("You have chosen: %s \n", url)

    fmt.Println("Scraping ...")
    Scrape(
        dbPath, 
        url,
        parentId,
    )
}

// Scrapes subpages from the provided parent documentation path
func Scrape(dbPath string, url string, parentId uint) {
    // Connect to the database
    db := InitDb(dbPath)

    // Declare new collector
    c := colly.NewCollector(
        colly.MaxDepth(1),
        colly.Async(true),
    )

    // Set parallelism to 1
    c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1})

    if parentId == 0 {
        // Scrape parent page
        scrapeParent(db, c)
    } else {
        // Scrape child page
        scrapeChild(db, c, parentId)
    }

    // Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

    // c.Visit("https://www.odoo.com/documentation/15.0/applications/inventory_and_mrp/inventory.html")
    c.Visit(url)

    c.Wait()
}

// Scrape parent page
func scrapeParent(db *gorm.DB, c *colly.Collector) {
    // On every a element which has href attribute call callback
	c.OnHTML(
        "#o_content > div[role=\"main\"] > section:first-child", 
        func(e *colly.HTMLElement) {
            var page *model.Page
            page = new(model.Page)

            // Set parent page title
            title := e.ChildText("h1")
            page.Title = strings.Trim(title, "¶")

            // Set parent page original url
            page.UrlOriginal = e.Request.AbsoluteURL(e.Request.URL.Path)

            // Set parent page texts
            pageText := ""
            e.ForEach("section > p", func (i int, el *colly.HTMLElement) {
                html, err := el.DOM.Html()
                if err != nil {
                    panic(err)
                }
                pageText = pageText + "<p>" + strings.TrimSpace(html) + "</p>"
            })

            // Fix image urls
            r, _ := regexp.Compile("(?:[\"])((../)+)")
            text := model.Localized{
                Lang: "en",
                Text: r.ReplaceAllString(pageText, "\"https://www.odoo.com/documentation/15.0/"),
            }

            page.Text = append(page.Text, text)

            // Persist page
            db.Create(page)
        })
}

// Scrape child page
func scrapeChild(db *gorm.DB, c *colly.Collector, parentId uint) {
    // On every a element which has href attribute call callback
	c.OnHTML(
        ".toctree-wrapper .toctree-l1", 
        func(e *colly.HTMLElement) {
            
            category := ""

            // Get category
            e.ForEach(".toctree-l1 > a", func(i int, elCat *colly.HTMLElement) {
                category = strings.TrimSpace(elCat.Text) 

                // Get subcategory
                e.ForEach(".toctree-l1 > ul li.toctree-l2", func(i int, elSubCat *colly.HTMLElement) {
                    subcategory := ""
                    elSubCat.ForEach("li.toctree-l2 > a", func(i int, elSubCatTitle *colly.HTMLElement) {
                        subcategory = strings.TrimSpace(elSubCatTitle.Text)
                        parent := elSubCatTitle.DOM.Parent()

                        // Define link selector and check if there are subcategories
                        linkSelector := "li.toctree-l3 a"
                        if len(parent.Find(linkSelector).Nodes) == 0 {
                            subcategory = ""
                            linkSelector = "li.toctree-l2 a"
                        } 

                        // Loop through links and run the scraper
                        parent.Find(linkSelector).Each(func(i int, link *goquery.Selection) {
                            href, _ := link.Attr("href")
                            requestUrl := e.Request.AbsoluteURL(href)

                            c.Visit(requestUrl)

                            // Scrape page
                            scrapePage(
                               db, 
                               c, 
                               parentId, 
                               requestUrl, 
                               category, 
                               subcategory, 
                            )
                        })
                    })
                }) 
            }) 
        })
}

// Scrape documentation pages
func scrapePage(
    db *gorm.DB, 
    c *colly.Collector, 
    parentId uint,
    requestUrl string,
    category string, 
    subcategory string,
) {
    // Scrape the main body
    c.OnHTML("article#o_content div[role=\"main\"]", func(e *colly.HTMLElement) {
        if e.Request.URL.String() == requestUrl {
            // Get title
            titleNode := e.DOM.Find("h1")

            if titleNode.Length() > 0 {
                var page *model.Page
                page = new(model.Page)

                // Get the title
                page.Title = strings.Trim(titleNode.First().Text(), "¶")
                
                // Remove the parent node
                titleNode.First().Remove()

                pageText := ""

                // Scrape the content
                e.ForEach("section", func(i int, h *colly.HTMLElement) {
                    html, err := h.DOM.Html()
                    if err != nil {
                        panic(err)
                    }

                    pageText = pageText + strings.TrimSpace(html)
                    page.UrlOriginal = requestUrl
                })

                // Fix image urls
                r, _ := regexp.Compile("(?:[\"])((../)+)")
                text := model.Localized{
                    Lang: "en",
                    Text: r.ReplaceAllString(pageText, "\"https://www.odoo.com/documentation/15.0/"),
                }

                page.Text = append(page.Text, text)
                page.Parent = parentId
                page.Category = category
                page.Subcategory = subcategory

                // Persist page
                db.Create(page)
            }
        }
    })
}

