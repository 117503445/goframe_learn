package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
