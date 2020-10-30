package postback

// Postback representation of postback into struct
type Postback struct {
	Endpoint	Endpoint   	`json:"endpoint"`
	Data   		[]Data 		`json:"data"`
}

// Endpoint representation of endpoint into struct
type Endpoint struct {
	Method 	string		`json:"method"`
	Url   	string 		`json:"url"`
}

// Data representation of data into struct
type Data struct {
	Key		string  `json:"key"`
	Value   string 	`json:"value"`
}

