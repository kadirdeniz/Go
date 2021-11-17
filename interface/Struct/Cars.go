package car

type Car struct {
	Id    int
	Year  int
	Model string
	Price int
}

var Cars = []Car{
	{Id: 1, Year: 2020, Model: "Mercedes", Price: 100000},
	{Id: 2, Year: 2020, Model: "BMW", Price: 200000},
	{Id: 3, Year: 2021, Model: "Porsche", Price: 199999},
	{Id: 4, Year: 2019, Model: "Volvo", Price: 250000},
}
