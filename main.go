package main

import (
	"fmt"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/seregant/decision-making-go/database"
)

func main() {
	database.CreateTable()
	fmt.Println("Selamat datang di konsultasi usaha!")
	fmt.Println("Isi kriteria berikut sesuai dengan keadaan anda !")

	var Pertanyaan []database.Pertanyaan
	var Decision []database.Decisions
	var PengetahuanKond []database.Pengetahuan

	var db = database.DbConnect()
	defer db.Close()

	db.Find(&Pertanyaan)
	db.Find(&Decision)
	db.Find(&PengetahuanKond)

	finalData := []int{0, 0, 0, 0, 0, 0}

	var choices []int
	var pvResult []database.Pengetahuan
	var pvResult2 []database.Pengetahuan_value

	for _, dataPertanyaan := range Pertanyaan {
		number := 1
		fmt.Println("---> " + dataPertanyaan.Pertanyaan)
		fmt.Println("Pilihan : ")
		db.Limit(2).Table("pengetahuan_value").Select("pengetahuan.pengetahuan_nama").Joins("INNER JOIN pengetahuan ON pengetahuan.pengetahuan_id = pengetahuan_value.pengetahuan_id").Where("pengetahuan_value.pertanyaan_id = ?", dataPertanyaan.ID).Find(&pvResult)

		for _, pengetahuan := range pvResult {
			numberStr := strconv.Itoa(number)
			fmt.Println(numberStr + "." + pengetahuan.Nama)
			number++
		}
		var input int
		fmt.Print("Pilihan anda : ")
		fmt.Scanln(&input)
		if input == 1 {
			choices = append(choices, 1)
			choices = append(choices, 0)
		} else {
			choices = append(choices, 0)
			choices = append(choices, 1)
		}
		fmt.Println()
	}

	for _, dataKeputusan := range Decision {
		db.Where("decision_id = ?", dataKeputusan.ID_Decision).Find(&pvResult2)
		loopScore := 0
		arrIndex := 0
		for _, hasil := range pvResult2 {
			if hasil.Yes == choices[arrIndex] {
				loopScore++
			}
			arrIndex++
		}

		if dataKeputusan.ID_Decision == 1 {
			finalData[0] += loopScore
		} else if dataKeputusan.ID_Decision == 2 {
			finalData[1] += loopScore
		} else if dataKeputusan.ID_Decision == 3 {
			finalData[2] += loopScore
		} else if dataKeputusan.ID_Decision == 4 {
			finalData[3] += loopScore
		} else if dataKeputusan.ID_Decision == 5 {
			finalData[4] += loopScore
		} else {
			finalData[5] += loopScore
		}
	}

	var finalDecsion []int

	start := 1
	finish := 1
	m := 0
	for _, e := range finalData {
		if e > m {
			m = e
			finish = start
			finalDecsion = nil
			finalDecsion = append(finalDecsion, finish)
		} else if e == m {
			finish = start
			finalDecsion = append(finalDecsion, finish)
		}
		start++
	}
	fmt.Println(finalDecsion)

	fmt.Println("Rekomendasi jenis usaha : ")
	for _, id := range finalDecsion {
		for _, dataOutput := range Decision {
			if dataOutput.ID_Decision == id {
				fmt.Print("- ")
				fmt.Println(dataOutput.Decision)
			}
		}
	}

}
