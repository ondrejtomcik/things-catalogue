# Public Things Catalogue

Evaluating GitHub Code Search API feasibility for the Thing Catalogue hosted on the GitHub.

## Test Scenarios

Two GitHub repositories were created. Public one with 3220 thing models and private one with 398 thing models.

Following search scenarios were tested:

- Search using a single keyword, e.g. manufacturer name

  `https://api.github.com/search/code?q=Mfg668b2fd6+in:path+repo:ondrejtomcik/things-catalogue`

- Search using a path combination, e.g. by manufacturer and family and name

  `https://api.github.com/search/code?q=Product2eeaaf74+in:file+path:Mfg668b2fd6/Family5e9dc410+repo:ondrejtomcik/things-catalogue`

- Search using a combination - code value and path value, e.g. by property and manufacturer
    `https://api.github.com/search/code?q=loadShedding+Mfg5b9b525f+in:file+language:json+repo:ondrejtomcik/things-catalogue`

- Combine multiple repositories, public one and private, search for property key

    `https://api.github.com/search/code?q=appleResistance+in:file+repo:ondrejtomcik/things-catalogue-Mfg35c56ebe+repo:ondrejtomcik/things-catalogue`

## Performance

All responses were returned in avg. 200-350ms.

## Rate Limiting

10 requests per minute.

## Authentication

Required even for public repositories.

## Conclusion

Idea of this test was to excercise the search API, as other aspects - interaction with the community, contributions, versioning, pipelines, etc. are working well. However, availability of the search API is not a hard requirement, as the catalogue should be by default shipped in form of a package with the app. Nevertheless, having the possibility to query the thing catalogue on the GitHub in case product was not found locally is a nice feature.

Overall, the search capabilities are fulfilling all the scenarios I identified. Even a combined search across more repositories, in case the company wants to keep some thing models private but combine search results with the public set, works very well. The big disadvantage, could be even a blocker, is the requirement for the access token in case of a public repository. Here I would talk to GitHub Team about this as the documentation contains information that public users should be able to access public repositories.
