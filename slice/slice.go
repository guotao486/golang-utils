package slice

// 切片头部插入
//
// 会引起内存的分配和复制操作
func SliceInsertHead(dst, src []interface{}) []interface{} {
	dst = append(src, dst...)
	return dst
}

// 切片尾部插入
//
// 若空间不足，会自动扩容并分配新的空间
func SliceInsertFooter(dst, src []interface{}) []interface{} {
	dst = append(dst, src...)
	return dst
}

// 切片中间插入
//
// 在中间（i）插入元素：
//
// 先通过append增长空间；
//
// 使用copy移动后面元素；
//
// 插入元素到指定位置
func SliceInsertIndex(dst, src []interface{}, i int) []interface{} {
	dst = append(dst, src)
	copy(dst[i+len(src):], dst[i:])
	copy(dst[i:], src)
	return dst
}

// 切片尾部删除
//
// l 长度
func SliceDeleteFooter(dst []interface{}, l int) []interface{} {
	if l >= len(dst) {
		return nil
	}

	dst = dst[:len(dst)-l]
	return dst
}

// 切片头部删除
//
// l 长度
//
// needle 改变指针
func SliceDeleteHead(dst []interface{}, l int, needle bool) []interface{} {
	if l >= len(dst) {
		return nil
	}

	if needle {
		return dst[l:]
	}

	dst = append(dst[:0], dst[l:]...)
	// dst = dst[:copy(dst, dst[l:])]
	return dst
}

// 切片指定位置删除
// i 删除位置
// l 删除长度
func SliceDeleteIndex(dst []interface{}, i, l int) []interface{} {
	dst = append(dst[:i], dst[i+l:]...)
	dst = dst[:i+copy(dst[i:], dst[i+l:])]
	return dst
}
