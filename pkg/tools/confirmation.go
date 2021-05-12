package tools

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func AskForConfirmation(s string, def bool) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		var res string
		if def {
			res = "Y/n"
		} else {
			res = "y/N"
		}

		log.Printf("%s [%s]", s, res)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		switch response {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			return def
		}
	}
}
