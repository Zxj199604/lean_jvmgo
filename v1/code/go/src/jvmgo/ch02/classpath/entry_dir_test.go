/**
 * @Author: Zxj
 * @Description:
 * @File: entry_dir_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/13 2:48 下午
 * @Software : GoLand
 */

package classpath

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	dirEntry := newDirEntry("/Users/zxj/IdeaProjects/hyperchainsdk-test/litesdk/MainSDK/target/classes/cn/hyperchain/vortex/utils")
	log.Println(dirEntry.String())
	data, _, _ := dirEntry.readClass("DatabaseUtil.class")
	log.Println(data)
}
