Performance 
============================


### compare JadeRabit with Grpc-Gateway


- CentOs X64 7.6 KVM
```$xslt
CPU:      1Core
MEMORY:   4G 
IP:       172.20.23.29
```

-  Grpc Gateway report, visit grpc RESTAPI
```$xslt

λ ab -n 10000 -c 500 -p p.txt -T application/x-www-form-urlencoded "http://172.20.23.29:8080/greeter/hello"
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 172.20.72.33 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        172.20.72.33
Server Port:            8080

Document Path:          /greeter/hello
Document Length:        20 bytes

Concurrency Level:      500
Time taken for tests:   104.913 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1740000 bytes
Total body sent:        1840000
HTML transferred:       200000 bytes
Requests per second:    95.32 [#/sec] (mean)
Time per request:       5245.666 [ms] (mean)
Time per request:       10.491 [ms] (mean, across all concurrent requests)
Transfer rate:          16.20 [Kbytes/sec] received
                        17.13 kb/s sent
                        33.32 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        3   10  95.0      6    3084
Processing:    47 5144 2121.1   4078   10705
Waiting:       22 2728 2223.4   2169   10145
Total:         52 5154 2123.4   4093   10711

Percentage of the requests served within a certain time (ms)
  50%   4093
  66%   6447
  75%   6790
  80%   6916
  90%   7463
  95%   9909
  98%  10070
  99%  10138
 100%  10711 (longest request)

```



- just use gin for normal visit
```$xslt

λ D:\xampp\apache\bin\ab.exe -n 10000 -c 500  -T application/x-www-form-urlencoded "http://172.20.72.33:8090/perman"
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 172.20.72.33 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        172.20.72.33
Server Port:            8090

Document Path:          /perman
Document Length:        26 bytes

Concurrency Level:      500
Time taken for tests:   94.499 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1490000 bytes
HTML transferred:       260000 bytes
Requests per second:    105.82 [#/sec] (mean)
Time per request:       4724.967 [ms] (mean)
Time per request:       9.450 [ms] (mean, across all concurrent requests)
Transfer rate:          15.40 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        3    9  79.3      6    3012
Processing:    88 4469 1582.7   3622    9622
Waiting:       15 2454 1714.7   2188    6669
Total:         94 4478 1583.6   3629    9626

Percentage of the requests served within a certain time (ms)
  50%   3629
  66%   4102
  75%   5093
  80%   6541
  90%   6631
  95%   6686
  98%   9559
  99%   9575
 100%   9626 (longest request)

```



### summary
grpc-gateway performan is close to gin(run the normal "echo" service)



- JadeRabit archit now 
```$xslt

λ D:\xampp\apache\bin\ab.exe -n 10000 -c 500  -T application/x-www-form-urlencoded "http://172.20.72.33:8090/echo"
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 172.20.72.33 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        172.20.72.33
Server Port:            8090

Document Path:          /echo
Document Length:        38 bytes

Concurrency Level:      500
Time taken for tests:   102.680 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1610000 bytes
HTML transferred:       380000 bytes
Requests per second:    97.39 [#/sec] (mean)
Time per request:       5134.020 [ms] (mean)
Time per request:       10.268 [ms] (mean, across all concurrent requests)
Transfer rate:          15.31 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        3   10  95.2      6    3152
Processing:   118 5009 2053.1   3692    9604
Waiting:       12 2682 2128.1   2140    9506
Total:        128 5019 2054.3   3699    9610

Percentage of the requests served within a certain time (ms)
  50%   3699
  66%   6467
  75%   6633
  80%   6730
  90%   7140
  95%   9514
  98%   9540
  99%   9555
 100%   9610 (longest request)

```





