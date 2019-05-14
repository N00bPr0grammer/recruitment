package compare

import "fmt"

func Compare() {
	command := "source" + "target"

	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
    		fmt.Println("Modified")
    		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}
