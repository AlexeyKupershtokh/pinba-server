package main

import "github.com/influxdb/influxdb/client"

type Writer struct{
	Client client.Client
	DB string
	RetentionPolicy string
}

func Convert(w *Writer, req Request) client.Point {

}

func Write(w *Writer, bps client.BatchPoints) {
	bps := client.BatchPoints{
		Points:          ,
		Database:        MyDB,
		RetentionPolicy: "default",
	}
	_, err := w.Client.write(bps)
}