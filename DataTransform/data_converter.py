#!/usr/bin/env python
# -*- coding: utf-8 -*-
#

import sqlite3
from datetime import datetime, date
import csv

databaspath = '../Car2GoDataMock/historyDatabase'

def get_data():
    conn = sqlite3.connect(databaspath)
    c = conn.cursor()
    query = "SELECT time, lonCurrent, latCurrent FROM HistoryData"
    c.execute(query)
    return c
    
def get_data_env():
    conn = sqlite3.connect(databaspath)
    c = conn.cursor()
    query = "SELECT monthDay, weekday, hour, isRush, isGoodWeather, isFallout, isWeekend, EventID, isPriceHigher, doesPolicyApply FROM EnvData"
    c.execute(query)
    return c

def squeeze_time(time):
    dt = datetime.fromtimestamp(time)
    return dt.hour

def squeeze_yearday(time):
    dayofyear = int((date.fromtimestamp(time) - date(2017,10,1)).days)
    print(dayofyear)
    return dayofyear

def in_zone(lon, lat, zone):
    x1,y1,x2,y2 = zone
    if x1 <= lat and lat <= x2:
        if y1 <= lon and lon <= y2:
            return True
    return False


def squeeze_zone(lon, lat):
    min_lat = 40.741549
    max_lat = 40.816033
    
    min_lon = -74.006140
    max_lon = -73.932326
    
    zone_grid = [2,4]
    
    x_step = (max_lat - min_lat) / zone_grid[0]
    y_step = (max_lon - min_lon) / zone_grid[1]
    
    zones = {}
    zone_idx = 0
    
    for x in range(zone_grid[0]):
        for y in range(zone_grid[1]):
            x1 = min_lat + x * x_step
            y1 = min_lon + y * y_step
            x2 = x1 + x_step
            y2 = y1 + y_step
            
            zone = (x1,y1,x2,y2)
            zones[zone_idx] = zone 
            zone_idx += 1
    
    for key, value in zones.iteritems():
        if in_zone(lon, lat, value):
            return key
    
    return None

def sum_data(rows):
    summed_data = {}
    for row in rows:
        time, lat, lon = row
        hour = squeeze_time(time)
        yearday = squeeze_yearday(time)
        zone = squeeze_zone(lon, lat)
        idx = (yearday, hour, zone)
        summed_data[idx] = summed_data.get(idx, 0) + 1
    return summed_data

def expand_value(val, count, minval = 0):
    lst = [0] * count
    lst[val - minval] = 1
    return lst

def gen_csv_env(rows):
    with open('env.csv', 'wb') as csvfile:
        csvwriter = csv.writer(csvfile, delimiter=',', quotechar='|', quoting=csv.QUOTE_MINIMAL)
        for row in rows:
            monthday, weekday, hour, isRush, isGoodWeather, isFallout, isWeekend, EventID = row
            events = expand_value(EventID, count = 3)
            csvwriter.writerow([monthday-1, min(weekday, 6), hour-1, isRush, isGoodWeather, isFallout, isWeekend] + events)
    
def gen_csv(processed_data):
    with open('requests.csv', 'wb') as csvfile:
        csvwriter = csv.writer(csvfile, delimiter=',', quotechar='|', quoting=csv.QUOTE_MINIMAL)
        for key, value in processed_data.iteritems():
            hour, zone = key
            requests = value
            csvwriter.writerow((hour, zone, requests))

def gen_csv_all(reqdata, envdata):
    with open('all.csv', 'wb') as csvfile:
        csvwriter = csv.writer(csvfile, delimiter=',', quotechar='|', quoting=csv.QUOTE_MINIMAL)
        for row in envdata:
            yearday, weekday, hour, isRush, isGoodWeather, isFallout, isWeekend, EventID, isPriceHigher, doesPolicyApply = row
            events = expand_value(min(EventID - 1, 0), count = 2)
            for zoneidx in range(8):
                idx = (yearday, hour, zoneidx)
                if idx in reqdata:
                    requests = reqdata[idx]
                else:
                    requests = 0
                csvwriter.writerow([yearday-1, min(weekday, 6), hour-1, isRush, isGoodWeather, isFallout, isWeekend] + events + [requests, zoneidx, isPriceHigher, doesPolicyApply])

def main():
    req_data = get_data()
    processed_req_data = sum_data(req_data)

    #gen_csv(processed_data)
    
    env_data = get_data_env()
    #gen_csv_env(env_data)
    gen_csv_all(processed_req_data, env_data)
    

if __name__ == "__main__":
    main()
