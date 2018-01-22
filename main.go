package url2str

// author : https://github.com/jf17
// микромодуль. ходит по ссылке и выкачивает сайт , превращая его в строку и отдает строку обратно.

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // считываем байты

	result := string(body) // кастим байты в строку

	return result

}
