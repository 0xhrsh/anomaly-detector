# Anomaly Detector
Helps in finding anomalies in GreedyGame app numbers.

### Installation:
#### Requirements:
- Golang 1.14

#### Procedure:
- Clone the Repository
```
   $ git clone https://bitbucket.org/harsh-not-haarsh/anomaly-detector/src/master/
```


- Navigate to the Repository
```
   $ cd anomaly-detector
```


- Install Dependencies
```
   $ go mod download
```

- BitBucket Client
```
   Create OAuth Consumer for bitbucket workspace
```

#### Running:
-  Compile Program
```
   $ go build *.go
```
- Run the Program
```
   $ USER_ID=<user id> AUTH_TOKEN=<auth token> ENDPOINT=<endpoint url> CLIENTID=<client_id> CLIENT_SECRET=<client_secret> OWNER=<owner-of workspace> DAU_REPOS=<slug1>,<slug2> (OTHER REPOS) DAU_SVC=<SVC-1>,<SVC-2> (SVC for other parameters)  ./detector
```

#### Running:
- Repos
```
DAU_REPOS = service-collector,service-thanos-2,service-apps,service-partner-collector
IMPRESSIONS_REPOS = service-campaign,service-adgroup,service-thanos-2,service-mystique,service-units
RESPONSES_REPOS = service-campaign,service-adgroup,service-thanos-2,service-mystique,service-units
REQUESTS_REPOS = service-campaign,service-adgroup,service-thanos-2,service-mystique,service-units,service-apps 
```
- Services for Activity logger
```
DAU_SVC = SVC-APPS, SVC-COLLECTOR, SVC-THANOS 
IMPRESSIONS_SVC = SVC-CAMPAIGN, SVC-ADGROUP, SVC-THANOS, SVC-UNITS ,SVC-MYSTIQUE 
RESPONSES_SVC = SVC-CAMPAIGN, SVC-ADGROUP, SVC-THANOS, SVC-UNITS, SVC-MYSTIQUE 
REQUESTS_SVC = SVC-CAMPAIGN, SVC-ADGROUP, SVC-THANOS, SVC-MYSTIQUE, SVC-UNITS, SVC-APPS
```
