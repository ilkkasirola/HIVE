package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
	"regexp"
)
type Passenger struct {
    BookingID string
    Name      string
    Seat      string
}

func ParsePassengerLine(line string) (Passenger, bool) {
	
	
	
	parts := strings.Split(line, "|")
	passenger := Passenger{}
	if len(parts) != 3 {
		return passenger, false
	}
	bookId := strings.TrimSpace(parts[0])
	name := strings.TrimSpace(parts[1])
	seat := strings.TrimSpace(parts[2])

	if bookId == "" || name == "" || seat == ""{
		return passenger, false
	}

	passenger.BookingID = bookId
	passenger.Name = name 
	passenger.Seat = seat 

	return passenger, true
}



func ParsePassengerManifest(block string) (map[string]Passenger, []int){

			lines := strings.Split(block, "\n")
			invalidLines := []int{}
			PassengerMap := make(map[string]Passenger)
			for i:=0 ; i < len(lines) ; i++{
				
				if strings.TrimSpace(lines[i]) == ""{
					continue
				}
				p, flag := ParsePassengerLine(lines[i])
				
				if !flag || !IsValidSeat(p.Seat) {
					invalidLines = append(invalidLines, i+1)
					continue
				}
				
				keyId := p.BookingID
				PassengerMap[keyId] = p
			} 

			return PassengerMap, invalidLines

}

func IsValidSeat(seat string)bool{

		seatreg := regexp.MustCompile("[A-F]")

		seatLetter := seatreg.FindAllString(seat, -1)

		if len(seatLetter) != 1{
			return false
		}

		vals := seatreg.Split(seat, -1)
		seatNumber , err := strconv.Atoi(vals[0]) 
			if err != nil || seatNumber < 1{
				return false
			}
		return true

}

//go enmaxint := Math.MaxInt

func main(){

	fmt.Println(ParsePassengerManifest(os.Args[1]))
	
}