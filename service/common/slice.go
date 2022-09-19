package service

import "gitlab.com/bracketnco/odoo-scraper/model"

func DeleteLocalizedByIndex(a model.LocalizedList, idx int) model.LocalizedList {
    return append(a[:idx], a[idx+1:]...)
}
