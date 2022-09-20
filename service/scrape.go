package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/narek-kerobian/odoo-scraper/model"
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
            title := fmt.Sprintf("%s", v.Title[0].Text)
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
    c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 3})


    // Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

    c.OnResponse(func(r *colly.Response) {
        // Scrape pages
        if parentId == 0 {
            // Scrape parent page
            scrapePage(
                db, 
                c, 
                0, 
                r.Request.URL.String(), 
                model.Localized{}, 
                model.Localized{},
            )
        } else {
            // Scrape child page
            scrapeChild(db, c, parentId)
        }
        
    })

    // c.Visit("https://www.odoo.com/documentation/15.0/applications/inventory_and_mrp/inventory.html")
    c.Visit(url)
    c.Wait()
}

// Scrape child page
func scrapeChild(db *gorm.DB, c *colly.Collector, parentId uint) {
    // On every a element which has href attribute call callback
	c.OnHTML(
        ".toctree-wrapper .toctree-l1", 
        func(e *colly.HTMLElement) {
            
            category := model.Localized{
                Lang: "en",
                Text: "",
            }

            // Get category
            e.ForEach(".toctree-l1 > a", func(i int, elCat *colly.HTMLElement) {
                category.Text = strings.TrimSpace(elCat.Text) 

                // Get subcategory
                e.ForEach(".toctree-l1 > ul li.toctree-l2", func(i int, elSubCat *colly.HTMLElement) {
                    subcategory := model.Localized{
                        Lang: "en",
                        Text: "",
                    }

                    elSubCat.ForEach("li.toctree-l2 > a", func(i int, elSubCatTitle *colly.HTMLElement) {
                        subcategory.Text = strings.TrimSpace(elSubCatTitle.Text)
                        parent := elSubCatTitle.DOM.Parent()

                        // Define link selector and check if there are subcategories
                        linkSelector := "li.toctree-l3 a"
                        if len(parent.Find(linkSelector).Nodes) == 0 {
                            subcategory.Text = ""
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
    category model.Localized, 
    subcategory model.Localized,
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
                pageTitle := model.Localized{
                    Lang: "en",
                    Text: strings.Trim(titleNode.First().Text(), "¶"),
                }
                page.Title = append(page.Title, pageTitle)

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
                if len(category.Text) > 0 {
                    page.Category = append(page.Category, category)
                }
                if len(subcategory.Text) > 0 {
                    page.Subcategory = append(page.Subcategory, subcategory)
                }

                // Persist page
                db.Create(page)
            }
        }
    })
}

