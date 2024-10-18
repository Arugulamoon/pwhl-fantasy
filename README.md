# pwhl-fantasy

Project with my wife to build out a PWHL (Professional Women's Hockey League) Fantasy League. Built using Golang. League starts in December. Intention to build out features during season. Start with Basic Goals/Assists; add in goalie metrics; additional player metrics over time to make more interesting.

## Functionality so far...
- Basic http web server
- Basic html templates for:
  - Listing Fantasy Teams
  - Creating a Fantasy Teams
  - Listing Players (hardcoded for now)
- Filesystem persistence

## TODOs:
- Ingest Team Player Lists from Training Camp rosters; update as rosters are solidified
- Add Player drafting
- Add Standings
- Make use of existing colly (html scraping package) code for fetching gamesheet data
- Explore options for ingesting Player and Gamesheets from APIs instead of webpage scraping (will cost some subscription money likely)
- Authentication and Authorization (use existing from another project)
- Client-side validation
- Richer css experience/Replace basic golang templating with something like React
- DB persistence

## Unit Tests
```
go test ./...
```

## Run
### On Linux/Mac/Windows Git Bash
```
go run ./cmd/web
```
### On Windows
```
go run .\cmd\web
```

Visit: http://localhost:9798/

## References
- Alex Edward's "Let's Go" online book: https://lets-go.alexedwards.net/
- PWHL Hockey: https://www.thepwhl.com/en/
