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
	var userName string
	var num int

	clientID, err := player.GetClientID()
	if err != nil {
		return
	}
	// fmt.Printf("%s\n", clientID) 
	fmt.Printf("Username: ")
	fmt.Scanf("%s\n", &userName) 
	fmt.Printf("---------------------------------\n")


	url := "https://api-v2.soundcloud.com/resolve?url=https://soundcloud.com/" + userName
	url = url + "&client_id="+ clientID
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	var user player.User
	json.Unmarshal([]byte(string(body)), &user)
	userID := user.ID

	url = "https://api-v2.soundcloud.com/users/" + strconv.Itoa(userID) + "/playlists?limit=50"
	url =  url + "&client_id=" + clientID
	resp, _ = http.Get(url)
	body, _ = ioutil.ReadAll(resp.Body)

	var playlists player.PlayLists
	json.Unmarshal([]byte(string(body)), &playlists)

	var playlistIDs[] int
	var playlistNames[] string

	for _, playlist := range playlists.Collection {
		playlistIDs =  append(playlistIDs, playlist.ID)	
		playlistNames = append(playlistNames, playlist.Title)	
	}

	for i, playlistName := range playlistNames {
		fmt.Printf("%d: %s\n", i+1, playlistName)
	}
	fmt.Printf("---------------------------------\n")
	fmt.Printf("Please select a playlist (Number): ")
	fmt.Scanf("%d\n", &num)
	if num > len(playlistIDs) {
		return
	}

	url = "https://api-v2.soundcloud.com/playlists/" + strconv.Itoa(playlistIDs[num-1])
	url =  url + "?client_id=" + clientID
	resp, _ = http.Get(url)
	body, _ = ioutil.ReadAll(resp.Body)


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
		url = "https://api-v2.soundcloud.com/tracks/" + strconv.Itoa(trackID)	
		url =  url +  "?client_id=" + clientID
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
	fmt.Printf("---------------------------------\n")
	fmt.Printf("Please select a song (Number): ")
	fmt.Scanf("%d\n", &num)
	if num > len(mediaURLTracks) {
		return
	}
	
	url = mediaURLTracks[num-1] + "?client_id=" + clientID
	resp, _ = http.Get(url)
	body, _ = ioutil.ReadAll(resp.Body)
	var stream player.Stream
	json.Unmarshal([]byte(string(body)), &stream)
	mp3URLTracks[num-1] = stream.URL
	
	resp,_ = http.Get(mp3URLTracks[num-1])
	rs := httprs.NewHttpReadSeeker(resp)
	defer rs.Close()

	streamer, format, err := mp3.Decode(rs)
	if err != nil {
		return
	}
	defer streamer.Close()

	player.Speaker(streamer,format)	
}




