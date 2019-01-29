package location

// A Location in the game
type Location struct {
	Name    string
	ID      int
	Terrain string
}

// GenerateLocations makes the map locations
func GenerateLocations(name string, id int) [10]Location {
	var locs [10]Location
	locationList := [10]string{"Detroit", "The ISS", "Castle Ravenloft", "Isengard", "15th Century Bulgaria", "315 North Grand Ave.", "Sesame Street", "Pound Town", "On A Boat", "Right Behind You"}
	for i := 0; i < 10; i++ {
		locs[i].Name = name
		locs[i].ID = i
		locs[i].Terrain = locationList[i]
	}
	return locs
}
