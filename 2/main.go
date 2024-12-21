package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var Pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scanln(n)
	for i := 0; i < *n; i++ {
		fmt.Scanln(&pustaka[i].id)
		fmt.Scanln(&pustaka[i].judul)
		fmt.Scanln(&pustaka[i].penulis)
		fmt.Scanln(&pustaka[i].penerbit)
		fmt.Scanln(&pustaka[i].eksemplar)
		fmt.Scanln(&pustaka[i].tahun)
		fmt.Scanln(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	maxRating := -1
	idxTerfavorit := -1
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			idxTerfavorit = i
		}
	}
	if idxTerfavorit != -1 {
		fmt.Println(pustaka[idxTerfavorit].judul, pustaka[idxTerfavorit].penulis, pustaka[idxTerfavorit].penerbit, pustaka[idxTerfavorit].tahun)
	}
}

func UrutBuku(pustaka *DaftarBuku, n *int) {
	for i := 1; i < *n; i++ {
		temp := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			fmt.Println(pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			found = true
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	DaftarkanBuku(&Pustaka, &nPustaka)
	CetakTerfavorit(Pustaka, nPustaka)
	UrutBuku(&Pustaka, &nPustaka)
	Cetak5Terbaru(Pustaka, nPustaka)
	var ratingCari int
	fmt.Scanln(&ratingCari)
	CariBuku(Pustaka, nPustaka, ratingCari)
}
