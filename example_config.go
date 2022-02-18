package main

import "github.com/humanmade/go-anonymize-mysqldump/pkg"

var ExampleWordPressConfig = pkg.Config{
	Patterns: []pkg.ConfigPattern{
		{
			TableName: "wp_users",
			Fields: []pkg.PatternField{
				{
					Field:       "user_login",
					Type:        "username",
					Position:    2,
					Constraints: nil,
				},
				{
					Field:       "user_pass",
					Type:        "password",
					Position:    3,
					Constraints: nil,
				},
				{
					Field:       "user_nicename",
					Type:        "username",
					Position:    4,
					Constraints: nil,
				},
				{
					Field:       "user_email",
					Type:        "email",
					Position:    5,
					Constraints: nil,
				},
				{
					Field:       "user_url",
					Type:        "url",
					Position:    6,
					Constraints: nil,
				},
				{
					Field:       "display_name",
					Type:        "name",
					Position:    10,
					Constraints: nil,
				},
			},
		},
		{
			TableName: "wp_usermeta",
			Fields: []pkg.PatternField{
				{
					Field:    "meta_value",
					Position: 4,
					Type:     "firstName",
					Constraints: []pkg.PatternFieldConstraint{
						{
							Field:    "meta_key",
							Position: 3,
							Value:    "first_name",
						},
					},
				},
				{
					Field:    "meta_value",
					Position: 4,
					Type:     "lastName",
					Constraints: []pkg.PatternFieldConstraint{
						{
							Field:    "meta_key",
							Position: 3,
							Value:    "last_name",
						},
					},
				},
				{
					Field:    "meta_value",
					Position: 4,
					Type:     "firstName",
					Constraints: []pkg.PatternFieldConstraint{
						{
							Field:    "meta_key",
							Position: 3,
							Value:    "nickname",
						},
					},
				},
				{
					Field:    "meta_value",
					Position: 4,
					Type:     "paragraph",
					Constraints: []pkg.PatternFieldConstraint{
						{
							Field:    "meta_key",
							Position: 3,
							Value:    "description",
						},
					},
				},
			},
		},
		{
			TableName: "wp_comments",
			Fields: []pkg.PatternField{
				{
					Field:       "comment_author",
					Type:        "username",
					Position:    3,
					Constraints: nil,
				},
				{
					Field:       "comment_author_email",
					Type:        "email",
					Position:    4,
					Constraints: nil,
				},
				{
					Field:       "comment_author_url",
					Type:        "url",
					Position:    5,
					Constraints: nil,
				},
				{
					Field:       "comment_author_IP",
					Type:        "ipv4",
					Position:    6,
					Constraints: nil,
				},
			},
		},
	},
}
