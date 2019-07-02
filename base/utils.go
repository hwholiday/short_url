/*
 * @Author: holiday
 * @Date: 2019-07-02 13:08:10
 * @Last Modified by: holiday
 * @Last Modified time: 2019-07-02 13:08:51
 */
package base

import "time"

func GetTime() int64 {
	return time.Now().Unix()
}
