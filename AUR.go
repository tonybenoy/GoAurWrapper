package AUR

import (
	"io/ioutil"
	"net/http"
)

func AurSearch(pkg string, queryby string) string {
	if queryby == "" {
		queryby = "name-desc"
	}
	list := []string{"name", "name-desc", "maintainer", "depends", "makedepends", "optdepends", "checkdepends"}
	for j, i := range list {
		if i == queryby {
			break
		} else {
			if j == len(list)-1 {
				panic("Unsupported Query")
			}
		}
	}
	url := "https://aur.archlinux.org/rpc.php/rpc/?v=5&type=search&by=" + queryby + "&arg=" + pkg
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
func AurInfo(args ...string) string {
	url := "https://aur.archlinux.org/rpc.php/rpc/?v=5&type=info"
	for _, i := range args {
		url += "&arg[]=" + i
	}
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
