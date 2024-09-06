package chaos

var singleton = New("")

func Set(chaos *Chaos) {
	if chaos == nil {
		return
	}
	singleton = chaos
}

// Fix freezes the chaos.
// When chaos is fixed, the values generated are always the same.
// That means the same method will always return the same value.
func Fix() {
	singleton.Fix()
}

// Unfix un-freezes the chaos.
// This results in new values being generated for each method call.
func Unfix() {
	singleton.Unfix()
}
