# rep1 - Radiochild Reporting Generator



### Pre-requisites
    go 1.18 (or above)  
    golangci-lint 1.48.0 (or above)  
    gnu make 3.81 (or above)  

### Build Steps
    go get github.com/radiochild/utils@v0.1.1  
    go get github.com/radiochild/repmeta@v0.1.0  
    make build  
  

## Parameters used by rep1


|              | CLI Usage           | Server Usage        | Lambda Usage        |
| ------------ | ------------------- | ------------------- | ------------------- |
| Environment  |                     | REPORT_OUTPUT_ROOT  | REPORT_OUTPUT_ROOT  |
|              | REPORT_DB_HOST      | REPORT_DB_HOST      |                     |
|              | REPORT_DB_USER      | REPORT_DB_USER      |                     |
|              | REPORT_DB_PASSWORD  | REPORT_DB_PASSWORD  |                     |
|              | REPORT_DB_DATABASE  | REPORT_DB_DATABASE  |                     |
|              | REPORT_DB_PORT      | REPORT_DB_PORT      |                     |
|              | REP1_LOG_LEVEL      | REP1_LOG_LEVEL      | REP1_LOG_LEVEL      |
|              |                     | REP1_PORT           | <none>              |
|              |                     |                     |                     |
| Vault        | REPORT_DB_HOST      | REPORT_DB_HOST      | REPORT_DB_HOST      |
|              | REPORT_DB_USER      | REPORT_DB_USER      | REPORT_DB_USER      |
|              | REPORT_DB_PASSWORD  | REPORT_DB_PASSWORD  | REPORT_DB_PASSWORD  |
|              | REPORT_DB_DATABASE  | REPORT_DB_DATABASE  | REPORT_DB_DATABASE  |
|              | REPORT_DB_PORT      | REPORT_DB_PORT      | REPORT_DB_PORT      |
|              |                     |                     |                     |
| CLI Args     | logLevel ('INFO')   | logLevel ('INFO')   | <none>              |
|              | cmd                 | port (9119)         |                     |
|              | version             |                     |                     |
|              | outputType ('TEXT') |                     |                     |
|              | outputFile (<none>) |                     |                     |
|              | hideDetails         |                     |                     |
|              | limit (<none>)      |                     |                     |
|              |                     |                     |                     |
| HTTP Request | <none>              | specs (<none>)      | specs (<none>)      |
|              |                     | limit (<none>)      | limit (<none>)      |
|              |                     | outputType ('JSON') | outputType ('JSON') |
|              |                     | outputFile (<none>) | outputFile (<none>) |
|              |                     | hideDetails (false) | hideDetails (false) |



