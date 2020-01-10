package all

import (
	_ "github.com/negbie/telegraf/plugins/inputs/cpu"
	_ "github.com/negbie/telegraf/plugins/inputs/disk"
	_ "github.com/negbie/telegraf/plugins/inputs/diskio"
	_ "github.com/negbie/telegraf/plugins/inputs/influxdb"
	_ "github.com/negbie/telegraf/plugins/inputs/influxdb_listener"
	_ "github.com/negbie/telegraf/plugins/inputs/logparser"
	_ "github.com/negbie/telegraf/plugins/inputs/prometheus"
	_ "github.com/negbie/telegraf/plugins/inputs/redis"
	_ "github.com/negbie/telegraf/plugins/inputs/snmp"
	_ "github.com/negbie/telegraf/plugins/inputs/snmp_trap"
	_ "github.com/negbie/telegraf/plugins/inputs/socket_listener"
	_ "github.com/negbie/telegraf/plugins/inputs/system"
	_ "github.com/negbie/telegraf/plugins/inputs/tail"
)
