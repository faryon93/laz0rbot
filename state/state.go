package state

// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type state struct {
	Aliases map[int]string `json:"aliases"`

	Ikr struct {
		Attendees IkrAttendeeList `json:"attendees"`
	} `json:"ikr"`
}


// ----------------------------------------------------------------------------------
//  global variables
// ----------------------------------------------------------------------------------

var State state


// ----------------------------------------------------------------------------------
//  initializer
// ----------------------------------------------------------------------------------

func init() {
	State.Aliases = make(map[int]string)
	State.Ikr.Attendees = make(IkrAttendeeList, 0)
}


// ----------------------------------------------------------------------------------
//  member functions
// ----------------------------------------------------------------------------------

func (this *state) Save() (error) {
	// TODO: implement marshalling to json file
}

func (this *state) AddIkrAttendee(id int, subscriptionType string) {
	this.Ikr.Attendees = append(this.Ikr.Attendees, IkrAttendee{
		TelegramId: id,
		SubscriptionType: subscriptionType,
	})
}

