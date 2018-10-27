package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

var infile string
var outfile string

//Config datatype
type Config struct {
	Settings struct {
		Gama     float64 `json:"gama"`
		Bright   float64 `json:"bright"`
		Contrast float64 `json:"contrast"`
	}
}

func main() {
	config, _ := LoadConfig("config.json")
	fmt.Println("Importing config file...")

	fmt.Print("Enter file to modify (raw.png): ")
	fmt.Scan(&infile)

	fmt.Print("Save file as (final.png): ")
	fmt.Scan(&outfile)

	src, err := imaging.Open(infile)
	if err != nil {
		log.Fatal("Failed to load image.", err)
	}

	gamacor := imaging.AdjustGamma(src, config.Settings.Gama)
	bright := imaging.AdjustBrightness(gamacor, config.Settings.Bright)
	contrast := imaging.AdjustContrast(bright, config.Settings.Contrast)

	err = imaging.Save(contrast, outfile)
	if err != nil {
		log.Fatal("Could not save the image.", err)
	}
}

//LoadConfig file
func LoadConfig(filename string) (Config, error) { // Load config file
	var config Config
	configFile, err := os.Open(filename)

	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}
