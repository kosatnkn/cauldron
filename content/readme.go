package content

import "fmt"

// GenerateReadme generates the content of the readme file.
func GenerateReadme(name string) string {

	return fmt.Sprintf(`# %s

---
Powered By [Catalyst](https://github.com/kosatnkn/catalyst)`, name)
}
