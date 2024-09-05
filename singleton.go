package chaos

var singleton = New("")

func Set(chaos *Chaos) {
	if chaos == nil {
		return
	}
	singleton = chaos
}

func Fix() {
	singleton.Fix()
}
