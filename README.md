1) Create .env in root
2) Create a variables:
   * `PORT = "8080"`
   * `KPI_API_BEARER = "TOKEN_HERE"`
   * `GEN_MOCK_REQ = "TRUE"` // TRUE - If you want to generate mock requests, FALSE - if you don't want to
3) In /cmd run command `go run main.go`

By default, the API handle will be accessible at http://localhost:8080/api/v1/proxy/save_fact. You will need to provide a bearer authentication token and a JSON body in order to access the API.

```json
  "period_start":"",
  "period_end":"",
  "period_key":"",
  "indicator_to_mo_id":"",
  "indicator_to_mo_fact_id":"",
  "value":"",
  "fact_time":"",
  "is_plan":"",
  "auth_user_id":"",
  "comment": ""
```

It is possible to generate mock requests in parallel with proxy handle bombing.

In this case, the buffer will accumulate all records.
