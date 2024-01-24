package basic

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/inigolabs/fezzik/examples/basic/gen/basic"
	"github.com/stretchr/testify/require"
)

type UnionSearchResultInLineSearch = basic.UnionSearchResultInLineSearch
type UnionSearchResultSearch = basic.UnionSearchResultSearch
type UnionSearchResultNoTypenameSearch = basic.UnionSearchResultNoTypenameSearch
type InterfaceGetCreatedGetCreated = basic.InterfaceGetCreatedGetCreated
type InterfaceGetCreatedNoTypenameGetCreated = basic.InterfaceGetCreatedNoTypenameGetCreated

func ptr(val string) *string {
	return &val
}

func TestJson(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		value    interface{}
		expected string
	}{
		{
			name: "UnionSearchResultInLineSearch Book",
			payload: `{
        "__typename":"Book",
        
        "Title":"Book-Title",
        "Updated":"Book-Updated",
        
        "Name":"Author-Name"
      }`,
			value: &UnionSearchResultInLineSearch{
				Typename: ptr("Book"),
				UnionSearchResultInLineSearchBook: &basic.UnionSearchResultInLineSearchBook{
					Title:   "Book-Title",
					Updated: "Book-Updated",
				},
			},
			expected: `{
        "__typename":"Book",
        
        "Title":"Book-Title",
        "Updated":"Book-Updated",
        
        "Name":""
      }`,
		},
		{
			name: "UnionSearchResultInLineSearch Author",
			payload: `{
		      "__typename":"Author",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &UnionSearchResultInLineSearch{
				Typename: ptr("Author"),
				AuthorFr: &basic.AuthorFr{
					Name:    "Author-Name",
					Updated: "Author-Updated",
				},
			},
			expected: `{
		      "__typename":"Author",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":""
		    }`,
		},
		{
			name: "UnionSearchResultInLineSearch Null",
			payload: `{
		      "__typename":null,

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &UnionSearchResultInLineSearch{},
			expected: `{
		      "__typename":null,

		      "Name":"",
		      "Updated":"",

		      "Title":""
		    }`,
		},
		{
			name: "UnionSearchResultInLineSearch Unknown",
			payload: `{
		      "__typename":"Unknown",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &UnionSearchResultInLineSearch{
				Typename: ptr("Unknown"),
			},
			expected: `{
		      "__typename":"Unknown",

		      "Name":"",
		      "Updated":"",

		      "Title":""
		    }`,
		},
		{
			name: "UnionSearchResultSearch Book",
			payload: `{
		      "__typename":"Book",

		      "Title":"Book-Title",
		      "Updated":"Book-Updated",

		      "Name":"Author-Name"
		    }`,
			value: &UnionSearchResultSearch{
				Typename: ptr("Book"),
				BookFr: &basic.BookFr{
					Title:   "Book-Title",
					Updated: "Book-Updated",
				},
			},
			expected: `{
		      "__typename":"Book",

		      "Title":"Book-Title",
		      "Updated":"Book-Updated",

		      "Name":""
		    }`,
		},
		{
			name: "UnionSearchResultSearch Author",
			payload: `{
		      "__typename":"Author",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &UnionSearchResultSearch{
				Typename: ptr("Author"),
				AuthorFr: &basic.AuthorFr{
					Name:    "Author-Name",
					Updated: "Author-Updated",
				},
				AuthorAnotherFr: &basic.AuthorAnotherFr{
					Name:    "Author-Name",
					Updated: "Author-Updated",
				},
			},
			expected: `{
		      "__typename":"Author",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":""
		    }`,
		},
		{
			name: "UnionSearchResultSearch Null",
			payload: `{
		      "__typename":null,

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &UnionSearchResultSearch{},
			expected: `{
		      "__typename":null,

		      "Name":"",
		      "Updated":"",

		      "Title":""
		    }`,
		},
		{
			name: "UnionSearchResultSearch Unknown",
			payload: `{
		      "__typename":"Unknown",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &UnionSearchResultSearch{
				Typename: ptr("Unknown"),
			},
			expected: `{
		      "__typename":"Unknown",

		      "Name":"",
		      "Updated":"",

		      "Title":""
		    }`,
		},
		{
			name: "InterfaceGetCreatedGetCreated Book",
			payload: `{
        "__typename":"Book",
        "Created":"Created",

        "Title":"Book-Title",
        "Updated":"Book-Updated",

        "Name":"Author-Name"
      }`,
			value: &InterfaceGetCreatedGetCreated{
				Typename: ptr("Book"),
				Created:  "Created",
				BookFr: &basic.BookFr{
					Title:   "Book-Title",
					Updated: "Book-Updated",
				},
			},
			expected: `{
        "__typename":"Book",
        "Created":"Created",

        "Title":"Book-Title",
        "Updated":"Book-Updated",

        "Name":""
      }`,
		},
		{
			name: "InterfaceGetCreatedGetCreated Author",
			payload: `{
		      "__typename":"Author",
          "Created":"Created",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &InterfaceGetCreatedGetCreated{
				Typename: ptr("Author"),
				Created:  "Created",
				AuthorFr: &basic.AuthorFr{
					Name:    "Author-Name",
					Updated: "Author-Updated",
				},
			},
			expected: `{
		      "__typename":"Author",
          "Created":"Created",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":""
		    }`,
		},
		{
			name: "InterfaceGetCreatedGetCreated Null",
			payload: `{
		      "__typename":null,
          "Created":"Created",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &InterfaceGetCreatedGetCreated{
				Created: "Created",
			},
			expected: `{
		      "__typename":null,
          "Created":"Created",

		      "Name":"",
		      "Updated":"",

		      "Title":""
		    }`,
		},
		{
			name: "InterfaceGetCreatedGetCreated Unknown",
			payload: `{
		      "__typename":"Unknown",
          "Created":"Created",

		      "Name":"Author-Name",
		      "Updated":"Author-Updated",

		      "Title":"Book-Title"
		    }`,
			value: &InterfaceGetCreatedGetCreated{
				Typename: ptr("Unknown"),
				Created:  "Created",
			},
			expected: `{
		      "__typename":"Unknown",
          "Created":"Created",

		      "Name":"",
		      "Updated":"",

		      "Title":""
		    }`,
		},
	}

	for _, test := range tests {
		objType := reflect.TypeOf(test.value)
		if objType.Kind() == reflect.Ptr {
			objType = objType.Elem()
		}

		obj := reflect.New(objType).Interface()
		require.NoError(t, json.Unmarshal([]byte(test.payload), &obj))
		res, err := json.Marshal(obj)

		require.NoError(t, err)
		require.EqualValues(t, test.value, obj)
		require.JSONEq(t, test.expected, string(res))
	}
}
