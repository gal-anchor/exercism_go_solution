package main

import (
	airportrobot "exercism_go_solution/airport-robot"
	"exercism_go_solution/armstrong-numbers"
	"exercism_go_solution/atbash-cipher"
	"exercism_go_solution/bob"
	"exercism_go_solution/booking-up-for-beauty"
	"exercism_go_solution/census"
	"exercism_go_solution/clock"
	jedlik "exercism_go_solution/elons-toys"
	"exercism_go_solution/expenses"
	greeting "exercism_go_solution/hello-world"
	lsproduct "exercism_go_solution/largest-series-product"
	"exercism_go_solution/meteorology"
	prime "exercism_go_solution/nth-prime"
	"exercism_go_solution/pangram"
	parsinglogfiles "exercism_go_solution/parsing-log-files"
	"exercism_go_solution/raindrops"
	"exercism_go_solution/reverse-string"
	"exercism_go_solution/sieve"
	"exercism_go_solution/simple-cipher"
	"exercism_go_solution/sorting-room"
	"exercism_go_solution/space-age"
	"exercism_go_solution/strain"
	thefarm "exercism_go_solution/the-farm"
	wordcount "exercism_go_solution/word-count"
	"fmt"
)

func main() {
	fmt.Println(greeting.HelloWorld())
	fmt.Println(bob.Hey("gal"))
	fmt.Println(clock.New(21, 2))
	fmt.Println(wordcount.WordCount("galgalglaglaglalgla adsadas adsadas"))
	fmt.Println(lsproduct.LargestSeriesProduct("123123123", 2))
	fmt.Println(prime.Nth(555))
	fmt.Println(space.Age(51548712, "Mercury"))
	fmt.Println(booking.Schedule("6/6/2005 10:30:00"))
	fmt.Println(atbash.Atbash("galgal414141"))
	fmt.Println(jedlik.NewCar(100, 500).CanFinish(300))
	var g airportrobot.Greeter = airportrobot.Italian{}
	fmt.Println(g.Greet("Alice"))
	fmt.Println(parsinglogfiles.IsValidLine("[DBG] - galgalglaglalgalga"))
	fmt.Println(sorting.DescribeNumber(42))
	fmt.Println(meteorology.MeteorologyData{})
	fmt.Println(thefarm.ValidateNumberOfCows(-1))
	residents := []*census.Resident{
		census.NewResident("Alice", 30, map[string]string{"street": "Main St"}),
		census.NewResident("Gal", 25, map[string]string{"street": ""})}
	fmt.Println(census.Count(residents))
	recs := []expenses.Record{
		{Day: 1, Amount: 10, Category: "food"},
		{Day: 2, Amount: 7.5, Category: "books"},
		{Day: 3, Amount: 5, Category: "food"},
	}
	totalFood, err := expenses.CategoryExpenses(recs, expenses.DaysPeriod{From: 1, To: 3}, "food")
	fmt.Println(totalFood, err)

	fmt.Println(pangram.IsPangram("The quick brown fox jumps over the lazy dog"))

	caesar := cipher.NewCaesar()
	encoded := caesar.Encode("hello")
	decoded := caesar.Decode(encoded)
	fmt.Println(encoded, decoded)

	fmt.Println(raindrops.Convert(28))
	fmt.Println(raindrops.Convert(30))

	fmt.Println(sieve.Sieve(30))
	fmt.Println(armstrong.IsNumber(153))
	fmt.Println(armstrong.IsNumber(154))
	fmt.Println(reverse.Reverse("Hello, 世界"))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(strain.Keep(nums, func(n int) bool { return n%2 == 0 }))
	fmt.Println(strain.Discard(nums, func(n int) bool { return n%2 == 0 }))

}
