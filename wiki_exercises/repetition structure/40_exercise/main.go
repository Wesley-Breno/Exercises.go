package main

import "fmt"

type City struct {
	cityCode                  float64
	numberOfPassengerVehicles float64
	numberOfTrafficAccidents  float64
}

func GetCityWithHighestAccidentRate(cities []City) City {
	var highestCity City
	var highestValue float64
	for _, city := range cities {
		if city.numberOfTrafficAccidents > highestValue {
			highestValue = city.numberOfTrafficAccidents
			highestCity = city
		}
	}
	return highestCity
}

func GetCityWithLowestAccidentRate(cities []City) City {
	var lowestCity City
	var lowestValue float64
	for i, city := range cities {
		if i == 0 {
			lowestValue = city.numberOfTrafficAccidents
			lowestCity = city
			continue
		}
		if city.numberOfTrafficAccidents < lowestValue {
			lowestValue = city.numberOfTrafficAccidents
			lowestCity = city
		}
	}
	return lowestCity
}

func GetAverageVehicles(cities []City) float64 {
	var sum float64
	var count float64
	for _, city := range cities {
		count += 1
		sum += city.numberOfPassengerVehicles
	}
	if count == 0 {
		return 0
	}
	return sum / count
}

func GetAverageAccidentsInCitiesWithLessThan2000Vehicles(cities []City) float64 {
	var sum float64
	var count float64
	for _, city := range cities {
		if city.numberOfPassengerVehicles < 2000 {
			count += 1
			sum += city.numberOfTrafficAccidents
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}

func main() {
	var cities []City
	for i := 0; i < 5; i++ {
		var cityCode float64
		var numberOfPassengerVehicles float64
		var numberOfTrafficAccidents float64

		fmt.Printf("Enter the code for city %v: ", i+1)
		fmt.Scan(&cityCode)

		fmt.Printf("\nEnter the number of passenger vehicles: ")
		fmt.Scan(&numberOfPassengerVehicles)

		fmt.Printf("\nEnter the number of traffic accidents with victims: ")
		fmt.Scan(&numberOfTrafficAccidents)

		cities = append(cities, City{
			cityCode:                  cityCode,
			numberOfPassengerVehicles: numberOfPassengerVehicles,
			numberOfTrafficAccidents:  numberOfTrafficAccidents,
		})
		fmt.Println("================================")
	}

	lowestCity := GetCityWithLowestAccidentRate(cities)
	highestCity := GetCityWithHighestAccidentRate(cities)

	fmt.Println("City with the highest accident rate:")
	fmt.Println("City code:", highestCity.cityCode)
	fmt.Println("Accidents:", highestCity.numberOfTrafficAccidents)

	fmt.Println("================================")

	fmt.Println("City with the lowest accident rate:")
	fmt.Println("City code:", lowestCity.cityCode)
	fmt.Println("Accidents:", lowestCity.numberOfTrafficAccidents)

	fmt.Println("================================")

	fmt.Println("Average number of vehicles across the 5 cities:")
	fmt.Println(GetAverageVehicles(cities))

	fmt.Println("================================")

	fmt.Println("Average number of traffic accidents in cities with less than 2000 passenger vehicles:")
	fmt.Println(GetAverageAccidentsInCitiesWithLessThan2000Vehicles(cities))
}
