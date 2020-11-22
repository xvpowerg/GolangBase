package utils

import (
	"fmt"
	"path/filepath"
)

//宣告常數
const (
	ROOT              string = `C:\MyFile\OnlineClass\it360\Golang\MaskCountWeb`
	ASSET_PATH        string = "asset"
	PHARMACY_CSV_FILE string = "pharmacy.csv"
	API_KEY           string = "AIzaSyBwjP1wROzaTDQwq1hDceP2x9v2L_7C27I"
	GEOCODING_API_URI string = "https://maps.googleapis.com/maps/api/geocode/json?key=%s&address=%s"
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
