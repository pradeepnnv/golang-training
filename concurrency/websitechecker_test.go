package concurrency

func mockWebSiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebSites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://apple.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"https://google.com":      true,
		"https://apple.com":       true,
		"waat://furhurterwe.geds": false,
	}
	got := CheckWebSites(mockWebSiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
