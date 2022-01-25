package main

type Musket struct {
	gun
}

func newMusket() iGun {
	return &Musket{
		gun: gun{
			name:  "Musket",
			power: 4,
		},
	}
}
