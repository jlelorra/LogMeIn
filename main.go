package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Player struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
}

type Deck struct {
    ID        string   `json:"id,omitempty"`
}

type Game struct {
    ID        string   `json:"id,omitempty"`
	DeckID    string   `json:"id"`
	PlayerID  string   `json:"id"`
}

 type Card struct {
	Suit  string `json:"suit"`
	Number int `json:"number"`
 }

//deck
var cards = []Card{
  {Suit: "Spades", Number: "2"},{Suit: "Spades", Number: "3"},{Suit: "Spades", Number: "4"},{Suit: "Spades", Number: "5"},{Suit: "Spades", Number: "6"},{Suit: "Spades", Number: "7"},{Suit: "Spades", Number: "8"},{Suit: "Spades", Number: "9"},{Suit: "Spades", Number: "10"},{Suit: "Spades", Number: "Jack"},{Suit: "Spades", Number: "Queen"},{Suit: "Spades", Number: "King"},{Suit: "Spades", Number: "Ace"},
  {Suit: "Hearts", Number: "2"},{Suit: "Hearts", Number: "3"},{Suit: "Hearts", Number: "4"},{Suit: "Hearts", Number: "5"},{Suit: "Hearts", Number: "6"},{Suit: "Hearts", Number: "7"},{Suit: "Hearts", Number: "8"},{Suit: "Hearts", Number: "9"},{Suit: "Hearts", Number: "10"},{Suit: "Hearts", Number: "Jack"},{Suit: "Hearts", Number: "Queen"},{Suit: "Hearts", Number: "King"},{Suit: "Hearts", Number: "Ace"},
  {Suit: "Diamonds", Number: "2"},{Suit: "Diamonds", Number: "3"},{Suit: "Diamonds", Number: "4"},{Suit: "Diamonds", Number: "5"},{Suit: "Diamonds", Number: "6"},{Suit: "Diamonds", Number: "7"},{Suit: "Diamonds", Number: "8"},{Suit: "Diamonds", Number: "9"},{Suit: "Diamonds", Number: "10"},{Suit: "Diamonds", Number: "Jack"},{Suit: "Diamonds", Number: "Queen"},{Suit: "Diamonds", Number: "King"},{Suit: "Diamonds", Number: "Ace"},
  {Suit: "Clubs", Number: "2"},{Suit: "Clubs", Number: "3"},{Suit: "Clubs", Number: "4"},{Suit: "Clubs", Number: "5"},{Suit: "Clubs", Number: "6"},{Suit: "Clubs", Number: "7"},{Suit: "Clubs", Number: "8"},{Suit: "Clubs", Number: "9"},{Suit: "Clubs", Number: "10"},{Suit: "Clubs", Number: "Jack"},{Suit: "Clubs", Number: "Queen"},{Suit: "Clubs", Number: "King"},{Suit: "Clubs", Number: "Ace"}
}

type deck []Card 
type players []Player


//shuffle
func Shuffle(slc []Card) {
    for i := 1; i < len(slc); i++ {
        r := rand.Intn(i + 1)
        if i != r {
            slc[r], slc[i] = slc[i], slc[r]
        }
    }
}

// Deal a specified amount of cards
func Deal(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < n; i++ {
		params := mux.Vars(r)
		var player Player
        _ = json.NewDecoder(r.Body).Decode(&player)
		for _, item := range player	{
            if item.ID == player_id
                player = append(player, d[i])
                json.NewEncoder(w).Encode(player)
	}
}

// REST API - PLAYER
func GetPlayers(w http.ResponseWriter, r *http.Request) {    
    json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {    
    json.NewEncoder(w).Encode(player)
}

func GetPlayerCard(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range player	{
    if item.ID == params["id"]
	}
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var player Player
    _ = json.NewDecoder(r.Body).Decode(&player)
    player.ID = params["id"]
    players = append(players, player)
    json.NewEncoder(w).Encode(players)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
   params := mux.Vars(r)
    for index, item := range player {
        if item.ID == params["id"] {
            players = append(players[:index])
            break
        }
        json.NewEncoder(w).Encode(players)
    }
} 

//REST API - GAME
func CreateGame(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var game Game
    _ = json.NewDecoder(r.Body).Decode(&game)
    game.ID = params["id"]
    game = append(game, deck)
    json.NewEncoder(w).Encode(game)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
   params := mux.Vars(r)
    for index, item := range game {
        if item.ID == params["id"] {
            game = append(game[:index])
            break
        }
        json.NewEncoder(w).Encode(game)
    }
}

// REST API - DECK
func CreateDeck(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var deck Deck
    _ = json.NewDecoder(r.Body).Decode(&deck)
    deck.ID = params["id"]
    deck = append(deck, card)
    json.NewEncoder(w).Encode(deck)
}

func GetDeck(w http.ResponseWriter, r *http.Request) {    
    json.NewEncoder(w).Encode(deck)
}

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/game/{id}", CreateGame).Methods("POST")
    players = append(players, Player{ID: "1", Firstname: "John", Lastname: "Doe"})
    players = append(players, Player{ID: "2", Firstname: "Bill", Lastname: "Gates"})
    players = append(players, Player{ID: "3", Firstname: "Steve", Lastname: "Jobs"})
    router.HandleFunc("/player/{id}", CreatePlayer).Methods("POST")
    router.HandleFunc("/player", GetPlayers).Methods("GET")
    router.HandleFunc("/player/{id}", GetPlayer).Methods("GET")
    router.HandleFunc("/deck/{id}", CreateDeck).Methods("POST")
    router.HandleFunc("/deck/{id}", GetDeck).Methods("GET")
    Shuffle(deck)
    router.HandleFunc("/player/{id}", Deal).Methods("POST")
    router.HandleFunc("/playerCards/{id}", GetPlayerCards).Methods("GET")
    router.HandleFunc("/playerCards/{id}", GetPlayerCards).Methods("GET")
    router.HandleFunc("/player/{id}", DeletePlayer).Methods("DELETE")	
    router.HandleFunc("/game/{id}", DeleteGame).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
