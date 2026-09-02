package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)


func AddExpense(db *sql.DB){
	reader := bufio.NewReader(os.Stdin)


	fmt.Print("Enter the name of the expense: ")
	text,err:= reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	name := strings.TrimSpace(text)

	var amount float64
	fmt.Print("Enter the amount of the expense: ")
	text,err = reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	amount,err = strconv.ParseFloat(strings.TrimSpace(text),64)
	if err != nil {
		fmt.Println("Invalid input: please enter a valid number")
		return
	}

	fmt.Print("Enter the category of the expense: ")
	text,err = reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	category := strings.TrimSpace(text)
	
	
	fmt.Print("Enter the date of expense(YYYY-MM-DD): ")
	text,err =reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	date := strings.TrimSpace(text)
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	

	_,err = db.Exec(`insert into expenses(name,amount,category,date) values(?,?,?,?)`, name, amount, category, date)
	if err != nil{
		fmt.Println("Error adding expense",err)
		return
	}

	fmt.Println("Expense added successfully")

}

func ViewAllExpenses(db *sql.DB){
	rows,err:=db.Query("select * from expenses")
	if err!=nil {
		fmt.Println("Error fetching expenses",err)
		return
	}
	defer rows.Close()

	fmt.Println("ID | Name | Amount | Category | Date")
	fmt.Println("----------------------------------------")
	
	for rows.Next(){
		var id int
		var name,category, date string
		var amount float64

		err:=rows.Scan(&id,&name,&amount,&category,&date)
		if err!=nil {
			fmt.Println("Error scanning row",err)
			return
		}
		fmt.Printf("%d | %s | %.2f | %s | %s\n",id,name,amount,category,date)
	}
}

func ViewExpensesByDate(db *sql.DB){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the startDate: ")
	text,err:= reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	startDate := strings.TrimSpace(text)

	fmt.Print("Enter the endDate: ")
	text,err = reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	endDate := strings.TrimSpace(text)

	rows,err:=db.Query("select * from expenses where date between ? and ?",startDate,endDate)
	if err!=nil {
		fmt.Println("Error fetching expenses",err)
		return
	}
	defer rows.Close()

	fmt.Println("ID | Name | Amount | Category | Date")
	fmt.Println("----------------------------------------")
	
	for rows.Next(){
		var id int
		var name,category, date string
		var amount float64

		err:=rows.Scan(&id,&name,&amount,&category,&date)
		if err!=nil {
			fmt.Println("Error scanning row",err)
			return
		}
		fmt.Printf("%d | %s | %.2f | %s | %s\n",id,name,amount,category,date)
	}
	
}

func ViewExpensesByCategory(db *sql.DB){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the category: ")
	text,err:= reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	category := strings.TrimSpace(strings.ToLower(text))

	rows,err:=db.Query("select * from expenses where category = ?",category)

	if err!=nil {
		fmt.Println("Error fetching expenses",err)
		return
	}
	defer rows.Close()

	fmt.Println("ID | Name | Amount | Category | Date")
	fmt.Println("----------------------------------------")
	
	for rows.Next(){
		var id int
		var name,category, date string
		var amount float64

		err:=rows.Scan(&id,&name,&amount,&category,&date)
		if err!=nil {
			fmt.Println("Error scanning row",err)
			return
		}
		fmt.Printf("%d | %s | %.2f | %s | %s\n",id,name,amount,category,date)
	}
	
}

func ViewExpenses(db *sql.DB){

	fmt.Println("1. View All Expenses")
	fmt.Println("2. View Expenses by Category")
	fmt.Println("3. View Expenses by Date")
	fmt.Println("4. Back to main menu")
	
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	choice := strings.TrimSpace(text)
	choiceInt, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("Error converting choice to integer")
		return
	}

	switch choiceInt{
		case 1:
			ViewAllExpenses(db)
		case 2:
			ViewExpensesByCategory(db)
		case 3:
			ViewExpensesByDate(db)
		case 4:
			main()
		default:
			fmt.Println("Invalid choice")
			return
	}
	
}

func DeleteExpense(db *sql.DB){
	reader:= bufio.NewReader(os.Stdin)
	text,err:=reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}

	id,err := strconv.Atoi(strings.TrimSpace(text))
	if err!=nil {
		fmt.Println("Error converting id to integer")
		return
	}

	_,err = db.Exec("delete from expenses where id = ? ",id)
	if err!=nil {
		fmt.Println("Error deleting expense",err)
		return
	}

	fmt.Println("Expense deleted successfully")

	
}




func main(){
	db,err:=sql.Open("sqlite","expenses.db")
	if err!=nil {
		fmt.Println("Error opening database",err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		log.Fatal("Error connecting to database",err)
	}
	// fmt.Println("Database connected successfully")
	

	schema,err:= os.ReadFile("schema.sql")
	if err!= nil {
		fmt.Println("Error reading schema",err)
		return
	}
	
	_,err=db.Exec(string(schema))
	if err != nil {
		log.Fatal("Error executing schema",err)
		return
	}
	// fmt.Println("Schema executed successfully")
	for{
		
		fmt.Println("Welcome to Kartike's Expense Tracker")
		fmt.Println("1. Add an Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Delete an Expense")
		fmt.Println("4. Exit")
	

	fmt.Print("Choose any of the above options: ")
	reader:= bufio.NewReader(os.Stdin)
	text,err:= reader.ReadString('\n')
	if err!=nil {
		fmt.Println("Error reading input")
		return
	}
	
	choice, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		fmt.Println("Invalid input: please enter a valid number")
		return
	}

	
	switch choice {
	case 1:

		AddExpense(db)
	case 2:
		fmt.Println("You chose 2")
		ViewExpenses(db)
	case 3:
		fmt.Println("You chose 3")
		DeleteExpense(db)
	case 4:
		fmt.Println("You chose 4")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
	
	
}
}