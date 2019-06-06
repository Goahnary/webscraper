# GoLang Web Scraper

### Getting started
1. Build the project with `go build webscraper.go` in the source directory
2. Run the webscraper with `./webscraper`
3. Enter the url of the site you want to scrape
```
Website to scrape: nytimes.com
```
4. Use the jquery selector convention to select an element or class of elements
e.g. will get the article tags from nytimes.com
```
Jquery Selector: article
```
5. Provide a filter term to reduce the results
```
Filter by: Trump

Filtered Results:
1. Trump’s Visit to Queen Mixes Pageantry With the PettyQueen Elizabeth II welcomed President Trump to Buckingham Palace with an honor guard and royal artillery salute.At the same time, Mr. Trump carried on an ugly dispute with the mayor of London, whom he called a “stone cold loser.”635 commentsPresident Trump and the first lady, Melania Trump, at Buckingham Palace with Queen Elizabeth II in London on Monday. Doug Mills/The New York Times

2. A Former Hotel Partner Alleges the Trumps Evaded Taxes in PanamaAccusations in a legal filing escalate a dispute with the president’s company and are fraught with potential diplomatic and legal complexities.The Trump Organization, the filing alleges, “made fraudulent and false claims to the Panamanian tax authorities” to “cover up its unlawful activities.”

3. The Editorial BoardTrump’s War on Worker RightsThe administration is making life easier for business owners and harder for workers.94 comments

4. Bernie SandersI Know Where I Came From. Does President Trump?

5. Letters: Lessons in Etiquette for Trump in Britain
```

Results may vary
