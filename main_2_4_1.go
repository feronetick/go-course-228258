package main

type Terminator struct {
	On    bool
	Ammo  int
	Power int
}

func (t *Terminator) Shoot() bool {
	if t.On && t.Ammo > 0 {
		t.Ammo--
		return true
	}

	return false
}

func (t *Terminator) RideBike() bool {
	if t.On && t.Power > 0 {
		t.Power--
		return true
	}

	return false
}
