package service

func PercentageCalculation(answers []int32) (int, int, int) {

	favorable := float32(0)
	neutral := float32(0)
	unfavorable := float32(0)
	for _, answer := range answers {

		if answer < 2 {
			favorable++
		} else if answer == 2 {
			neutral++
		} else if neutral > 2 && neutral < 5 {
			unfavorable++
		}
	}
	total := favorable + neutral + unfavorable

	favorable = (favorable / total) * 100
	neutral = (neutral / total) * 100
	unfavorable = (unfavorable / total) * 100

	return int(favorable), int(neutral), int(unfavorable)

}

// func PercentageAnswrsQuestions()
