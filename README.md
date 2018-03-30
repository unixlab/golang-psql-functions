# golang-psql-functions
PostgreSQL extension written in Golang

# build
```build
go build -buildmode=c-shared -o golang-psql-functions.so
```

# load
```load
psql DATABASE -f functions.psql
```

# use
lastWeekday
```lastWeekday
felix=# select last_weekday();
 last_weekday 
--------------
 2018-03-29
(1 row)

```

nsToTime
```nsToTime
felix=# select ns_to_time(28000000000000);
   ns_to_time    
-----------------
 07:46:40.000000
(1 row)
```
