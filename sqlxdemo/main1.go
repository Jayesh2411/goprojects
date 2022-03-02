package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	//
	// file1, err := os.Create("leave.csv")
	//
	// if err != nil {
	// 	log.Fatalln("failed to open file", err)
	// }
	//
	// csvwriter := csv.NewWriter(file1)
	// leavesArr := [][]string{}
	// leavesArr = append(leavesArr, []string{"leaveid", "empid", "supervisorid", "name"})
	// for i := 0; i < 10000000; i++ {
	//
	// 	var leave1 ApplyLeave1
	// 	leave1.EmpId = strconv.Itoa(i)
	// 	leave1.LeaveId = i
	// 	leave1.Supervisorid = strconv.Itoa(i)
	// 	leave1.Name = "name"
	// 	row := []string{leave1.EmpId, strconv.Itoa(leave1.LeaveId), leave1.Supervisorid, leave1.Name}
	// 	leavesArr = append(leavesArr, row)
	// }
	//
	// csvwriter.WriteAll(leavesArr)
	// csvwriter.Flush()
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	// //================
	scan := []string{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		s := sc.Text() // GET the line string
		scan = append(scan, s)

	}

	scan = scan[1:]
	db, err := sql.Open("postgres", "host=localhost user=postgres password=root dbname=demo port=5432 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	v := "company"
	_, _ = db.Exec(fmt.Sprintf("delete from %s;", v))

	div := (1 << 16 / 5) - 1
	// t := len(scan) / div
	start := time.Now().Second()
	for i := 0; i < len(scan); i += div {
		end := i + div
		if end > len(scan) {
			end = len(scan)
		}
		scan1 := scan[i:end]
		insert(scan1, db)

	}
	end := time.Now().Second()
	fmt.Println(end - start)
	//
	// query := `INSERT INTO TABLENAME(leaveid, empid, supervisorid, name)
	// 		  VALUES(:leaveid, :empid, :supervisorid, :name)`

	// s := time.Now()
	// div := (1 << 16 / 4) - 1
	// t := len(leaves) / div
	// var wg sync.WaitGroup
	// wg.Add(t + 1)
	// for i := 0; i < len(leaves); i += div {
	// 	end := i + div
	// 	if end > len(leaves) {
	// 		end = len(leaves)
	// 	}
	// 	fmt.Println(i, ",", end)
	//
	// 	go newFunction(&wg, db, query, leaves[i:end])
	//
	// }
	// wg.Wait()
	// e := time.Now()
	// d := e.Second() - s.Second()
	// fmt.Println(d)

}

func insert(scan1 []string, db *sql.DB) {
	valueStrings := make([]string, 0, len(scan1))
	valueArgs := make([]interface{}, 0, len(scan1)*5)
	i := 0
	for _, post := range scan1 {
		leave := strings.Split(string(post), ",")
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d,$%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, leave[0])
		valueArgs = append(valueArgs, leave[1])
		valueArgs = append(valueArgs, leave[2])
		valueArgs = append(valueArgs, leave[3])
		valueArgs = append(valueArgs, leave[4])
		i++
	}

	stmt := fmt.Sprintf("INSERT INTO company(name,other_employer_name_required,organization_type,industry,category) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		log.Fatalln(err)
	}
}

// func newFunction(wg *sync.WaitGroup, db *sql.DB, query string, leaves []ApplyLeave1) {
//
// 	_, e := db.ExecContext(context.Background(), query, leaves)
// 	if e != nil {
// 		_, _ = db.Exec(fmt.Sprintf("rollback"))
// 		log.Fatalln(e)
// 	}
// 	wg.Done()
// }

// type ApplyLeave1 struct {
// 	LeaveId      int    `csv:"leaveid"`
// 	EmpId        string `csv:"empid"`
// 	Supervisorid string `csv:"supervisorid"`
// 	Name         string `csv:"name"`
// }
