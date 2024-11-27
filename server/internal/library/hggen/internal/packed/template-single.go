package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
