## Create

    request(post): https://host:port/create
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
        "Header" : {
            "Content-Type" : "application/json",
            "Status" : 200,
        },
        "Body" : {
            "url_short" : "cut.er/hashid",
        } 
    }

    bad
    {
        "Header" : {
            "Status" : не 200,
        },
        "Body" : {
            "Error" : "texterror",
        } 
    }

## Redirect 
    
    request(get)
        https://host:port/hashid
