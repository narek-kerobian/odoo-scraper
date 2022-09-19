# Odoo Documentation Scraper

# Uses
You can use the tool to both scrape odoo documentation pages and serve the scraped data to be able to translate them to other languages using a simple web ui.   

## Scrape
In order to start scraping, grab a url of the page you want to scrape and head to the terminal emulator of your choice.   

Run the program with the `-scrape` flag and you will be presented with a prompt asking you if the page is first level or not (like [this one](https://www.odoo.com/documentation/15.0/applications/inventory_and_mrp/inventory.html)).   
```
First level? (y/n): [y]
```
The default selection is `y` (yes), which means no sub-pages will be scraped.   

If you want to scrape sub-pages of a category, choose `n` (no) which will then prompt you with the list of first level pages that you've already scraped.   
Provide the corresponding number of the first level page you want newly scraped pages to be associated with.   
```
You have chosen: n
Please choose a parent page from the list:
1. Inventory
1
You have chosen: 1, with ID: 1
Please provide the url: https://www.odoo.com/documentation/15.0/applications/inventory_and_mrp/inventory.html
You have chosen: https://www.odoo.com/documentation/15.0/applications/inventory_and_mrp/inventory.html
Scraping ...
```

## Serve
You can use the built web server to view and interact with the scraped pages in a convenient web ui.
Run the program with `-serve` flag to start the server under the port specified in `APP_PORT` env variable.   

## Flags
| Flag      | Description                     |
|-----------|---------------------------------|
| `-scrape` | Scrape pages by providing a url |
| `-serve`  | Serve scraped pages in web ui   |

## ENV Variables
| Variable        | Default value     | Description                                                |
|-----------------|-------------------|------------------------------------------------------------|
| `APP_PORT`      | 8088              | Specifiec which port to use when `-serve` flag is selected |
| `APP_DB_PATH`   | data/odoo_docs.db | Location of the SQLite database                            |
| `APP_LANGUAGES` | "en,hy"           | Coma separated list of supported languages                 |

