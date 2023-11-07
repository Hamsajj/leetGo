package accounts_merge

import (
	"github.com/Hamsajj/leetGo/utils"
	"testing"
)

func TestAccountsMerge(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{
			name: "testCase 1",
			input: [][]string{
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
		},
		{
			name: "testCase 2",
			input: [][]string{
				{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
				{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
				{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
				{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
			},
			expected: [][]string{
				{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
				{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := accountsMerge(tt.input)
			for _, accounts := range got {
				name := accounts[0]
				for _, expectedAccounts := range tt.expected {
					if expectedAccounts[0] == name {
						if !utils.AreSlicesEqual(accounts, expectedAccounts) {
							t.Errorf("expected %v but got %v", expectedAccounts, accounts)
						}
					}
				}
			}
		})
	}
}
