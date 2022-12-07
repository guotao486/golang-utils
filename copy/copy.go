/*
 * @Author: GG
 * @Date: 2022-12-07 16:50:20
 * @LastEditTime: 2022-12-07 16:54:58
 * @LastEditors: GG
 * @Description: 深拷贝
 * @FilePath: \golang-utils\copy\copy.go
 *
 */
package copy

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

/**
 * 有三种方法
 * 一是用gob序列化成字节序列再反序列化生成克隆对象；
 * 二是先转换成json字节序列，再解析字节序列生成克隆对象；
 * 三是针对具体情况，定制化拷贝。
 * 前两种方法虽然比较通用但是因为使用了reflex反射，性能比定制化拷贝要低出2个数量级，所以在性能要求较高的情况下应该尽量避免使用前两者。
 */

/**
 * @Description: 利用gob进行深拷贝
 */
func DeepCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}

/**
 * @Description: 利用json进行深拷贝
 */
func DeepCopyByJson(src, dst interface{}) error {
	if tmp, err := json.Marshal(&src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}
