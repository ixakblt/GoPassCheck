/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
 ElitHatz.com
------------------------------------
*/

package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, _ := os.Create("ixHayabusa.txt")      //ixhayabusa adında boş bir txt oluşturduk
	dosya, _ := ioutil.ReadFile("alice.txt") //içerisinde listemizin bulunduğu alice.txt listesini okuduk
	//Satır satır döngüye aldık
	for _, i := range strings.Split(string(dosya), "\n") {
		//eğer dönen veri varsa işleme al dedik
		if i != "" {
			a := strings.Split(i, ":")[1] // gelen satırı : ile ayırdık ve :'den sonrasını aldık
			if checkLen(a) && isBig(a) && isKck(a) && isNum(a) {
				f.WriteString(i) //yukarda  belirttiğimiz şartlara uygunsa ixhayabusa dosyasına yazdırdık
			}
		}
	}
}

//isNum fonksiyonu ile içerisinde sayı var mı kontrol ettik
func isNum(veri string) bool {
	num := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	kontrol := false
	//Gelen veri değişkenindeki her harfi alıp teker teker num dizisi ile kıyaslamakta
	//eğer içerisinde varsa kontrol değişkeni true dönmekte yoksa false kalmakta
	for _, i := range veri {
		for _, j := range num {
			if string(i) == j {
				kontrol = true
			}
		}
	}
	return kontrol
}

//Küçük harf kontrolü
func isKck(veri string) bool {
	kck := [...]string{"a", "b", "c", "ç", "d", "e", "f", "g", "ğ", "h", "ı", "i", "j", "k", "l", "m", "n", "o", "ö", "p", "q", "r", "s", "ş", "t", "u", "ü", "v", "w", "x", "y", "z"}
	kontrol := false
	for _, i := range veri {
		for _, j := range kck {
			if string(i) == j {
				kontrol = true
			}
		}
	}
	return kontrol
}

//Büyük harf kontrolü
func isBig(veri string) bool {
	kontrol := false
	byk := [...]string{"A", "B", "C", "Ç", "D", "E", "F", "G", "Ğ", "H", "I", "İ", "J", "K", "L", "M", "N", "O", "Ö", "P", "Q", "R", "S", "Ş", "T", "U", "Ü", "V", "W", "X", "Y", "Z"}
	for _, i := range veri {
		for _, j := range byk {
			if string(i) == j {
				kontrol = true
			}
		}
	}
	return kontrol
}

//uzunluk kontrolü
func checkLen(veri string) bool {
	kontrol := false
	if len(veri) > 7 {
		kontrol = true
	}
	return kontrol
}
