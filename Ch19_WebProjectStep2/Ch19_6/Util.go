package utils

import (
	"fmt"
	"path/filepath"
	"time"
)

//宣告常數
const (
	ROOT              string = `C:\MyFile\OnlineClass\it360\Golang\MaskCountWeb`
	ASSET_PATH        string = "asset"
	PHARMACY_CSV_FILE string = "pharmacy.csv"
	API_KEY           string = "AIzaSyBwjP1wROzaTDQwq1hDceP2x9v2L_7C27I"
	GEOCODING_API_URI string = "https://maps.googleapis.com/maps/api/geocode/json?key=%s&address=%s"

	//position json檔
	POSITION_JSON_FILE string = "position.json"
	//log路徑
	LOG_PATH string = "log"
)

//取得GeocodingApi網址
func GetGeocodingApiUrl(address string) string {
	return fmt.Sprintf(GEOCODING_API_URI, API_KEY, address)
}

//取得AssetPath
func getAssetPath() string {
	return filepath.Join(ROOT, ASSET_PATH)
}

//取得PharmacyCsvPath
func GetPharmacyCsvPath() string {
	return filepath.Join(getAssetPath(), PHARMACY_CSV_FILE)
}

//取得Position.jspn存取路徑
func GetPositionJsonPath() string {
	return filepath.Join(getAssetPath(), POSITION_JSON_FILE)
}

//取得Log的資料夾路徑
func getLogDirPath() string {
	return filepath.Join(ROOT, LOG_PATH)
}

//取得Log的檔案名稱格式為年月日
func getLogFilePath(logName string) string {
	tmpLogName := fmt.Sprintf("%s%s.log",
		logName,
		time.Now().Format("-2006-01-02")) //表示日期格式年月日
	return filepath.Join(getLogDirPath(), tmpLogName)
}
