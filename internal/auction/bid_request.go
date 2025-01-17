package auction

type BidRequest struct {
    ID   string   `json:"id"`
    Imp  []Imp    `json:"imp"`
    Site *Site    `json:"site,omitempty"`
    Device *Device `json:"device,omitempty"`
    User *User    `json:"user,omitempty"`
    At   int      `json:"at"`
    TMax int      `json:"tmax"`
    Cur  []string `json:"cur"`
}

type Imp struct {
    ID       string  `json:"id"`
    Banner   *Banner `json:"banner,omitempty"`
    BidFloor float64 `json:"bidfloor"`
}

type Banner struct {
    W int `json:"w"`
    H int `json:"h"`
}

type Site struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Domain string `json:"domain"`
}

type Device struct {
    UA string `json:"ua"`
    IP string `json:"ip"`
}

type User struct {
    ID string `json:"id"`
}
