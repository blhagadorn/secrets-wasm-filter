## Security Model of Secrets WASM Filter

Pros:
 - Container escape doesn't have credential inside on file system or environment variable
 
Cons: 
 - Request hijack could potentially intercept the request to a different domain either at the application or egress level
