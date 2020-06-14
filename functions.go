package namegenerator

// NewSession will provide a session for the user to use.
func NewSession() *Session {
	s := Session{}
	return &s
}

//GetName provides a random last name
func (s *Session) GetName() (resp string, err error) {
	s.lastNames.l.Lock()
	defer s.lastNames.l.Unlock()
	//FROM CACHE FIRST
	resp, err = s.lastNames.getNameFromCache()
	if err == nil {
		return
	}
	//TRIES TO USE THE API
	body, err := request(EndpointLastNames())
	if err != nil {
		return
	}

	response := parseRequest(body)
	s.lastNames.addNamesToCache(response)

	resp, err = s.lastNames.getNameFromCache()
	if err == nil {
		return
	}

	return
}

//GetCityName provides a random city name
func (s *Session) GetCityName() (resp string, err error) {
	s.cityNames.l.Lock()
	defer s.cityNames.l.Unlock()
	//FROM CACHE FIRST
	resp, err = s.cityNames.getNameFromCache()
	if err == nil {
		return
	}
	//TRIES TO USE THE API
	body, err := request(EndpointCities())
	if err != nil {
		return
	}

	response := parseRequest(body)
	s.cityNames.addNamesToCache(response)

	resp, err = s.cityNames.getNameFromCache()
	if err == nil {
		return
	}

	return
}

//GetAdjectiveOne provides a random adjective (type 1)
func (s *Session) GetAdjectiveOne() (resp string, err error) {
	s.adjectiveOne.l.Lock()
	defer s.adjectiveOne.l.Unlock()
	//FROM CACHE FIRST
	resp, err = s.adjectiveOne.getNameFromCache()
	if err == nil {
		return
	}
	//TRIES TO USE THE API
	body, err := request(EndpointAdjectivesOne())
	if err != nil {
		return
	}

	response := parseRequest(body)
	s.adjectiveOne.addNamesToCache(response)

	resp, err = s.adjectiveOne.getNameFromCache()
	if err == nil {
		return
	}

	return
}

//GetAdjectiveTwo provides a random adjective (type 2)
func (s *Session) GetAdjectiveTwo() (resp string, err error) {
	s.adjectiveTwo.l.Lock()
	defer s.adjectiveTwo.l.Unlock()
	//FROM CACHE FIRST
	resp, err = s.adjectiveTwo.getNameFromCache()
	if err == nil {
		return
	}
	//TRIES TO USE THE API
	body, err := request(EndpointAdjectivesTwo())
	if err != nil {
		return
	}

	response := parseRequest(body)
	s.adjectiveTwo.addNamesToCache(response)

	resp, err = s.adjectiveTwo.getNameFromCache()
	if err == nil {
		return
	}

	return
}

//GetKingdomName provides a random kingdom name
func (s *Session) GetKingdomName() (resp string, err error) {
	s.kingdomNames.l.Lock()
	defer s.kingdomNames.l.Unlock()
	//FROM CACHE FIRST
	resp, err = s.kingdomNames.getNameFromCache()
	if err == nil {
		return
	}
	//TRIES TO USE THE API
	body, err := request(EndpointKingdomNames())
	if err != nil {
		return
	}

	response := parseRequest(body)
	s.kingdomNames.addNamesToCache(response)

	resp, err = s.kingdomNames.getNameFromCache()
	if err == nil {
		return
	}

	return
}

//GetTownName provides a random town name
func (s *Session) GetTownName() (resp string, err error) {
	s.townNames.l.Lock()
	defer s.townNames.l.Unlock()
	//FROM CACHE FIRST
	resp, err = s.townNames.getNameFromCache()
	if err == nil {
		return
	}
	//TRIES TO USE THE API
	body, err := request(EndpointTownNames())
	if err != nil {
		return
	}

	response := parseRequest(body)
	s.townNames.addNamesToCache(response)

	resp, err = s.townNames.getNameFromCache()
	if err == nil {
		return
	}

	return
}

//GetCountryName provides a random country name
func (s *Session) GetCountryName() (resp string, err error) {
	s.countryNames.l.Lock()
	defer s.countryNames.l.Unlock()
	//FROM CACHE FIRST
	resp, err = s.countryNames.getNameFromCache()
	if err == nil {
		return
	}
	//TRIES TO USE THE API
	body, err := request(EndpointCountryNames())
	if err != nil {
		return
	}

	response := parseRequest(body)
	s.countryNames.addNamesToCache(response)

	resp, err = s.countryNames.getNameFromCache()
	if err == nil {
		return
	}

	return
}
