package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"gitlab.com/bracketnco/odoo-scraper/model"
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
        colly.Async(false),
    )

    // Set parallelism to 1
    c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1})

    // On every a element which has href attribute call callback
	c.OnHTML(
        ".toctree-l1.o_menu_applications_inventory_and_mrp_inventory_management a[href].reference.internal", 
        func(e *colly.HTMLElement) {
            link := e.Attr("href")

            // Print link
            if link != "#" {
                fmt.Printf("Link found: %q -> %s\n", e.Text, link)
                // Visit link found on page
                // Only those links are visited which are in AllowedDomains
                c.Visit(e.Request.AbsoluteURL(link))

                // Scrape the main body
                c.OnHTML("article#o_content", func(e *colly.HTMLElement) {
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
                            page.UrlOriginal = link
                        })

                        // Fix image urls
                        r, _ := regexp.Compile("(?:[\"])((../)+)")
                        text := model.Localized{
                            Lang: "en",
                            Text: r.ReplaceAllString(pageText, "\"https://www.odoo.com/documentation/15.0/"),
                        }

                        page.Text = append(page.Text, text)
                        page.Parent = parentId

                        // Persist page
                        db.Create(page)
                    }
                })
            }
	})

    // Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

    // c.Visit("https://www.odoo.com/documentation/15.0/applications/inventory_and_mrp/inventory.html")
    c.Visit(url)
}
