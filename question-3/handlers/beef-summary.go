package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func BeefSummary(c *fiber.Ctx) error {

	// fetch api
	req, _ := http.NewRequest(http.MethodGet, "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text", nil)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// read response
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// clean data
	data := strings.ReplaceAll(string(b), ".", "")
	data = strings.ReplaceAll(data, ",", "")
	dataClean := strings.Fields(data)

	// count beef name
	mapBeefName := map[string]int{}
	for _, v := range dataClean {
		beef := strings.ToLower(v)
		if val, ok := mapBeefName[beef]; !ok {
			mapBeefName[beef] = 1
		} else {
			n := val + 1
			mapBeefName[beef] = n
		}
	}

	return c.JSON(fiber.Map{
		"beef": mapBeefName,
	})
}
