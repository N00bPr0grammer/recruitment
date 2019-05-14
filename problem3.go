package compare

import "fmt"

//comparing modified change
func Compare() {
	command := "source" + "target"

	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
    		fmt.Println("Modified")
    		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}
