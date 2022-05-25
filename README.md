# C10K

## Environment

### Server

- Cloud Service Provider (Azure)
- Virtual machine (Standard F2s v2 (2 vcpus, 4 GiB memory))
- System (Linux c10k-server 5.15.0-1005-azure #6-Ubuntu SMP Wed Apr 20 09:27:47 UTC 2022 x86_64 x86_64 x86_64 GNU/Linux)
- Go (go version go1.18.2 linux/amd64)
- UPX (3.96-3)

### Client

- Cloud Service Provider (Azure)
- Virtual machine (Standard F2s v2 (2 vcpus, 4 GiB memory))
- System (Linux c10k-server 5.15.0-1005-azure #6-Ubuntu SMP Wed Apr 20 09:27:47 UTC 2022 x86_64 x86_64 x86_64 GNU/Linux)
- Apache benchmark tool (2.4.52-1ubuntu4)

### Result (Nginx)

```text
ubuntu@c10k-client:~$ ab -n 20000 -c 200 "http://10.2.0.4/"
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.2.0.4 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        nginx/1.18.0
Server Hostname:        10.2.0.4
Server Port:            80

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      200
Time taken for tests:   0.986 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      3260000 bytes
HTML transferred:       220000 bytes
Requests per second:    20280.91 [#/sec] (mean)
Time per request:       9.861 [ms] (mean)
Time per request:       0.049 [ms] (mean, across all concurrent requests)
Transfer rate:          3228.31 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    4   1.4      4      27
Processing:     2    6   1.0      6      12
Waiting:        0    4   1.0      4      10
Total:          6   10   1.7     10      35

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     10
  75%     10
  80%     11
  90%     11
  95%     12
  98%     14
  99%     16
 100%     35 (longest request)
```

### Result (FastHTTP)

```text
ubuntu@c10k-client:~$ ab -n 20000 -c 200 "http://10.2.0.4/"
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.2.0.4 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        fasthttp
Server Hostname:        10.2.0.4
Server Port:            80

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      200
Time taken for tests:   1.050 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      3300000 bytes
HTML transferred:       220000 bytes
Requests per second:    19045.57 [#/sec] (mean)
Time per request:       10.501 [ms] (mean)
Time per request:       0.053 [ms] (mean, across all concurrent requests)
Transfer rate:          3068.87 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    4   1.4      4      27
Processing:     2    6   1.1      6      13
Waiting:        0    4   1.1      4      10
Total:          3   10   1.8     10      35

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     11
  75%     11
  80%     11
  90%     12
  95%     13
  98%     14
  99%     17
 100%     35 (longest request)
```

### Result (Nginx and BBR)

```text
ubuntu@c10k-client:~$ ab -n 20000 -c 200 "http://10.2.0.4/"
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.2.0.4 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        nginx/1.18.0
Server Hostname:        10.2.0.4
Server Port:            80

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      200
Time taken for tests:   1.036 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      3260000 bytes
HTML transferred:       220000 bytes
Requests per second:    19313.74 [#/sec] (mean)
Time per request:       10.355 [ms] (mean)
Time per request:       0.052 [ms] (mean, across all concurrent requests)
Transfer rate:          3074.36 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    4   1.2      4      27
Processing:     2    6   1.1      6      12
Waiting:        0    5   1.1      5      10
Total:          5   10   1.5     10      34

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     11
  75%     11
  80%     11
  90%     11
  95%     12
  98%     13
  99%     14
 100%     34 (longest request)
```

### Result (FastHTTP and BBR)

```text
ubuntu@c10k-client:~$ ab -n 20000 -c 200 "http://10.2.0.4/"
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.2.0.4 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        fasthttp
Server Hostname:        10.2.0.4
Server Port:            80

Document Path:          /
Document Length:        11 bytes

Concurrency Level:      200
Time taken for tests:   1.058 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      3300000 bytes
HTML transferred:       220000 bytes
Requests per second:    18902.82 [#/sec] (mean)
Time per request:       10.580 [ms] (mean)
Time per request:       0.053 [ms] (mean, across all concurrent requests)
Transfer rate:          3045.87 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    4   1.4      4      25
Processing:     0    6   1.5      6      19
Waiting:        0    5   1.6      5      17
Total:          1   10   1.7     10      31

Percentage of the requests served within a certain time (ms)
  50%     10
  66%     11
  75%     11
  80%     11
  90%     12
  95%     13
  98%     14
  99%     16
 100%     31 (longest request)
```

### Result (FastHTTP and BBR and Postgres)

```text
# TODO
```

## LICENSE

[MIT License](LICENSE)
