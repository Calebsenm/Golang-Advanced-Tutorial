package main

type nullDepartament struct {
	numberOfProfesors int
}

func (c *nullDepartament) getNumberOfProfessors() int {
	return 0
}

func (c *nullDepartament) getName() string {
	return "nullDepartment"
}
