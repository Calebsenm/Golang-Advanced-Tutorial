package main

import "fmt"

func main() {
	college1 := createCollege1()
	college2 := createCollege2()

	totalProfessors := 0
	dertmentArray := []string{"computerscience", "mechanical", "civil", "civil"}

	for _, departmentName := range dertmentArray {
		d := college1.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college is %d\n", totalProfessors)

	// reset the professorcount
	totalProfessors = 0
	for _, departmentName := range dertmentArray {
		d := college2.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *college {
	college := &college{}
	college.addDepartament("computerscience", 4)
	college.addDepartament("mechanical", 5)

	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartament("computerscience", 2)
	return college
}
