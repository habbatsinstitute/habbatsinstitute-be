package packet

import "time"

type Message struct {
	Room int		`json:"room"`
	Role string		`json:"role"`
	From int		`json:"from"`
	To   int		`json:"to"`
	Text string		`json:"text"`
	Time time.Time	`json:"time"`
}