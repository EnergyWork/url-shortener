## Create

    request(post): https://host:port
    {
        "Header" : {
            "Request" : "url/create",
            "Content-Type" : "application/json",
            "Signature" : "ndasdasuidaSdnasdaoi;djasdaskldmasdadaj",
        },
        "Body" : {
            "url_long" : "https://www.some_host.ru/"
        }
        
    }

    response:
    good
    {
        "Status" : 200,
        "ShortUrl" : "cut.er/hashid"
    }
    bad
    {
        "Status" : HttpCode,
        "Error" : "text of error"
    }

## Redirect 
    
    request(get)
        https://host:port/hashid

## List of links


## Expande

    request(post): https://host:port/hashid
    {
        "shortUrl" : "https://www.some_host.ru/"
    }

    response:
    good
    {
        "status" : 200,
        "longUrl" : "https://www.some_host.ru/"
    }
    bad
    {
        "status" : HttpCode,
        "error" : "text of error"
    }