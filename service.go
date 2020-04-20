package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	resources, err := QueryAllServices()
	if err != nil {
		Log.Error(err)
		return
	}
	wr := csv.NewWriter(w)
	w.Header().Set("Content-Type", "text/csv")
	filename := fmt.Sprintf("attachment;filename=resources-%s.csv", time.Now().Format("2006-01-02 15:04:05"))
	w.Header().Set("Content-Disposition", filename)
	wr.Write([]string{"id", "url", "protocol", "method", "firstpath", "ip", "created_time", "updated_time"})
	for i := range *resources {
		resource := (*resources)[i]
		record := []string{strconv.Itoa(int(resource.Id)), resource.Url, resource.Protocol, resource.Method,
			resource.Firstpath, resource.Ip, resource.CreatedTime.Format("2006-01-02 15:04:05"),
			resource.UpdatedTime.Format("2006-01-02 15:04:05")}
		err := wr.Write(record)
		if err != nil {
			Log.Error(err)
			return
		}
	}
	wr.Flush()
}

func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	resources, err := QueryAllServices()
	if err != nil {
		Log.Error(err)
		return
	}
	data, err := json.Marshal(resources)
	if err != nil {
		Log.Error(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	assets, err := QueryAllAssets()
	if err != nil {
		Log.Error(err)
		return
	}
	data, err := json.Marshal(assets)
	if err != nil {
		Log.Error(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DownloadResources(w http.ResponseWriter, r *http.Request) {
	assets, err := QueryAllAssets()
	if err != nil {
		Log.Error(err)
	}
	wr := csv.NewWriter(w)
	w.Header().Set("Content-Type", "text/csv")
	filename := fmt.Sprintf("attachment;filename=assets-%s.csv", time.Now().Format("2006-01-02 15:04:05"))
	w.Header().Set("Content-Disposition", filename)
	wr.Write([]string{"id", "url", "method", "params", "created_time", "updated_time"})
	for i := range *assets {
		asset := (*assets)[i]
		record := []string{strconv.Itoa(int(asset.Id)), asset.Url, asset.Method, asset.Params, asset.CreatedTime.Format("2006-01-02 15:04:05"),
			asset.UpdatedTime.Format("2006-01-02 15:04:05")}
		err := wr.Write(record)
		if err != nil {
			Log.Error(err)
			return
		}
	}
	wr.Flush()
}

func SetDownloadService() {
	http.HandleFunc("/download-resources", DownloadHandler)
	http.HandleFunc("/download-assets", DownloadResources)
	http.HandleFunc("/get-resources", ResourcesHandler)
	http.HandleFunc("/get-assets", AssetsHandler)
	Log.Info(http.ListenAndServe(":80", nil))
}
