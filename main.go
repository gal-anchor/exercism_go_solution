package main

import (
	"exercism_go_solution/acronym"
	airportrobot "exercism_go_solution/airport-robot"
	"exercism_go_solution/animal-magic"
	"exercism_go_solution/annalyns-infiltration"
	"exercism_go_solution/armstrong-numbers"
	"exercism_go_solution/atbash-cipher"
	"exercism_go_solution/bird-watcher"
	"exercism_go_solution/blackjack"
	"exercism_go_solution/bob"
	"exercism_go_solution/booking-up-for-beauty"
	"exercism_go_solution/card-tricks"
	"exercism_go_solution/cars-assemble"
	"exercism_go_solution/census"
	"exercism_go_solution/chessboard"
	"exercism_go_solution/clock"
	"exercism_go_solution/collatz-conjecture"
	"exercism_go_solution/darts"
	"exercism_go_solution/difference-of-squares"
	jedlik "exercism_go_solution/elons-toys"
	"exercism_go_solution/expenses"
	"exercism_go_solution/gigasecond"
	"exercism_go_solution/gophers-gorgeous-lasagna"
	"exercism_go_solution/gross-store"
	"exercism_go_solution/hamming"
	greeting "exercism_go_solution/hello-world"
	"exercism_go_solution/interest-is-interesting"
	"exercism_go_solution/isogram"
	lsproduct "exercism_go_solution/largest-series-product"
	lasagnamaster "exercism_go_solution/lasagna-master"
	"exercism_go_solution/leap"
	"exercism_go_solution/logs-logs-logs"
	"exercism_go_solution/luhn"
	"exercism_go_solution/meteorology"
	speed "exercism_go_solution/need-for-speed"
	prime "exercism_go_solution/nth-prime"
	"exercism_go_solution/nucleotide-count"
	"exercism_go_solution/pangram"
	parsinglogfiles "exercism_go_solution/parsing-log-files"
	"exercism_go_solution/party-robot"
	"exercism_go_solution/phone-number"
	primefactors "exercism_go_solution/prime-factors"
	"exercism_go_solution/raindrops"
	"exercism_go_solution/resistor-color"
	"exercism_go_solution/reverse-string"
	"exercism_go_solution/rna-transcription"
	"exercism_go_solution/scrabble-score"
	"exercism_go_solution/series"
	"exercism_go_solution/sieve"
	"exercism_go_solution/simple-cipher"
	"exercism_go_solution/sorting-room"
	"exercism_go_solution/space-age"
	"exercism_go_solution/strain"
	"exercism_go_solution/sublist"
	thefarm "exercism_go_solution/the-farm"
	"exercism_go_solution/two-fer"
	"exercism_go_solution/vehicle-purchase"
	"exercism_go_solution/weather-forecast"
	"exercism_go_solution/welcome-to-tech-palace"
	wordcount "exercism_go_solution/word-count"
	"fmt"
	"time"
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
	fmt.Println(darts.Score(0, 0))
	fmt.Println(darts.Score(2, 2))
	fmt.Println(darts.Score(11, 0))
	fmt.Println(phonenumber.Format("+1 (613)-995-0253"))
	fmt.Println(lasagna.ElapsedTime(3, 20))
	fmt.Println(annalyn.CanFreePrisoner(false, false, true, false))
	fmt.Println(annalyn.CanFreePrisoner(true, false, false, true))
	fmt.Println(partyrobot.AssignTable("Gal", 7, "Alice", "left", 3.5))
	fmt.Println(weather.Forecast("Goblinopolis", "sunny"))
	fmt.Println(blackjack.FirstTurn("ace", "king", "five"))
	fmt.Println(cars.CalculateWorkingCarsPerHour(221, 80))
	fmt.Println(cars.CalculateWorkingCarsPerMinute(221, 80))
	fmt.Println(cars.CalculateCost(37))
	fmt.Println(purchase.NeedsLicense("car"))
	fmt.Println(purchase.ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
	fmt.Println(purchase.CalculateResellPrice(10000, 4))
	units := gross.Units()
	bill := gross.NewBill()
	gross.AddItem(bill, units, "eggs", "dozen")
	qty, _ := gross.GetItem(bill, "eggs")
	fmt.Println(qty)
	fmt.Println(lasagnamaster.PreparationTime([]string{"noodles", "sauce", "meat"}, 0))
	car := speed.NewCar(5, 2)
	track := speed.NewTrack(100)
	fmt.Println(speed.CanFinish(car, track))
	fmt.Println(birdwatcher.TotalBirdCount([]int{2, 5, 0, 7, 4, 1}))
	fmt.Println(cards.FavoriteCards())
	fmt.Println(logs.Application("❗ recommended search"))
	cb := chessboard.Chessboard{
		"A": {true, false, true, false, true, false, true, false},
		"B": {false, true, false, true, false, true, false, true},
		"C": {false, false, false, false, false, false, false, false},
		"D": {false, false, false, false, false, false, false, false},
		"E": {false, false, false, false, false, false, false, false},
		"F": {false, false, false, false, false, false, false, false},
		"G": {false, false, false, false, false, false, false, false},
		"H": {false, false, false, false, false, false, false, false},
	}
	fmt.Println(chessboard.CountInFile(cb, "A"))
	fmt.Println(chessboard.CountInRank(cb, 1))
	msg := techpalace.AddBorder(techpalace.WelcomeMessage("gal"), 20)
	fmt.Println(msg)
	fmt.Println(techpalace.CleanupMessage(msg))
	fmt.Println(interest.InterestRate(200.75))
	fmt.Println(interest.AnnualBalanceUpdate(200.75))
	fmt.Println(chance.RollADie())
	fmt.Println(chance.ShuffleAnimals())
	fmt.Println(twofer.ShareWith("Gal"))
	fmt.Println(twofer.ShareWith(""))
	fmt.Println(collatzconjecture.CollatzConjecture(12))
	fmt.Println(gigasecond.AddGigasecond(time.Date(2011, 4, 25, 0, 0, 0, 0, time.UTC)))
	fmt.Println(hamming.Distance("GAGCCT", "CATCGT"))
	fmt.Println(scrabble.Score("cabbage"))
	fmt.Println(leap.IsLeapYear(1996))
	fmt.Println(leap.IsLeapYear(1900))
	fmt.Println(isogram.IsIsogram("lumberjacks"))
	fmt.Println(isogram.IsIsogram("six-year-old"))
	fmt.Println(diffsquares.Difference(10))
	fmt.Println(luhn.Valid("4539 1488 0343 6467"))
	fmt.Println(luhn.Valid("8273 1232 7352 0569"))
	fmt.Println(resistorcolor.Colors())
	fmt.Println(resistorcolor.ColorCode("green"))
	fmt.Println(dna.DNA("ACGTACGT").Counts())
	fmt.Println(strand.ToRNA("ACGTGGTCTTAA"))
	fmt.Println(primefactors.Factors(360))
	fmt.Println(series.All(3, "49142"))
	fmt.Println(series.UnsafeFirst(3, "49142"))
	fmt.Println(sublist.Sublist([]int{1, 2, 3}, []int{0, 1, 2, 3, 4}))
	fmt.Println(acronym.Abbreviate("Portable Network Graphics"))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(strain.Keep(nums, func(n int) bool { return n%2 == 0 }))
	fmt.Println(strain.Discard(nums, func(n int) bool { return n%2 == 0 }))

}
