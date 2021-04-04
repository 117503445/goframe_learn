package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
