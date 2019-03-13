# gostresser

* Go version of simpler, config based locust <a href="https://locust.io">locust.io</a>

## Internal

* It builds a web service based on [gin](https://github.com/gin-gonic/gin), store stress test data in mongo. When you start a stress test, it create some modified [hey](https://github.com/rakyll/hey) workers, and the workers swarm requests to the targets urls. The hey workers transport request stat result to stress server through Grpc, periodically.

## Screenshots

* Stress Test Config List

![](./image/test_conf_list.png)

* Edit Stress Config

![](./image/config_edit.png)

* Start Stress Test

![](./image/start_stress.png)


## Credit:

* <a href="https://github.com/rakyll/hey">hey</a>: HTTP load generator, ApacheBench (ab) replacement, formerly known as rakyll/boom
