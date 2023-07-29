package sockets

type Clients interface {
	AddClient(client Client)
	MessageClientsForRaceId(raceId int)
	MessageClientsForRaceIdStarting(raceId, remainingTime int, hoot bool)
}

type clients struct {
	clients map[int][]Client
}

func NewClients() *clients {
	return &clients{
		clients: map[int][]Client{},
	}
}

func (c *clients) AddClient(client Client) {
	raceId := client.getRace().GetId()
	clientForRace, ok := c.clients[raceId]
	if !ok {
		c.clients[raceId] = []Client{client}
	}
	c.clients[raceId] = append(clientForRace, client)
	if client.getRace().GetStartSequence().IsComplete() {
		message := message{
			Phase: "started",
			Time:  client.getRace().GetElapsedTime(),
		}
		client.getConn().WriteJSON(message)
		return
	}
	message := message{
		Phase: "starting",
		Time:  client.getRace().GetStartSequence().GetRemainingTime(),
		Hoot:  false,
	}
	client.getConn().WriteJSON(message)
}

type message struct {
	Phase string `json:"phase"`
	Time  int    `json:"time"`
	Hoot  bool   `json:"hoot"`
}

func (c *clients) MessageClientsForRaceId(raceId int) {
	raceClients := c.getClientsForRaceId(raceId)
	for _, c := range raceClients {
		message := message{
			Phase: "started",
			Time:  c.getRace().GetElapsedTime(),
		}
		c.getConn().WriteJSON(message)
	}
}

func (c *clients) MessageClientsForRaceIdStarting(raceId, remainingTime int, hoot bool) {
	raceClients := c.getClientsForRaceId(raceId)
	for _, c := range raceClients {
		message := message{
			Phase: "starting",
			Time:  remainingTime,
			Hoot:  hoot,
		}
		c.getConn().WriteJSON(message)
	}
}

func (c *clients) getClientsForRaceId(raceId int) []Client {
	return c.clients[raceId]
}
