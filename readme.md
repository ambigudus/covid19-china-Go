## COVID-19 China Go Server

Refined Stats From COVID-19 Official Website

`Data is periodontally collected from Official Website`

#### API Server:



#### API Endpoints:

1. `GET /covid19/all`
2. `GET /covid19/date/{date}`
3. `GET /covid19/dateRange/{startDate}/{endDate}`
4. `GET /covid19/formattedData` (Optional query parameter: `startDate` and `endDate`).

   - `/covid19/formattedData?startDate=15-05-2020` (Optional endDate. Default: `Today`)
   - `/covid19/formattedData?endDate=21-05-2020` (Optional startDate. Default: `03-04-2020`)
   - `/covid19/formattedData?startDate=15-05-2020&endDate=21-05-2020`
   

Note:

1. Accepted date format: `DD-MM-YYYY` (eg. `18-04-2020`)
2. Data available from `26-06-2020`

Frontend Code: https://github.com/ambigudus/covid19-china-React



[MIT Licence](./LICENCE)
