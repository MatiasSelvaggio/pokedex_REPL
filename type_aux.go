package main

type Pokemon struct {
	Weight         int
	Height         int
	BaseExperience int
	Id             int
	Name           string
	Count          int
	Stats          []struct {
		BaseStat int
		Effort   int
		Stat     struct {
			Name string
			Url  string
		}
	}
	Types []struct {
		Slot int
		Type struct {
			Name string
			Url  string
		}
	}
}
