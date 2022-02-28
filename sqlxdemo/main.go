package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	file, err := os.Create("leave.csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	csvwriter := csv.NewWriter(file)
	leavesArr := [][]string{}
	leavesArr = append(leavesArr, []string{"leaveid", "empid", "supervisorid", "name"})
	for i := 0; i < 1589800; i++ {

		var leave1 ApplyLeave1
		leave1.EmpId = strconv.Itoa(i)
		leave1.LeaveId = i
		leave1.Supervisorid = strconv.Itoa(i)
		leave1.Name = "name"
		row := []string{leave1.EmpId, strconv.Itoa(leave1.LeaveId), leave1.Supervisorid, leave1.Name}
		leavesArr = append(leavesArr, row)
	}

	csvwriter.WriteAll(leavesArr)
	csvwriter.Flush()
	file, err = os.Open("leave.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	var leaves []ApplyLeave1
	err = gocsv.Unmarshal(file, &leaves)
	if err != nil {
		panic(err)
	}

	// //================

	db, err := sqlx.Connect("postgres", "host=localhost user=postgres password=root dbname=demo port=5432 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	v := "TABLENAME"
	_, _ = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", v))
	_, _ = db.Exec(fmt.Sprintf("commit"))
	query := `INSERT INTO TABLENAME(leaveid, empid, supervisorid, name)
			  VALUES(:leaveid, :empid, :supervisorid, :name)`

	s := time.Now()
	div := (1 << 16 / 4)-1
	t := len(leaves) / div
	var wg sync.WaitGroup
	wg.Add(t+1)
	for i := 0; i < len(leaves); i += div {
		end := i + div
		if end > len(leaves) {
			end = len(leaves)
		}
		fmt.Println(i, ",", end)

		go newFunction(&wg, db, query, leaves[i:end])

	}
	wg.Wait()
	e := time.Now()
	d := e.Second() - s.Second()
	fmt.Println(d)

}

func newFunction(wg *sync.WaitGroup, db *sqlx.DB, query string, leaves []ApplyLeave1) {

	_, e := db.NamedExec(query, leaves)
	if e != nil {
		log.Fatalln(e)
	}
	wg.Done()
}

type ApplyLeave1 struct {
	LeaveId      int    `csv:"leaveid"`
	EmpId        string `csv:"empid"`
	Supervisorid string `csv:"supervisorid"`
	Name         string `csv:"name"`
}
