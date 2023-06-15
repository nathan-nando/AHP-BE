# Decision Support System using AHP Services

# PRE RUN Required:
1. pairwise.json in `assets/`
2. env file in `config/`

# HOW TO RUN:
1. Run in local mode `go run main.go` same as `go run main.go app --appMode=local --dbMode=msysql`
2. Run in prod mode `go run main.go --appMode=prod --dbMode=mysql` 
3. Run migration `go run main.go migration --dbMode=mysql`