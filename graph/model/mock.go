package model

func stringToPtr(s string) *string {
	return &s
}

func intToPtr(i int) *int {
	return &i
}

func floatToPtr(f float64) *float64 {
	return &f
}

var EventsMock []*Event = []*Event{
	{
		ID:   2524864,
		Name: stringToPtr("Del Potro J.M.-Tsonga J-W."),
		Markets: []*Market{
			{
				ID:   2524864,
				Name: stringToPtr("Winner"),
				Odds: []*Odd{
					{
						ID:    2319,
						Name:  stringToPtr("1"),
						Value: floatToPtr(1.45),
					},
					{
						ID:    2320,
						Name:  stringToPtr("2"),
						Value: floatToPtr(2.40),
					},
				},
			},
			{
				ID:   2528173,
				Name: stringToPtr("1. Set"),
				Odds: []*Odd{
					{
						ID:    2359,
						Name:  stringToPtr("1"),
						Value: floatToPtr(1.55),
					},
					{
						ID:    2360,
						Name:  stringToPtr("2"),
						Value: floatToPtr(2.20),
					},
				},
			},
		},
	},
	{
		ID:   2383848,
		Name: stringToPtr("TCViboV-RPA Perugia"),
		Markets: []*Market{
			{
				ID:   2383848,
				Name: stringToPtr("Winner"),
				Odds: []*Odd{
					{
						ID:    200,
						Name:  stringToPtr("1"),
						Value: floatToPtr(1.70),
					},
					{
						ID:    201,
						Name:  stringToPtr("2"),
						Value: floatToPtr(1.90),
					},
				},
			},
		},
	},
}

