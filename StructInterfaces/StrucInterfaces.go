package main

import (
	"fmt"
)

type company interface {
	Show()
}

type employee struct {
	Name, EmpId, Gender, Department, Role string
}

type project struct {
	Title       string
	Leader      employee
	TeamMembers []string
}

func (emp employee) show() {
	fmt.Println("Name :", emp.Name)
	fmt.Println("Employee ID : ", emp.EmpId, "\nGender : ", emp.Gender, "\nDepartment : ", emp.Department, "\nRole : ", emp.Role)
}

func (proj project) show() {
	fmt.Println("Title : ", proj.Title, "\nLeader : ")
	proj.Leader.show()
}

func main() {
	emp1 := employee{
		Name:       "Sachin Tendulkar",
		EmpId:      "emp110",
		Gender:     "Male",
		Department: "Product Development",
		Role:       "Software Developer"}

	proj1 := project{
		Title:       "ProjectX",
		Leader:      emp1,
		TeamMembers: []string{"John Doe", "Nora Bing", "Steven Smith"}}

	proj1.show()

}
