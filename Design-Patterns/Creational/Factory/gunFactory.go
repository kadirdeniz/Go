package main

func NewGun(env string) iGun {
	if env == "ak47" {
		return newAk47()
	} else if env == "musket" {
		return newMusket()
	}
	return nil
}
