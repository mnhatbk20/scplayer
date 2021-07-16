package main

import(
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/jfbus/httprs"
	"github.com/faiface/beep/mp3"
	"github.com/mnhatbk20/scplayer/player"
)

func main() {
	var urlPlayList string
	var num int

	fmt.Printf("Please Wait...\n")

	clientID, err := player.GetClientID()
	if err != nil {
		return
	}

	fmt.Printf("URL Playlist: ")
	fmt.Scanf("%s\n", &urlPlayList) 
	fmt.Printf("Please Wait...\n")

	url := "https://api-v2.soundcloud.com/resolve?url="
	url = url + urlPlayList + "&client_id="+ clientID
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	var playlist player.PlayList
	json.Unmarshal([]byte(string(body)), &playlist)
	var trackIDs []int
	for _, track := range playlist.Tracks {
		trackIDs =  append(trackIDs, track.ID)	
	}

	var mediaURLTracks []string
	var trackNames []string
	var mp3URLTracks  []string
	for _, trackID := range trackIDs {
		url = "https://api-v2.soundcloud.com/tracks/"	
		url =  url + strconv.Itoa(trackID) + "?client_id=" + clientID
		resp, _ = http.Get(url)
		body, _ = ioutil.ReadAll(resp.Body)
		var track player.Track
		json.Unmarshal([]byte(string(body)), &track)
		mediaURLTracks = append(mediaURLTracks, track.Media.Transcodings[1].URL )
		trackNames =  append(trackNames, track.Title)	
		mp3URLTracks = append(mp3URLTracks , "" )
	}
	for i, trackName := range trackNames {
		fmt.Printf("%d: %s\n", i+1, trackName)
	}
	fmt.Printf("Please select a song (Number): ")
	fmt.Scanf("%d\n", &num)
	if num > len(mp3URLTracks) {
		return
	}
	fmt.Printf("Please Wait...\n")

	url = mp3URLTracks[num-1]	+ "?client_id=" + clientID
	resp, _ = http.Get(url)
	body, _ = ioutil.ReadAll(resp.Body)
	var stream player.Stream
	json.Unmarshal([]byte(string(body)), &stream)
	mp3URLTracks[num-1] = stream.URL
	
	resp,_ = http.Get(mp3URLTracks [num-1])
	rs := httprs.NewHttpReadSeeker(resp)
	defer rs.Close()

	streamer, format, err := mp3.Decode(rs)
	if err != nil {
		return
	}
	defer streamer.Close()

	player.Speaker(streamer,format)	
}




