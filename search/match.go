package search

import (
	"log"
)

// Result contains the result of a search.
//Results will be shown in struct with two fields, field and content both strings
type Result struct {
	Field   string
	Content string
}

var TransferResults []Result

//here is creation of Mathcer interface, this interface implements method Search with *Feed and string
//it returns []*Result and an error
// Matcher defines the behavior required by types that want
// to implement a new search type.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//Here is function Mathc it takes 4 parameters, Matcher, *Feed, string and chan<- *Result
//it saves two variables(one error) from matcher.Search(feed, searchTerm) this iss possible thanks
//to interface
//on populated variable range loop is launched which sends result data into results chanell

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		//fmt.Println("From Match:", result)
		TransferResults = append(TransferResults, *result)
		results <- result

	}
	//fmt.Println("Pakovanje za html", transferResults)

}

//Here is function Display which takes chan *Result
//it uses range over results and log.Printf Result strucot borh fileds
// Display writes results to the console window as they
// are received by the individual goroutines.
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.

	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)

	}

}
