package tools

import "tw.com.maskweb/obj"

func LoadingMaskList(positionList []*obj.Position, latlonP1 *obj.LatLng) []*obj.Mask {
	result := make([]*obj.Mask, 0, 7000)
	maskCountMap := CsvToMaskCountMap()
	for _, position := range positionList {
		maskCount, ok := maskCountMap[position.ID]
		if ok {
			//表示此藥局存在於我們的藥局資訊內
			//剩餘的map表示 不存在於我們position.json藥局資料內的藥局
			delete(maskCountMap, position.ID)
		} else {
			//表示口罩售完
			maskCount = &obj.MaskCount{
				Adult: 0,
				Ahild: 0,
			}
		}

	}
	return result
}
