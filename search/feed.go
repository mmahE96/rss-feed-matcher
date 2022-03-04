package search

import (
	"encoding/json"
	"os"
)

//create constat that will use data.json file from data folder
const dataFile = "data/data.json"

// Feed contains information we need to process a feed.
//create Feed struct that will store data and add struct tags to be ready for json encoding and etc.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.

//create function RetrieveFeeds, that has no parameters, and returns slice of pointers to feed and an error
//it uses Open from os package and takes dataFile as an argument, checks for an error

func RetrieveFeeds() ([]*Feed, error) {

	/*
		const dataFile = "data/data.json" is just a llocation
		fmt.Println("Here is dataFile:", dataFile)
		we get this:
		 Here is dataFile: data/data.json
	*/

	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	/*
		if we print file here:
		fmt.Println("Here is rare file:", file)
		we get:
		Here is rare file: &{0xc0000b2240}

	*/

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	//fmt.Println("decoded feeds into the Feed struct, here as a slice:", feeds)

	/*
				decoded feeds into the Feed struct, here as a slice:
				 [0xc000100db0 0xc000100f00 0xc000100f30 0xc000100f60 0xc000100fc0 0xc000100ff0 0xc000101020
				  0xc000101050 0xc000101080 0xc0001010b0 0xc0001010e0 0xc000101110 0xc000101140 0xc000101170
				   0xc0001011a0 0xc0001011d0 0xc000101200 0xc000101230 0xc000101260 0xc000101290 0xc0001012c0
				    0xc0001012f0 0xc000101320 0xc000101350 0xc000101380 0xc0001013b0 0xc0001013e0 0xc000101410
					 0xc000101440 0xc000101470 0xc0001014a0 0xc0001014d0 0xc000101500 0xc000101530 0xc000101560
					  0xc000101590 0xc0001015c0]

			with range we have this:
			for i, d := range feeds {

				fmt.Println(i, d)
			}

			0 &{npr http://www.npr.org/rss/rss.php?id=1001 rss}
		1 &{npr http://www.npr.org/rss/rss.php?id=1008 rss}
		2 &{npr http://www.npr.org/rss/rss.php?id=1006 rss}
		3 &{npr http://www.npr.org/rss/rss.php?id=1007 rss}
		4 &{npr http://www.npr.org/rss/rss.php?id=1057 rss}
		5 &{npr http://www.npr.org/rss/rss.php?id=1021 rss}
		6 &{npr http://www.npr.org/rss/rss.php?id=1012 rss}
		7 &{npr http://www.npr.org/rss/rss.php?id=1003 rss}
		8 &{npr http://www.npr.org/rss/rss.php?id=2 rss}
		9 &{npr http://www.npr.org/rss/rss.php?id=3 rss}
		10 &{npr http://www.npr.org/rss/rss.php?id=5 rss}
		11 &{npr http://www.npr.org/rss/rss.php?id=13 rss}
		12 &{npr http://www.npr.org/rss/rss.php?id=46 rss}
		13 &{npr http://www.npr.org/rss/rss.php?id=7 rss}
		14 &{npr http://www.npr.org/rss/rss.php?id=10 rss}
		15 &{npr http://www.npr.org/rss/rss.php?id=39 rss}
		16 &{npr http://www.npr.org/rss/rss.php?id=43 rss}
		17 &{bbci http://feeds.bbci.co.uk/news/rss.xml rss}
		18 &{bbci http://feeds.bbci.co.uk/news/business/rss.xml rss}
		19 &{bbci http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml rss}
		20 &{cnn http://rss.cnn.com/rss/cnn_topstories.rss rss}
		21 &{cnn http://rss.cnn.com/rss/cnn_world.rss rss}
		22 &{cnn http://rss.cnn.com/rss/cnn_us.rss rss}
		23 &{cnn http://rss.cnn.com/rss/cnn_allpolitics.rss rss}
		24 &{cnn http://rss.cnn.com/rss/cnn_crime.rss rss}
		25 &{cnn http://rss.cnn.com/rss/cnn_tech.rss rss}
		26 &{cnn http://rss.cnn.com/rss/cnn_health.rss rss}
		27 &{cnn http://rss.cnn.com/rss/cnn_topstories.rss rss}
		28 &{foxnews http://feeds.foxnews.com/foxnews/opinion?format=xml rss}
		29 &{foxnews http://feeds.foxnews.com/foxnews/politics?format=xml rss}
		30 &{foxnews http://feeds.foxnews.com/foxnews/national?format=xml rss}
		31 &{foxnews http://feeds.foxnews.com/foxnews/world?format=xml rss}
		32 &{nbcnews http://feeds.nbcnews.com/feeds/topstories rss}
		33 &{nbcnews http://feeds.nbcnews.com/feeds/usnews rss}
		34 &{nbcnews http://rss.msnbc.msn.com/id/21491043/device/rss/rss.xml rss}
		35 &{nbcnews http://rss.msnbc.msn.com/id/21491571/device/rss/rss.xml rss}
		36 &{nbcnews http://rss.msnbc.msn.com/id/28180066/device/rss/rss.xml rss}


	*/

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
