package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetUrl(date string, location string, lang string) string {
	return fmt.Sprintf("https://www.stw.uni-heidelberg.de/external-tools/speiseplan/speiseplan.php?lang=%s&mode=Mensa+Bildungscampus+Heilbronn&date=%s", date, date)
}

func GetMenuData(date string, location string, lang string) *Menu {

	url := GetUrl(date, location, lang)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var menu Menu
	menu.MenuData.Dishes = make([]Dish, 0)

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		var dish Dish

		dish.Name = s.Find("td.block-mobile").Text()
		dish.Type = strings.Split(s.Find("td.inline-block-mobile").Eq(0).Text(), ":")[1]
		dish.Price.Student = strings.Split(s.Find("td.inline-block-mobile").Eq(1).Text(), ":")[1]
		dish.Price.Worker = strings.Split(s.Find("td.inline-block-mobile").Eq(2).Text(), ":")[1]
		dish.Price.Guest = strings.Split(s.Find("td.inline-block-mobile").Eq(3).Text(), ":")[1]

		menu.MenuData.Dishes = append(menu.MenuData.Dishes, dish)
	})

	menu.MenuData.Lang = lang
	menu.Date = date

	return &menu
}
