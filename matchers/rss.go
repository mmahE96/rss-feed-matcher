package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"../search"
)

type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher implements the Matcher interface.
type rssMatcher struct{}

// init registers the matcher with the program.

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

//create function that has ident name Search and accepts *search.Feed and string(term)
//this function has rssMatcher as receiver and returns []*search.Result and error
//create var reuslts of type []*search.Result

//log printf, feed.TYPE, feed.Name, feed.URI

//retreive data with rssMatcher.retreive(*search.Feed) and check for error and save data to document var

//iterate over document.Channel.Item.. save regexp.MatchString(string(term). channelItem.Title) to var
//and check for the errors

//if that variable is true append results to the results var with Field:   "Title",
//Content: channelItem.Title,

//check the description of the feed for search term similar to title

//if results are find, save them on same way as title results

// Search looks at the document for the specified search term.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// Retrieve the data to search.
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// Check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// retrieve performs a HTTP Get request for the rss feed and decodes the results.

//create function that accepts rssMatcher as reciever, is of retrieve identifier that accpets *search.Feed
//as aprameter, returns *rssDocument and error

//it checks if feed.URI is emty string and if true returns nil, errors.New("No rss feed uri provided")

//fetch URIs with http.Get(.......), checks for error and saves fetched data into the var

//it defers closing of response body

//checks for resposne status 200, with goal to  check if proper response is received

//it creates var of rrsDocument type
//decodes it with xml.NewDecoder(resp.body).decode(&document)
//it returns &document and err
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	/*
		Resp(response) content:
		Fresh from URL: &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, max-age=0]
		Connection:[keep-alive]
		Content-Type:[text/xml; charset=UTF-8]
		Date:[Mon, 28 Feb 2022 09:31:42 GMT]
		Etag:[x/Vikwsi4fUYfaXwpj5qcXRwlNk]
		Expires:[Mon, 28 Feb 2022 09:31:42 GMT]
		Last-Modified:[Mon, 28 Feb 2022 09:05:29 GMT]
		Server:[GSE] Vary:[Accept-Encoding]
		X-Content-Type-Options:[nosniff]
		X-Xss-Protection:[1; mode=block]] 0xc0003b2100 -1 [] false true map[] 0xc0003a0300 <nil>}
	*/
	/*
			fmt.Println("REsponse body:", reflect.TypeOf(resp.Body))
		REsponse body: *http.gzipReader
		REsponse body: *http.gzipReader
		REsponse body: *http.gzipReader
		REsponse body: *http.gzipReader
		REsponse body: *http.gzipReader
		REsponse body: *http.gzipReader
		REsponse body: *http.http2gzipReader
		REsponse body: *http.http2gzipReader
	*/

	// Close the response once we return from the function.
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	fmt.Println("REsponse body:", &document)

	/*

			PART OF DOCUMENT DECODED FILE:
			document freshly populated {{ } {{ }          {{ }   } []}}
		2022/02/28 10:38:32 expected element type <rss> but have <html>
		document freshly populated {{ rss} {{ channel} Politics :
		NPR NPR's expanded coverage of U.S. and world politics,
		the latest news from Congress and the White House, and elections.
		https://www.npr.org/templates/story/story.php?storyId=1014
		Sun, 27 Feb 2022 17:13:36 -0500
			  en   {{ image} https://media.npr.org/images/podcasts/primary/npr_generic_image_300.jpg?s=200
			  Politics https://www.npr.org/templates/story/story.php?storyId=1014}
			  [{{ item} Sun, 27 Feb 2022 17:13:36 -0500 Montgomery,
			  Ala., mayor on leading the city through the voting rights battle NPR's Michel Martin speaks with
			  Steven Reed, the first Black mayor of Montgomery, Ala.
			  https://www.npr.org/2022/02/27/1083379284/montgomery-ala-mayor-on-leading-the-city-through-the-voting-rights-battle
			  https://www.npr.org/2022/02/27/1083379284/montgomery-ala-mayor-on-leading-the-city-through-the-voting-rights-battle }
			  {{ item} Sun, 27 Feb 2022 17:13:36 -0500 Sen.
			  Sullivan supports sending more military aid to Ukraine NPR's
			  Michel Martin speaks with Sen. Dan Sullivan (R-Alaska) about the Russian
			  invasion of Ukraine.
			  https://www.npr.org/2022/02/27/1083379277/sen-sullivan-supports-sending-more-military-aid-to-ukraine https://www.npr.org/2022/02/27/1083379277/sen-sullivan-supports-sending-more-military-aid-to-ukraine
			  } {{ item} Sun, 27 Feb 2022 08:24:16 -0500
			  Co-chair of Senate Ukraine Caucus calls for stricter sanctions against Russia
			  Sarah McCammon asks Sen. Dick Durbin, D-Ill., about the crisis in Ukraine and the
			  American government's response.

	*/
	//fmt.Println("document freshly populated", document)
	return &document, err
}
