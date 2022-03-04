package search

import (
	"fmt"
	"log"
	"sync"
)

// A map of registered matchers for searching.
//1.Create map called matcher which has key as string and Matcher as user-defined type value
var matchers = make(map[string]Matcher)

// Run performs the search logic.
//2.Create function Run that accepts searchTerm a string as an argument
//then, run function RetrieveFeeds from different file within same package and save returned values
//as variable feeds and check for possible errors. log Fatal error if there is one
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through.
	feeds, err := RetrieveFeeds()
	if err != nil {
		fmt.Println(err)
	}

	// Create an unbuffered channel to receive match results to display. called results, pointer to Result
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	//3. WaitGroup add, with argument of feeds length
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	//4. Range over feeds slice and use feed variable, then save to variables matcher, exists
	//everything of matchers[feed.Type]
	//if there is notihing in variable exists, set matcher to the matchers["default"]
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		/*
					This is feed &{npr http://www.npr.org/rss/rss.php?id=1001 rss}
			This is matcher {}
			This is matcher, after if statement {}


		*/

		// Launch the goroutine to perform the search.
		//to go routine function pass two arguments matcher, and feed, use Match function with
		//four agruments matcher, feed, searchTerm, results
		//then create waitGroup and set it to Done
		//initialize anonymous function with passing two argumnets matcher and feed
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)

			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	Display(results)

}

// Register is called to register a matcher for use by the program.
//create function Register that takes two inputs feedType as string and matchers as Matcher
//create loop with _, exists that sets matchers[feedType]; exists
//log fatalln with feedType and MAtcher already registered
//log println "Register", feedType, matcher
//set matchers[feedType] to matcher

func Register(feedType string, matcher Matcher) {
	fmt.Println("Registering a matcher:", matchers)
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
