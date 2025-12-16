package secret

func Handshake(code uint) []string {
	actions := []string{}
	if code&1 != 0 {
		actions = append(actions, "wink")
	}
	if code&2 != 0 {
		actions = append(actions, "double blink")
	}
	if code&4 != 0 {
		actions = append(actions, "close your eyes")
	}
	if code&8 != 0 {
		actions = append(actions, "jump")
	}
	if code&16 != 0 {
		// Reverse the order
		for i, j := 0, len(actions)-1; i < j; i, j = i+1, j-1 {
			actions[i], actions[j] = actions[j], actions[i]
		}
	}
	return actions
}
