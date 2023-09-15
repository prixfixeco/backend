package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	queryOutput := map[string][]*Query{
		"admin.sql":                      buildAdminQueries(),
		"webhooks.sql":                   buildWebhooksQueries(),
		"users.sql":                      buildUsersQueries(),
		"households.sql":                 buildHouseholdsQueries(),
		"household_user_memberships.sql": buildHouseholdUserMembershipsQueries(),
	}

	for filePath, queries := range queryOutput {
		existingFile, err := os.ReadFile(path.Join("internal", "database", "postgres", "sqlc_queries", filePath))
		if err != nil {
			panic(err)
		}

		var fileContent string
		for i, query := range queries {
			if i != 0 {
				fileContent += "\n"
			}
			fileContent += query.Render()
		}

		fileOutput := ""
		for _, line := range strings.Split(strings.TrimSpace(fileContent), "\n") {
			fileOutput += strings.TrimSuffix(line, " ") + "\n"
		}

		if string(existingFile) != fileOutput {
			fmt.Printf("files don't match: %s\n", filePath)
		}
	}

	fmt.Println("done")
}