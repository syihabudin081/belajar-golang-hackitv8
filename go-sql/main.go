package main

import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "123"
	db_name = "db_go_sql"
)
var (
	db *sql.DB
	err error
)

type Employee struct{
	ID int	
	Full_name string
	Email string
	Age int
	Division string
}

func main(){
	pSQLInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)

	db,err = sql.Open("postgres", pSQLInfo)
	if err != nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Successfully connected!")
	// CreateEmployee()
	GetEmployees()
	// updateEmployee()
	// DeleteEmployee()
}

func CreateEmployee(){
	// Create a new employee
var employee = Employee{}

sqlStatement := `INSERT INTO employees (full_name, email, age, division) VALUES ($1, $2, $3, $4) RETURNING *`

err = db.QueryRow(sqlStatement,"Airell Jordan","jordan@gmail.com", 23, "IT").
Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

if err != nil{
	panic(err)
}
fmt.Println("New employee Data: ", employee)

}

func GetEmployees(){
	var result = []Employee{}
	sqlStatement := `SELECT * FROM employees`
	rows, err := db.Query(sqlStatement)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		var employee Employee
		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
		if err != nil{
			panic(err)
		}
		result = append(result, employee)
	}

	fmt.Println("All employees: ", result)
}


func updateEmployee(){
	sqlStatement := `UPDATE employees SET full_name = $2, email = $3, division = $4, age= $5 WHERE id = $1`
	res, err := db.Exec(sqlStatement, 1, "Airell Jordan", "jordan2@gmail.com", "Accounting", 25)
	if err != nil{
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Updated Data Amount:", count)

}

func DeleteEmployee(){
	sqlStatement := `DELETE FROM employees WHERE id = $1`
	res, err := db.Exec(sqlStatement, 1)
	if err != nil{
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil{
		panic(err)
	}
	fmt.Printf("Deleted Data Amount:", count)
}