package namegenerator

// Known API Endpoints.
const (
	MainEndpoint = "https://www.namegenerator2.com/application/p.php"
)

//All the URLs
var (
	EndpointLastNames = func() string {
		return MainEndpoint + "?type=1&id=last_names&spaceflag=false"
	}
	EndpointCities = func() string {
		return MainEndpoint + "?type=1&id=fantasy_city_names&spaceflag=false"
	}
	EndpointAdjectivesOne = func() string {
		return MainEndpoint + "?type=1&id=adjectives&spaceflag=false"
	}
	EndpointAdjectivesTwo = func() string {
		return MainEndpoint + "?type=3&id=adjectives&id2=youtube_names&spaceflag=false"
	}
	EndpointKingdomNames = func() string {
		return MainEndpoint + "?type=4&id=fantasy_male_names&id2=fantasy_female_names&spaceflag=false"
	}
	EndpointTownNames = func() string {
		return MainEndpoint + "?type=1&id=fantasy_town_names&spaceflag=false"
	}
	EndpointCountryNames = func() string {
		return MainEndpoint + "?type=1&id=fantasy_country_names&spaceflag=false"
	}
)
