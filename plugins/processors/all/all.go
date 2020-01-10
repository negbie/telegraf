package all

import (
	_ "github.com/negbie/telegraf/plugins/processors/clone"
	_ "github.com/negbie/telegraf/plugins/processors/converter"
	_ "github.com/negbie/telegraf/plugins/processors/date"
	_ "github.com/negbie/telegraf/plugins/processors/enum"
	_ "github.com/negbie/telegraf/plugins/processors/override"
	_ "github.com/negbie/telegraf/plugins/processors/parser"
	_ "github.com/negbie/telegraf/plugins/processors/pivot"
	_ "github.com/negbie/telegraf/plugins/processors/printer"
	_ "github.com/negbie/telegraf/plugins/processors/regex"
	_ "github.com/negbie/telegraf/plugins/processors/rename"
	_ "github.com/negbie/telegraf/plugins/processors/strings"
	_ "github.com/negbie/telegraf/plugins/processors/tag_limit"
	_ "github.com/negbie/telegraf/plugins/processors/topk"
	_ "github.com/negbie/telegraf/plugins/processors/unpivot"
)
