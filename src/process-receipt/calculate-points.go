package processreceipt

import (
	"fmt"
	"log/slog"
	"math"
	"strings"
	"time"
	"unicode"
)

var PointValues = map[string]float64{
	"POINTS_PER_RETAILER_LETTER":  1,
	"POINTS_IF_ROUND_DOLLAR":      50,
	"POINTS_IF_25_MULTIPLE":       25,
	"POINTS_PER_TWO_ITEMS":        5,
	"MULTIPLE_IF_DESC_DIVIS_BY_3": 0.2,
	"POINTS_IF_ODD_DAY":           6,
	"POINTS_IF_BTWN_2_AND_4":      10,
}

func getRetailerNamePoints(name string) int {
	count := 0
	for _, char := range name {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	points := count * int(PointValues["POINTS_PER_RETAILER_LETTER"])
	slog.Info("| " + fmt.Sprint(points) + " points from retailer name")
	return points
}

func getRoundDollarTotalPoints(total float64) int {
	if math.Mod(total, 1) == 0 {
		points := int(PointValues["POINTS_IF_ROUND_DOLLAR"])
		slog.Info("| " + fmt.Sprint(points) + " points for round dollar amount")
		return points
	}
	return 0
}

func getMultipleOf25Points(total float64) int {
	if math.Mod(total, 0.25) == 0 {
		points := int(PointValues["POINTS_IF_25_MULTIPLE"])
		slog.Info("| " + fmt.Sprint(points) + " points for being multiple of .25")
		return points
	}
	return 0
}

func getPointsPerTwoItems(numItems int) int {
	numPairs := int(numItems / 2)
	points := numPairs * int(PointValues["POINTS_PER_TWO_ITEMS"])
	slog.Info("| " + fmt.Sprint(points) + " points for item count")
	return points
}

func getItemDescPoints(desc string, price float64) int {
	length := len(strings.TrimSpace(desc))
	if length%3 == 0 {
		points := int(math.Ceil(price * PointValues["MULTIPLE_IF_DESC_DIVIS_BY_3"]))
		slog.Info("| " + fmt.Sprint(points) + " points from item description")
		return points
	}
	return 0
}

func getOddDayPoints(datetime time.Time) int {
	if datetime.Day()%2 == 1 {
		points := int(PointValues["POINTS_IF_ODD_DAY"])
		slog.Info("| " + fmt.Sprint(points) + " points for being an odd day")
		return points
	}
	return 0
}

func getTimeRangePoints(datetime time.Time) int {
	hour := datetime.Hour()
	if hour >= 14 && hour <= 16 {
		points := int(PointValues["POINTS_IF_BTWN_2_AND_4"])
		slog.Info("| " + fmt.Sprint(points) + " points for being between 2-4pm")
		return points
	}
	return 0
}
