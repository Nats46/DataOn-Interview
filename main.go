package main

import (
	"bufio"
	"database/sql"
	answer1 "dataon_interview/Question1"
	answer3 "dataon_interview/Question3"
	answer4 "dataon_interview/Question4"
	answer5 "dataon_interview/Question5"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please insert the question number (1, 3, 4, 5, or 0 to exit): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		questionNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		switch questionNumber {
		case 1:
			//please run this first go get -u github.com/denisenkom/go-mssqldb
			connectionString := "server=localhost;database=DataOn;trusted_connection=yes"
			db, err := sql.Open("sqlserver", connectionString)
			if err != nil {
				log.Fatal("Failed to connect to the database:", err)
			}
			defer db.Close()
			fmt.Println("You chose question 1.")
			fmt.Print("please input the option (read/create/migrate/exit): ")
			option, _ := reader.ReadString('\n')
			option = strings.TrimSpace(option)
			switch option {
			case "migrate":
				result := answer1.Migrate(db)
				if !result {
					fmt.Print("fail")
				} else {
					fmt.Print("success crate")
				}
			case "read":
				answer1.Read(db)
			case "create":
				
				fmt.Print("please input the option (auto/jenis/nama/kualitas/exit): ")
				options, _ := reader.ReadString('\n')
				options = strings.TrimSpace(options)
				switch options{
				case "auto":
					result:=answer1.Create(db)
					if !result{
						fmt.Println("fail create data")
					}else{
						fmt.Println("Success create data")
					}
				case "jenis":
					fmt.Print("please input jenis kain: ")
					jenis, _ := reader.ReadString('\n')
					result:=answer1.InsertJenisKain(db,jenis)
					if !result{
						fmt.Println("fail create data")
					}else{
						fmt.Println("Success create data")
					}
				case "nama":
					fmt.Print("please input nama kain: ")
					nama, _ := reader.ReadString('\n')
					fmt.Print("please input id jenis kain: ")
					jenis, _ := reader.ReadString('\n')
					jeniss,_ := strconv.Atoi(jenis)
					result:=answer1.InsertNamaKain(db,nama,jeniss)
					if !result{
						fmt.Println("fail create data")
					}else{
						fmt.Println("Success create data")
					}
				case "kualitas":
					fmt.Print("please input kualitas kain: ")
					kualitas, _ := reader.ReadString('\n')
					kualitass,_:= strconv.Atoi(kualitas)
					fmt.Print("please input nama kualitas kain: ")
					namakualitas, _ := reader.ReadString('\n')
					fmt.Print("please input id harga kain: ")
					harga, _ := reader.ReadString('\n')
					fmt.Print("please input id nama kain: ")
					nama, _ := reader.ReadString('\n')
					namaa,_ := strconv.Atoi(nama)
					result:=answer1.InsertKualitasKain(db,kualitass,namakualitas,harga,namaa)
					if !result{
						fmt.Println("Fail create data")
					}else{
						fmt.Println("Success create data")
					}
				case "exit":
					fmt.Println("exiting...")
					return
				}			
			case "exit":
				print("Exiting...")
				return
			}
		case 3:
			fmt.Println("You chose question 3.")
			answer3.EmailString()
		case 4:
			fmt.Println("You chose question 4.")
			answer4.SplitInput()
		case 5:
			fmt.Println("You chose question 5.")
			answer5.LenString()
		case 0:
			fmt.Println("Exiting...")
			
		default:
			fmt.Println("Invalid question number. Please enter 1, 3, 4, 5, or 0 to exit.")
		}
	}
}
