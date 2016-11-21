package state

// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type IkrAttendeeList []IkrAttendee

type IkrAttendee struct {
	TelegramId int
	SubscriptionType string
}


// ----------------------------------------------------------------------------------
//  initializer
// ----------------------------------------------------------------------------------

func (this IkrAttendeeList) Contains(needle int) (bool) {
	for _, haystack := range this {
		if needle == haystack.TelegramId {
			return true
		}
	}

	return false
}