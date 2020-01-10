package all

import (
	_ "github.com/influxdata/telegraf/plugins/inputs/cpu"
	_ "github.com/influxdata/telegraf/plugins/inputs/disk"
	_ "github.com/influxdata/telegraf/plugins/inputs/diskio"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/logparser"
	_ "github.com/influxdata/telegraf/plugins/inputs/prometheus"
	_ "github.com/influxdata/telegraf/plugins/inputs/redis"
	_ "github.com/influxdata/telegraf/plugins/inputs/snmp"
	_ "github.com/influxdata/telegraf/plugins/inputs/snmp_trap"
	_ "github.com/influxdata/telegraf/plugins/inputs/tail"
	_ "github.com/influxdata/telegraf/plugins/inputs/tcp_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/udp_listener"
)
