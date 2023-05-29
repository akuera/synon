package synon_test

import (
	"github.com/akuera/synon"
	"github.com/stretchr/testify/assert"
	"testing"
)

type EmbeddedStr struct {
	Name string
}

type testUser struct {
	Name    string
	Phone   string
	Email   string
	private string
	More    *EmbeddedStr
	//Me    *Me
}

func Test_rnvl(t *testing.T) {
	tests := []struct {
		Name string
		Val1 interface{}
		Val2 interface{}
		Want interface{}
	}{
		{
			"[struct] missing value",
			&testUser{
				Name:  "test1",
				Phone: "+12345678",
				Email: "test1@firstmail.com",
			},
			&testUser{
				Name:  "",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
			&testUser{
				Name:  "test1",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
		},
		{
			"[struct] missing field",
			&testUser{
				Name: "test1",
			},
			&testUser{
				Name:  "",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
			&testUser{
				Name:  "test1",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
		},
		{
			"[struct] private field",
			&testUser{
				Name: "test1",
			},
			&testUser{
				Name:    "",
				Phone:   "+98765432",
				Email:   "test2@testmail.com",
				private: "private",
			},
			&testUser{
				Name:  "test1",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
		},
		{
			"[struct] mismatch types",
			"user1",
			&testUser{
				Name:  "",
				Phone: "+12345678",
			},
			nil,
		},
		{
			"[string] simple replace",
			"user1",
			"user2",
			"user2",
		},
		{
			"[string] simple replace",
			"",
			"user2",
			"user2",
		},
		{
			Name: "[string] nil value",
			Val2: "user2",
			Want: nil,
		},
		{
			"[int] simple replace",
			2,
			1,
			1,
		},
		{
			"[int] same values",
			1,
			1,
			1,
		},
		{
			Name: "[int] both nil",
			Want: nil,
		},
		{
			"[struct] nil complex embedded",
			&testUser{
				Name:  "test1",
				Phone: "+12345678",
				Email: "test1@firstmail.com",
			},
			&testUser{
				Name:  "",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
			&testUser{
				Name:  "test1",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
		},
		{
			"[struct] nil second param embedded",
			&testUser{
				Name:  "test1",
				Phone: "+12345678",
				Email: "test1@firstmail.com",
				More:  &EmbeddedStr{Name: "test"},
			},
			&testUser{
				Name:  "",
				Phone: "+98765432",
				Email: "test2@testmail.com",
			},
			&testUser{
				Name:  "test1",
				Phone: "+98765432",
				Email: "test2@testmail.com",
				More:  &EmbeddedStr{Name: "test"},
			},
		},
		{
			"[struct] nil first param embedded",
			&testUser{
				Name:  "test1",
				Phone: "+12345678",
				Email: "test1@firstmail.com",
			},
			&testUser{
				Name:  "",
				Phone: "+98765432",
				Email: "test2@testmail.com",
				More:  &EmbeddedStr{Name: "test"},
			},
			&testUser{
				Name:  "test1",
				Phone: "+98765432",
				Email: "test2@testmail.com",
				More:  &EmbeddedStr{Name: "test"},
			},
		},
		{
			"[map] simple case",
			map[string]string{
				"foo": "bar",
			},
			map[string]string{},
			map[string]string{
				"foo": "bar",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := synon.Synon(tt.Val1, tt.Val2)
			assert.Equal(t, tt.Want, got)
		})
	}
}
