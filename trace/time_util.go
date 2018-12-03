/*
 * Licensed to the OpenSkywalking under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenSkywalking licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package trace

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"
)

func getMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func getHostName() string{
	host, _ := os.Hostname()
	return  host
}


func getIpV4() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	var ip []string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip=append(ip, ipnet.IP.String())
			}
		}
	}
	return ip
}

func getOsName() string{
	return  runtime.GOOS
}





