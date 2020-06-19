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
- 
```
   $ go build *.go
```
- 
```
   $ USER_ID=<user id> AUTH_TOKEN=<auth token> ENDPOINT=<endpoint url> CLIENTID=<client_id> CLIENT_SECRET=<client_secret> OWNER=<owner-of workspace> DAU_REPOS=<slug1>,<slug2> (OTHER REPOS) DAUSVC=<SVC-1>,<SVC-2> (SVC for other parameters)  ./detector
```
