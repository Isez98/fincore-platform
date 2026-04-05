package services

func AccountService() []Account {
	return []Account{
		{
			ID:       "acc_001",
			Name:     "Checking Account",
			Balance:  1250.75,
			Currency: "USD",
		},
		{
			ID:       "acc_002",
			Name:     "Savings Account",
			Balance:  5400.00,
			Currency: "USD",
		},
	}
}
