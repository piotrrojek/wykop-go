# wykop-go

Przykładowe zastosowanie (pobranie 5 najlepszych gorących z ostanich 6 godzin):
```go
package main

import (
	"fmt"

	"github.com/piotrrojek/wykop-go"
)

const (
	appsecret = "twójSekret"
	appkey    = "Twójappkey"
)

func main() {
	w := wykop.NewSession(appkey, appsecret)
	w.ClearOutput = true

	entries, err := w.GetStreamHot(1, 6)
	if err != nil {
		fmt.Println("Wystąpił błąd:", err)
	}

	for i, x := range entries[0:4] {
		fmt.Printf("%v:\n", i+1)
		fmt.Println("Autor:", x.Author)
		fmt.Println("Data:", x.Date)
		fmt.Println("URL:", x.URL)
		fmt.Println("Treść:", x.Body)
		fmt.Println("Ilość komentarzy:", x.CommentCount)
	}

}

```

i output:
```
1:
Autor: repiv
Data: 2015-06-04 18:57:05
URL: http://www.wykop.pl/wpis/12986589/total-biscuit-przemowil-w-sprawie-afery-ze-zbyt-bi/
Treść: Total Biscuit przemówił w sprawie afery ze zbyt białym Wiedźminem 

https://twitter.com/Totalbiscuit/status/606497905948565504

#wiedzmin3 #gry #bekazlewactwa
Ilość komentarzy: 20
2:
Autor: TOKSYDROM
Data: 2015-06-04 18:19:37

...
```