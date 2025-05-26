package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/xtt28/nutrislice/network"
)

const urlFormat = "https://bergen.api.nutrislice.com/menu/api/weeks/school/bergen-academy/menu-type/lunch/%d/%d/%d"

// Create, Grill, Veg Out
var wantedStations = []int{3085, 3087, 48840}

func main() {
	day := time.Now()
	
	fmt.Printf("Lunch for %d %s %d\n", day.Day(), day.Month(), day.Year())
	url := fmt.Sprintf(urlFormat, day.Year(), day.Month(), day.Day())
	data, err := network.GetMenuWeekData(url)
	if err != nil {
		panic(err)
	}

	dayOfWeek := day.Weekday()
	menuDayData := data.Days[dayOfWeek]
	if len(menuDayData.MenuItems) == 0 {
		fmt.Println("There is no lunch today, you stupid pig.")
		return
	}

	for _, item := range menuDayData.MenuItems {
		if !slices.Contains(wantedStations, item.StationID) {
			continue;
		}
		if item.IsSectionTitle || item.IsStationHeader {
			fmt.Printf("\n== %s ==\n", item.Text)
			continue;
		}
		food := item.Food
		if item.Category != "entree" && item.Category != "meat" {
			fmt.Print("\t")
		}
		fmt.Println(food.Name)
	}
}
